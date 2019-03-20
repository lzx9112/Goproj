package scheduler

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
