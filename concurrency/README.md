## Concurrency in Golang

Resources.[^1]

### Why concurrency?

Distinction: In Go you talk to channels.

### Concurrency vs. Parallellism

**Concurrency**: A composition of independently executing processes. ***Dealing** with lots of things at once*.

**Parallellism**: Simultaneous executing of multiple things; *possibly related, possbily not*. ***Doing** with lots of things at once*.

Concurrency need well-structured communication to work.

Gophers are cheaper than threads - Google in practice uses millions at the same time!

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
> "*"Don't communicate by sharing memory, share memory by communicating.*"
 - Rob Pike

"Patterns"
 - Function that returns a channel

### Multiplexing

Fan-in solutions let's several goroutines count together in lockstep. *Lockstep means no one talks at the exact same time*. It decouples the execution of each goroutine's task.

### Sequencing

Pass a channel to the channel that tells the processes to wait. *Like a wait-channel*.

### Select

Unique to concurrency. Each case in select-case is a communication. It starts at the top like a switch-case, and then evaluates all the channels it is communicating with and blocks until one of them is ready to communicate. A default case can be implemented and is executed if all other cases are blocking, but without it the slect-case will continuously block. The order of messages received simultaneously are picked pseudo-randomly.

```
select {
case variable1 := <-channel1:
    fmt.Printf("Received: %v from channel 1.\n", variable1)
case variable2 := <-channel2:
    fmt.Printf("Received: %v from channel 1.\n", variable1)
default:
    fmt.Printf("Waiting for messages.")
}
```

Things that can be done with **select**:
- Wrap it with a for{} to listen forever.
- Timeout the select-statement with "case <-time.After()".
- Quit the select with a case listening to a channel then executing a return. *Send a value to the exit channel that the other part can read like an error message*.

### Work around timeouts with replication AKA *spamming the server with requests*

Timeouts can be very useful when working with communication in case something takes too long or there are deadlines. However, what if the timeout discards important data due to slow servers? You can send more request simultaneously and choose the fastest one that returns. Replicas of the search reduces *tail latency*.












### Resources
[^1]: Go Concurrency Patterns: https://www.youtube.com/watch?v=f6kdp27TYZs