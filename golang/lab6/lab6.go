package lab6

import "fmt"

type Plane struct {
	Name  string
	Color string
	Speed int
}

func NewPlane(name, color string, speed int) *Plane {
	plane := new(Plane)
	plane.Name = name
	plane.Color = color
	plane.Speed = speed
	return plane
}

func (p Plane) GetName() string {
	return p.Name
}

func (p Plane) GetColor() string {
	return p.Color
}

func (p *Plane) SetSpeed(speed int) {
	p.Speed = speed
}

func Lab6() {
	boeing := NewPlane("Boeing", "Белый", 120)
	boeing.SetSpeed(150)
	fmt.Println(boeing.GetName())
	fmt.Println(boeing.GetColor())
}
