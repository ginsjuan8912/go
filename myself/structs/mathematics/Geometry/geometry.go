package Geometry

import (
	"errors"
	"math"
)

type IShapeLoader interface {
	LoadShapes(string) []Shape
}

//Declare an interface as a shape
type Shape interface {
	Name() string
	Edges() int
	Perimeter() float64
	Area() float64
}

//An interface can also be empty to represent an object, in this case represents a figure
type Figure interface{}

/*This type represents a rectangle*/
type Rectangle struct {
	Length float64
	Width  float64
}

/*This to load shapes from a string*/
type ShapeLoader struct {
	Shapes []Shape
}

func (sh ShapeLoader) LoadShapes(shapes string) error {
	return errors.New("not implemented")
}

func (r Rectangle) Name() string {
	return "Rectangle"
}

func (r Rectangle) Edges() int {
	return 4
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length * r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

/*This type represents a circle*/
type Circle struct {
	Radius   float64
	Diameter float64
}

func (c Circle) Perimeter() float64 {
	return (2 * math.Pi) * c.Radius
}

func (c Circle) Name() string {
	return "Circle"
}

func (c Circle) Area() float64 {
	return math.Pow(c.Perimeter(), 2)
}

func (c Circle) Edges() float64 {
	return 0
}
