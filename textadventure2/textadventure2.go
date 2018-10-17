package main

// possible TODO
// NPC (non player characters) - talk to them, fight
// NPC move around the graph
// items that can be picked up or placed down
// accept natural language as input
// build other text-based games

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Println(choice.cmd, choice.description)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Sorry, I didn't understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func printArray(a []string) {
	for _, e := range a {
		fmt.Println(e)
	}
}

func main() {

	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
	You are in large chamber, deep underground
	You see three passages leading out. A north passage leads into darkness.
	To the south, a passage appears to head upward. The eastern passages appears
	flat and well traveled.`} // unset fields are nil

	darkRoom := storyNode{text: "Is is pitch black. You cannot see a thing."}
	darkRoomLit := storyNode{text: "The dark passage is now lit by your lantern. You can continue north or head back south."}
	grue := storyNode{text: "While stumbling in the darkness, you are eaten by a grue."}
	trap := storyNode{text: "You head down the well traveled path when suddenly a trap door opens and you fall into a pit."}
	treasure := storyNode{text: "You arrive at a small chamber, filled with treasure!"}

	start.addChoice("N", "Go north", &darkRoom)
	start.addChoice("S", "Go south", &darkRoom)
	start.addChoice("E", "Go east", &trap)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on your lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go north", &treasure)
	darkRoomLit.addChoice("S", "Go south", &start)

	start.play()

	fmt.Println()
	fmt.Println("The End.")
}
