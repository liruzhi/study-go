package service

import (
	"fmt"
	"time"
)

var chan1 = make(chan int, 1)
var chan2 = make(chan int)
var finishChan= make(chan bool)

func OutputStrAndNum() {
	chan1 <- 1
	go f1()
	go f2()
	<-finishChan
}

func f1() {
	for i:=1;i<11;i++ {
		<-chan1
		fmt.Println(i)
		chan2<-1
	}
}

func f2() {
	str := "ABCDEFGHIJ"
	for i:=0;i<10;i++ {
		<-chan2
		fmt.Println(string(str[i]))
		chan1<-1
	}
	close(chan1)
	close(chan2)
	close(finishChan)
}

//两个协程交替打印奇偶数，一个channel

func DumpNum() {
	ch := make(chan int)

	go func() {
		for i:=1;i<=100;i++ {
			ch<-1
			if i%2 == 1 {
				fmt.Println("g1---", i)
			}
		}
	}()

	go func() {
		for i:=1;i<=100;i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println("g2---", i)
			}
		}
	}()

	time.Sleep(time.Second)
}


