package account

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
)

type NewPassword struct {
	Password string `json:"password"`
}

func (pass *NewPassword) OutputData() {
	color.Cyan("%s\n", pass.Password)
	//fmt.Printf("%s\n", pass.password)
}

func (pass *NewPassword) GeneratePassword(num int) string {
	simbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, num)
	for i := range result {
		result[i] = simbols[rand.Intn(len(simbols))]
	}
	pass.Password = string(result)
	//passTag, _ := reflect.TypeOf(pass).Elem().FieldByName("password")
	//fmt.Println(passTag)
	return pass.Password
}

func (pass *NewPassword) ToBytes() ([]byte, error) {
	file, err := json.Marshal(pass)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return file, nil
}
