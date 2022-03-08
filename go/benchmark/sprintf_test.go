package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkStrConcatSprintf(b *testing.B) {
	for i := 0; i < n; i++ {
		_ = fmt.Sprintf("a random string:, %v", randomString(10))
	}
}

func BenchmarkStrConcatPlus(b *testing.B) {
	for i := 0; i < n; i++ {
		_ = "a random string:" + randomString(10)
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
