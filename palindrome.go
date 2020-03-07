package main

import "fmt"
import "unicode"

func main(){
	temp:="^^KADAK^%^"
	isPal:=true
	for i,j:=0,len(temp)-1;i<j; {
		if !unicode.IsLetter(rune(temp[i])) {
			i=i+1
			continue
		}
		if !unicode.IsLetter(rune(temp[j])) {
			j=j-1
			continue
		}		

		if temp[i] != temp[j]{
			isPal = false
		}
		i=i+1
		j=j-1		
	}	
	fmt.Println("Given word is palindrome:", isPal)
}



