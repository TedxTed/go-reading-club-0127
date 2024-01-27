package main

import "fmt"

type Person struct {
	name string
}

type Car struct{}

type Motorcycle struct{}

type ElectricCar struct {
	Car
}

type HybridElectricCar struct {
	Car
}

type StandardMotorcycle struct {
	Motorcycle
}

type ManualMotorcycle struct {
	Motorcycle
}

type CarryPassenger interface {
	Carry(person Person)
}

func (c Car) Carry(person Person) {
	fmt.Println("Car is carrying a person")
}

func (c Car) StartEngine() {
	fmt.Println("Car engine started")
}

func (c Car) TurnSteeringWheel() {
	fmt.Println("Turning the car's steering wheel")
}

func (c Car) Refuel() {
	fmt.Println("Refueling the car")
}

func (c Car) Brake() {
	fmt.Println("Car is braking")
}

func (m Motorcycle) Carry(person Person) {
	fmt.Println("Motorcycle is carrying a person")
}

func (m Motorcycle) StartEngine() {
	fmt.Println("Motorcycle engine started")
}

func (m Motorcycle) TwistThrottle() {
	fmt.Println("Twisting the motorcycle's throttle")
}

func (m Motorcycle) Drive() {
	fmt.Println("Driving the motorcycle")
}

func (m Motorcycle) Lock() {
	fmt.Println("Motorcycle is locked")
}

func (m Motorcycle) Brake() {
	fmt.Println("Motorcycle is braking")
}

func (e ElectricCar) Carry(person Person) {
	fmt.Println("Electric car is carrying a person")
}

func (h HybridElectricCar) Carry(person Person) {
	fmt.Println("Hybrid electric car is carrying a person")
}

func (s StandardMotorcycle) Carry(person Person) {
	fmt.Println("Standard motorcycle is carrying a person")
}

func (m ManualMotorcycle) Carry(person Person) {
	fmt.Println("Manual motorcycle is carrying a person")
}

func main() {
	var person Person = Person{name: "ted"}
	var car Car
	var motorcycle Motorcycle

	car.Carry(person)
	car.StartEngine()
	car.TurnSteeringWheel()
	car.Refuel()
	car.Brake()

	motorcycle.Carry(person)
	motorcycle.StartEngine()
	motorcycle.TwistThrottle()
	motorcycle.Drive()
	motorcycle.Lock()
	motorcycle.Brake()
}
