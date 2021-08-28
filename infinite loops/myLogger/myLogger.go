package myLogger

import "log"

func ListenForLog(ch chan string) {
	for {
		str := <-ch
		log.Println(str)
	}
}
