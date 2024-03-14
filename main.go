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

	result, err := battle.Fight(&hero, &snake)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Println(result)
}
