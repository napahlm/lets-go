# lets-go
## Learning the Go idiomacy *and hopefully utilizing it*

And abhorrent puns. So let's Go.

### Priorities:

- Interfaces
- Channels
- Concurrency
- Working with files
- Packages and modules
- Testing

To make the code simple and readible it seems like spending time with interfaces and channels are worth spending time on early on.

### Descriptions of certain concepts

**Package:** Collection o fsource files in the same directory compiled together.

**Module:** Collection of related Go packages that are released together. The go.mod file declares the module path.

### Testing

Go has ready made testing capabilities:

- Test files are called the same as the file to be tested but appended with "_test.go". So the test for "example.go" becomes "example_test.go".
- Create the test :eyes:
- Perform the test by calling "go test" in the terminal in the same directory as the test.

## "Effective Go" - Some notes