package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	CalculateArea() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func NewCircle(r float64) (*Circle, error) {
	if r <= 0 {
		return nil, errors.New("некорректный радиус для круга")
	}

	fmt.Printf("Круг: радиус %.f\n", r)
	return &Circle{
		Radius: r,
	}, nil
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

func NewRectangle(w, h float64) (*Rectangle, error) {
	if w <= 0 || h <= 0 {
		return nil, errors.New("некорректные стороны для прямоугольника")
	}

	fmt.Printf("Прямоугольник: ширина %.f, высота %.f\n", w, h)
	return &Rectangle{
		Width:  w,
		Height: h,
	}, nil
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) CalculateArea() float64 {
	area := 0.5 * t.Base * t.Height

	return area
}

func NewTriangle(b, h float64) (*Triangle, error) {
	if b <= 0 || h <= 0 {
		return nil, errors.New("некорректные стороны для треугольника")
	}

	fmt.Printf("Треугольник: основание %.f, высота %.f\n", b, h)
	return &Triangle{
		Base:   b,
		Height: h,
	}, nil
}

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не является фигурой")
	}

	return shape.CalculateArea(), nil
}

func main() {
	circle, err := NewCircle(5)
	if err != nil {
		fmt.Println("Ошибка при создании круга:", err)
	}

	areaCircle, err := calculateArea(circle)
	if err == nil {
		fmt.Printf("Площадь: %.14f\n", areaCircle)
	} else {
		fmt.Println("Ошибка:", err)
	}

	rectangle, err := NewRectangle(10, 5)
	if err != nil {
		fmt.Println("Ошибка при создании прямоугольника", err)
	}

	areaRectangle, err := calculateArea(rectangle)
	if err == nil {
		fmt.Printf("Площадь: %.f\n", areaRectangle)
	} else {
		fmt.Println("Ошибка:", err)
	}

	triangle, err := NewTriangle(8, 6)
	if err != nil {
		fmt.Println("Ошибка при создании треугольника", err)
	}

	areaTriangle, err := calculateArea(triangle)
	if err == nil {
		fmt.Printf("Площадь: %.f\n", areaTriangle)
	} else {
		fmt.Println("Ошибка:", err)
	}

	randomObject := "неизвестный объект"
	_, err = calculateArea(randomObject)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
