package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// var data string = "2x²+3x-5"
	data := "-5+2x²+3x"
	// var data string

	fmt.Print("Informe a equação de segundo grau: ")
	// fmt.Scanln(&data)
	fmt.Println(data)
	x1, x2 := calculaBhaskara(data)
	fmt.Printf("x1 = %v, x2 = %v \n", x1, x2)
}

func calculaBhaskara(data string) (x1 float64, x2 float64) {
	a, b, c, err := setItens(data)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Printf("a = %v, b = %v, c = %v \n", a, b, c)

	delta := calculateDelta(a, b, c)

	x1 = (b + math.Sqrt(delta)) / (2 * a)
	x2 = (b - math.Sqrt(delta)) / (2 * a)

	return
}

func setItens(data string) (a float64, b float64, c float64, err error) {
	if len(data) <= 0 {
		return 0, 0, 0, errors.New("informe uma equação")
	}

	equation := splitOnMinus(data)

	equation = splitOnPlus(equation)

	for _, e := range equation {
		if strings.Contains(e, "x²") {
			a, err = strconv.ParseFloat(strings.TrimRight(e, "x²"), 64)
			if err != nil {
				return
			}
			continue
		}
		if strings.Contains(e, "x") {
			b, err = strconv.ParseFloat(strings.TrimRight(e, "x"), 64)
			if err != nil {
				return
			}
			continue
		}

		c, err = strconv.ParseFloat(e, 64)
		if err != nil {
			return
		}
		continue

	}

	return
}

func splitOnPlus(equation []string) (res []string) {
	for _, s := range equation {
		res = append(res, strings.Split(s, "+")...)
	}

	return
}

func splitOnMinus(data string) (res []string) {
	firstCharacterIsMinus := strings.HasPrefix(data, "-")

	res = strings.Split(data, "-")

	if firstCharacterIsMinus {
		res = res[1:]
	}

	for i, s := range res {
		res[i] = "-" + s
	}

	return
}

func calculateDelta(a float64, b float64, c float64) (delta float64) {
	delta = math.Pow(b, 2) - 4*a*c

	return
}
