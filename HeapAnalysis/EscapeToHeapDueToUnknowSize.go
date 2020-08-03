package main

func main(){
	b := make([]byte,32)
	ReadNoEscape(b)
	ReadEscape()
}
func ReadNoEscape(b []byte){
}
func ReadEscape()[]byte{
	b := make([]byte,32)
	return b
}