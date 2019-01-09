package main

import (
	"bufio"
	"context"
	"fmt"
	"kafka-go"
	"os"
	"strings"
	"time"
)

func main()  {


	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		conn := producer("t",0)
		conn.SetWriteDeadline(time.Now().Add(time.Second))
		conn.WriteMessages(
			kafka.Message{Value: []byte(text)})
		conn.Close()
	}

}


func producer(topic string, partition int) (*kafka.Conn){
	 conn , _ := kafka.DialLeader(context.Background(), "tcp", "192.168.45.11:9092", topic, partition)
	 return conn
}