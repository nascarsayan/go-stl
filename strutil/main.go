// Copyright 2020 Sayan Naskar <nascarsayan@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strutil

// Reverse takes a string as input and returns its reverse.
func Reverse(str string) string {
	rev := []rune(str)
	for i := 0; i < len(rev)/2; i++ {
		rev[i], rev[len(rev)-i-1] = rev[len(rev)-i-1], rev[i]
	}
	return string(rev)
}
