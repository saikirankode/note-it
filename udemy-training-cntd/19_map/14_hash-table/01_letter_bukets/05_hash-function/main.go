package main

import "fmt"

func main() {
	n := hashbucket("Go", 12)
	fmt.Println(n)
}

func hashbucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
