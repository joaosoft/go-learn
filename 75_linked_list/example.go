package main

import (
	"errors"
	"fmt"
)

type Post struct {
	body        string
	publishDate int64 // Unix timestamp
	next        *Post // link to the next Post
}

type Feed struct {
	length int
	start  *Post
	end    *Post
}

func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		lastPost := f.end
		lastPost.next = newPost
		f.end = newPost
	}
	f.length++
}

func (f *Feed) Remove(publishDate int64) {
	if f.length == 0 {
		panic(errors.New("feed is empty"))
	}

	var previousPost *Post
	currentPost := f.start

	for currentPost.publishDate != publishDate {
		if currentPost.next == nil {
			panic(errors.New("no such Post found."))
		}

		previousPost = currentPost
		currentPost = currentPost.next
	}
	previousPost.next = currentPost.next

	f.length--
}

func (f *Feed) Insert(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
	} else {
		var previousPost *Post
		currentPost := f.start

		for currentPost.publishDate < newPost.publishDate {
			previousPost = currentPost
			currentPost = previousPost.next
		}

		previousPost.next = newPost
		newPost.next = currentPost
	}
	f.length++
}

func main() {
	f := &Feed{}
	p1 := Post{
		body: "lorem ipsum",
	}
	f.Append(&p1)

	fmt.Printf("length: %v\n", f.length)
	fmt.Printf("first: %v\n", f.start)

	p2 := Post{
		body: "dolor sit amet",
	}
	f.Append(&p2)

	fmt.Printf("length: %v\n", f.length)
	fmt.Printf("first: %v\n", f.start)
	fmt.Printf("second: %v\n", f.start.next)
}
