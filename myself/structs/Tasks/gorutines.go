package Tasks

import (
	"fmt"
	"structs/mathematics/Operations"
	"sync"
)

// ExecuteWaitGroup /*This implementation proceed with a parallel sum by using gorutines and using a wait group*/
func ExecuteWaitGroup(numbers ...int) {
	const numberOfParallelTasks = 3

	//The following process use a waitGroup to wait until all the task are completed
	wg := sync.WaitGroup{}
	wg.Add(numberOfParallelTasks)

	//Use a for loop to iterate through the tasks
	for i := 0; i < numberOfParallelTasks; i++ {
		task := i
		//use a gorutine to execute the task
		go func() {
			//Mark the task done before exiting the anonymous function
			defer wg.Done()
			//Execute the task here
			fmt.Println("Executing task ", task)
		}()
	}

	//wait until all the tasks are completed
	wg.Wait()
	fmt.Println("All the tasks were completed")

}

// ParallelSum /*The following operation is executed in a parallel enviroment, this operation
//uses mutex to avoid race conditioning over the gorutine, and locks and unlock the shared operation*/
func ParallelSum(v []int) int {

	wg := sync.WaitGroup{}
	wg.Add(1)

	totalSum := 0

	ch := make(chan int, len(v))

	go func() {

		defer wg.Done()

		for _, number := range v {
			ch <- number
		}

		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		totalSum += Operations.Sum(ch)
	}()

	//Wait for the operation to complete
	wg.Wait()

	return totalSum
}
