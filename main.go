package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

var hashIn string
var hashOut string
var search string
var query string

func main() { //Simple IO here
	// header()
	fmt.Println("Input the hash")
	fmt.Scanf("%s", &search)
	hashIn, hashOut, query = solver(search)
	fmt.Printf("%s  |  %s  |  %s", hashIn, hashOut, query)

}

func header() {
	fmt.Println("Hello")
}

func solver(search string) (string, string, string) {
	query = search
	hashIn = md5Solver(query)

	var i int
	for i = 0; i > -1; i++ {
		query = strconv.Itoa(i)
		hashOut = md5Solver(query)
		if hashOut == hashIn {
			println("Found The value")
			break
		} else {
			println(query)
		}
	}
	return hashOut, hashIn, query
}

func md5Solver(query string) string {

	md5 := md5.New()
	md5.Write([]byte(string(query)))
	hash := hex.EncodeToString(md5.Sum(nil))
	return hash
}
