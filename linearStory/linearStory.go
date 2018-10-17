package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory(scanner *bufio.Scanner) {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
		scanner.Scan()
	}
}

func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

func deletePageAfter(page *storyPage) {
	page.nextPage = page.nextPage.nextPage
}

func insertPageAfter(pageBefore, insertedPage *storyPage) {
	insertedPage.nextPage = pageBefore.nextPage
	pageBefore.nextPage = insertedPage
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	page1 := storyPage{"It was a dark and stormy night.", nil}
	page1.addToEnd("You are alone, and you need to find the sacred helmet before the bad guys do")
	page1.addToEnd("You see a troll ahead")

	var page2Ptr *storyPage = page1.nextPage

	deletePageAfter(&page1)           // delete second page
	insertPageAfter(&page1, page2Ptr) // add second page back where it was

	// fmt.Println(reflect.TypeOf(scanner)) // *bufio.Scanner
	// fmt.Printf("%T\n", scanner)

	page1.addAfter("Testing addAfter()")
	page1.playStory(scanner)
}
