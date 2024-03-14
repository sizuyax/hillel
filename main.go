package main

import (
	"github.com/sirupsen/logrus"
	"hillel/battle"
	"hillel/models"
)

func main() {
	hero := models.NewHero()
	snake := models.NewSnake()

	defer func() {
		if r := recover(); r != nil {
			logrus.Error(r)
		}
	}()

	result := battle.Fight(&hero, &snake)

	logrus.Println(result)
}
