package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Shapewitch
	Shapewitchperemetr
}
type Shapewitch interface{
	Area() float32
}
type Shapewitchperemetr interface {
	Peremetr()  float32
}

type Sguare struct {
	sideLenght float32
}

func (s Sguare) Area() float32 {
	return s.sideLenght * s.sideLenght
}

func (s Sguare) Peremetr() float32{
	return s.sideLenght * 4
}

type Cricle struct {
	radius float32
}
	  
func (c Cricle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Sguare{5}
	circle := Cricle{8}

	Printshapearea(square)
//	Printshapearea(circle)
//Printshapearea(square)
//Printshapearea(circle)
printinterface("dasd")
printinterface(23)
printinterface(square)
printinterface(circle)
}

func Printshapearea(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Peremetr())
}

//func printshapearea(shape Shape) {
	//fmt.Println(shape.Area())
//}

func printinterface(i interface{}) {
	//switch value := i.(type) {
	//case int:
	//	fmt.Println("int", value)

	//case bool:
	//	fmt.Println("bool", value)
	//default:

	//	fmt.Println("unkowe tupe", value)
	//}
  	str, ok := i.(string)
	if !ok {
		fmt.Println("hahaha lox")
		return
	}
	fmt.Println(len(str))
	//fmt.Println(i)
}


