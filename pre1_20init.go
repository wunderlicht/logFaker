// originally written for go1.17. Semantics of rand.Seed changed in go 1.20
// which makes initialization of the random generator using rand.Seed obsolete.
// Will not be included in builds with version >= go1.20
//go:build !go1.20

package main

import (
	"math/rand"
	"time"
)

// only executed in go versions < 1.20
func init() {
	rand.Seed(time.Now().UnixNano())
}
