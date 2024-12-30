package account

import (
	"github.com/fatih/color"
	"math/rand"
)

type Password struct {
	password string
}

func (pass *Password) OutputData() {
	color.Cyan("%s\n", pass.password)
	//fmt.Printf("%s\n", pass.password)
}

func (pass *Password) GeneratePassword(num int) string {
	simbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, num)
	for i := range result {
		result[i] = simbols[rand.Intn(len(simbols))]
	}
	pass.password = string(result)
	return pass.password
}
