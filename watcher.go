package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// Watcher represents a file watcher
type Watcher struct {
	fsWatcher *fsnotify.Watcher
	source    string
	dest      string
	isDir     bool
}

// NewWatcher creates a new file watcher
func NewWatcher(source, dest string) (*Watcher, error) {
	// Check if source exists
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return nil, err
	}

	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		fsWatcher: fsWatcher,
		source:    source,
		dest:      dest,
		isDir:     sourceInfo.IsDir(),
	}, nil
}

// Close closes the file watcher
func (w *Watcher) Close() {
	w.fsWatcher.Close()
}

// IsSourceDir returns true if the source is a directory
func (w *Watcher) IsSourceDir() bool {
	return w.isDir
}

// StartWatching starts watching for file changes
func (w *Watcher) StartWatching() error {
	// Add the source to the watcher
	if w.isDir {
		err := w.addDirToWatch(w.source)
		if err != nil {
			return err
		}
		log.Printf("Watching directory: %s", w.source)
	} else {
		err := w.fsWatcher.Add(w.source)
		if err != nil {
			return err
		}
		log.Printf("Watching file: %s", w.source)
	}

	// Start watching for events
	go w.handleEvents()

	return nil
}

// PerformInitialCopy performs an initial copy of all files if source is a directory
func (w *Watcher) PerformInitialCopy() error {
	if w.isDir {
		err := CopyFilesInDir(w.source, w.dest)
		if err != nil {
			log.Printf("Error during initial copy: %v", err)
			return err
		}
		log.Printf("Initial copy completed from %s to %s", w.source, w.dest)
	}
	return nil
}

// handleEvents processes file change events
func (w *Watcher) handleEvents() {
	for {
		select {
		case event, ok := <-w.fsWatcher.Events:
			if !ok {
				return
			}
			w.processEvent(event)
		case err, ok := <-w.fsWatcher.Errors:
			if !ok {
				return
			}
			log.Printf("Watcher error: %v", err)
		}
	}
}

// processEvent processes a single file event
func (w *Watcher) processEvent(event fsnotify.Event) {
	if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
		// We only care about write and create events
		info, err := os.Stat(event.Name)
		if err != nil {
			log.Printf("Error accessing changed file %s: %v", event.Name, err)
			return
		}

		// Skip directories for write events
		if info.IsDir() {
			// If a new directory is created, watch it
			if event.Has(fsnotify.Create) {
				w.addDirToWatch(event.Name)
			}
			return
		}

		// Get the relative path if source is a directory
		var relPath string
		if w.isDir {
			relPath, err = filepath.Rel(w.source, event.Name)
			if err != nil {
				log.Printf("Error getting relative path: %v", err)
				return
			}
		} else {
			// If source is a file, just use the filename
			relPath = filepath.Base(event.Name)
		}

		// Construct destination path
		destPath := filepath.Join(w.dest, relPath)

		// Ensure destination directory exists
		destDir := filepath.Dir(destPath)
		err = os.MkdirAll(destDir, 0755)
		if err != nil {
			log.Printf("Error creating destination directory %s: %v", destDir, err)
			return
		}

		// Copy the file
		err = CopyFile(event.Name, destPath)
		if err != nil {
			log.Printf("Error copying %s to %s: %v", event.Name, destPath, err)
			return
		}

		log.Printf("Copied %s to %s", event.Name, destPath)
	}
}

// addDirToWatch adds a directory and all its subdirectories to the watcher
func (w *Watcher) addDirToWatch(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return w.fsWatcher.Add(path)
		}
		return nil
	})
}
