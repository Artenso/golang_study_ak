package main

import "fmt"

type TVer interface {
	switchOFF() bool
	switchOn() bool
	GetStatus() bool
	GetModel() string
}

type Sumsunger interface {
	TVer
	SamsungHub() string
}

type LGer interface {
	TVer
	LGHub() string
}

type LgTV struct {
	status bool
	model  string
}

func (tv *LgTV) switchOFF() bool {
	tv.status = false
	return true
}

func (tv *LgTV) switchOn() bool {
	tv.status = true
	return true
}

func (tv *LgTV) GetStatus() bool {
	return tv.status
}

func (tv *LgTV) GetModel() string {
	return tv.model
}

func (tv *LgTV) LGHub() string {
	return "LGHub"
}

type SamsungTV struct {
	status bool
	model  string
}

func (tv *SamsungTV) switchOFF() bool {
	tv.status = false
	return true
}

func (tv *SamsungTV) switchOn() bool {
	tv.status = true
	return true
}

func (tv *SamsungTV) GetStatus() bool {
	return tv.status
}

func (tv *SamsungTV) GetModel() string {
	return tv.model
}

func (tv *SamsungTV) SamsungHub() string {
	return "SumsungHub"
}

func main() {
	tv := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	fmt.Println(tv.GetStatus())  // true
	fmt.Println(tv.GetModel())   // Samsung XL-100500
	fmt.Println(tv.SamsungHub()) // SamsungHub
	fmt.Println(tv.GetStatus())  // false
	fmt.Println(tv.switchOn())   // true
	fmt.Println(tv.GetStatus())  // true
	fmt.Println(tv.switchOFF())  // true
	fmt.Println(tv.GetStatus())  // false
}
