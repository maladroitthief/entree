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

func Switch(children []Node) (Status, error) {
	for i := 0; i < len(children); i += 2 {
		if i == len(children)-1 {
			// default case
			return children[i].Tick()
		}

		status, err := children[i].Tick()
		if err != nil {
			return Failure, err
		}

		if status == Running {
			return Running, nil
		}

		if status == Success {
			return children[i+1].Tick()
		}

	}

	return Success, nil
}
