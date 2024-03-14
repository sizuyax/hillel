package battle

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"hillel/models"
	"math/rand"
	"time"
)

func HeroHit(snake *models.Snake) int {
	rand.Seed(time.Now().UnixNano())

	logrus.Println("Богатырь рубает головы змею и...\n\n")

	randStrongAttack := rand.Intn(2)

	hit := 1
	switch randStrongAttack {
	case 0:
		hit = rand.Intn(6) + 1
	case 1:
		hit = rand.Intn(21) + 1
	}

	if snake.Heads-hit < 0 {
		hit = snake.Heads
	}

	snake.Heads -= hit

	logrus.Printf("Богатырь срубил %d голов, осталось %d\n", hit, snake.Heads)
	return hit
}

func GrowHeads(snake *models.Snake, hit int) {
	rand.Seed(time.Now().UnixNano())

	growth := 0
	for i := 0; i < hit; i++ {
		probability := rand.Intn(100)
		switch {
		case probability < 40:
			// ничего не вырастает
		case probability < 70:
			growth += 1
		case probability < 90:
			growth += 2
		default:
			growth += 3
		}
	}

	snake.Heads += growth

	logrus.Printf("Змей вырос на %d голов, теперь у него %d голов\n\n", growth, snake.Heads)
}

func Fight(hero *models.Hero, snake *models.Snake) (string, error) {
	battle := models.NewBattle()
	for {
		battle.Round++

		if battle.Round > 600 {
			time.Sleep(1 * time.Second)
			logrus.Println("Богатырь: чел, давай на юзефа, кто выиграл тот и выиграл! А то затянулось это..")
			time.Sleep(5 * time.Second)
			logrus.Println("Змей: а давай!\n\n")
			time.Sleep(5 * time.Second)

			usefa := []string{"камень", "ножницы", "бумага"}

			logrus.Println("Богатырь выбрал..")
			time.Sleep(2 * time.Second)
			randUsefaHero := rand.Intn(3)
			logrus.Println(usefa[randUsefaHero], "\n\n")

			logrus.Println("Змей выбрал..")
			time.Sleep(2 * time.Second)
			randUsefaSnake := rand.Intn(3)

			logrus.Println(usefa[randUsefaSnake], "\n\n")

			if randUsefaHero == randUsefaSnake {
				panic("К сожалению ничья поэтому побеждает Влад!")
			}
			if (randUsefaHero == 0 && randUsefaSnake == 1) || (randUsefaHero == 1 && randUsefaSnake == 2) || (randUsefaHero == 2 && randUsefaSnake == 0) {
				logrus.Println("Богатырь: ХААА лашок я выиграл")
				return fmt.Sprintf("Победил %s! Благодаря игре Камень, Ножницы, Бумага. У него был %s", hero.Name, usefa[randUsefaHero]), nil
			}

			logrus.Println("Змей: СЮДААА, заносик")

			snakeWinner := fmt.Sprintf("Змей Феодосий выиграл! Благодаря игре Камень, Ножницы, Бумага! У него был %s. Поэтому богатырь запаниковал и убежал", usefa[randUsefaSnake])

			panic(snakeWinner)
		}

		hit := HeroHit(snake)

		if snake.Heads <= 0 {
			return fmt.Sprintf("Победил %s! Это заняло %d раундов!", hero.Name, battle.Round), nil
		}
		if snake.Heads >= 200 {
			panic("Змей Феодосий выиграл! Богатырь паникует и убегает")
		}
		if snake.Heads == 5 {
			return "", errors.New("Змей: ААААА у меня фобия на 5 голов")
		}

		GrowHeads(snake, hit)
	}
}
