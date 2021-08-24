package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
	"time"
)

//const endpoints =[]string{"127.0.0.1:2379"}
//const Err =errors.New("sb")
func main() {
	cli, err := clientv3.New(clientv3.Config{
	Endpoints: []string{"127.0.0.1:2379"},
	DialTimeout: time.Second*5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	defer cli.Close()

	// 创建两个单独的会话用来演示锁竞争
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	// 会话s1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 等待直到会话s1释放了/my-lock/的锁
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")
	//ctx , cancel:= context.WithTimeout(context.Background(),time.Second*3)
	//_,err=cli.Put(ctx,"lyz","sb")
	//if err != nil {
	//	fmt.Printf("set from etcd failed, err:%v\n", err)
	//	return
	//}
	//cancel()
	//// get
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	//resp, err := cli.Get(ctx, "lyz")
	//cancel()
	//if err != nil {
	//	fmt.Printf("get from etcd failed, err:%v\n", err)
	//	return
	//}
	//
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	//}

}

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
