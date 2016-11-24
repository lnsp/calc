package lib

import (
	"errors"
	"fmt"
)

var EmptyQueueError = errors.New("Queue is empty.")

type OperatorQueue []string

func (queue *OperatorQueue) Append(item string) {
	*queue = append(*queue, item)
}

func (queue *OperatorQueue) Poll() string {
	if len(*queue) < 1 {
		return ""
	}

	value := (*queue)[0]
	*queue = (*queue)[1:]
	return value
}

func (queue *OperatorQueue) IsEmpty() bool {
	return len(*queue) < 1
}

func (queue *OperatorQueue) String() string {
	return fmt.Sprint("Queue", *queue)
}

func NewQueue() OperatorQueue {
	return make(OperatorQueue, 0)
}
