package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	data, err := os.ReadFile("password.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func WriteFile(pass []byte, nameFile string) {
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File created")
}
