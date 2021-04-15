# Notes

## On TDD

A few takeaways.

### Speed

The approach emphasizes to get working software *as quickly as possible*. This builds up momentum
and enables work to be split in small iterative improvements.

Why?

1. Keep the problem space small at any given time — divide and conquer
2. Avoid going down in rabbit holes
3. If ever getting stuck, there are tests to back-up; minimizes rework

> Make the test work quickly, committing **whatever sins** necesary in the process.
> Beck, K.

In theoretical terms: Get out of red as soon as you can :-)


## HTTP

Ref: https://golang.org/pkg/net/http/

Ref: https://golang.org/pkg/net/http/httptest/

### HandlerFunc

The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

```go
type HandlerFunc func(ResponseWriter, *Request)
```

### ListenAndServe

ListenAndServe takes a port to listen on a Handler. If there is a problem the web server will return an error, an example of that might be the port already being listened to. For that reason we wrap the call in log.Fatal to log the error to the user.


