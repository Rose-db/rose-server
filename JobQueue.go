package main

import (
	"sync"
	"time"
)

type FsJobQueue struct {
	Limit int
	Open int
	Lock *sync.RWMutex
}

type Job struct {
	Id uint
	Value *[]byte
}

func NewJobQueue(limit int) *FsJobQueue {
	return &FsJobQueue{
		Limit: limit,
		Lock: &sync.RWMutex{},
		Open: 0,
	}
}

func (jq *FsJobQueue) Run(j *Job) {
	jq.Lock.Lock()
	jq.Open++
	jq.Lock.Unlock()

	if jq.Open >= jq.Limit {
		time.Sleep(2000 * time.Millisecond)
	}

	FsWrite(j.Id, j.Value)

	jq.Lock.Lock()
	jq.Open--
	jq.Lock.Unlock()
}
