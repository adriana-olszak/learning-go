# Go The Big Picture

[Go: The Big Picture | Pluralsight](https://app.pluralsight.com/course-player?clipId=86312c75-3968-462a-8875-90fd6d074c34)


Content of this directory includes:
- Notes taken during going through the course.
- Code exnhanced with comments that I have found relevant for the best results of keeping the knowledge in. 

# Notes from video content 
Each H2 denotes a Section and each H3 denotes a lesson.

## **What is go?**

Created by engineers working at Google. It was designed from 2007 to 2012 when the first version was released.

The course that I’m watching the current version is 1.12.

My installed version is 1.19.5.

Fun fact: Docker is written in GoLang.

### **Language Characteristics**

- **Strong, static type system** the decision was made to put much emphasis on the compiler to help you with types. The syntax for Types is not so verbose thanks to that.
- **C-inspired syntax -** with reduced clutter
- **Multi paradigm -** Procedural or Object-oriented
- **Garbage collected** - you don’t need to manage memory allocation. Highly optimized to make it fast.
- **Fully compiled -** down to the executable binary.
- **Rapid compilation -** it allows for staying in the loop when using TDD thanks to almost no time of compilation.
- **Single binary output -** simple deployments
---
## **How Can Go Help You?**

### **Overview of Philosophy and Values**

- **Simplicity,** it was designed with that in mind. It should be easy to learn and work in.
- **Network aware and concurrent apps** - Designed to be run on multiple threads that are currently widespread
- **Out-of-the-box experience -** all of the tools are built into and easy to use.
- **Cross-platform** - you can use it in all operation systems
- **Backward compatibility is** honored as much as possible

### **Embracing the Power of Simplicity**

Looping in Go = using for loop with a combination of parameters or lack of it.

### **Focusing on Network-Aware**

In the core, we have the following packages

- net and net/http packages.
- goroutines - For concurrency, we have goroutines which we can easily spawn
- channels - it allows us to safely communicate between concurrent types in an event-based, thread-share way so no memory-sharing problems.

### **Out-of-the-box Experience**

Standard Library is focused on having core concepts that are required for significant applications. It includes:

- testing,
- network APIs,
- data compression, file, and string manipulation.

Go CLI gives you everything without the need to use 3rd party package. build process, code generation, project initialization, dependency retrieval, running tests, profiling, documentation, and report language bugs.

### **Simplifying Cross-platform Development**

Go has built into the standard library build GOOS architecture to build executables for Windows, OS X, Android, Linux whatever technology.

### **Backward Compatibility**

Commitment to always working applications when using the same Major version of the language. The only exceptions are:

- Security
- Unspecified behavior
- Specification errors
- Bugs

### **Go’s primary Use Cases**

__Used for__

- Web services - moving data through the network
- Web applications - rendering HTML
- DevOps - low-level system tools
- GIU / Thick-client - cross-platform builds
- Machine learning - because of the high performance of Go

__Shine in__

Web services, web applications, DevOps

---
## **What is it like to use Go?**
This section is covered by code. You can find it in this directory. 