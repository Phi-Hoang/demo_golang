package main

import "fmt"

type Ball struct {
	Radius int
}

func (b Ball) Bounce() {
	fmt.Println(b, b.Radius)
}

type BasketBall struct {
	Ball
	Weight int
	Radius int // added
}

type BasketBall2 struct {
	*Ball
	Weight int
	Radius int
}

func main() {
	b := Ball{23}
	bb := BasketBall{Weight: 7, Ball: b, Radius: 30}
	b = Ball{7}
	bb.Bounce()
	fmt.Println(bb.Radius)
	fmt.Println(bb.Ball.Radius)

	bb2 := BasketBall2{Weight: 7, Ball: &b, Radius: 30}
	b = Ball{12}
	bb2.Bounce()
	fmt.Println(bb2.Radius)
	fmt.Println(bb2.Ball.Radius)

	// embeded interface
	vb := VolleyBall(&Ball{5}, 50)

	//worked
	vb.Bounce()

	//do not work
	r := vb.Radius
}

type Action interface {
	Bounce()
}

type VolleyBall struct {
	Action
	Weight int
}
