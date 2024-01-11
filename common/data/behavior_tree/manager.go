/*
   Copyright 2021 Joseph Cumines

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package behaviortree

import (
	"errors"
	"sync"

	"github.com/maladroitthief/entree/common/data"
)

type (
	Manager interface {
		Ticker
		Add(Ticker) error
	}

	manager struct {
		mutex   sync.RWMutex
		once    sync.Once
		worker  data.Worker
		done    chan struct{}
		stop    chan struct{}
		tickers chan managerTicker
		errs    []error
	}

	managerTicker struct {
		Ticker Ticker
		Done   func()
	}

	errManagerTicker  []error
	errManagerStopped struct{ error }
)

var (
	ErrManagerStopped error = errManagerStopped{error: errors.New("manager is stopped")}
)

func NewManager() Manager {
	return &manager{
		done:    make(chan struct{}),
		stop:    make(chan struct{}),
		tickers: make(chan managerTicker),
	}
}

func (m *manager) Done() <-chan struct{} {
	return m.done
}

func (m *manager) Stop() {
	m.once.Do(
		func() {
			close(m.stop)
			m.start()()
		},
	)
}

func (m *manager) Err() error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if len(m.errs) != 0 {
		return errManagerTicker(m.errs)
	}
	return nil
}

func (m *manager) Add(ticker Ticker) error {
	if ticker == nil {
		return errors.New("manager add nil ticker")
	}

	done := m.start()
	select {
	case <-m.stop:
	default:
		select {
		case <-m.stop:
		case m.tickers <- managerTicker{Ticker: ticker, Done: done}:
			return nil
		}
	}
	done()

	err := m.Err()
	if err != nil {
		return errManagerStopped{error: err}
	}
	return ErrManagerStopped
}

func (m *manager) start() (done func()) {
	return m.worker.Run(m.run)
}

func (m *manager) run(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			select {
			case <-m.stop:
				select {
				case <-m.done:
				default:
					close(m.done)
				}
			default:
			}
			return
		case t := <-m.tickers:
			go m.handle(t)
		}
	}
}

func (m *manager) handle(t managerTicker) {
	select {
	case <-t.Ticker.Done():
		t.Ticker.Stop()
	case <-m.stop:
		t.Ticker.Stop()
		<-t.Ticker.Done()
	}

	err := t.Ticker.Err()
	if err != nil {
		m.mutex.Lock()
		m.errs = append(m.errs, err)
		m.mutex.Unlock()
		m.Stop()
	}
	t.Done()
}

func (e errManagerTicker) Error() string {
	var b []byte
	for i, err := range e {
		if i != 0 {
			b = append(b, ' ', '|', ' ')
		}
		b = append(b, err.Error()...)
	}

	return string(b)
}

func (e errManagerTicker) Is(target error) bool {
	for _, err := range e {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}

func (e errManagerStopped) Unwrap() error {
	return e.error
}

func (e errManagerStopped) Is(target error) bool {
	switch target.(type) {
	case errManagerStopped:
		return true
	default:
		return false
	}
}
