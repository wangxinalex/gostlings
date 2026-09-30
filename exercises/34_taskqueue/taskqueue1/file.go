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

// FileStore persists a queue as one JSON document.
type FileStore struct {
	Path string
}

// Save writes the whole queue, or nothing when the queue is empty.
func (f FileStore) Save(tasks []Task) error {
	// TODO: encode the queue and write it with an explicit permission mode; an
	//       empty queue must leave no file behind.
	return nil
}

// Load reads the queue, treating a missing file as an empty queue.
func (f FileStore) Load() ([]Task, error) {
	// TODO: read and decode the document, returning an empty queue instead of an
	//       error when the file does not exist yet.
	return nil, nil
}
