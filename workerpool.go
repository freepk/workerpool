package workerpool

type Task func()

type worker struct {
	queue chan chan Task
	task  chan Task
}

func (w *worker) start() {
	for {
		w.queue <- w.task
		select {
		case task := <-w.task:
			task()
		}
	}
}

func newWorker(queue chan chan Task) *worker {
	return &worker{queue: queue, task: make(chan Task)}
}

type Pool struct {
	queue chan chan Task
	task  chan Task
}

func NewPool(num int) *Pool {
	queue := make(chan chan Task, num)
	for i := 0; i < num; i++ {
		w := newWorker(queue)
		go w.start()
	}
	return &Pool{queue: queue, task: make(chan Task)}
}

func (p *Pool) Start() {
	for {
		select {
		case task := <-p.task:
			workerTask := <-p.queue
			workerTask <- task
		}
	}
}

func (p *Pool) Run(task Task) {
	p.task <- task
}

func (p *Pool) Stop() {
}
