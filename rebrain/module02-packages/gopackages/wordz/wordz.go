package wordz

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func init() {
	fmt.Println("Func Init -->>")
}

var Hello = "This is wordz package"
var Prefix = "Random word: "
var Wordz = []string{"one", "two", "three", "four", "five"}

func Random() string {
	max := len(Wordz)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}

func get(index int64) string {
	return Prefix + Wordz[index]
}
