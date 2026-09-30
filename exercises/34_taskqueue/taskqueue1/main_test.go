package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestRunAddsAndListsInOrder(t *testing.T) {
	var out bytes.Buffer
	queue := NewMemory()
	if err := run([]string{"add", "-priority", "2", "write docs"}, &out, queue); err != nil {
		t.Fatalf("add with priority: %v", err)
	}
	if err := run([]string{"add", "review pr"}, &out, queue); err != nil {
		t.Fatalf("add without priority: %v", err)
	}

	out.Reset()
	if err := run([]string{"list"}, &out, queue); err != nil {
		t.Fatalf("list: %v", err)
	}
	const want = "1 [open] write docs\n2 [open] review pr\n"
	if out.String() != want {
		t.Fatalf("list output = %q, want %q", out.String(), want)
	}
}

func TestRunCompletesATask(t *testing.T) {
	var out bytes.Buffer
	queue := NewMemory()
	if err := run([]string{"add", "ship it"}, &out, queue); err != nil {
		t.Fatalf("add: %v", err)
	}
	if err := run([]string{"done", "1"}, &out, queue); err != nil {
		t.Fatalf("done: %v", err)
	}

	out.Reset()
	if err := run([]string{"list"}, &out, queue); err != nil {
		t.Fatalf("list: %v", err)
	}
	const want = "1 [done] ship it\n"
	if out.String() != want {
		t.Fatalf("list output = %q, want %q", out.String(), want)
	}
}

func TestRunReportsARejectedTask(t *testing.T) {
	var out bytes.Buffer
	err := run([]string{"add", "-priority", "-1", "  "}, &out, NewMemory())
	if !errors.Is(err, ErrTitleRequired) {
		t.Fatalf("add with a blank title error = %v, want ErrTitleRequired", err)
	}
	if !errors.Is(err, ErrPriorityNegative) {
		t.Fatalf("add with a negative priority error = %v, want ErrPriorityNegative", err)
	}
	if out.Len() != 0 {
		t.Fatalf("a rejected task wrote %q to the command output", out.String())
	}
}

func TestRunReportsBadCommands(t *testing.T) {
	var out bytes.Buffer
	if err := run([]string{"done", "7"}, &out, NewMemory()); !errors.Is(err, ErrNotFound) {
		t.Fatalf("done 7 error = %v, want ErrNotFound", err)
	}
	if err := run([]string{"done", "abc"}, &out, NewMemory()); err == nil {
		t.Fatal("done with a non-numeric id returned no error")
	}
	if err := run([]string{"explode"}, &out, NewMemory()); err == nil {
		t.Fatal("an unknown subcommand returned no error")
	}
	if err := run(nil, &out, NewMemory()); err == nil {
		t.Fatal("running without a subcommand returned no error")
	}
}

func TestValidateReportsEveryProblem(t *testing.T) {
	err := Task{Title: "   ", Priority: -1}.Validate()
	if !errors.Is(err, ErrTitleRequired) {
		t.Fatalf("Validate() error = %v, want ErrTitleRequired", err)
	}
	if !errors.Is(err, ErrPriorityNegative) {
		t.Fatalf("Validate() error = %v, want ErrPriorityNegative", err)
	}
	if err := (Task{Title: "fine"}).Validate(); err != nil {
		t.Fatalf("Validate() on a good task = %v, want nil", err)
	}
}

func TestUnknownStatusRendersDistinctly(t *testing.T) {
	unknown := Status("archived").String()
	if unknown == "" || unknown == Open.String() || unknown == Done.String() {
		t.Fatalf("unknown status rendered as %q, want something distinct", unknown)
	}
}

func TestMemoryListReturnsACopy(t *testing.T) {
	queue := NewMemory()
	if _, err := queue.Add("keep me", 0); err != nil {
		t.Fatalf("add: %v", err)
	}
	listed := queue.List()
	if len(listed) != 1 {
		t.Fatalf("List() returned %d tasks, want 1", len(listed))
	}
	listed[0].Title = "changed"
	if got := queue.List()[0].Title; got != "keep me" {
		t.Fatalf("List() handed out the store's own slice: the title is now %q", got)
	}
}

func TestFileStoreRoundTrips(t *testing.T) {
	file := FileStore{Path: filepath.Join(t.TempDir(), "queue.json")}
	tasks := []Task{
		{ID: 1, Title: "ship it", Priority: 3, Status: Done},
		{ID: 2, Title: "write docs", Status: Open},
	}
	if err := file.Save(tasks); err != nil {
		t.Fatalf("Save: %v", err)
	}
	loaded, err := file.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !reflect.DeepEqual(loaded, tasks) {
		t.Fatalf("Load() = %+v, want %+v", loaded, tasks)
	}
}

func TestFileStoreTreatsAMissingFileAsAnEmptyQueue(t *testing.T) {
	file := FileStore{Path: filepath.Join(t.TempDir(), "queue.json")}
	loaded, err := file.Load()
	if err != nil {
		t.Fatalf("Load() on a missing file = %v, want no error", err)
	}
	if len(loaded) != 0 {
		t.Fatalf("Load() on a missing file = %+v, want an empty queue", loaded)
	}
}

func TestFileStoreWritesNothingForAnEmptyQueue(t *testing.T) {
	path := filepath.Join(t.TempDir(), "queue.json")
	if err := (FileStore{Path: path}).Save(nil); err != nil {
		t.Fatalf("Save(nil): %v", err)
	}
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("saving an empty queue created %s", path)
	}
}
