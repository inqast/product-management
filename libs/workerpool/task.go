package workerpool

// Task сущность задачи
type Task[Input any, Output any] struct {
	// Входные данные
	input Input
	// Выходные данные
	Output Output
	// Ошибка выполнения
	Err error
}
