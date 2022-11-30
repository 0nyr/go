# GO

### Useful links

[download go](https://golang.org/dl/)

[go std library](https://golang.org/pkg/)

### Useful tuto

[Get started with Go](https://golang.org/doc/tutorial/getting-started)

[A tour of Go](https://tour.golang.org/list) ⭐️ (nice)

## TODO

* [ ] **TODO (2020-11-22) :** Finish the [Go tour tutorial](https://golang.org/doc/tutorial/create-module).

## Commands

`go run [FILE].go`: run the specified program. Go build commands are designed to locate the modules required for packages you import.

`go help`: provides short essentials info about go.

`go mod init [NAME]`: inits a go module with provided `[NAME]`.

`go mod tidy`: update go imported modules.

### Notes

##### first tests with go

Following the official [Get started with Go](https://golang.org/doc/tutorial/getting-started) :

```shell
 ❮ onyr ★  kenzae❯ ❮ test1❯❯ go run hello.go 
Hello, World!
 ❮ onyr ★  kenzae❯ ❮ test1❯❯ go mod init hello
go: creating new go.mod: module hello
 ❮ onyr ★  kenzae❯ ❮ test1❯❯ go run hello.go 
go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
Hello, World!
Don't communicate by sharing memory, share memory by communicating.
```

### Exported Packages

> In Go, a name is exported IF IT BEGINS WITH A CAPITAL LETTER!!!
