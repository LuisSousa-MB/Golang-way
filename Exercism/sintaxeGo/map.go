package sintaxego

import "fmt"

func main() {

	cidade := make(map[string]int)
	AddToMap(cidade, "Sao paulo", 12_000_000)
	fmt.Println(cidade)
	DeleteFromMap(cidade, "Sao paulo")
	fmt.Println(cidade)

}

func AddToMap(m map[string]int, c string, p int) {
	m[c] = p

}

func DeleteFromMap(m map[string]int, c string) {
	delete(m, c)
}
