package models

import "gorm.io/gorm"

type IPersonagem interface {
	SetID(id uint)
	SetDefesa(defesa int)
	SetClass(class string)
	SetName(name string)
	SetHP(hp int)
	SetMP(mp int)
	SetPoder(poder int)
	SetForca(forca int)
	SetAgilidade(agilidade int)
	SetLevel(level int)
	SetXP(xp int)
	GetClass() string
	GetName() string
	GetHP() int
	GetMP() int
	GetPoder() int
	GetForca() int
	GetAgilidade() int
	GetLevel() int
	GetXP() int
	GetID() uint
	GetDefesa() int
}

type Personagem struct {
	gorm.Model
	Class, Name                                        string
	Defesa, HP, MP, Poder, Forca, Agilidade, Level, XP int
}

func (p *Personagem) SetID(id uint) {
	p.ID = id
}
func (p *Personagem) GetID() uint {
	return p.ID
}
func (p *Personagem) SetClass(class string) {
	p.Class = class
}

func (p *Personagem) SetName(name string) {
	p.Name = name
}
func (p *Personagem) SetHP(hp int) {
	p.HP = hp
}
func (p *Personagem) SetMP(mp int) {
	p.MP = mp
}
func (p *Personagem) SetPoder(poder int) {
	p.Poder = poder
}
func (p *Personagem) SetForca(forca int) {
	p.Forca = forca
}
func (p *Personagem) SetAgilidade(agilidade int) {
	p.Agilidade = agilidade
}
func (p *Personagem) SetLevel(level int) {
	p.Level = level
}
func (p *Personagem) SetXP(xp int) {
	p.XP = xp
}
func (p *Personagem) GetClass() string {
	return p.Class
}
func (p *Personagem) GetName() string {
	return p.Name
}
func (p *Personagem) GetHP() int {
	return p.HP
}
func (p *Personagem) GetMP() int {
	return p.MP
}
func (p *Personagem) GetPoder() int {
	return p.Poder
}
func (p *Personagem) GetForca() int {
	return p.Forca
}
func (p *Personagem) GetAgilidade() int {
	return p.Agilidade
}
func (p *Personagem) GetLevel() int {
	return p.Level
}
func (p *Personagem) GetXP() int {
	return p.XP
}
func (p *Personagem) SetDefesa(defesa int) {
	p.Defesa = defesa
}
func (p *Personagem) GetDefesa() int {
	return p.Defesa
}
