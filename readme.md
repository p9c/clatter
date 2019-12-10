# clatter
 
*noun* 

1. A loud percussive noise that is strident usually undesirable and 
   disruptive
2. One of the name of a group of Jackdaws (corvus monedula) referring
   to the clipped chirping noises they make in large groups from time
   to time.

## What is this thing?

Clatter is a simple framework for creating finite state automata used
on the Concurrent Sequential Processes model used in Newsqueak and Go
channels.

Clatter operates similarly to modular routing systems like used in 
software sound mixing systems, but for any kind of data at all, and 
inspired by the Newsqueak GUI.

## Architecture

### Sources

A source is an event listener or generator, either tied to a hardware
input queue or derivative of one (such as tickers)

### Sinks

A sink is a function which accepts messages from Sources, across a Wire

### Wires

A wire is something like a subscriber in the Pub/Sub model, and something like a priority queue. It is implemented using a dynamic array (slice) of Sinks, which can be attached and detached to a Source and delivers the message to a Sink.

## Implementation notes

Due to the 'limitations' of Go with regard to generics, instead of 
generics a lightweight, decode-on-demand interface based on the 
'simplebuffer' found in github.com/p9c/pod/pkg/simplebuffer

The initial alpha target is a fully functional, non-optimised 
implementation that does not tinker with unsafe functions but in 
Beta it is intended that chunks of bytes can be imputed into their 
type by creating a pointer to that type at the location of the first 
byte of the value. So 4 bytes would be interpreted there as a 32 
bit integer, for example, and its type imputed as such.

The rationale for this is that the complexity of generics anyway is excessive and has been left out of Go intentionally from its 1.0, as, underneath all the other reasons, they cannot be implemented with a single pass compiler like Go, and this strategy allows much faster compilation times, a key goal of Go.

It is the author's opinion that a methodology and set of small simple tools can replace the utility of generics without requiring a two pass compilation process, and will use `go:generate` in a novel way, to both build generated generic implementations and later, the zero copy unsafe re-typing code that will allow all of the features of generics to be used while combining the generate/bundle/compile process into one command:

```go generate ./...```

and depending on an environment variable places it in a path or as with `go build` in the pwd.

## Neurobiomimicry

I believe I am coining a word here. Up until recently, the dominant 
model used for computing is based on the illusory unity of processing
inside computers, illusory because it is simply such short delays 
between many processes that it appears to be instantaneous, but 
being that computing circuits have now hit the physical limit of 
frequency vs energy vs resistance, concurrency of processing, and
problems caused by failure of synchronisation have become very 
visible in computer network systems.

Biological signalling and computation systems have long ago solved 
these problems by various means and to various degrees. Slow 
converging consensus, more intelligence at the edge of the system, 
and all those Distributed Systems ideas are really, exactly, as 
I am calling it - Neurobiomimicry.

The definition extends beyond the internal pathways inside individual
animals, and includes the synchronisation and coordination systems 
that you find in social networks and ecosystems.

### A 'new' model for building Finite State Automata

It is not my original idea at all. In the 1985 paper by Rob Pike on 
the implementation of the language Newsqueak (no relation to 
Smalltalk, it is a language designed for concurrent functional 
programming specifically aimed at user input/output processing).

I only recently became fully competent in understanding and building algorithms based on the CSP model used in Go, and after reading about the implementation of Newsqueak a whole swathe of new ideas connected themselves together.

Finite State Automata are basically something along similar lines to the Rube Goldberg Machine - complex chains of devices and objects that trigger reactions that can split into parallel paths, merge, and of course in some pathway, finally terminate.

The conventional model for this kind of processing, as seen in most language translators and compilers is complicated, requires complex, domain specific grammars, and is thus the preserve of those with the luxury of time needed to acquire these skills and resources.

However, the process of compilation is a multi-stage, branched tree that requires different responses depending on the results of previous processing. The linear, procedural syntax of most languages used to build compilers is quite incompatible with this and creates excessive complexity, thus the special languages and requirement for, as the grammar of the language grows more complex, the use of conventionally artificial intelligence techniques to allow disambiguation, it requires expensive, multi-pass compilation processes.

## Enter Newsqueak

Newsqueak was basically one of the first implementations of the 
concept of Concurrent Sequential Process model of programming. It is 
defined by the use of atomic load/unload FIFO queues (channels) and
coroutines, which are a variant of subroutine that allows processing
to move quickly from one task to another without preemption and 
with a lower administrative overhead as found in preemptive 
multitasking schedulers.

It does depend on correctness in the code, but the payoff has proven
to be so much greater than you would naively expect. In certain 
kinds of computation systems, synchronisation becomes the biggest 
performance limiter. Notably, distributed network systems. But the 
same concurrency exists in every computing system, the difference 
between them is only latency, not in this fundamental feature. 
   
Something that jumped out at me when I was reading [the Newsqueak 
Implementation paper](https://swtch.com/~rsc/thread/newsquimpl.pdf) 
was that I clearly saw that an ingredient missing from CPU designs 
that facilitates fast implementations of an interpretive language 
system with concurrency is Compare and Swap operations. These 
operations take place in one instruction cycle and switch a value 
from one to another without an intermediate read/decide step in 
between. The implementations of race free channels and mutexes is nowadays often made with these operations, and it can essentially eliminate the high cost of context switching, ameliorating the performance hit while switching from one process to another, if the entire scheduling system is based on CSP.

## Clatter is not just for GUIs

Clatter can be used thus to implement a compiler, a GUI, a kernel process scheduler, and even low overhead memory allocation.

Go is not a perfect language for this, as such an implementation basically will require replacing infix operators with methods, and to implement Copy on Write passing of values, which underneath the syntax of the language, requires the use of CaS and transforms ordinary variable declaration and assigment semantics from the convenient infix operator `=` to use functions instead. It can be done in Go, however, especially for mathematics, results in a syntax that resembles Lisp rather than C. Infix operators are not native to CPU designs, as they are either verb subject object or subject object verb, similar to the distinction between German and English grammars - English uses infixes (known as prepositions, and, or, to from etc) where german has a fairly rigid format of subject object verb, which makes Lisp like German and C is like English.

Localisation of computer programming languages is a thorny problem and results in an unfair advantage of users of natural prepositional languages versus those that use lisp-like notations which also are found in Reverse Polish Notation, in the opposite sequence. 

All of these syntax elements are essentially concurrent. Until you have the three either explicitly or implicitly, it is not a statement or expression. So concurrent models have a power that allows you to combine sequential operations in ways that are intuitive to our brains, but for some people foreign because of their native language, which changes the way in which you serialize the concurrency of the three elements. 
