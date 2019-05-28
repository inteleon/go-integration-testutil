package async

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPassesAfterTwoSeconds(t *testing.T) {
	err := AwaitSeconds(Until200MillisHavePassed(), time.Millisecond*300, time.Millisecond*50)
	assert.NoError(t, err)
}

func TestFailsAfter200MillisSeconds(t *testing.T) {
	err := AwaitSeconds(Until200MillisHavePassed(), time.Millisecond*100, time.Millisecond*50)
	assert.Error(t, err)
}

func Until200MillisHavePassed() func() error {
	end := time.Now().Add(time.Millisecond * 200)
	return func() error {
		if time.Now().After(end) {
			return nil
		}
		return fmt.Errorf("200 ms have not passed yet")
	}
}
