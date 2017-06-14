package main

import "fmt"
func main(){
	jch := jpungAuth("aminov")
	gch := gongsterAuth("aminov")

	var (
		jvalid  bool
		gvalid  bool
	)

	for i := 0; i < 2; i++{
		select {
		case jvalid =  <-jch:
		case gvalid = <-gch:
		}
	}

	if jvalid || gvalid {
		fmt.Println("Auth pass")
	}else{
		fmt.Println("Auth fail")
	}
}

func jpungAuth(name string) chan bool {
	jch := make(chan bool)
	go func() {
		jch <- false
	}()
	return jch
}

func gongsterAuth(name string) chan bool{
	gch := make(chan bool)
	go func(){
		gch <- true
	}()
	return gch
}