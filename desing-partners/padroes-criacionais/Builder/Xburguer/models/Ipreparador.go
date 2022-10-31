package models

type IPreparador interface {
	SetNomeCombo()
	SetHamburguer()
	SetBebida()
	SetAcompanhamento()
	SetSobremesa()
	SetBrinde()
	SetDescricao()
	GetCombo() Combo
}

func GetPreparador(tipoDeCombo string) IPreparador {
	switch tipoDeCombo {
	case "ComboX":
		return novoComboX()
	case "ComboY":
		return novoComboY()
	case "ComboZ":
		return novoComboZ()
	case "ComboS":
		return novoComboS()
	case "ComboMonsterX":
		return novoComboMonsterX()
	case "OnlyFries":
		return novoOnlyFries()
	case "IceCloudX":
		return novoIceCloudX()
	default:
		return novoComboX()
	}
}
