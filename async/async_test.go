package async

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPassesAfterThreeSeconds(t *testing.T) {
	err := AwaitSeconds(UntilThreeSecondsHavePassed(), time.Second*5)
	assert.NoError(t, err)
}

func TestFailsAfterTwoSeconds(t *testing.T) {

	err := AwaitSeconds(UntilThreeSecondsHavePassed(), time.Second*2)
	assert.Error(t, err)
}

func UntilThreeSecondsHavePassed() func() error {
	end := time.Now().Add(time.Second * 3)
	return func() error {
		if time.Now().After(end) {
			return nil
		}
		return fmt.Errorf("three seconds have not passed yet")
	}
}
