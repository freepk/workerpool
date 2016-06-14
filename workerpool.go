package workerpool

type Task interface {
	Run()
}

type Queue chan Task

type worker struct {
	p chan Queue
	q Queue
}

func newWorker(p chan Queue) *worker {
	return &worker{p: p, q: make(Queue)}
}

func (w *worker) run(t Task) {
	t.Run()
}

func (w *worker) start() {
	for {
		w.p <- w.q
		select {
		case t := <-w.q:
			w.run(t)
		}
	}
}

type Pool struct {
	p chan Queue
	q Queue
}

func NewPool(n int) *Pool {
	p := make(chan Queue, n)
	q := make(Queue)
	for i := 0; i < n; i++ {
		w := newWorker(p)
		go w.start()
	}
	return &Pool{p: p, q: q}
}

func (p *Pool) Start() {
	for {
		select {
		case t := <-p.q:
			q := <-p.p
			q <- t
		}
	}
}

func (p *Pool) Stop() {
}

func (p *Pool) Run(t Task) {
	p.q <- t
}
