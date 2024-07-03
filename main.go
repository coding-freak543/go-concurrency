// waitgroups project main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	Birthday()
}

func Birthday() {
	fmt.Println("\n*************  Starting birthday Function   *************")

	var wg sync.WaitGroup //waitgroup to monitor the function

	wg.Add(2)

	go passingParcel(&wg)
	go musicalChair(&wg)

	wg.Wait()
	cutcake()

}

func passingParcel(wg *sync.WaitGroup) {
	fmt.Println("\n ************* Group1 started playing PassingParcel ************* \n*")

	for i := 0; i < 5; i++ {

		fmt.Println(" PassingParcel is in progress")
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\n************* PassingParcel done *************")

	wg.Done()

}

func musicalChair(wg *sync.WaitGroup) {
	fmt.Println("\n************* Group2 started playing musicalChair *************\n")

	for i := 0; i < 5; i++ {

		fmt.Println(" MusicalChair is in progress")
		time.Sleep(500 * time.Millisecond)

	}
	fmt.Println("\n************* muscial chairs done *************")
	wg.Done()

}

func cutcake() {

	fmt.Println("\n ******************** Games Over, Time to cut the cake ********************")
	fmt.Print(`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡖⠙⡢⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣿⡏⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⣿⡿⣿⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⠶⠞⠛⠿⣿⣿⡿⣿⣇⡀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡾⠋⢠⠄⠀⠀⠀⠀⠈⠉⡉⠈⠛⢦⡀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣤⠀⠀⠀⠀⠀⠻⢦⣴⣀⣀⠀⠀⠀⠀⢈⠁⣀⡄⠀⠹⣆⠀⠀⠀
⠀⠀⠀⠀⢰⠓⡄⠀⠀⠀⠀⠀⢸⡇⠉⠉⠙⠛⠛⠛⠛⠛⠉⠀⠀⠀⠹⣆⠀⠀
⢠⣒⣊⢍⣩⢙⣩⣍⡩⡇⠀⠀⢸⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⡀⠀
⣸⣀⣈⣁⣀⣉⣀⣀⣠⣇⠀⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡆⠘⣇⠀
⠉⠉⠉⠛⠿⠻⣯⠉⠉⠉⠀⢀⣾⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⢿⠀
⠀⠀⠀⠀⠀⠀⠈⢳⣄⠀⢠⡞⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⢸⡇
⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠋⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠁⠀⠘⣇
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⠀⠀⠀⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠟⠁⠀⠀⢻
`)
}
