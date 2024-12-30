package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	return
}

func WriteFile(pass string, nameFile string) {
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(pass)
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("File created")
	file.Close()
}
