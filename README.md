# go-flag-demo
Go flag package demo

This is a demo application which explores go's `flag` package.

The application implements the scenario where 1st argument is a "function" and the subsequent arguments are arguments of that function:
```
myapp function [--arg1 value][--arg2 value]...
```

Argument names can be prepended with single dash:
```
myapp function [-arg1 value][-arg2 value]...
```

An assignment sign (`=`) can appear between the argument name and value:
```
myapp function [--arg1 = value][--arg2 = value]...
```



Usage examples:

```
$ go run ./main.go doSomething1
$ go run ./main.go doSomething1 -arg11=11
$ go run ./main.go doSomething1 -arg11=11 -arg12=12

$ go run ./main.go doSomething2
$ go run ./main.go doSomething2 -arg21=21
$ go run ./main.go doSomething2 -arg21=21 -arg22=22

$ go run ./main.go doSomething2 --arg21=21 --arg22=22

$ go run ./main.go doSomething2 --arg21 21 --arg22 22
```
