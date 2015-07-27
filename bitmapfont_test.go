// Copyright 2015 Jacques Supcik, HEIA-FR
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

package bitmapfont

import (
	"fmt"
	"testing"
)

func TestBasic8(t *testing.T) {
	fmt.Println(getCharMSB(65, F88))
	fmt.Println(getCharLSB(65, F88))
	fmt.Println(getCharMLSB(65, F88))
	fmt.Println(getCharLMSB(65, F88))
}

func TestBasic6(t *testing.T) {
	fmt.Println(getCharMSB(65, F68))
	fmt.Println(getCharLSB(65, F68))
	fmt.Println(getCharMLSB(65, F68))
	fmt.Println(getCharLMSB(65, F68))
}
