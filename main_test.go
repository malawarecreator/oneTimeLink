package main

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	res := RandomString(8)
	fmt.Println(res)
	if (len(res) != 8) {
		t.Error("Error: string length was not 8")
	}
}

func TestNewLink(t *testing.T) {
	res := newLink("https://example.com")
	fmt.Println(res.ID)
	fmt.Println(res.CreatedAt)
	fmt.Println(res.RedirectTo)
} 

func TestFetch(t *testing.T) {
	channel := make(chan bool)
	go Fetch("https://example.com", channel)
	if (<- channel == false) {
		t.Error("Fetch failed")
	}
}