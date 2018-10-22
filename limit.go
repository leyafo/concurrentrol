package concurrentrol

import (
	"sync"
)

//Process is concurrent function.
type Process func(i int) error

//Run the concurrent tasks.
func Run(maxJobs, taskNum int, process Process) error {
	concurrentChan := make(chan error, maxJobs) //initialize the concurrent channel for task result
	for i := 0; i != maxJobs; i++ {
		concurrentChan <- nil
	}

	wg := &sync.WaitGroup{}
	var err error
	for i := 0; i != taskNum; i++ {
		middleErr := <-concurrentChan //limited the maxJobs for concurrent running
		if middleErr != nil {
			err = middleErr
			break //If error happend, break all tasks
		}
		wg.Add(1)
		go func(taskNumber int) {
			defer wg.Done()
			concurrentChan <- process(taskNumber) //When a task has finished, back the result.
		}(i)
	}
	wg.Wait()       //waitting for all tasks finished.
	if err == nil { //check the remain task results
	loopCheck:
		for {
			select {
			case e := <-concurrentChan:
				err = e
				if err != nil {
					break loopCheck
				}
			default:
				break loopCheck
			}
		}
	}
	close(concurrentChan)
	if err != nil {
		return err
	}
	return nil
}
