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
	"context"
	"errors"
	"sync"
	"time"
)

type (
	Ticker interface {
		Done() <-chan struct{}
		Err() error
		Stop()
	}

	ticker struct {
		ctx    context.Context
		cancel context.CancelFunc
		node   Node
		ticker *time.Ticker
		done   chan struct{}
		stop   chan struct{}
		once   sync.Once
		mutex  sync.Mutex
		err    error
	}

	tickerStopOnFail struct {
		Ticker
	}
)

var (
	ErrExitOnFailure = errors.New("ticker exit on failure")
)

func NewTicker(ctx context.Context, duration time.Duration, node Node) Ticker {
	if ctx == nil {
		panic(errors.New("NewTicker nil context"))
	}

	if duration <= 0 {
		panic(errors.New("NewTicker duration <= 0"))
	}

	if node == nil {
		panic(errors.New("NewTicker duration nil Node"))
	}

	ticker := &ticker{
		node:   node,
		ticker: time.NewTicker(duration),
		done:   make(chan struct{}),
		stop:   make(chan struct{}),
	}

	ticker.ctx, ticker.cancel = context.WithCancel(ctx)

	go ticker.run()

	return ticker
}

func (t *ticker) run() {
	var err error
TickLoop:
	for err == nil {
		select {
		case <-t.ctx.Done():
			err = t.ctx.Err()
			break TickLoop
		case <-t.stop:
			break TickLoop
		case <-t.ticker.C:
			_, err = t.node.Tick()
		}

	}

	t.mutex.Lock()
	t.err = err
	t.mutex.Unlock()
	t.Stop()
	t.cancel()
	close(t.done)
}

func (t *ticker) Done() <-chan struct{} {
	return t.done
}

func (t *ticker) Err() error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.err
}

func (t *ticker) Stop() {
	t.once.Do(
		func() {
			t.ticker.Stop()
			close(t.stop)
		},
	)
}
