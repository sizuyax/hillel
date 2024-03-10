package main

import (
	"flag"
	"fmt"
)

type Animal struct {
	runSpeed int
	name     string
	voice    string
}

type Turtle struct {
	animal       Animal
	carapaceSize int
	length       int
}

type Tiger struct {
	animal    Animal
	teethSize int
	paws      int
}

type Fish struct {
	animal Animal
	fin    int
	color  string
}

// Race AnimalCry - нужно для того чтобы выводить крик победителя и проигравшего
type Race struct {
	Distance        int
	Turtle          Turtle
	Tiger           Tiger
	Fish            Fish
	TurtleRemaining int
	TigerRemaining  int
	FishRemaining   int
	AnimalCry       Animal
}

// WinnersAndLosers - структура для хранения победителей и проигравших а также количества итераций
type WinnersAndLosers struct {
	Counts []int
	Names  []string
}

// WinnerSay - выводит крик победителя
func (a Animal) WinnerSay() {
	fmt.Println("**Cool winner's cry - " + a.name + "**\n")
}

// LoserSay - выводит плачь проигравшего
func (a Animal) LoserSay() {
	fmt.Println("**Loser cry - " + a.name + "**\n")
}

// Start - начало гонки
func (r *Race) Start() WinnersAndLosers {

	winnersAndLosers := WinnersAndLosers{}

	count := 0

	fmt.Println("Race started!")
	fmt.Println("Distance: ", r.Distance)
	fmt.Println("--------------------------------------")
	fmt.Println("Turtle speed: ", r.Turtle.animal.runSpeed)
	fmt.Println("Tiger speed: ", r.Tiger.animal.runSpeed)
	fmt.Println("Fish speed: ", r.Fish.animal.runSpeed)
	fmt.Println("--------------------------------------")

	// инициализация оставшегося расстояния для каждого животного
	r.TurtleRemaining = r.Distance
	r.TigerRemaining = r.Distance
	r.FishRemaining = r.Distance

	for len(winnersAndLosers.Names) < 3 {
		count++

		if r.TurtleRemaining > 0 {
			r.TurtleRemaining -= r.Turtle.animal.runSpeed
			if r.TurtleRemaining <= 0 {
				winnersAndLosers.Names = append(winnersAndLosers.Names, r.Turtle.animal.name)
				winnersAndLosers.Counts = append(winnersAndLosers.Counts, count)
			}
		}

		if r.TigerRemaining > 0 {
			r.TigerRemaining -= r.Tiger.animal.runSpeed
			if r.TigerRemaining <= 0 {
				winnersAndLosers.Names = append(winnersAndLosers.Names, r.Tiger.animal.name)
				winnersAndLosers.Counts = append(winnersAndLosers.Counts, count)
			}
		}

		if r.FishRemaining > 0 {
			r.FishRemaining -= r.Fish.animal.runSpeed
			if r.FishRemaining <= 0 {
				winnersAndLosers.Names = append(winnersAndLosers.Names, r.Fish.animal.name)
				winnersAndLosers.Counts = append(winnersAndLosers.Counts, count)
			}
		}
	}
	return winnersAndLosers
}

// CreateTeam - создание команды
func (r Race) CreateTeam(turtle Turtle, tiger Tiger, fish Fish, distance int) Race {
	race := Race{
		Distance: distance,
		Turtle:   turtle,
		Tiger:    tiger,
		Fish:     fish,
	}
	return race
}

// InitAnimal - инициализация животных
func InitAnimal(runSpeedTurtle, runSpeedTiger, runSpeedFish int) (Turtle, Tiger, Fish) {
	turtle := Turtle{
		animal: Animal{
			runSpeed: runSpeedTurtle,
			name:     "Turtle",
			voice:    "I'm a turtle",
		},
		carapaceSize: 10,
		length:       20,
	}
	tiger := Tiger{
		animal: Animal{
			runSpeed: runSpeedTiger,
			name:     "Tiger",
			voice:    "I'm a tiger",
		},
		teethSize: 8,
		paws:      4,
	}
	fish := Fish{
		animal: Animal{
			runSpeed: runSpeedFish,
			name:     "Fish",
			voice:    "I'm a fish",
		},
		fin:   2,
		color: "red",
	}
	return turtle, tiger, fish
}

// PrintResultOfRace - выводит результаты гонки
func PrintResultOfRace(winnersAndLosers WinnersAndLosers, race Race) {
	// пробегаемся по всем именам всех животных
	for _, name := range winnersAndLosers.Names {

		if name == winnersAndLosers.Names[0] {
			fmt.Printf("%s finished the race in %d iteration(s).\n", name, winnersAndLosers.Counts[0])
		} else if name == winnersAndLosers.Names[1] {
			fmt.Printf("%s finished the race in %d iteration(s).\n", name, winnersAndLosers.Counts[1])
		} else {
			fmt.Printf("%s finished the race in %d iteration(s).\n", name, winnersAndLosers.Counts[2])
		}

		// присваиваем имя животного для вывода крика
		race.AnimalCry.name = name

		if name == winnersAndLosers.Names[0] {
			race.AnimalCry.WinnerSay()
		} else if name == winnersAndLosers.Names[1] {
			// тут я не добавляю и пропускаю вывод, потому что если будет else и без else if в условии то оно выведет LoserSay() и для второго места
			fmt.Print("\n")
		} else {
			race.AnimalCry.LoserSay()
		}

	}
}

func main() {

	// инициализация флагов
	var runSpeedTurtle int
	var runSpeedTiger int
	var runSpeedFish int
	var distance int

	// инициализация флагов
	flag.IntVar(&runSpeedTurtle, "runSpeedTurtle", 10, "Run speed of the turtle")
	flag.IntVar(&runSpeedTiger, "runSpeedTiger", 100, "Run speed of the tiger")
	flag.IntVar(&runSpeedFish, "runSpeedFish", 555, "Run speed of the fish")
	flag.IntVar(&distance, "distance", 1000, "Distance")

	// парсим флаги
	flag.Parse()

	// инициализация животных
	tiger, fish, turtle := InitAnimal(runSpeedTurtle, runSpeedTiger, runSpeedFish)

	race := Race{}
	// создание команды, передаем всех животных и дистанцию из флага
	race = race.CreateTeam(tiger, fish, turtle, distance)

	// начало гонки
	winnersAndLosers := race.Start()

	// вывод результатов гонки
	PrintResultOfRace(winnersAndLosers, race)

}
