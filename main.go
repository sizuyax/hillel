package main

import (
	"errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Generator interface {
	Generate() interface{}
	GenerateSlice() []interface{}
	GenerateWithParam(param int) (interface{}, error)
}

type IntGenerator struct{}

func (IntGenerator) Generate() interface{} {
	return rand.Int()
}

func (IntGenerator) GenerateSlice() []interface{} {
	rand.Seed(time.Now().UnixNano())

	b := make([]interface{}, 10)

	for i := range b {
		randNum := rand.Intn(10)
		b[i] = randNum
	}

	return b
}

func (IntGenerator) GenerateWithParam(param int) (interface{}, error) {
	if param < 0 {
		return nil, errors.New("param should be non-negative")
	}
	return rand.Intn(param), nil
}

type StringGenerator struct{}

func (StringGenerator) Generate() interface{} {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (StringGenerator) GenerateSlice() []interface{} {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]interface{}, 10)
	for i := range b {
		randomString := make([]rune, 4)
		for j := range randomString {
			randomString[j] = letters[rand.Intn(len(letters))]
		}
		b[i] = string(randomString)
	}
	return b
}

func (StringGenerator) GenerateWithParam(param int) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())

	if param < 0 {
		return nil, errors.New("param should be non-negative")
	}

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, param)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b), nil
}

type BoolGenerator struct{}

func (BoolGenerator) Generate() interface{} {
	return rand.Intn(2) == 1
}

func (BoolGenerator) GenerateSlice() []interface{} {
	rand.Seed(time.Now().UnixNano())

	b := make([]interface{}, 10)
	for i := range b {
		b[i] = rand.Intn(2) == 1
	}
	return b
}

func (BoolGenerator) GenerateWithParam(param int) (interface{}, error) {
	if param < 0 || param > 1 {
		return nil, errors.New("param should be 0 or 1")
	}

	randBool := rand.Intn(param)
	if randBool == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

type FloatGenerator struct{}

func (FloatGenerator) Generate() interface{} {

	val := rand.Float64() * float64(100)

	val = math.Round(val*100) / 100

	symbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomSymbol := symbols[rand.Intn(len(symbols))]

	valStr := strconv.FormatFloat(val, 'f', -1, 64)

	valStr += string(randomSymbol)

	return valStr
}

func (FloatGenerator) GenerateSlice() []interface{} {
	rand.Seed(time.Now().UnixNano())

	var result []interface{}

	for i := 0; i < 10; i++ {
		val := rand.Float64() * 100
		val = math.Round(val*100) / 100

		symbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		randomSymbol := symbols[rand.Intn(len(symbols))]

		valStr := strconv.FormatFloat(val, 'f', -1, 64)
		valStr += string(randomSymbol)

		result = append(result, valStr)
	}

	return result
}

func (FloatGenerator) GenerateWithParam(param int) (interface{}, error) {
	if param < 0 {
		return nil, errors.New("param should be non-negative")
	}

	val := rand.Float64() * float64(param)

	val = math.Round(val*100) / 100

	symbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomSymbol := symbols[rand.Intn(len(symbols))]

	valStr := strconv.FormatFloat(val, 'f', -1, 64)

	valStr += string(randomSymbol)

	return valStr, nil
}

func GenerateAndPrint(g Generator) {
	logrus.Println("Используем Generate:")
	logrus.Println("Тип:", g.Generate(), "\n\n")
	logrus.Println("Используем GenerateSlice:")
	logrus.Println("Тип:", g.GenerateSlice(), "\n\n")

	val, err := g.GenerateWithParam(10)
	if err != nil {
		zap.Error(err)
	} else {
		logrus.Println("Используем GenerateWithParam (param = 10):")
		logrus.Println("Тип:", val)
	}

	switch val := val.(type) {
	case int:
		logrus.Println("Generated type is int: ", val)
	case float64:
		logrus.Println("Generated type is float64: ", val)
	case bool:
		logrus.Println("Generated type is bool: ", val)
	case string:
		logrus.Println("Generated type is string: ", val)
	default:
		logrus.Println("Unknown generated type")
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	zapLogger := logger.Sugar()

	_, err = g.GenerateWithParam(-1)
	if err != nil {
		zapLogger.Errorw("Error generating with param", "error", err)
	}

	_, err = g.GenerateWithParam(2)
	if err != nil {
		zapLogger.Errorw("Error generating with param", "error", err)
	}
}

func main() {
	intGen := IntGenerator{}
	stringGen := StringGenerator{}
	boolGen := BoolGenerator{}
	floatGen := FloatGenerator{}

	GenerateAndPrint(intGen)
	GenerateAndPrint(stringGen)
	GenerateAndPrint(boolGen)
	GenerateAndPrint(floatGen)
}
