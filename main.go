package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	age      int
	position string
}

func oldestEmployee(employees []Employee) string {
	oldest := employees[0]
	for _, emp := range employees {
		if emp.age > oldest.age {
			oldest = emp
		}
	}
	return oldest.name
}

type Circle struct {
	radius float64
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return c.Circumference()
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Vehicle interface {
	Start()
	Stop()
}

type Car struct {
	model string
}

func (c Car) Start() {
	fmt.Println(c.model, "đang khởi động.")
}

func (c Car) Stop() {
	fmt.Println(c.model, "đang dừng lại.")
}

type Bike struct {
	model string
}

func (b Bike) Start() {
	fmt.Println(b.model, "đang khởi động.")
}

func (b Bike) Stop() {
	fmt.Println(b.model, "đang dừng lại.")
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func increment(x *int) {
	*x++
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	}

	fib := fibonacci(n - 1)
	return append(fib, fib[len(fib)-1]+fib[len(fib)-2])
}

func safeDivision(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Lỗi:", r)
		}
	}()
	if b == 0 {
		panic("không thể chia cho 0")
	}
	fmt.Println("Kết quả:", a/b)
}

func greetAndPanic() {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	panic("Một lỗi đã xảy ra!")
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func max(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	maxValue := numbers[0]
	for _, num := range numbers {
		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue
}

func main() {
	employees := []Employee{
		{"John", 30, "Manager"},
		{"Alice", 25, "Engineer"},
		{"Bob", 40, "Director"},
	}
	fmt.Println("Nhân viên lớn tuổi nhất:", oldestEmployee(employees))

	circle := Circle{radius: 5}
	fmt.Printf("Chu vi: %.2f, Diện tích: %.2f\n", circle.Circumference(), circle.Area())

	shapes := []Shape{
		Circle{radius: 5},
		Rectangle{width: 4, height: 6},
	}
	for _, shape := range shapes {
		fmt.Printf("Diện tích: %.2f, Chu vi: %.2f\n", shape.Area(), shape.Perimeter())
	}

	vehicles := []Vehicle{
		Car{"Tesla Model 3"},
		Bike{"Yamaha"},
	}
	for _, vehicle := range vehicles {
		vehicle.Start()
		vehicle.Stop()
	}

	a, b := 3, 5
	fmt.Println("Trước khi hoán đổi:", a, b)
	swap(&a, &b)
	fmt.Println("Sau khi hoán đổi:", a, b)

	x := 10
	increment(&x)
	fmt.Println("Giá trị x sau khi tăng:", x)

	fmt.Println("Giai thừa của 5:", factorial(5))
	fmt.Println("Dãy Fibonacci 6 số đầu:", fibonacci(6))

	safeDivision(10, 0) 
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Lỗi trong greetAndPanic:", r)
			}
		}()
		greetAndPanic()
	}()

	fmt.Println("Tổng:", sum(1, 2, 3, 4))
	fmt.Println("Giá trị lớn nhất:", max(3, 5, 7, 2))
}