package main

import (
	"fmt"
	"strconv"
	"strings"
	Tasks2 "structs/Tasks"
	"structs/cache"
	"structs/io"
	"structs/mathematics/Geometry"
)

func main() {

	ioHandler := io.IO{}

	contents, err := ioHandler.ReadFile("figure.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Loaded figures: %s", contents)

	//An empty interface could represent an object, that could be cast then to any implementation
	figures := []Geometry.Figure{
		Geometry.Circle{Radius: 4, Diameter: 8}, Geometry.Rectangle{Width: 8, Length: 8},
		Geometry.Rectangle{Width: 10, Length: 10},
	}

	sb := strings.Builder{}

	numbers, err := io.LoadNumbers()

	if err == nil {
		result := Tasks2.ParallelSum(numbers)
		fmt.Printf("the sum of all 100000 numbers is %d\n", result)
	}

	Tasks2.ExecuteWaitGroup(1, 2, 3)

	//Save a figure under the cache
	figureCache := cache.Create(20)

	counter := 0
	//iterate through all figures
	for _, figure := range figures {

		counter++
		//The following validation should be implemented to cast btw an empty interface to a specific object
		if rectangle, ok := figure.(Geometry.Rectangle); ok {
			fmt.Printf("\nThis figure %+v is a %v, the area is: %v and the perimeter is %v\n", rectangle, rectangle.Name(), rectangle.Area(), rectangle.Perimeter())

			key := rectangle.Name() + strconv.Itoa(counter)

			if err := figureCache.Set(key, rectangle); err != nil {
				figureCache.HandleError(err, key)
			}

			sb.WriteString(fmt.Sprintf("\nFigure: %s;a:%f;p:%f", rectangle.Name(), rectangle.Area(), rectangle.Perimeter()))
		}

		//Save the figure in a file

		//The following validation should be implemented to cast btw an empty interface to a specific object
		if circle, ok := figure.(Geometry.Circle); ok {

			key := circle.Name() + strconv.Itoa(counter)

			if err := figureCache.Set(key, circle); err != nil {
				figureCache.HandleError(err, key)
			}

			sb.WriteString(fmt.Sprintf("\nFigure: %s;a:%f;p:%f", circle.Name(), circle.Area(), circle.Perimeter()))
		}
	}

	if err := ioHandler.CreateFile("figure.txt", sb.String()); err != nil {
		fmt.Println(err)
	}

}
