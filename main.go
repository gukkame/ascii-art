package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "io"
    "log"
)
func main() {

	if len(os.Args)  > 1 {

		incomingSentence := os.Args[1]
		
		


		//Printing words so that \n marks the start of a new line
		input_slice := strings.Split(incomingSentence,"\\n")
		
		for _, word := range input_slice {

			var ch, position rune	
			var arr []int 
		
			for _, ch = range word {
				position = (ch - 32) * 9
				arr = append(arr, int(position)) // array of NUMBERS of first lines of input charachters
			}
			if word == ""{
				fmt.Println("")
			}else{
			Printer(word, "standard.txt", arr)
			}
		}
	} else {
		fmt.Println("Please enter more than 0 arguments")
	}
}

func Printer(input, fontName string, originalArray[] int) {
	for lineCounter := 1; lineCounter < 9; lineCounter++ { //loop through the heigth of the ascii characters
        for i,_ := range originalArray{
			oneLineofOneLetter, _ := ReadLine(fontName, originalArray[i]+lineCounter)
			fmt.Print(oneLineofOneLetter) //just print a bunch of same "height" lines of different letters in a row 
		}
		fmt.Print(" ")
		fmt.Print("\n") //and start a new line again
	}
}



func ReadLine(fontName string, startingLineEach int) (line string, err error) {

    var lastLineOfScanner int
    FontFile, err := os.Open(fontName) //open file
	if err != nil { //check for errors
		log.Fatal(err)
	}
	defer FontFile.Close() //close file
	scanner := bufio.NewScanner(FontFile) // make a scaner machine aka not doing anything yet

	for scanner.Scan() { //loop through lines until you reached the wished line aka start machine 

        if lastLineOfScanner == startingLineEach {
			return scanner.Text(), scanner.Err() //scanner.Text() is where you store the result of scanner you can print it aswell if you make it equal to a value         
		}
		lastLineOfScanner++
	}
	return line, io.EOF // it will return the line when the scanner mathces reading to the input line number aka gets to it
}
