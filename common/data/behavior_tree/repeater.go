package behaviortree

import (
	"time"
)

func Repeater(duration time.Duration, tick Tick) Tick {
	if tick == nil {
		return nil
	}

	if duration <= 0 {
		return nil
	}

	ticker := time.NewTicker(duration)

	return func(children []Node) (Status, error) {
		var err error
		var status Status
	RepeaterLoop:
		for err == nil {
			select {
			case <-ticker.C:
				break RepeaterLoop
			default:
				status, err = tick(children)
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
