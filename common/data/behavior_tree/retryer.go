package behaviortree

import (
	"time"
)

func Retryer(duration, frequency time.Duration, tick Tick) Tick {
	if tick == nil {
		return nil
	}

	if duration <= 0 {
		return nil
	}

	if frequency <= 0 {
		return nil
	}

	ticker := time.NewTicker(frequency)
	stopwatch := time.NewTicker(duration)

	return func(children []Node) (Status, error) {
		var err error
		var status Status

	RetryerLoop:
		for err == nil {
			select {
			case <-ticker.C:
				status, err = tick(children)
				if status == Success {
					break RetryerLoop
				}
			case <-stopwatch.C:
				status = Failure
				break RetryerLoop
			}
		}

		return status, err
	}
}
