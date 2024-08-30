package chat
import (
	"io"
	"log"
	"os"
)
func OpenTxtFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening txt file")
	}
	defer file.Close()
	result, _ := io.ReadAll(file)
	return result
}
func WelcomeMessage() []byte {
	txt := OpenTxtFile("NameIcon.txt")
	return txt
}