package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"hillel/battle"
	"hillel/errors"
	"hillel/models"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	hero := models.NewHero()
	snake := models.NewSnake()

	defer func() {
		if r := recover(); r != nil {
			logrus.Error("Помилка: ", r)
			if err, ok := r.(error); ok {
				errors.LogError(err)
			}
		}
	}()

	result := battle.Fight(&hero, &snake)
	fmt.Println(result)
}
