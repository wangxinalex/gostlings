// Concept: persistence is a second Store implementation, not a special case.
// Task: implement Save and Load on the file store.
// Expected behavior: focused tests pass (run `go test ./exercises/34_taskqueue/taskqueue1`)
// Hint: the file holds one JSON document for the whole queue. Writing uses an
//       explicit permission mode, an empty queue is not worth a file, and reading a
//       path that does not exist yet is an empty queue rather than a failure,
//       because the first command a user runs should not be a special case.
// Stuck?: encoding/json turns the slice into bytes; os.ErrNotExist separates
//       "no file yet" from a real read failure.

package main

import (
	"encoding/json"
	"errors"
	"os"
)

// FileStore persists a queue as one JSON document.
type FileStore struct {
	Path string
}

// Save writes the whole queue, or nothing when the queue is empty.
func (f FileStore) Save(tasks []Task) error {
	if len(tasks) == 0 {
		return nil
	}
	encoded, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.Path, append(encoded, '\n'), 0o644)
}

// Load reads the queue, treating a missing file as an empty queue.
func (f FileStore) Load() ([]Task, error) {
	encoded, err := os.ReadFile(f.Path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var tasks []Task
	if err := json.Unmarshal(encoded, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
