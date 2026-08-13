# Errors: error handling progression

Work through the exercises in order. They are ordered by how often each
error-handling technique appears in real code: first return and check errors,
then preserve the cause across layers, and finally practice validation,
retries, and application-boundary conversion.

| Exercise | Focus | Goal |
| --- | --- | --- |
| errors1 | Return and check error | Callers must handle the failure branch |
| errors2 | Input validation | Use a non-nil error for invalid input |
| errors3 | sentinel + `%w` + `errors.Is` | Wrap an error while keeping its classification |
| errors4 | Custom error + `errors.As` | Extract structured error information |
| errors5 | Standard-library error types + `errors.As` | Preserve the underlying `strconv.NumError` |
| errors6 | `errors.Join` | Report several independent validation errors at once |
| errors7 | Custom wrapper + `Unwrap` | Let business errors participate in `errors.Is` |
| errors8 | Application-boundary conversion | Map internal errors to stable user-facing semantics |
| errors9 | Retryable errors | Retry only failures that explicitly allow it |
| errors10 | Error composition capstone | Recognize multiple causes after cross-layer wrapping |

While solving, always keep two things separate: the error string is for logs
and diagnostics, and `errors.Is`/`errors.As` is for program logic. Unless an
exercise explicitly asks you to show text, do not classify errors by comparing
`err.Error()`.

Reference: [Go Error Handling](https://go.dev/doc/effective_go#errors),
[`errors` package](https://pkg.go.dev/errors).
