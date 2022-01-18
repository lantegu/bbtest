package scripttool

import (
	scripttool "bbtest/scripttool"
	"context"
	"fmt"
	"testing"
	"time"
)
func TestPatternFilter(t *testing.T) {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	batchData := make([][]string, 0)
	for _, answer := range answers {
		batchData = append(batchData, []string{answer})
	}
	filterData, _ := scripttool.PatternFilter(batchData, scripttool.Percent, 50)
	fmt.Print(filterData,"\n")
	fmt.Print(len(batchData),"\n")
	fmt.Print(len(filterData),"\n")
}

func TestConsumer(t *testing.T) {
	testData := []int{1}
	intStream := make(chan int)

	go func() {
		for i := range intStream {
			fmt.Printf("recieve num:%v\n", i)
			time.Sleep(4*time.Second)
		}
	}()

	for _, data := range testData {
		intStream <- data
		fmt.Printf("send num%v\n",data)
	}
	fmt.Print("not ready")
	close(intStream)
	time.Sleep(3*time.Second)
}

func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	cancel()
	ctx.Done()
}