package main
// generate a shit ton of spam to make website shit
// fuck you i wrote it in go
import "os"

func main() {
	f, _ := os.Create("spam")
	defer f.Close()
	for i:=0; i<99999999999; i++ {
		f.WriteString("u")
	}
}