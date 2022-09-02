package main

import "fmt"

func main() {
	arc := car{
		owner_firstname: "er",
		owner_lastname:  "as",
		hgs_no:          1234567890,
		car_class:       "otomobil",
		balance:         20,
	}
	arc2 := car{
		owner_firstname: "er",
		owner_lastname:  "as",
		hgs_no:          1234567890,
		car_class:       "otobüs",
		balance:         30,
	}

	counter := counter{
		cost: class{class1: 14, class2: 17, class3: 20},
	}
	report := managementreport{
		dailyearning: 0,
	}
	counter.Pay(&arc)
	counter.Pay(&arc2)
	report.Dailyearning(counter)
	fmt.Printf("Gün içinde geçen araçlar: %v \n", counter.passedcarsinaday)
	fmt.Printf("Gün içinde geçen araç sayısı: %v \n", counter.numberofpassedcarsinaday)
	fmt.Printf("günlük kazanç:%d", report.dailyearning)

}
