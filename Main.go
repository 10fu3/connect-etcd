package main

import (
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

func main() {
	fmt.Println("A")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	})
	fmt.Println("B")
	if err != nil {
		panic(err)
	}
	fmt.Println("C")
	
	if err != nil {
		panic(err)
	}
	concurrency.NewSTM(cli, func(stm concurrency.STM) error {
		r := stm.Get("test")
		fmt.Println("F", r)
		stm.Put("test", r+"/test123")
		return nil
	})

	cli.Close()
	fmt.Println("D")
}
