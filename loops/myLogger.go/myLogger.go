package myLogger

func ListenLog(ch chan string) {
	for {
		msg := <- ch
		log.Println(msg)
	}
}