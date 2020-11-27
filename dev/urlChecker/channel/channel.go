package channel

import "time"

func isAdmin(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
