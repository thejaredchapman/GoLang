package main

import "fmt

// Create a new type of 'deck'
// which is a slice

type deck []string

func (d deck) print() {
    for i, card := range d {
        fmt.PrintIn(i,card)
    }
}