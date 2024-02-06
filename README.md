# go-example

## What is it?
这是一个 golang 的练习库。用于熟悉一些 golang 的机制。
基于 [cobra](https://github.com/spf13/cobra)，和 [Go by Example](https://gobyexample.com/)

## How to use?
```bash
# list all example
go run main.go list

# run a certain example
go run main.go run <GoExample>
```
TODO: add doc for certain go example

## How to build?
```bash
./scripts/build.sh
```
you may find executable file in (project path)/output/

## How to develop?
1. create a path for your example
```bash
mkdir pkg/goexpl/foo/foo.go
```
2. Implement your owner GoExample  
    a. Implement interface GoExample  
    b. define a name in pkg/common/constats  
    c. add your go example in cmd.AllGoExample  
    you may refer to former commit.
3. run for test
```bash
go run main.go run foo
```
