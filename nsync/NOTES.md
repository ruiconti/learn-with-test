# NOTES

## Sync

Ref: https://golang.org/pkg/sync/

Ref: https://github.com/golang/go/wiki/MutexOrChannel

This module introduces the `sync` package. It provides basic synchronization primitives between concurrent executions. Such as mutual exclusion locks (aka mutexes).

Mutexes' way of communication can be a contrast to CSP (channels), which is the preferred and more elegant way. In Pike's words:

> Share memory by communicating, don't communicate by sharing memory.

However, it is noted that a [common Go newbie mistake](https://github.com/golang/go/wiki/MutexOrChannel) is to over-use channels. You shouldn't be afraid to use a mutex if that fits the problem best.

Therefore, Go does have a general opinon on when to use channels and when to use locks.

1. Channels: passing ownership of data, distributing units of work and communicating async results
2. Mutex: caches, state


### WaitGroup

It's a counter that keeps track of a collection of goroutines to finish. You can set how many goroutines you wish to wait for with `Add`. Then, as each goroutine finishes its job, it calls `Done`, which decrements a unit on the counter.

Finally, you can call `Wait()` to wait until counter reaches 0. Similar to Python's `gather`. Also, semantically related to the `wait()` system call. Related because the system call does that implementing a queue on forked childs.

### Mutexes

Nevertheless, we can use mutexes to coordinate the write on a common structure.

Thumb rules:

>  A mutex must not be copied after its first use.

This means that it must always be passed by reference instead of as value. Because go automatically tries to copy when passing by value.

This also is true for every container (`type`, etc) that holds a mutex.
