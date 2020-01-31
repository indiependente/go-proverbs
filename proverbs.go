package goproverbs

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	// from https://go-proverbs.github.io/
	proverbs = []string{
		"Don't communicate by sharing memory, share memory by communicating",
		"Concurrency is not parallelism",
		"Channels orchestrate; mutexes serialize",
		"The bigger the interface, the weaker the abstraction",
		"Make the zero value useful",
		"interface{} says nothing",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite",
		"A little copying is better than a little dependency",
		"Syscall must always be guarded with build tags",
		"Cgo must always be guarded with build tags",
		"Cgo is not Go",
		"With the unsafe package there are no guarantees",
		"Clear is better than clever",
		"Reflection is never clear",
		"Errors are values",
		"Don't just check errors, handle them gracefully",
		"Design the architecture, name the components, document the details",
		"Documentation is for users",
		"Don't panic",
	}
)

// Proverb returns a random Go proverb.
func Proverb() string {
	return proverbs[rand.Intn(len(proverbs))]
}
