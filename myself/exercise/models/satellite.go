package models

import "strings"

type Satellite struct {
	Id      string
	X       float32
	Y       float32
	Message []string
	sb      strings.Builder
}

func (s Satellite) Initialize(x, y float32, id string) {
	s.X = x
	s.Y = y
	s.Id = id
	s.sb = strings.Builder{}
}

func (s Satellite) GetLocation(distances ...float32) (x, y float32) {
	if s.Y == 0 && s.X == 0 {
		panic("Satellite initial coordinates are not known")
	}

	if len(distances) < 2 {
		panic("Provide and x, and y distance relative to the origin")
	}

	var xOrigin, YOrigin float32
	xOrigin = s.X - distances[0]
	YOrigin = s.Y - distances[1]

	return xOrigin, YOrigin
}

func (s Satellite) SetMessage(messages chan string) Satellite {
	for word := range messages {

		length := len(word)
		for i := 0; i < length; i++ {

			if len(word) <= 0 {
				s.sb.WriteString("")
				continue
			}

			s.sb.WriteString(word)
		}
	}

	return s
}

func (s Satellite) GetMessages() string {
	return s.sb.String()
}
