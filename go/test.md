## check test coverage per method

ref: https://blog.golang.org/cover

```.sh
$ go test -coverprofile=coverage.out
$ go tool cover -func=coverage.out
```

## run a particular test case

https://golang.org/pkg/cmd/go/internal/test/

>   -run regexp
>       Run only those tests and examples matching the regular expression.

```.sh
go test ./package -run TestName
```
