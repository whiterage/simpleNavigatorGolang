package queue

import (
	"errors"
	"testing"
)

func TestQueue_PushAndFront_Correct(t *testing.T) {
	q := NewQueue()

	// Тест 1: Push и проверка Front
	testValues := []int{10, 20, 30, 40, 50}

	for _, val := range testValues {
		q.Push(val)
		front, err := q.Front()
		if err != nil {
			t.Errorf("Unexpected error on Front(): %v", err)
		}
		if front != testValues[0] {
			t.Errorf("Front() = %d, want %d", front, testValues[0])
		}
	}

	// Проверка размера
	if q.Size() != len(testValues) {
		t.Errorf("Size() = %d, want %d", q.Size(), len(testValues))
	}
}

func TestQueue_Pop_Correct(t *testing.T) {
	q := NewQueue()

	// Добавляем элементы
	q.Push(100)
	q.Push(200)
	q.Push(300)

	// Проверяем порядок извлечения (FIFO)
	expected := []int{100, 200, 300}

	for i, exp := range expected {
		val, err := q.Pop()
		if err != nil {
			t.Errorf("Unexpected error on Pop() #%d: %v", i+1, err)
		}
		if val != exp {
			t.Errorf("Pop() #%d = %d, want %d", i+1, val, exp)
		}
	}

	// После всех извлечений очередь должна быть пустой
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty after all pops, but Size() = %d", q.Size())
	}
}

func TestQueue_Back_Correct(t *testing.T) {
	q := NewQueue()

	// Добавляем элементы и проверяем последний
	q.Push(5)
	back, err := q.Back()
	if err != nil {
		t.Errorf("Unexpected error on Back(): %v", err)
	}
	if back != 5 {
		t.Errorf("Back() = %d, want 5", back)
	}

	q.Push(15)
	back, err = q.Back()
	if err != nil {
		t.Errorf("Unexpected error on Back(): %v", err)
	}
	if back != 15 {
		t.Errorf("Back() = %d, want 15", back)
	}

	q.Push(25)
	back, err = q.Back()
	if err != nil {
		t.Errorf("Unexpected error on Back(): %v", err)
	}
	if back != 25 {
		t.Errorf("Back() = %d, want 25", back)
	}

	// Проверяем, что Front остался неизменным
	front, _ := q.Front()
	if front != 5 {
		t.Errorf("Front() = %d, want 5 (should not change after Back())", front)
	}
}

// НЕКОРРЕКТНЫЕ ТЕСТЫ (ОШИБКИ)

func TestQueue_Pop_EmptyQueue(t *testing.T) {
	q := NewQueue()

	// Попытка извлечь из пустой очереди
	val, err := q.Pop()

	if err == nil {
		t.Error("Expected error on Pop() from empty queue, got nil")
	}

	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("Expected ErrEmptyQueue, got %v", err)
	}

	if val != 0 {
		t.Errorf("Expected zero value on error, got %d", val)
	}
}

func TestQueue_Front_EmptyQueue(t *testing.T) {
	q := NewQueue()

	// Попытка получить передний элемент из пустой очереди
	val, err := q.Front()

	if err == nil {
		t.Error("Expected error on Front() from empty queue, got nil")
	}

	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("Expected ErrEmptyQueue, got %v", err)
	}

	if val != 0 {
		t.Errorf("Expected zero value on error, got %d", val)
	}
}

func TestQueue_Back_EmptyQueue(t *testing.T) {
	q := NewQueue()

	// Попытка получить задний элемент из пустой очереди
	val, err := q.Back()

	if err == nil {
		t.Error("Expected error on Back() from empty queue, got nil")
	}

	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("Expected ErrEmptyQueue, got %v", err)
	}

	if val != 0 {
		t.Errorf("Expected zero value on error, got %d", val)
	}
}

func TestQueue_MixedOperationsWithErrors(t *testing.T) {
	q := NewQueue()

	// 1. Проверяем ошибку на пустой очереди
	_, err := q.Front()
	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("Expected ErrEmptyQueue on Front(), got %v", err)
	}

	// 2. Добавляем элемент
	q.Push(42)

	// 3. Проверяем, что ошибки больше нет
	front, err := q.Front()
	if err != nil {
		t.Errorf("Unexpected error after Push(): %v", err)
	}
	if front != 42 {
		t.Errorf("Front() = %d, want 42", front)
	}

	// 4. Извлекаем элемент
	_, err = q.Pop()
	if err != nil {
		t.Errorf("Unexpected error on Pop(): %v", err)
	}

	// 5. Снова проверяем ошибку на пустой очереди
	_, err = q.Back()
	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("Expected ErrEmptyQueue on Back() after Pop(), got %v", err)
	}

	// 6. Проверяем, что очередь действительно пуста
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty after all operations")
	}
}
