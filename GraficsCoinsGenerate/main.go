package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Criamos um gráfico e definimos o título e os rótulos dos eixos
	p := plot.New()

	p.Title.Text = "Volatilidade de Preço das Criptomoedas"
	p.X.Label.Text = "Tempo (em horas)"
	p.Y.Label.Text = "Preço (em dólares)"

	// Geramos alguns dados aleatórios para o gráfico
	n := 24
	precos := make(plotter.XYs, n)
	for i := 0; i < n; i++ {
		precos[i].X = float64(i)
		precos[i].Y = 16000 + rand.Float64()*1000
	}

	// Adicionamos os dados ao gráfico e salvamos o gráfico em um arquivo PNG
	l, err := plotter.NewLine(precos)
	if err != nil {
		panic(err)
	}
	p.Add(l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "precos.png"); err != nil {
		panic(err)
	}
}
