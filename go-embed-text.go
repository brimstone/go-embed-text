package embedtext

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	// Token is used to delimite the code from the actul message
	Token = "FINDME"
	// PartialCount is the number of times the token will apper in normal code before actually delimiting code from message
	PartialCount = 2
	// Exe is the path to the exe to modify
	Exe string
)

func init() {
	Exe, _ = os.Executable()
}

// Embed stores a message in the Exe
func Embed(text string) error {
	elf, _ := ioutil.ReadFile(Exe)
	partials := strings.SplitN(string(elf), Token, 3)
	if len(partials) < PartialCount+1 {
		partials = append(partials, "")
	}
	partials[PartialCount] = text
	err := ioutil.WriteFile(Exe+"new", []byte(strings.Join(partials, Token)), 0755)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
	os.Rename(Exe+"new", Exe)

	return nil
}

// Read returns the message stored in the Exe
func Read() (message string) {
	elf, _ := ioutil.ReadFile(Exe)
	partials := strings.SplitN(string(elf), Token, 3)
	if len(partials) < PartialCount+1 {
		partials = append(partials, "")
	}

	return partials[PartialCount]
}
