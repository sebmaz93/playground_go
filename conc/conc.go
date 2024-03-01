package conc

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", phrase)
	doneChan <- true
	close(doneChan)
}

func Main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	go greet("nice to meet you", done)
	go greet("huh ?", done)
	go slowGreet("Are you there ???", done)
	go greet("......", done)
	<-done

	for range done {
		//fmt.Println(doneChan)
	}

	// you can use select to pick the first channel that sends back result (done, err)
	items := []int{1, 2, 3, 4}
	doneChans := make([]chan bool, len(items))
	errorChans := make([]chan error, len(items))

	for index := range items {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("done!")
		}
	}
	// defer will call the desired thing once the surrounding function ends the execution

}
