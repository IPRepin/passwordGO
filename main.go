package main

import (
	"fmt"
	"passwordGO/account"
	"passwordGO/files"
)

func newPassword() *account.NewPassword {
	return &account.NewPassword{}
}

func main() {
	CreatePassword()
}

func CreatePassword() {
	lenPass := promptData()
	myPass := newPassword()
	myPass.GeneratePassword(lenPass)
	myPass.OutputData()
	file, err := myPass.ToBytes()
	if err != nil {
		fmt.Println(err)
	}
	files.WriteFile(file, "data.json")
}

func promptData() int {
	fmt.Printf("Длинна пароля: \n")
	var lenPass int
	fmt.Scanln(&lenPass)
	return lenPass
}
