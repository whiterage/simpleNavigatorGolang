package stack

import (
	"errors"
	"testing"
)

func TestStack_PushAndTop_Correct(t *testing.T) {
	s := NewStack()

	// Тест 1: Push и проверка Top
	testValues := []int{10, 20, 30, 40, 50}

	for _, val := range testValues {
		s.Push(val)
		top, err := s.Top()
		if err != nil {
			t.Errorf("Unexpected error on Top(): %v", err)
		}
		if top != val {
			t.Errorf("Top() = %d, want %d", top, val)
		}
	}

	// Проверка размера
	if s.Size() != len(testValues) {
		t.Errorf("Size() = %d, want %d", s.Size(), len(testValues))
	}
}

func TestStack_Pop_Correct(t *testing.T) {
	s := NewStack()

	// Добавляем элементы
	s.Push(100)
	s.Push(200)
	s.Push(300)

	// Проверяем порядок извлечения (LIFO)
	expected := []int{300, 200, 100}

	for i, exp := range expected {
		val, err := s.Pop()
		if err != nil {
			t.Errorf("Unexpected error on Pop() #%d: %v", i+1, err)
		}
		if val != exp {
			t.Errorf("Pop() #%d = %d, want %d", i+1, val, exp)
		}
	}

	// После всех извлечений стек должен быть пустым
	if !s.IsEmpty() {
		t.Errorf("Stack should be empty after all pops, but Size() = %d", s.Size())
	}
}

func TestStack_Top_Correct(t *testing.T) {
	s := NewStack()

	// Добавляем элементы и проверяем верхний
	s.Push(5)
	top, err := s.Top()
	if err != nil {
		t.Errorf("Unexpected error on Top(): %v", err)
	}
	if top != 5 {
		t.Errorf("Top() = %d, want 5", top)
	}

	s.Push(15)
	top, err = s.Top()
	if err != nil {
		t.Errorf("Unexpected error on Top(): %v", err)
	}
	if top != 15 {
		t.Errorf("Top() = %d, want 15", top)
	}

	s.Push(25)
	top, err = s.Top()
	if err != nil {
		t.Errorf("Unexpected error on Top(): %v", err)
	}
	if top != 25 {
		t.Errorf("Top() = %d, want 25", top)
	}
}

// НЕКОРРЕКТНЫЕ ТЕСТЫ (ОШИБКИ)

func TestStack_Pop_EmptyStack(t *testing.T) {
	s := NewStack()

	// Попытка извлечь из пустого стека
	val, err := s.Pop()

	if err == nil {
		t.Error("Expected error on Pop() from empty stack, got nil")
	}

	if !errors.Is(err, ErrEmptyStack) {
		t.Errorf("Expected ErrEmptyStack, got %v", err)
	}

	if val != 0 {
		t.Errorf("Expected zero value on error, got %d", val)
	}
}

func TestStack_Top_EmptyStack(t *testing.T) {
	s := NewStack()

	// Попытка получить верхний элемент из пустого стека
	val, err := s.Top()

	if err == nil {
		t.Error("Expected error on Top() from empty stack, got nil")
	}

	if !errors.Is(err, ErrEmptyStack) {
		t.Errorf("Expected ErrEmptyStack, got %v", err)
	}

	if val != 0 {
		t.Errorf("Expected zero value on error, got %d", val)
	}
}

func TestStack_MixedOperationsWithErrors(t *testing.T) {
	s := NewStack()

	// 1. Проверяем ошибку на пустом стеке
	_, err := s.Top()
	if !errors.Is(err, ErrEmptyStack) {
		t.Errorf("Expected ErrEmptyStack on Top(), got %v", err)
	}

	// 2. Добавляем элемент
	s.Push(42)

	// 3. Проверяем, что ошибки больше нет
	top, err := s.Top()
	if err != nil {
		t.Errorf("Unexpected error after Push(): %v", err)
	}
	if top != 42 {
		t.Errorf("Top() = %d, want 42", top)
	}

	// 4. Извлекаем элемент
	_, err = s.Pop()
	if err != nil {
		t.Errorf("Unexpected error on Pop(): %v", err)
	}

	// 5. Снова проверяем ошибку на пустом стеке
	_, err = s.Top()
	if !errors.Is(err, ErrEmptyStack) {
		t.Errorf("Expected ErrEmptyStack on Top() after Pop(), got %v", err)
	}

	// 6. Проверяем, что стек действительно пуст
	if !s.IsEmpty() {
		t.Errorf("Stack should be empty after all operations")
	}
}
