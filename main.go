//Here below is a example of using goroutines,channel,waitgroup

package main

import (
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()
	id := getUserByName("varma")
	log.Println(id)
	//waitgroups
	wg := &sync.WaitGroup{}
	//channel
	ch := make(chan *Message, 2)

	wg.Add(2)

	// chats := getUserByChats(id)
	go getUserByChats(id, ch, wg)

	// friends := getUserFriends(id)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}
	// log.Println(chats)
	// log.Println(friends)
	log.Println(time.Since(now))

}

// func getUserFriends(id string,ch chan<-*Message) []string {
// 	time.Sleep(time.Second * 1)
// 	return []string{
// 		"varama",
// 		"singh",
// 		"patil",
// 		"ravi",
// 		"kumar",
// 	}
// }

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	ch <- &Message{
		friends: []string{
			"varama",
			"singh",
			"patil",
			"ravi",
			"kumar",
		},
	}
	//close(ch)
	wg.Done()
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)
	return name

}

// func getUserByChats(id string,ch chan<-*Message) []string {
// 	return []string{
// 		"varama",
// 		"singh",
// 		"patil",
// 	}
// }

func getUserByChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	ch <- &Message{
		chats: []string{
			"varama",
			"singh",
			"patil",
		},
	}
	//close(ch)
	wg.Done()
}
