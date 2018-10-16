/*
	TODO:
	1. Add a function that will insert a new page, after a given page
	2. Add a function that will delete a page
*/

package main

import (
	"fmt"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func playStory(page *storyPage) {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	page1 := storyPage{"It was a dark and stormy night.", nil}
	page2 := storyPage{"You are alone, and you need to find the sacred helmet before the bad guys do", nil}
	page3 := storyPage{"You see a troll ahead", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page1)
}
