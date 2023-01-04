package main

import (
	"fmt"
	"sync"
)

func entranceexam(Wg *sync.WaitGroup) {
	fmt.Println("entrance exam")
	Wg.Done()
}
func score(Wg *sync.WaitGroup) {
	fmt.Println("score")
	Wg.Done()
}
func engineeringcollege(Wg *sync.WaitGroup) {
	fmt.Println("engineering college")
	Wg.Done()
}
func stream(Wg *sync.WaitGroup) {
	fmt.Println("stream")
	Wg.Done()
}
func friends(Wg *sync.WaitGroup) {
	fmt.Println("friends")
	Wg.Done()
}
func nickname(Wg *sync.WaitGroup) {
	fmt.Println("nickname")
	Wg.Done()
}
func funjokes(Wg *sync.WaitGroup) {
	fmt.Println("fun jokes")
	Wg.Done()
}
func troublepeople(Wg *sync.WaitGroup) {
	fmt.Println("trouble people")
	Wg.Done()
}
func fighting(Wg *sync.WaitGroup) {
	fmt.Println("fighting")
	Wg.Done()
}

func main() {
	fmt.Println("I LOVE INDIA")

	var Wg sync.WaitGroup
	Wg.Add(9)

	go entranceexam(&Wg)
	
	go score(&Wg)
	
	go engineeringcollege(&Wg)
	
	go stream(&Wg)
	
	go friends(&Wg)
	
	go nickname(&Wg)

	go funjokes(&Wg)
	
	go troublepeople(&Wg)
	
	go fighting(&Wg)
	
	Wg.Wait()
	fmt.Println("good luck")
}






