package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Add version flag
	versionFlag := flag.Bool("version", false, "Show version information")
	flag.Parse()

	// Show version if requested
	if *versionFlag {
		fmt.Printf("cpw version %s\n", GetVersionInfo())
		return
	}

	// Check command-line arguments
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: cpw <source> <destination>")
		fmt.Println("       cpw -version")
		os.Exit(1)
	}

	source := args[0]
	destination := args[1]

	// Check if source exists
	sourceInfo, err := os.Stat(source)
	if err != nil {
		log.Fatalf("Error accessing source %s: %v", source, err)
	}

	// Check if destination exists, create if it's a directory
	destInfo, err := os.Stat(destination)
	if err != nil {
		if os.IsNotExist(err) {
			// If source is a dir, create destination dir
			if sourceInfo.IsDir() {
				err = os.MkdirAll(destination, 0755)
				if err != nil {
					log.Fatalf("Error creating destination directory %s: %v", destination, err)
				}
			}
		} else {
			log.Fatalf("Error accessing destination %s: %v", destination, err)
		}
	} else if sourceInfo.IsDir() && !destInfo.IsDir() {
		log.Fatalf("Source is a directory but destination is not")
	}

	// Create new watcher
	watcher, err := NewWatcher(source, destination)
	if err != nil {
		log.Fatalf("Error creating watcher: %v", err)
	}
	defer watcher.Close()

	// Start watching
	err = watcher.StartWatching()
	if err != nil {
		log.Fatalf("Error starting watcher: %v", err)
	}

	// Perform initial copy
	err = watcher.PerformInitialCopy()
	if err != nil {
		log.Printf("Warning: Initial copy had errors: %v", err)
	}

	// Wait forever
	fmt.Printf("CPW %s | File watcher started. Press Ctrl+C to stop.\n", Version)
	<-make(chan struct{})
}
