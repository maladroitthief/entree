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

import "errors"

const (
	_ Status = iota
	Running
	Success
	Failure
)

type (
	Status int
	Node   func() (Tick, []Node)
	Tick   func(children []Node) (Status, error)
)

func New(tick Tick, children ...Node) Node {
	return nodeFactory(tick, children)
}

var nodeFactory = func(tick Tick, children []Node) (node Node) {
	node = func() (Tick, []Node) {
		return tick, children
	}
	return
}

func (n Node) Tick() (Status, error) {
	if n == nil {
		return Failure, errors.New("behavior tree cannot tick a nil node")
	}

	tick, children := n()
	if tick == nil {
		return Failure, errors.New("behavior tree cannot tick a node with a nil tick")
	}

	return tick(children)
}

func (s Status) Status() Status {
	switch s {
	case Running:
		return Running
	case Success:
		return Success
	default:
		return Failure
	}
}
