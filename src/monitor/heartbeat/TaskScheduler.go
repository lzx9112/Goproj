package heartbeat

import (
	"sync"
	"time"
)

type Job interface {
	Run()
}

type Task struct {
	job     Job
	addTime int64
	// 留作扩展
}

func (t *Task) AddTime() int64 {
	return t.addTime
}

func (t *Task) SetAddTime(addTime int64) {
	t.addTime = addTime
}

func (t *Task) Job() Job {
	return t.job
}

func (t *Task) SetJob(job Job) {
	t.job = job
}

type JobCallback func()

func (f JobCallback) Run() { f() }

type TaskScheduler struct {
	tasks    []*Task
	duration time.Duration
	lock     sync.Mutex
	stop     chan struct{}
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:    make([]*Task, 0),
		duration: time.Second,
		stop:     make(chan struct{}),
	}
}

func (scheduler *TaskScheduler) SetDuration(duration time.Duration) {
	scheduler.duration = duration
}

func (scheduler *TaskScheduler) Duration() time.Duration {
	return scheduler.duration
}

func (scheduler *TaskScheduler) CreateTask(f func()) *Task {
	return &Task{
		job:     JobCallback(f),
		addTime: time.Now().UnixNano(),
	}
}

func (scheduler *TaskScheduler) AddTask(task *Task) {
	// 暂时不考虑并发
	scheduler.lockQ()
	defer scheduler.unLockQ()
	scheduler.tasks = append(scheduler.tasks, task)
}

func (scheduler *TaskScheduler) lockQ() {
	scheduler.lock.Lock()
}

func (scheduler *TaskScheduler) unLockQ() {
	scheduler.lock.Unlock()
}

func (scheduler *TaskScheduler) getTask() *Task {
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

func (scheduler *TaskScheduler) Run() {
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

func (scheduler *TaskScheduler) Stop() {
	scheduler.stop <- struct{}{}
}
