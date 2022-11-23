package main

import (
	"fmt"
	"math/rand"
)

func RangeOfHash(x int) int {
	if x >= 58 && x <= 64 {
		x += 7
		return x
	}
	if x >= 91 && x <= 96 {
		x += 6
		return x
	}
	return x
}

func GenerateHashByte(pass string) string {
	hash := ""
	cofre := make([]int, 5)
	var segredo int
	for _, h := range pass {
		cofre = append(cofre, int(h))
	}
	for i := 0; i < len(cofre); i++ {
		segredo += cofre[i]
	}
	rand.Seed(int64(segredo))
	mySena()
	for i := 0; i < len(cofre); i++ {
		x := rand.Intn(122-48+1) + 48
		x = RangeOfHash(x)
		//println("X: ", x)
		hash += string(x)
	}

	println("hash: ", hash)
	println("len: ", len(hash))
	if len(hash) > 32 {
		cut := len(hash) - 32
		hash = hash[:len(hash)-cut]
		println("new len: ", len(hash))
		println("hash: ", hash)
	}
	if len(hash) < 32 {
		need := 32 - len(hash)
		for i := 0; i < need; i++ {
			x := rand.Intn(122-48+1) + 48
			x = RangeOfHash(x)
			hash += string(x)

		}
		println("new len: ", len(hash))
		println("hash: ", hash)
	}

	return string(hash)
}

func VerificarSenhaSecreta(pass string) bool {
	if pass == "wEBHqEHkPoMngxeGBWMEaFcAfXEwfhKR" {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("feito!")
	if VerificarSenhaSecreta(GenerateHashByte("Se for da vontade de Deus!")) {
		println("Autorizado!")
	} else {
		println("Acesso negado!")
	}
}
func mySena() {
	for m := 0; m < 10; m++ {
		for s := 0; s < 6; s++ {
			numbers := make([]int32, 1)
			number := rand.Int31n(59) + 1
			numbers = append(numbers, number)
			for loop := 0; loop < 1; {
				for j := 0; j < len(numbers); j++ {
					//fmt.Println("Number is: ", number, "\nV is: ", v)
					//time.Sleep(time.Second)
					if numbers[j] == number {

						number = rand.Int31n(59) + 1
						numbers[j] = number
					} else {
						loop++
					}
				}
			}

			fmt.Print(" | ", number, " | ")
		}
		println()
	}
}
