package workerpool

import (
	"context"
	log "route256/libs/logger"
	"sync"
)

// Worker контролирует выполнение задач
type Worker[Input any, Output any] struct {
	// идентификатор воркера для логгирования
	id int
	// Канал получения задач на исполнение
	taskChan <-chan *Task[Input, Output]
	// Канал отправки результатов исполнения
	resultsChan chan<- *Task[Input, Output]
	// Действие исполняемое над задачами
	f func(context.Context, Input) (Output, error)
}

// newWorker возвращает новый экземпляр worker-а
func newWorker[Input any, Output any](
	id int,
	f func(context.Context, Input) (Output, error),
	tasksChan <-chan *Task[Input, Output],
	resultsChan chan<- *Task[Input, Output],
) *Worker[Input, Output] {
	return &Worker[Input, Output]{
		id:          id,
		f:           f,
		taskChan:    tasksChan,
		resultsChan: resultsChan,
	}
}

// запуск worker-a
func (wr *Worker[Input, Output]) Start(ctx context.Context, wg *sync.WaitGroup) {
	// Оповещаем wg что запустился воркер
	wg.Add(1)

	go func() {
		// Оповещаем wg что выключился воркер
		defer wg.Done()

		log.Infof("start worker %d", wr.id)
		defer log.Infof("end worker %d", wr.id)

		// Ждем таски из канала
		for {
			select {
			// на случай завершения программы во время работы воркера
			case <-ctx.Done():
				return
			case task, ok := <-wr.taskChan:
				// выключаем воркер при закрытии канала
				if !ok {
					return
				}
				// делаем полезности
				wr.process(ctx, task)
			}
		}
	}()
}

func (wr *Worker[Input, Output]) process(ctx context.Context, task *Task[Input, Output]) {
	// записываем в сущность задачи результаты ее выполнения
	task.Output, task.Err = wr.f(ctx, task.input)

	select {
	// на случай завершения программы во время работы воркера
	case <-ctx.Done():
		return
	case wr.resultsChan <- task:
	}
}
