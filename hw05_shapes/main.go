package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	CalculateArea() (float64, error)
}

type Circle struct {
	Radius float64
}

func (c Circle) CalculateArea() (float64, error) {
	if c.Radius <= 0 {
		return 0, fmt.Errorf("переданный объект не является фигурой")
	}
	return math.Pi * c.Radius * c.Radius, nil
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) CalculateArea() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return r.Width * r.Height, nil
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) CalculateArea() (float64, error) {
	if t.Base <= 0 || t.Height <= 0 {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return 0.5 * t.Base * t.Height, nil
}

func calculateArea(s Shape) (float64, error) {
	return s.CalculateArea()
}

func main() {
	circle := Circle{5}
	areaCircle, err := calculateArea(circle)
	if err == nil {
		fmt.Printf("Круг: радиус %.f\n", circle.Radius)
		fmt.Printf("Площадь: %.14f\n", areaCircle)
	} else {
		fmt.Println("Ошибка:", err)
	}

	rectangle := Rectangle{10, 5}
	areaRectangle, err := calculateArea(rectangle)
	if err == nil {
		fmt.Printf("Прямоугольник: ширина %.f, высота %.f\n", rectangle.Width, rectangle.Height)
		fmt.Printf("Площадь: %.f\n", areaRectangle)
	} else {
		fmt.Println("Ошибка:", err)
	}

	triangle := Triangle{8, 6}
	areaTriangle, err := calculateArea(triangle)
	if err == nil {
		fmt.Printf("Треугольник: основание %.f, высота %.f\n", triangle.Base, triangle.Height)
		fmt.Printf("Площадь: %.f\n", areaTriangle)
	} else {
		fmt.Println("Ошибка:", err)
	}

	circle = Circle{0}
	areaCircle, err = calculateArea(circle)
	if err == nil {
		fmt.Printf("Круг: радиус %.f\n", circle.Radius)
		fmt.Printf("Площадь: %.14f\n", areaCircle)
	} else {
		fmt.Println("Ошибка:", err)
	}
}
