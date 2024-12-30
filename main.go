package main

import (
	"fmt"
	"passwordGO/account"
	"passwordGO/files"
)

func newPassword() *account.Password {
	return &account.Password{}
}

func main() {
	files.ReadFile()
	lenPass := promptData()
	myPass := newPassword()
	password := myPass.GeneratePassword(lenPass)
	myPass.OutputData()
	files.WriteFile(password, "password.txt")
}

func promptData() int {
	fmt.Printf("Длинна пароля: \n")
	var lenPass int
	fmt.Scanln(&lenPass)
	return lenPass
}
