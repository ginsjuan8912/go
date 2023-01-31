package Tasks

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"structs/mathematics/Operations"
	"structs/tree"
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

func ToTree(fileName string) (tree.BinaryTree, error) {

	var bTree = tree.BinaryTree{}
	bTree.Initialize(0)

	wg := sync.WaitGroup{}

	var ch = make(chan int, 10)

	//Read the numbers from the file
	wg.Add(1)
	go func() {
		defer wg.Done()

		file, err := os.Open(fileName)

		if err != nil {
			panic("cannot read the requested file")
		}

		defer file.Close()

		reader := bufio.NewReader(file)

		for {
			line, _, err := reader.ReadLine()

			if err == io.EOF {
				break
			}

			for _, s := range strings.Split(string(line), ",") {

				if s == "0 " || s == "\n" {
					continue
				}

				number, convErr := strconv.Atoi(s)

				if convErr != nil {
					fmt.Printf("cannot convert the number %d", number)
				}

				ch <- number
			}

		}
		close(ch)
	}()

	//Add each stream of the channel to the binary tree
	wg.Add(1)
	go func() {

		defer wg.Done()

		for number := range ch {
			err := bTree.Root.Insert(int64(number))

			if err != nil {
				fmt.Printf("Cannot insert %d", number)
				continue
			}
		}

	}()

	wg.Wait()

	return bTree, nil
}
