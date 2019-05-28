# go-integration-testutil
Common utilities for integration tests

### Packages

##### env

Provides functions related to reading environment variables from the OS

**GetEnv** usage example:

    animal := GetEnv("MY_ANIMAL", "horses")

##### async

Provides functions related to waiting for a given state to be fulfilled before a timeout is triggered.

**AwaitSeconds** usage example:

    func myFunc() {
        err := AwaitSeconds(UntilThreeSecondsHavePassed(), time.Second*5)
        if err != nil {
            // do something...
        }
    }
    
    // Make-believe function that will return a nil error if invoked more than three seconds since its inception.
    // This function is typically checking for a given state in a DB, an external mock or a message queue.
    func UntilThreeSecondsHavePassed() func() error {
    	end := time.Now().Add(time.Second * 3)
    	return func() error {
    		if time.Now().After(end) {
    			return nil
    		}
    		return fmt.Errorf("three seconds have not passed yet")
    	}
    }
    
# LICENSE
See [LICENSE](LICENSE)