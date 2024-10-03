// Mostly copied from go source at tip, commit d922c0a.
//
// Copyright 2015 The Go Authors. All rights reserved.

package doc

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gnolang/gno/gno/pkg/gnomod"
)

// A bfsDir describes a directory holding code by specifying
// the expected import path and the file system directory.
type bfsDir struct {
	importPath string // import path for that dir
	dir        string // file system directory
}

// dirs is a structure for scanning the directory tree.
// Its Next method returns the next Go source directory it finds.
// Although it can be used to scan the tree multiple times, it
// only walks the tree once, caching the data it finds.
type bfsDirs struct {
	scan   chan bfsDir // Directories generated by walk.
	hist   []bfsDir    // History of reported Dirs.
	offset int         // Counter for Next.
}

// newDirs begins scanning the given stdlibs directory.
// dirs are "gopath-like" directories, such as @/gno/stdlibs, whose path
// relative to the root specify the import path.
// modDirs are user directories, expected to have gno.mod files
func newDirs(dirs []string, modDirs []string) *bfsDirs {
	d := &bfsDirs{
		hist: make([]bfsDir, 0, 256),
		scan: make(chan bfsDir),
	}

	roots := make([]bfsDir, 0, len(dirs)+len(modDirs))
	for _, dir := range dirs {
		roots = append(roots, bfsDir{
			dir:        dir,
			importPath: "",
		})
	}

	for _, mdir := range modDirs {
		gm, err := gnomod.ParseGnoMod(filepath.Join(mdir, "gno.mod"))
		if err != nil {
			log.Printf("%v", err)
			continue
		}
		roots = append(roots, bfsDir{
			dir:        mdir,
			importPath: gm.Module.Mod.Path,
		})
		roots = append(roots, getGnoModDirs(gm)...)
	}

	go d.walk(roots)
	return d
}

func getGnoModDirs(gm *gnomod.File) []bfsDir {
	// cmd/go makes use of the go list command, we don't have that here.

	dirs := make([]bfsDir, 0, len(gm.Require))
	for _, r := range gm.Require {
		mv := gm.Resolve(r)
		path := gnomod.PackageDir("", mv)
		if _, err := os.Stat(path); err != nil {
			// only give directories which actually exist and don't give
			// an error when accessing
			if !os.IsNotExist(err) {
				log.Println("open source directories from gno.mod:", err)
			}
			continue
		}
		dirs = append(dirs, bfsDir{
			importPath: mv.Path,
			dir:        path,
		})
	}

	return dirs
}

// Reset puts the scan back at the beginning.
func (d *bfsDirs) Reset() {
	d.offset = 0
}

// Next returns the next directory in the scan. The boolean
// is false when the scan is done.
func (d *bfsDirs) Next() (bfsDir, bool) {
	if d.offset < len(d.hist) {
		dir := d.hist[d.offset]
		d.offset++
		return dir, true
	}
	dir, ok := <-d.scan
	if !ok {
		return bfsDir{}, false
	}
	d.hist = append(d.hist, dir)
	d.offset++
	return dir, ok
}

// walk walks the trees in the given roots.
func (d *bfsDirs) walk(roots []bfsDir) {
	for _, root := range roots {
		d.bfsWalkRoot(root)
	}
	close(d.scan)
}

// bfsWalkRoot walks a single directory hierarchy in breadth-first lexical order.
// Each Go source directory it finds is delivered on d.scan.
func (d *bfsDirs) bfsWalkRoot(root bfsDir) {
	root.dir = filepath.Clean(root.dir)

	// this is the queue of directories to examine in this pass.
	this := []bfsDir{}
	// next is the queue of directories to examine in the next pass.
	next := []bfsDir{root}

	for len(next) > 0 {
		this, next = next, this[:0]
		for _, dir := range this {
			fd, err := os.Open(dir.dir)
			if err != nil {
				log.Print(err)
				continue
			}

			// read dir entries.
			entries, err := fd.ReadDir(0)
			fd.Close()
			if err != nil {
				log.Print(err)
				continue
			}

			// stop at module boundaries
			if dir.dir != root.dir && containsGnoMod(entries) {
				continue
			}

			hasGnoFiles := false
			for _, entry := range entries {
				name := entry.Name()
				// For plain files, remember if this directory contains any .gno
				// source files, but ignore them otherwise.
				if !entry.IsDir() {
					if !hasGnoFiles && strings.HasSuffix(name, ".gno") {
						hasGnoFiles = true
					}
					continue
				}
				// Entry is a directory.

				// Ignore same directories ignored by the go tool.
				if name[0] == '.' || name[0] == '_' || name == "testdata" {
					continue
				}
				// Remember this (fully qualified) directory for the next pass.
				next = append(next, bfsDir{
					dir:        filepath.Join(dir.dir, name),
					importPath: path.Join(dir.importPath, name),
				})
			}
			if hasGnoFiles {
				// It's a candidate.
				d.scan <- dir
			}
		}
	}
}

func containsGnoMod(entries []os.DirEntry) bool {
	for _, entry := range entries {
		if entry.Name() == "gno.mod" && !entry.IsDir() {
			return true
		}
	}
	return false
}

// findPackage finds a package iterating over d where the import path has
// name as a suffix (which may be a package name or a fully-qualified path).
// returns a list of possible directories. If a directory's import path matched
// exactly, it will be returned as first.
func (d *bfsDirs) findPackage(name string) []bfsDir {
	d.Reset()
	candidates := make([]bfsDir, 0, 4)
	for dir, ok := d.Next(); ok; dir, ok = d.Next() {
		// want either exact matches or suffixes
		if dir.importPath == name || strings.HasSuffix(dir.importPath, "/"+name) {
			candidates = append(candidates, dir)
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		// prefer shorter paths -- if we have an exact match it will be of the
		// shortest possible pkg path.
		ci := strings.Count(candidates[i].importPath, "/")
		cj := strings.Count(candidates[j].importPath, "/")
		if ci != cj {
			return ci < cj
		}
		// use alphabetical ordering otherwise.
		return candidates[i].importPath < candidates[j].importPath
	})
	return candidates
}

// findDir determines if the given absdir is present in the Dirs.
// If not, the nil slice is returned. It returns always at most one dir.
func (d *bfsDirs) findDir(absdir string) []bfsDir {
	d.Reset()
	for dir, ok := d.Next(); ok; dir, ok = d.Next() {
		if dir.dir == absdir {
			return []bfsDir{dir}
		}
	}
	return nil
}
