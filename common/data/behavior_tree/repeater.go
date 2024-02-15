package behaviortree

import (
	"time"
)

func Repeater(duration, frequency time.Duration, tick Tick) Tick {
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
	RepeaterLoop:
		for err == nil {
			select {
			case <-ticker.C:
				status, err = tick(children)
			case <-stopwatch.C:
				break RepeaterLoop
			}
		}

		if err != nil {
			return Failure, err
		}

		switch status {
		case Running:
			return Running, nil
		case Success:
			return Success, nil
		default:
			return Failure, nil
		}
	}
}
