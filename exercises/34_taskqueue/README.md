# Task queue: the multi-file capstone

Every other chapter asks for a function. This one asks for a program: a small
command-line task queue built from five files that have to agree with each other.
It is the only exercise in the repository where the learner designs more than one
layer, wires them together, and tests the result end to end.

## What to build

```text
taskqueue add -priority 2 "write docs"   prints the created task
taskqueue list                           prints every task, oldest first
taskqueue done 1                         marks task 1 complete
```

## Layers

| File | Owns | Depends on |
| --- | --- | --- |
| `task.go` | the task type, its statuses, and validation | nothing in the project |
| `store.go` | ordering, id assignment, the copy boundary, and lookup failure | `task.go` |
| `file.go` | persisting the queue as one JSON document | `task.go` |
| `main.go` | flag parsing, subcommand dispatch, and output | all of the above |
| `main_test.go` | end-to-end behaviour through the command layer | as a user would |

Each file holds its own `TODO` seams. The dependency direction is the point: the
domain file knows nothing about the store, and the command layer duplicates
neither validation nor lookup rules — it reports what the lower layers return.

## What the tests check

The focused tests cover the behaviour a user would notice: listing in insertion
order, completing a task, both validation causes surviving as `errors.Is`
targets, a missing id reported as a distinct error, a listing that cannot be
used to rewrite stored state, a JSON round trip through `t.TempDir`, and reading
a file that does not exist yet as an empty queue instead of a failure.

## Why one package

Go forbids importing an `internal` package from outside its parent tree, and
`scripts/verify_solutions.sh` mirrors every exercise into `solutions/`, so a
sub-package would have one import path in the exercise tree and a different one
in the solution tree — the overlaid test could not compile in both. The layers
are therefore files in one package rather than packages. Teaching package
boundaries properly needs a harness change and is tracked as follow-up work.

## Commands

```sh
sh check.sh exercises/34_taskqueue/taskqueue1
go test -race ./exercises/34_taskqueue/taskqueue1
sh scripts/verify_exercise_starters.sh exercises/34_taskqueue
```
