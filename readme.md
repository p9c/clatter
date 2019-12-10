# clatter
 
*noun* 

1. A loud percussive noise that is strident usually undesirable and 
   disruptive
2. One of the names of a group of Jackdaws (corvus monedula) referring
   to the clipped chirping noises they make in large groups from time
   to time.

## What is this thing?

Clatter is a code generator framework to allow the easier implementation of the programming model described in Rob Pike's paper [The Implementation of Newsqueak](doc/impl-new-TR.pdf) published in July 1990

The name refers to the idea of a programming model that resembles the rapid chatter of Jackdaws especially in the crepuscular hours. Channels and goroutines in Go are very light and efficient, so they can be used exactly the same way as function and return... lots of little fast messages flying through and being mutated in each node to generate the outputs they forward to the next stage in processing. If it had a sound, it would be the 'caw' of a jackdaw, highly concurrent and directional.

Further implementation to the full spec is planned but for now this application applies text transformations to encapsulate functions inside a wrapper called a 'node' and instead of having parameters and return values, these are implicitly transformed into a single slot buffered channel which has the property of allowing one message to be sent that can theoretically be also received in the same goroutine because the buffer allows it to not block. This will be part of the node wrapper syntax, and from the parameters it generates a list of channels names as global variables in a single package containing only the node's code.

From there, the package hierarchy is generated, the standard requirement of imports based on the lower case version of the node's name, and Node() (no parameters or return as all in and out goes through the package global channels)

The reason for this structure and the expedience of making it a generator instead of a full interpreter, for now, is this is only one step away from Go so there is a ready market for adoption. It also assists in the future development plans of Plan 9 Crypto, which is a realisation of something like the Plan 9 operating system, built on Newsqueak, with its implicit pure function, no pointers, no shared data and copies only when needed.

There is no doubt that functional programming methodology is a fad at this time with a JVM targeting Elixir, and every language now including closures and functions as variables. However, less emergent is concurrent and parallel programming. The biggest amount of use of these models is in back end servers that run the 'cloud' through systems like Docker and Kubernetes, which encapsulate sometimes very small services neatly and allow greater numbers of options for what to do with this.

Clatter is based on the same principle, to put everything into small neat boxes that are easy to understand, complete, and scalable from one to a billion concurrent threads given the memory. Channels incur some overhead but they compensate for this in their latency properties and reduction of the need for preemptive context switching.

## Roadmap

### v0.1.0
 Implement code generator that simply extends Go syntax with a new type of function that runs as a goroutine in parallel with all other nodes in the application, and default designated 'main' package of course defines command entry points and thus units of compilation, also ensuring that generated binaries have no unused code in them by default. 
 
 Thus every node is a package, and inside one can ostensibly create multiple functions, or closures, but the pattern of how a node executes helps discourage excessively long algorithms, as it is my opinion that 500 lines is about as long as a single file should get.

### v0.2.0 
Create a full showcase implementation of a simple but usable multi-platform GUI that transparently can run on any hardware natively for a target available to the Go compiler, now including WebAssembly, with an extension that creates http/json RPC endpoints attached to channels through an additional type qualifier symbol.
 
 For native, an extension to the syntax to cause separate binaries to be created and attached using a standard input/output pipe for highly parallel workloads (streaming, cryptocurrency mining, and so on), and with the same syntax as the remote for the WebAssembly target, except with package paths.
 
### v0.3.0
Implement an additional extension to the code generation that implements the pure by-value channel messaging between nodes, and a switchable debug and profile to create a tree of the mutations of each variable as it passes between nodes, and the timestamp of entry into a node in the latter.

Eliminate the need for explicit variable declaration by symbol categorisation to generate appropriate declarations, and creating a node context structure with all of its variables. Implicitly a one to one relationship between names inside a node elimiates shadowing bugs which equate to reinitialising a the same symbol. This will include everything inside closures as well, which can be given names by standard func prefix instead of a variable declaration, as is already possible with const.

Extend and generalize socket implementation to allow custom transports and filter chains to be created for various kinds of protocol. Nodes channels then attach to these with a prefix declaration naming the handler node that does the connection type, over a bidirectional channel that is sent and listens on both sides in two goroutines.

### ...
That's probably enough for now.
