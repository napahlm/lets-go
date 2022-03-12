## Concurrency in Golang

Resources.[^1]

### Why concurrency?

Distinction: In Go you talk to channels.

### Goroutine
Goroutines exit if nor deferred. Independently executed function (which mean you must communicate between them).
- Own stack which grows and shrinks automatically
- Prqactical to have thousands even
- Not a thread

> go func

### Channels
Channels are how one communicates between goroutines.

Communication through channels can be thought of as a synchronization operation. *Channels blocks both ways; a receiver waiting for channel data or a channel waiting for someone to receive.*
Initialization:

> var c chan int
> c = make(chan int)

Receive data on channel:

> c <- data

Send data on channel:

> data := <-c

*But what about **buffered** channels?*

- They don't synchronize.
- Usueful for very certain problems.

### The Go approach to concurrent software
> *Don't communicate by sharing memory, share memory by communicating." - Rob Pike

"Patterns"
: Function that returns a channel














### Resources
[^1]: https://www.youtube.com/watch?v=f6kdp27TYZs