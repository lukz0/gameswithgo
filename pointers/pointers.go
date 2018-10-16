package main

import (
	"fmt"
	"reflect"
)

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func addOne(x *int) {
	*x = *x + 1
}

func whereIsBadGuy(b *badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {
	x := 5
	fmt.Println(x)

	var xPtr *int = &x
	fmt.Println(xPtr)
	fmt.Println(reflect.TypeOf(xPtr))

	addOne(&x)
	fmt.Println(x)

	p := position{4, 2}
	badguy := badGuy{"Jabba the Hut", 100, p}
	whereIsBadGuy(&badguy)
}
