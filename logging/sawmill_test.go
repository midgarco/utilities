package logging

import (
	"fmt"
	"sync"
	"testing"
)

func TestIncludeGlobalFields(t *testing.T) {
	logger := NewLogger()

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			logger.IncludeGlobalFields(Fields{fmt.Sprintf("%d", i): i})
			logger.WithFields(Fields{"with": "fields"}).Info("Test1-", i)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			logger.IncludeGlobalFields(Fields{fmt.Sprintf("%d", i): i})
			logger.WithFields(Fields{"with": "fields"}).Info("Test2-", i)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			logger.IncludeGlobalFields(Fields{fmt.Sprintf("%d", i): i})
			logger.WithFields(Fields{"with": "fields"}).Info("Test3-", i)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			logger.IncludeGlobalFields(Fields{fmt.Sprintf("%d", i): i})
			logger.WithFields(Fields{"with": "fields"}).Info("Test4-", i)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			logger.IncludeGlobalFields(Fields{fmt.Sprintf("%d", i): i})
			logger.WithFields(Fields{"with": "fields"}).Info("Test5-", i)
		}(i)
	}

	logger.IncludeGlobalFields(Fields{fmt.Sprintf("%d", 0): 0})
	logger.WithFields(Fields{"with": "fields"}).Info("TestX-", 0)

	wg.Wait()
}
