package vectors

import "math"

//This file shows how to implement structures

//Declare a struct
type Vector struct {
	X float64
	Y float64
}

//Incrust a struct from a parent to a child
type Vector3 struct {
	Vector
	Z float64
}

//This will take the Vector type and add recievers to this function
/* Get the direction of the vector */
func (v Vector) GetDirection() float64 {
	return math.Atan(v.Y / v.X)
}

// CalculateDistance /*Calculate the distance btw two vectors*/
func (v Vector) CalculateDistance(end Vector) float64 {
	x := end.X - v.X
	y := end.Y - end.Y
	res := math.Pow(x, 2) + math.Pow(y, 2)
	return math.Sqrt(res)
}
