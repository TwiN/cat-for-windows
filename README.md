# cat-for-windows

A Go port of the popular 'cat' utility used to concatenate file(s) to standard output.
It can be compiled for most modern operating systems, but this is intended for use on Windows, since 
Linux and OSX already have the `cat` utility.


## Usage

```bash
cat <file>...
```

## Example

### Printing a single file to standard output

```bash
cat main.go
```


### Printing multiple files to standard output

```bash
cat main.go README.md
```


## Building the executable

You can easily build the executable for your architecture with the following command:

```
go build
```

Alternatively, there's a pre-built executable for Windows.
