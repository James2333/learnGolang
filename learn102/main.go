package main

import "fmt"

type Job interface {
	Do()
}

type Worker struct {
	JobQueue chan Job
}

//创建工作队列，用于阻塞接job
func NewWorker() *Worker {
	return &Worker{JobQueue: make(chan Job)}
}

func (w Worker) Run(wq chan chan Job) {
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
			}
		}
	}()
}

type WorkerPool struct {
	workerlen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		workerlen:   workerlen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerlen),
	}
}

func (wp *WorkerPool)Run()  {
	fmt.Println("初始化worker")
	//，创建多个工作队列（channel）
	for i:=1;i<=wp.workerlen;i++{
		w:=NewWorker()
		w.Run(wp.WorkerQueue)
	}
	go func() {
		for {
			select {
			case job:=<-wp.JobQueue:
				work:=<-wp.WorkerQueue
				work<-job
			}
		}
	}()
}