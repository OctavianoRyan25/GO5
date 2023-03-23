package main

import (
	"fmt"
	"sync"
)

// type coba interface {
// 	bisa()
// 	coba()
// }

// type struct1 struct{}

// func (Struct1 struct1) coba() {
// 	variabel := "coba1 coba2 coba3"
// 	for i := 0; i < 4; i++ {
// 		fmt.Println(variabel)
// 	}
// }
// func (Struct1 struct1) bisa() {
// 	variabel := "bisa1 bisa2 bisa3"
// 	for i := 0; i < 4; i++ {
// 		fmt.Println(variabel)
// 	}
// }

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	variabel := []string{"coba1, coba2, coba3"}
	variabel2 := []string{"bisa1, bisa2, bisa3"}

	for i := 0; i <= 3; i++ {
		wg.Add(2)
		go func() {
			fmt.Printf("%s %v \n", variabel, i)
			wg.Done()
		}()
		ii := i
		go func() {
			fmt.Printf("%s %v \n", variabel2, ii)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Menggunakan Lock")

	for i := 0; i <= 3; i++ {
		wg.Add(2)
		mx.Lock()
		go func() {
			fmt.Printf("%s %v \n", variabel, i)
			mx.Unlock()
			wg.Done()
		}()
		ii := i
		mx.Lock()
		go func() {
			fmt.Printf("%s %v \n", variabel2, ii)
			wg.Done()
			mx.Unlock()
		}()
	}
	wg.Wait()
}



