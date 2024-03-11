package main

import (
	"fmt"
	"time"
)

func main(){

	count_ramadhan := time.Now().AddDate(0, 1, 0)
	if time.Now().Before(count_ramadhan) {
		fmt.Println("Selamat Menunaikan Ibadah Puasa Ramadhan 1445 H")
	} else {
		fmt.Println("Selamat Hari Raya Idul Fitri 1445 H")
	}
}