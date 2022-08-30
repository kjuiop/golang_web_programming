package practice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
	"time"
)

// golang 학습 테스트
func TestGolang(t *testing.T) {
	t.Run("string test", func(t *testing.T) {
		str := "Ann,Jenny,Tom,Zico"
		actual := strings.Split(str, ",")
		expected := []string{"Ann", "Jenny", "Tom", "Zico"}
		assert.Equal(t, expected, actual)
	})

	t.Run("goroutine에서 slice에 값 추가해보기", func(t *testing.T) {
		var numbers []int
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func(i int) {
				defer wg.Done()

				mu.Lock()
				numbers = append(numbers, i)
				mu.Unlock()
			}(i)
		}
		wg.Wait()

		var expected []int // actual : [0 1 2 ... 99]
		for i := 0; i < 100; i++ {
			expected = append(expected, i)
		}
		assert.ElementsMatch(t, expected, numbers)
	})

	t.Run("fan out, fan in", func(t *testing.T) {

		inputCh := generate()
		outputCh := make(chan int)
		go func() {
			defer close(outputCh)
			for {
				select {
				case value, ok := <-inputCh:
					if !ok {
						return
					}
					outputCh <- value * 10
				}
			}
		}()

		var actual []int
		for value := range outputCh {
			actual = append(actual, value)
		}
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, actual)
	})

	t.Run("context timeout", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx := context.TODO()
		timeoutCtx, _ := context.WithTimeout(ctx, add)

		var endTime time.Time
		select {
		case <-timeoutCtx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context deadline", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx := context.TODO() // TODO 3초후에 종료하는 timeout context로 만들어주세요.
		deadlineCtx, _ := context.WithDeadline(ctx, time.Now().Add(add))

		var endTime time.Time
		select {
		case <-deadlineCtx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context value", func(t *testing.T) {

		ctx := context.Background()
		ctx = context.WithValue(ctx, "key", "value")
		assert.Equal(t, "value", ctx.Value("key"))
		assert.Nil(t, ctx.Value("key1"))
	})
}

func generate() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	return ch
}
