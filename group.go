package parallelizer

import (
	"errors"
	"sync"
)

const (
	nilFunctionError = "nil function"
)

type Job func(workerId int)

// NewGroup create a new group of workers
func NewGroup(options ...GroupOption) *Group {
	groupOptions := newGroupOptions(options...)

	group := &Group{
		jobsChannel: make(chan Job, groupOptions.JobQueueSize),
		waitGroup:   &sync.WaitGroup{},
	}

	for i := 0; i < groupOptions.PoolSize; i++ {
		go group.worker(i)
	}
	return group
}

// Group a group of workers executing functions concurrently
type Group struct {
	jobsChannel chan Job
	waitGroup   *sync.WaitGroup
}

// Add adds function to queue of jobs to execute
func (g *Group) Add(job Job) error {
	if job == nil {
		return errors.New(nilFunctionError)
	}

	g.waitGroup.Add(1)
	g.jobsChannel <- job
	return nil
}

// Wait waits until workers finished the jobs in the queue
func (g *Group) Wait(options ...WaitOption) error {
	waitOptions := newWaitOptions(options...)

	channel := make(chan bool)
	go func() {
		g.waitGroup.Wait()
		close(channel)
	}()

	select {
	case <-waitOptions.Context.Done():
		return waitOptions.Context.Err()
	case <-channel:
		return nil
	}
}

// Close closes resources
func (g *Group) Close() {
	close(g.jobsChannel)
}

func (g *Group) worker(id int) {
	for job := range g.jobsChannel {
		job(id)
		g.waitGroup.Done()
	}
}
