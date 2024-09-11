package chan_more_work

import (
	"fmt"
	"sync"
	"time"
)

// 需要执行的任务
type Task struct {
	ID int
}

type TaskQueue chan Task

type Worker struct {
	// 比如从数据库中查询些数据
	ID        int       // 工作流ID
	TaskQueue TaskQueue // 任务队列
	QuitChan  chan bool // 是否退出的信号通道
	WG        *sync.WaitGroup
}

// 创建一个Woker对象
func NewWorker(id int, taskQueue chan Task, quitChan chan bool, wg *sync.WaitGroup) *Worker {
	return &Worker{
		ID:        id,
		TaskQueue: taskQueue,
		QuitChan:  quitChan,
		WG:        wg,
	}
}

func (w *Worker) Start() {
	w.WG.Add(1)
	go func() {
		defer w.WG.Done()
		for {
			select {
			case task := <-w.TaskQueue: // 获取工作流中的人物并且执行
				fmt.Printf("Worker  %d started task %d\n", w.ID, task)
				time.Sleep(time.Second)
				fmt.Printf("Woker %d Finished task %d\n", w.ID, task)

			case <-w.QuitChan: // 关闭该工作流
				return
			}
		}
	}()
}

func AdjustWorker(taskQueue TaskQueue, wg *sync.WaitGroup, minWorker, maxWorker int) {
	var wokers []*Worker
	for i := 0; i < minWorker; i++ {
		quitChan := make(chan bool)
		wokers = append(wokers, NewWorker(i, taskQueue, quitChan, wg))
		wokers[i].Start()
	}

	// 监控队列长度，如果长度变化，调整work数量
	ticker := time.Tick(time.Second * 2)
	for range ticker {
		// 计算出需要的worker数量 比如查询ES发现数据有10w条，目前只启动了 1个goroutine ，需要再启动按照某种函数处理的数据
		targetWorks := len(wokers)
		if len(taskQueue) > 10 { // 队列长度超过10000 增加一个worker处理
			if targetWorks < maxWorker {
				targetWorks++
			}
		} else if len(taskQueue) < 5 { // 队列长度小于10000 减少一个woker处理
			targetWorks--
		}

		// 调整worker数量
		if targetWorks != len(wokers) {
			if targetWorks > len(wokers) {
				// 循环添加worker直到目标数量
				for len(wokers) < targetWorks {
					quitChan := make(chan bool)
					newWorker := NewWorker(len(wokers), taskQueue, quitChan, wg)
					wokers = append(wokers, newWorker)
					newWorker.Start()
				}
			} else {
				// 减少worker ，可能需要执行安全关闭 当前线程 实现的话信号量 ，如何安全的关闭一个goroutine
				// 创建一个信号量， 发送给当前worker 告诉他执行完之后关闭
				worker := wokers[len(wokers)-1]
				worker.QuitChan <- true
				// 将被移除的worker移出出 workers
				wokers = wokers[:len(wokers)-2]
			}
		}

	}
}

func main() {
	var wg *sync.WaitGroup
	taskQueue := make(TaskQueue, 10)

	go AdjustWorker(taskQueue, wg, 2, 10)

	// 往队列里面扔任务
	for i := 0; ; i++ {
		taskQueue <- Task{ID: 1}

		// 如果超出队列停止？ 或者说队列满了等待？
		if i > 20 {
			break
		}
		time.Sleep(500 * time.Second)
	}

	// 关闭所有Worker的QuitChan来优雅地停止它们（注意：这里的逻辑并未在AdjustWorkers中实现）
	// 由于上面的AdjustWorkers示例中没有实现Worker的优雅停止，这里只是注释掉
	// 在实际应用中，你需要设计一种机制来安全地关闭Worker

	// 等待所有Worker完成
	wg.Wait()
	fmt.Println("All tasks processed and workers stopped.")
	// 注意：上面的代码示例中，Worker并没有真正地根据任务队列的长度来停止，
	// 这只是为了展示如何动态地创建Worker。在实际应用中，你需要实现一种机制来
	// 优雅地关闭不再需要的Worker，并确保它们不会再处理新的任务。
	// 这通常涉及到关闭Worker的QuitChan，并在Worker内部检查这个Chan来安全地退出循环。
}
