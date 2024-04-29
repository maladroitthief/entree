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

func Background(tick func() Tick) Tick {
	if tick == nil {
		return nil
	}

	var nodes []Node
	return func(children []Node) (Status, error) {
		for i, node := range nodes {
			status, err := node.Tick()
			if err == nil && status == Running {
				continue
			}

			// TODO replace with linked list
			copy(nodes[i:], nodes[i+1:])
			nodes[len(nodes)-1] = nil
			nodes = nodes[:len(nodes)-1]

			return status, err
		}

		node := New(tick(), children...)
		status, err := node.Tick()
		if err != nil || status != Running {
			return status, err
		}

		nodes = append(nodes, node)
		return Running, nil
	}
}
