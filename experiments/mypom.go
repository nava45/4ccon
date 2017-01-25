package main

import (
	"fmt"
	"time"
)

const POM_TIME int = 60
const LON_INT int = 10
const SHO_INT int = 5

var tot int = 3
var count int = 1

func shortBreak(c chan int) {
	fmt.Println("Take 3 sec Break")
	time.Sleep(time.Second * 3)
	c <- 1
}

func longBreak(c chan int) {
	fmt.Println("LONG break - 7 secs")
	time.Sleep(time.Second * 7)
	c <- 1
}

func pomStart(_in chan int) {
	fmt.Println("FOCUS ON YOUR WORK NOW.")
	time.Sleep(time.Second * 10)
	_in <- 1
}

func startAgain() string {
	fmt.Println("Again start Y/N?")
	var input string
	fmt.Scanln(&input)
	return input
}

func interver_start(pom, sh, lo, exit chan int) {
	for {
		select {
		case _ = <-pom:
			fmt.Println("Pomodoro completed count:", count)
			if count < tot {
				go shortBreak(sh)
			} else if count == tot {
				go longBreak(lo)
			}
			count += 1

		case _ = <-sh:
			go pomStart(pom)

		case _ = <-lo:
			res := startAgain()
			if res == "Y" {
				count = 1
				go pomStart(pom)
			} else {
				exit <- 1
			}
		}
	}

}

func main() {
	fmt.Println("Our Pomodoro!!")

	pom_flag := make(chan int)
	shot_flag := make(chan int)
	long_flag := make(chan int)
	can_exit := make(chan int)

	go pomStart(pom_flag)
	go interver_start(pom_flag, shot_flag, long_flag, can_exit)

	<-can_exit
}
