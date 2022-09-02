package main

import "time"

type car struct {
	hgs_no          int
	owner_firstname string
	owner_lastname  string
	car_class       string
	balance         int
	time            string
}
type class struct {
	class1 int
	class2 int
	class3 int
}

type counter struct {
	cost                     class
	passedcarsinaday         []car
	numberofpassedcarsinaday int
}
type managementreport struct {
	dailyearning int
}

func (g *counter) Pay(car *car) {
	car.time = time.Now().Format("2006-01-02 15:04:05 GMT MST")
	if car.car_class == "otomobil" {
		car.balance -= g.cost.class1
	} else if car.car_class == "minibüs" {
		car.balance -= g.cost.class2
	} else if car.car_class == "otobüs" {
		car.balance -= g.cost.class3
	}
	g.passedcarsinaday = append(g.passedcarsinaday, *car)
	g.numberofpassedcarsinaday = len(g.passedcarsinaday)
}
func (yb *managementreport) Dailyearning(counter counter) {
	for _, i := range counter.passedcarsinaday {
		if i.car_class == "otomobil" {
			yb.dailyearning += counter.cost.class1

		} else if i.car_class == "minibüs" {
			yb.dailyearning += counter.cost.class2

		} else if i.car_class == "otobüs" {
			yb.dailyearning += counter.cost.class3

		} else {
			yb.dailyearning += 0
		}

	}
}
