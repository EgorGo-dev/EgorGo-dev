package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func Retry(operation func() error, maxRetries int, baseDelay int) error {
	var err error
	for i := 0; i <= maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		if i == maxRetries {
			break
		}
		delay := time.Duration(baseDelay*(1<<i)) * time.Millisecond
		time.Sleep(delay)
	}
	return fmt.Errorf("исчерпаны все попытки, последняя ошибка: %w", err)
}

func Timeout(operation func() error, timeout int) error {
	ch := make(chan error, 1)
	go func() {
		ch <- operation()
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(time.Duration(timeout) * time.Millisecond):
		return errors.New("превышено время ожидания (timeout)")
	}
}

type DeadLetterQueue struct {
	mu       sync.Mutex
	messages []string
}

func NewDeadLetterQueue() *DeadLetterQueue {
	return &DeadLetterQueue{
		messages: make([]string, 0),
	}
}

func (dlq *DeadLetterQueue) Add(msg string) {
	dlq.mu.Lock()
	defer dlq.mu.Unlock()
	dlq.messages = append(dlq.messages, msg)
}

func (dlq *DeadLetterQueue) GetMessages() []string {
	dlq.mu.Lock()
	defer dlq.mu.Unlock()
	res := make([]string, len(dlq.messages))
	copy(res, dlq.messages)
	return res
}

func ProcessWithDLQ(messages []string, processFunc func(msg string) error, dlq *DeadLetterQueue) {
	for _, msg := range messages {
		if err := processFunc(msg); err != nil {
			dlq.Add(msg)
		}
	}
}

func main() {}
