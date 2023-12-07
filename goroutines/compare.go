package goroutines

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return b
}

const homePath = "C:\\Users\\anilc\\golang\\goroutines\\comparefiles"

func generateFiles() {

	for i := 1; i <= 2; i++ {
		folderPath := homePath + "\\folder" + fmt.Sprint(i)
		err := os.MkdirAll(folderPath, 0777)
		if err != nil {
			fmt.Println("Error Creating a folder")
		} else {
			fmt.Println("Folder Created with name: ", folderPath)
		}
		for j := 1; j <= 10; j++ {
			text := generateRandomString(10)
			err := os.WriteFile(folderPath+"\\file"+fmt.Sprint(j)+".txt", text, 077)
			if err != nil {
				fmt.Println("Error Creating a file")
			} else {
				fmt.Println(j, " file(s) Created")
			}
		}

	}

}

func Compare() {

	generateFiles()
	compareFolders()

}

func compareFolders() {

	channnel := make(chan bool)
	entries, err := os.ReadDir(homePath + "\\folder1")
	if err != nil {
		fmt.Println("error reading directory")
	}

	compareFiles(channnel, entries)

}

func compareFiles(c chan bool, entries []fs.DirEntry) {

	for i, e := range entries {

		file1 := homePath + "\\folder1\\" + e.Name()
		file2 := homePath + "\\folder2\\" + e.Name()

		file1Content, err1 := os.ReadFile(file1)
		if err1 != nil {
			fmt.Println("error reading file")
		}
		file2Content, err2 := os.ReadFile(file2)
		if err2 != nil {
			fmt.Println("error reading file")
		}

		go compareFile(c, string(file1Content), string(file2Content), i, len(entries))
		// if string(file1Content) == string(file2Content) {
		// 	c <- true
		// } else {
		// 	c <- false
		// }

	}
	//time.Sleep(1 * time.Second)
	// close(c)
	for v := range c {

		fmt.Println("Received ", v)
	}

}

func compareFile(c chan bool, file1Content string, file2Content string, index, l int) {

	time.Sleep(5 * time.Second)
	if string(file1Content) == string(file2Content) {
		c <- true
	} else {
		c <- false
	}
	if (index + 1) == l {
		fmt.Println(index, l)
		close(c)
	}
	//close(c)
}
