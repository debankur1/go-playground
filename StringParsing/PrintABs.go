package main

import "fmt"

func main()  {
	txt := "ABABAB??AB??AB"
	i := 0
	var startChar string
	var startCharIndetification bool
	var result string
	for i< len(txt){
		startCharIndetification = false
		exclamationCount := 0
		 	 for string(txt[i])=="?"{
		 	 	if startCharIndetification==false {
					startChar = string(txt[i-1])
				}
		 	 	 exclamationCount++
		 	 	 i++
		 	 	 startCharIndetification = true
		 	 }
		        if exclamationCount > 0 {
					  if startChar != string(txt[i]) && exclamationCount%2 != 0 {
					  	  result = "NA"
						 // fmt.Println("Not possible")
						  break
					  }
					  if startChar != string(txt[i]) && exclamationCount%2 == 0 {
					  	  result = "P"
						//  fmt.Println("possible")
					  }
					  if startChar == string(txt[i]) && exclamationCount%2 != 0 {
					  	 result = "P"
						//  fmt.Println("possible")
					  }
					  if startChar == string(txt[i]) && exclamationCount%2 == 0 {
					  	  result = "NA"
						//  fmt.Println("Not possible")
						  break
					  }
			    }
           i++
	}
	fmt.Println(result)
}
