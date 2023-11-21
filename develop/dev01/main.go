package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

func main() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, fmterr := fmt.Fprint(os.Stderr, err.Error())
		if fmterr != nil {
			log.Fatal(fmterr)
		}
		os.Exit(1)
	}

	ntpTimeFormatted := ntpTime.Format(time.UnixDate)

	fmt.Printf("Network time: %v\n", ntpTime)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
	fmt.Println("-------------------------------------------")
	timeFormatted := time.Now().Local().Format(time.UnixDate)
	fmt.Printf("System time: %v\n", time.Now())
	fmt.Printf("Unix Date System time: %v\n", timeFormatted)

}
