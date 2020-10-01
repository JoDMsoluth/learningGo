package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {

	// 초기화 하지 않은 map은 따로 변경할 수 없다. 때문에 make 사용
	var results = make(map[string]string)

	c := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}

	// go 루틴 사용해서 병렬적으로 실행 할 수 있다.
	// go 루틴은 main 함수가 실행되는 동안 유용하다.
	// main function은 go 루틴을 기다려주지 않고 끝낸다.
	go sexyCount("nico")
	sexyCount("jo")

	c2 := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c2)
	}

	// 채널에서 답이 올때까지 기다려준다.
	// 채널은 동기적으로 처리됨
	// 따라서 응답을 받지 못하면 교착상태가 됨 하단의 패턴 사용
	// requestResult := <-c
	// fmt.Println(requestResult)
	// fmt.Println(<-c)

	// 이 패턴으로 교착상태 방지
	for i := 0; i < len(people); i++ {
		fmt.Print("wating for", i, "\n")
		fmt.Println(<-c2)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is Sexy"
}

func hitURL(url string, c chan<- requestResult) {
	fmt.Println("checking", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}

}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
