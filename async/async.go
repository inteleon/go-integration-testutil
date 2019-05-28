package async

import (
	"time"
)

// AwaitSeconds will try to execute the passed function once per second, either until it returns a nil error or the
// timeoutSeconds has passed.
func AwaitSeconds(until func() error, timeoutSeconds time.Duration) error {
	stop := time.After(timeoutSeconds)
	var localerror error
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-stop:
			return localerror
		default:
			localerror = until()
			if localerror == nil {
				return nil
			}
		}
	}
}
