# Go Getting Started

[Go: Getting Started | Pluralsight](https://app.pluralsight.com/course-player?clipId=86312c75-3968-462a-8875-90fd6d074c34)


Content of this directory includes:
- Notes taken during going through the course.
- Code exnhanced with comments that I have found relevant for the best results of keeping the knowledge in. 

# Notes from video content 
Each H2 denotes a Section and each H3 denotes a lesson.

## Starting a project

### Installing the Go tools 
Installing Go can be done from golang.org side witch houses packages for all platforms. Go through the installation process.

To check if go command was successfully installed hit `go version`

### Overview of the Go Command
All of critical tools are prepackaged in the go command (getting packages, running tests, installing dependencies).

`go doc` will provide extensive docs for different parts of the language. It basically menas that you can get documentaion straight in your terminal. 

### Setting up an Enditor
Note taking omitted.

### Creating a Project 
We need to initialize the go module
`go mod init <name>/<of-the-module>` where name of the module is usually the path to where module is housed.

Running can be done using `go run .` it will pick the closes main.go that has `package main` defined. Another option is to run it using the name of the module used during initializaztion (it can be found in go.mod `module` section) like so: `go run <name-of-the-module>`

## Working with Primitive Data

### Declaring variables with Primitive Data types
The lesson is fully done in code.

### Working with Pointers
We have a Pointed Date type in go language. The variable keeps the pooiinter to the memory where the value is stored. 

### Creating Constraints 

### usint Iota and Constant Expression