package errorhandling

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func Test() {
	err := os.WriteFile("test.txt", []byte("test text"), 000)
	var epErr *fs.PathError
	if err != nil && errors.As(err, &epErr) {
		fmt.Println("Error Creating a file: ", err)
	} else {
		fmt.Println(" file(s) Created")
	}

	f, err := os.Open("test.txt")
	//defer f.Close()
	if err != nil {
		var pErr *os.PathError
		if errors.As(err, &pErr) {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
