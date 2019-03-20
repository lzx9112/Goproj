package scheduler

import (
	"sync"
	"time"
)

type scheduler struct {
	tasks    []*Task
	duration time.Duration
	lock     sync.Mutex
	stop     chan struct{}
}

func NewScheduler() *scheduler {
	return &scheduler{
		tasks:    make([]*Task, 0),
		duration: time.Second,
		stop:     make(chan struct{}),
	}
}

func (scheduler *scheduler) SetDuration(duration time.Duration) {
	scheduler.duration = duration
}

func (scheduler *scheduler) Duration() time.Duration {
	return scheduler.duration
}

func (scheduler *scheduler) CreateTask(f func()) *Task {
	return &Task{
		job:     JobCallback(f),
		addTime: time.Now().UnixNano(),
	}
}

func (scheduler *scheduler) AddTask(task *Task) {
	// 暂时不考虑并发
	scheduler.lockQ()
	defer scheduler.unLockQ()
	scheduler.tasks = append(scheduler.tasks, task)
}

func (scheduler *scheduler) lockQ() {
	scheduler.lock.Lock()
}

func (scheduler *scheduler) unLockQ() {
	scheduler.lock.Unlock()
}

func (scheduler *scheduler) getTask() *Task {
	scheduler.lockQ()
	defer scheduler.unLockQ()

	length := len(scheduler.tasks)
	if length == 0 {
		return nil
	}
	first := scheduler.tasks[0]
	scheduler.tasks = append(scheduler.tasks[:0], scheduler.tasks[1:]...)
	return first
}

func (scheduler *scheduler) Run() {
	for {
	restart:
		task := scheduler.getTask()
		if task == nil || task.AddTime() <= 0 {
			time.Sleep(scheduler.Duration())
			continue
		}
		timer := time.NewTimer(scheduler.Duration())
		for {
			select {
			case <-timer.C:
				go task.Job().Run()
				timer.Stop()
				goto restart
			case <-scheduler.stop:
				timer.Stop()
				return
			}
		}
	}
}

func (scheduler *scheduler) Stop() {
	scheduler.stop <- struct{}{}
}
