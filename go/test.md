## check test coverage per method

ref: https://blog.golang.org/cover

```.sh
$ go test -coverprofile=coverage.out
$ go tool cover -func=coverage.out
```
