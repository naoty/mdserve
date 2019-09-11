package contents

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// Watcher represents a watcher of changes in contents.
type Watcher struct {
	*fsnotify.Watcher
	logger *log.Logger
}

// NewWatcher returns a new Watcher.
func NewWatcher(logger *log.Logger) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{Watcher: watcher, logger: logger}, nil
}

// AddAll adds contents in passed dir to watch list.
func (w *Watcher) AddAll(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		w.Watcher.Add(path)
		return nil
	})
}

// Start starts watching contents.
func (w *Watcher) Start() {
	w.logger.Println("watching contents")

	go func() {
		for {
			select {
			case event, ok := <-w.Events:
				if !ok {
					return
				}
				w.logger.Println(event)
			case err, ok := <-w.Errors:
				if !ok {
					return
				}
				w.logger.Println(err)
			}
		}
	}()
}
