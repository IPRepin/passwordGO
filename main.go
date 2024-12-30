package main

import (
	"fmt"
	"passwordGO/account"
)

func newPassword() *account.Password {
	return &account.Password{}
}

func main() {
	lenPass := promptData()
	myPass := newPassword()
	myPass.GeneratePassword(lenPass)
	myPass.OutputData()
}

func promptData() int {
	fmt.Printf("Длинна пароля: \n")
	var lenPass int
	fmt.Scanln(&lenPass)
	return lenPass
}
