package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero struct {
	Name string
}

func NewHero() Hero {
	rand.Seed(time.Now().UnixNano())

	heroAdj := []string{"Прекрасный ", "Чудесный ", "Смешной ", "Тупой "}
	heroAdj2 := []string{"Опухлый ", "Выпивший ", "(побитый Бомжом) ", "Нанюханый "}
	heroSub := []string{"Владислав", "Марик", "Илья", "Мушка"}

	randAdj := heroAdj[rand.Intn(len(heroAdj))]
	randAdj2 := heroAdj2[rand.Intn(len(heroAdj2))]
	randSub := heroSub[rand.Intn(len(heroSub))]

	name := fmt.Sprintf("%s %s %s", randAdj, randAdj2, randSub)

	return Hero{Name: name}
}

type Snake struct {
	Name  string
	Heads int
}

func NewSnake() Snake {
	rand.Seed(time.Now().UnixNano())

	heads := rand.Intn(101) + 50

	return Snake{Name: "Накуренный Змей Феодосий", Heads: heads}
}
