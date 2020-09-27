package main

import (
	"sync"
)

type FsJobQueue struct {
	Limit int
	JobBuffer []*Job
}

type Job struct {
	Id uint
	Value *[]byte
}

func (jq *FsJobQueue) LimitReached() bool {
	return len(jq.JobBuffer) >= jq.Limit
}

func (jq *FsJobQueue) Add(j *Job) {
	jq.JobBuffer = append(jq.JobBuffer, j)
}

func (jq *FsJobQueue) Consume() {
	wg := &sync.WaitGroup{}
	wg.Add(len(jq.JobBuffer))

	for _, j := range jq.JobBuffer {
		go func(j *Job, wg *sync.WaitGroup) {
			FsWrite(j.Id, j.Value)

			wg.Done()
		}(j, wg)
	}

	jq.JobBuffer = nil
	jq.JobBuffer = []*Job{}

	wg.Wait()
}
