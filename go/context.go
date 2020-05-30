package main

import (
	"context"
	"fmt"
)

func main() {
	// this is probably a bad idea
	// "go-lint: should not use basic type string as key in context.WithValue"
	ctx := context.WithValue(context.Background(), "accountID", 1234)
	foo(ctx)

	// do something like this
	// https://stackoverflow.com/questions/40891345/fix-should-not-use-basic-type-string-as-key-in-context-withvalue-golint
	ctx = context.WithValue(context.Background(), contextKeyAccountID1, 1234)
	bar(ctx)

	// we could also customize our String method
	// https://medium.com/@matryer/context-keys-in-go-5312346a868d
	ctx = context.WithValue(context.Background(), contextKeyAccountID2, 1234)
	baz(ctx)
}

func foo(ctx context.Context) {
	accID := ctx.Value("accountID")
	fmt.Println(accID)
}

type contextKeyInt int

const (
	contextKeyAccountID1 = iota
	//...
)

func bar(ctx context.Context) {
	accID := ctx.Value(contextKeyAccountID1)
	fmt.Println(accID)
}

type contextKeyString string

const (
	contextKeyAccountID2 = contextKeyString("accountID")
	//...
)

func (c contextKeyString) String() string {
	return "context key " + string(c)
}

func baz(ctx context.Context) {
	accID := ctx.Value(contextKeyAccountID2)
	fmt.Println(accID)
}
