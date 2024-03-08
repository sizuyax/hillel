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

type Race struct {
	Distance        int
	Turtle          Turtle
	Tiger           Tiger
	Fish            Fish
	TurtleRemaining int
	TigerRemaining  int
	FishRemaining   int
	Animal          Animal
}

func (a *Animal) WinnerSay() {
	fmt.Println("**Cool winner's cry - " + a.name + "**\n")
}

func (a *Animal) LoserSay() {
	fmt.Println("**Loser cry - " + a.name + "**\n")
}

func (r *Race) Start() (map[string]int, []string) {

	names := []string{}
	count := 0
	mapa := make(map[string]int)

	fmt.Println("Race started!")
	fmt.Println("Distance: ", r.Distance)
	fmt.Println("--------------------------------------")
	fmt.Println("Turtle speed: ", r.Turtle.animal.runSpeed)
	fmt.Println("Tiger speed: ", r.Tiger.animal.runSpeed)
	fmt.Println("Fish speed: ", r.Fish.animal.runSpeed)
	fmt.Println("--------------------------------------")

	r.TurtleRemaining = r.Distance
	r.TigerRemaining = r.Distance
	r.FishRemaining = r.Distance

	for len(names) < 3 {
		count++

		if r.TurtleRemaining > 0 {
			r.TurtleRemaining -= r.Turtle.animal.runSpeed
			if r.TurtleRemaining <= 0 {
				names = append(names, r.Turtle.animal.name)
				mapa[r.Turtle.animal.name] = count
			}
		}

		if r.TigerRemaining > 0 {
			r.TigerRemaining -= r.Tiger.animal.runSpeed
			if r.TigerRemaining <= 0 {
				names = append(names, r.Tiger.animal.name)
				mapa[r.Tiger.animal.name] = count
			}
		}

		if r.FishRemaining > 0 {
			r.FishRemaining -= r.Fish.animal.runSpeed
			if r.FishRemaining <= 0 {
				names = append(names, r.Fish.animal.name)
				mapa[r.Fish.animal.name] = count
			}
		}
	}
	return mapa, names
}

func (r *Race) CreateTeam(carapaceSize, length, teethSize, paws, fin, runSpeedTurtle, runSpeedTiger, runSpeedFish, distance int, color string) *Race {
	team := &Race{
		Distance: distance,
		Turtle: Turtle{
			animal: Animal{
				runSpeed: runSpeedTurtle,
				name:     "Turtle",
				voice:    "I'm a turtle! And i'm the best!",
			},
			carapaceSize: carapaceSize,
			length:       length,
		},
		Tiger: Tiger{
			animal: Animal{
				runSpeed: runSpeedTiger,
				name:     "Tiger",
				voice:    "I'm a tiger! And i'm the best!",
			},
			teethSize: teethSize,
			paws:      paws,
		},
		Fish: Fish{
			animal: Animal{
				runSpeed: runSpeedFish,
				name:     "Fish",
				voice:    "I'm a fish! And i'm the best!",
			},
			fin:   fin,
			color: color,
		},
	}
	return team
}

func main() {

	var runSpeedTurtle int
	var runSpeedTiger int
	var runSpeedFish int
	var distance int

	flag.IntVar(&runSpeedTurtle, "runSpeedTurtle", 10, "Run speed of the turtle")
	flag.IntVar(&runSpeedTiger, "runSpeedTiger", 100, "Run speed of the tiger")
	flag.IntVar(&runSpeedFish, "runSpeedFish", 50, "Run speed of the fish")
	flag.IntVar(&distance, "Distance", 1000, "Distance")

	flag.Parse()

	team := &Race{}
	team = team.CreateTeam(10, 20, 8, 4, 2, runSpeedTurtle, runSpeedTiger, runSpeedFish, distance, "red")
	mapa, names := team.Start()

	for _, name := range names {
		time, ok := mapa[name]

		if ok {
			fmt.Printf("%s finished the race in %d iteration(s).\n", name, time)
			if name == names[1] {
				fmt.Print("\n")
			}
		}

		team.Animal = Animal{name: name}
		if name == names[0] {
			team.Animal.WinnerSay()
		} else if name == names[1] {
			//
		} else {
			team.Animal.LoserSay()
		}

	}
}
