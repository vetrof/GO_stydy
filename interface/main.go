package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Office interface {
	PrintEmployersInfo()
	Hash() string
}

type Employers struct {
	Name string
	Id   int
}

func (e *Employers) PrintEmployersInfo() {
	fmt.Println(e.Id, e.Name)
}

func (e *Employers) Hash() string {
	str := e.Name + strconv.Itoa(e.Id)
	hash := sha256.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func main() {
	e := Employers{
		Name: "assa",
		Id:   99,
	}
	e.PrintEmployersInfo()
	hash := e.Hash()
	fmt.Println(e, hash)
}
