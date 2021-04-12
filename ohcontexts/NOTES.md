# NOTES

## Context

Ref: https://blog.golang.org/context

Ref: https://golang.org/pkg/context/

It is a concurrent pattern that is generally used to manage long-lived processes.

It was developed at Google to pass request-scoped values (cancelaton signals, deadlines) across API boundaries to all goroutines involved in the span of that request.

How to use them?

- Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
- The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue.
- When a Context is canceled, all Contexts derived from it are also canceled.
