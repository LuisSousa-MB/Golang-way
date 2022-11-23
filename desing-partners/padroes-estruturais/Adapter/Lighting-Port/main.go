package main

import "LighitingPort/model"

func main() {

	cliente := &model.Cliente{}
	mac := &model.Mac{}
	//win := &model.Windows{}

	cliente.InserirNoConectorLightingDoComputador(mac)

	computadorWindows := &model.Windows{}
	adaptadorParaWindows := &model.AdaptadorWindows{
		ComputadorWindows: computadorWindows,
	}

	cliente.InserirNoConectorLightingDoComputador(adaptadorParaWindows)
}
