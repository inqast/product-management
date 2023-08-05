package workerpool

import (
	"context"
	"sync"
)

// Pool воркеров
type Pool[Input any, Output any] struct {
	ctx    context.Context
	cancel context.CancelFunc

	// Список задач на исполнение
	tasks []*Task[Input, Output]
	// Действие производимое над задачами
	f func(context.Context, Input) (Output, error)
	// Количество воркеров
	workerCount int

	// Для отслеживания выключения воркеров
	wg sync.WaitGroup
	// Канал связи с воркерами
	tasksChan chan *Task[Input, Output]
	// Канал c результатами выполнения
	resultsChan chan *Task[Input, Output]
}

// NewPool инициализирует новый пул с заданным действием
func NewPool[Input any, Output any](
	ctx context.Context,
	f func(context.Context, Input) (Output, error),
	workerCount int,
) *Pool[Input, Output] {
	ctx, cancel := context.WithCancel(ctx)

	return &Pool[Input, Output]{
		ctx:         ctx,
		cancel:      cancel,
		f:           f,
		workerCount: workerCount,
		tasksChan:   make(chan *Task[Input, Output]),
		resultsChan: make(chan *Task[Input, Output]),
	}
}

// AddTask Добавляет задачу в очередь на исполнение
func (p *Pool[Input, Output]) AddTask(in Input) {
	p.tasks = append(p.tasks, &Task[Input, Output]{
		input: in,
	})
}

// Run запускает всю работу в Pool и блокирует ее до тех пор,
// пока она не будет закончена.
func (p *Pool[Input, Output]) Run() {
	// Запуск воркеров
	for i := 1; i <= p.workerCount; i++ {
		worker := newWorker(i, p.f, p.tasksChan, p.resultsChan)
		worker.Start(p.ctx, &p.wg)
	}

	go func() {
		// Посылка задач на исполнение
		for _, task := range p.tasks {
			select {
			case <-p.ctx.Done():
				return
			case p.tasksChan <- task:
			}
		}
		// Закрываем канал что бы они вырубились когда закончат работу
		close(p.tasksChan)
	}()

	go func() {
		//Ждем готовых результатов
		p.wg.Wait()

		close(p.resultsChan)
	}()
}

// GetResults средство получения результатов выполнения
func (p *Pool[Input, Output]) GetResults() <-chan *Task[Input, Output] {
	return p.resultsChan
}

func (p *Pool[Input, Output]) Stop() {
	p.cancel()
}
