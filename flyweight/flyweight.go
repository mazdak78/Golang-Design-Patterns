package flyweight

import "fmt"

type ShapeType string
const (
	CircleType ShapeType = "Circle"
	RectangleType ShapeType = "Rectangle"
)

type IShape interface{
	DrawShape()
	ShapeType() ShapeType
}

type Circle struct{}

func (c *Circle) DrawShape(){
	fmt.Printf("Drawing Circle\n")
}

func (c *Circle) ShapeType()  ShapeType{
	return CircleType
}


type Rectangle struct{}

func (r *Rectangle) DrawShape(){
	fmt.Printf("Drawing Rectangle\n")
}

func (r *Rectangle) ShapeType() ShapeType{
	return RectangleType
}

//Flyweight factory
type ShapeFactory struct {
	ObjectCount int
	shapeObjects map[ShapeType]IShape
}

func NewShapeFactory() *ShapeFactory{
	return &ShapeFactory{
		ObjectCount:  0,
		shapeObjects: make(map[ShapeType]IShape),
	}
}

func (s *ShapeFactory) GetShapeToDisplay(_shape ShapeType) IShape{
	if _, ok := s.shapeObjects[_shape]; ok{
		return s.shapeObjects[_shape]
	}else{
		switch _shape {
		case CircleType:
			s.shapeObjects[_shape] = &Circle{}
			s.ObjectCount++
		case RectangleType:
			s.shapeObjects[_shape] = &Rectangle{}
			s.ObjectCount++
		}
	}

	return s.shapeObjects[_shape]
}
