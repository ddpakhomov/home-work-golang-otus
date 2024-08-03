package main

import (
	"testing"
)

func TestCalculateAreaCircle(t *testing.T) {
	var expectedArea, radius float64 = 78.53981633974483, 5

	circle := Circle{Radius: radius}
	area := circle.CalculateArea()

	if area != expectedArea {
		t.Errorf("Неверная площадь. Ожидаемая площадь: %f, полученная площадь: %f", expectedArea, area)
	}
}

func TestNewCircleWithInvalidRadius(t *testing.T) {
	var invalidRadius float64 = -5

	_, err := NewCircle(invalidRadius)
	if err == nil {
		t.Error("Ожидалась ошибка при создании круга с некорректным радиусом")
	}
}

func TestCalculateAreaRectangle(t *testing.T) {
	var expectedArea, width, height float64 = 50, 10, 5

	rectangle := Rectangle{
		Width:  width,
		Height: height,
	}
	area := rectangle.CalculateArea()

	if area != expectedArea {
		t.Errorf("Неверная площадь. Ожидаемая площадь: %f, полученная площадь: %f", expectedArea, area)
	}
}

func TestNewRectangleWithInvalidWidth(t *testing.T) {
	var invalidWidth, height float64 = -10, 5

	_, err := NewRectangle(invalidWidth, height)
	if err == nil {
		t.Error("Ожидалась ошибка при создании прямоугольника с некорректной шириной")
	}
}

func TestNewRectangleWithInvalidHeight(t *testing.T) {
	var width, invalidHeight float64 = 10, -5

	_, err := NewRectangle(width, invalidHeight)
	if err == nil {
		t.Error("Ожидалась ошибка при создании прямоугольника с некорректной высотой")
	}
}

func TestCalculateAreaTriangle(t *testing.T) {
	var expectedArea, base, height float64 = 24, 8, 6

	triangle := Triangle{
		Base:   base,
		Height: height,
	}
	area := triangle.CalculateArea()

	if area != expectedArea {
		t.Errorf("Неверная площадь. Ожидаемая площадь: %f, полученная площадь: %f", expectedArea, area)
	}
}

func TestNewTriangleWithInvalidBase(t *testing.T) {
	var invalidBase, height float64 = -8, 6

	_, err := NewTriangle(invalidBase, height)
	if err == nil {
		t.Error("Ожидалась ошибка при создании треугольника с некорректным основанием")
	}
}

func TestNewTriangleWithInvalidHeight(t *testing.T) {
	var base, invalidHeight float64 = 8, -6

	_, err := NewTriangle(base, invalidHeight)
	if err == nil {
		t.Error("Ожидалась ошибка при создании прямоугольника с некорректной высотой")
	}
}
