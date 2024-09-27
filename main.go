package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("Original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reverse again: %q\n", doubleRev)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yo, this is %s!", r.URL.Path[1:])
}