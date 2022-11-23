package model

import "fmt"

type AdaptadorWindows struct {
	ComputadorWindows *Windows
}

func (w *AdaptadorWindows) InserirNoConectorLighting() {
	fmt.Println("O Adaptador converteu o sinal do Lighting para USB.")
	w.ComputadorWindows.InserirNaPortaUSB()
}
