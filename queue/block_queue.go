package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// 实现高并发的阻塞队列 消费者竞争问题 应该怎么实现
// 有几种情况， 1. 生产者生产速度< 消费者， 使用信号量阻塞 ，sync.Cond ， 或者原子操作

type BlockingQueue struct {
	queue   []interface{} // 队列
	lock    sync.Mutex    // 加锁
	cond    *sync.Cond    // 信号量阻塞
	maxSize int           // 队列最大长度
}

func NewBlockingQueue(maxSize int) *BlockingQueue {
	q := &BlockingQueue{
		queue:   make([]interface{}, 0, maxSize),
		lock:    sync.Mutex{},
		maxSize: maxSize,
	}
	q.cond = sync.NewCond(&q.lock)
	return q
}

// 入队操作
func (q *BlockingQueue) EnQueue(item interface{}) error {
	// 先获取锁锁住
	q.lock.Lock()
	defer q.lock.Unlock()

	// 判断是否能发送数据到queue里面, 满了不能往里面加入数据
	if len(q.queue) >= q.maxSize {
		return errors.New("队列满啦")
	}

	q.queue = append(q.queue, item)
	q.cond.Signal()
	return nil
}

// 出队列操作
func (q *BlockingQueue) DeQueue() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for len(q.queue) == 0 { // 一定是用for循环轮训查询， cond.Wait() 被唤醒之后是从他的下一行进行执行，如果是虚假唤醒比如其他线程调用，或者操作系统用调用终端，影响到goroutine， 会错误的获取数据
		q.cond.Wait() // cond.Wait() 是会进入阻塞阶段，直到被 cond.Signal() (唤醒单个goroutine) 或者cond.Broadcast() (唤醒所有goroutine)，
	}

	data := q.queue[0]
	q.queue = q.queue[1:]
	return data, nil

}

// 消费队列中的数据

func main() {
	bQueue := NewBlockingQueue(10)

	// 单个生产者
	go func() {
		for i := 0; i < 10; i++ {
			err := bQueue.EnQueue(i)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Produced:", i)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		data, err := bQueue.DeQueue()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Consumer:", data)
	}

	time.Sleep(time.Second * 10)
}
