// Copyright (c) 2019, Cody Opel <codyopel@gmail.com>
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

// Package base64uid generates URL safe uids.
package base64uid

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// Generates a base64 encoded string of a specfied number of random bytes.
// Arbitrarily limited to uint8.
func New(numBytes uint8) (string, error) {
	if numBytes < 1 {
		return "", fmt.Errorf("numBytes must be greater than or equal to 1")
	}
	randBytes := make([]byte, numBytes)
	if _, err := rand.Read(randBytes[:]); err != nil {
		return "", err
	}
	return string(base64.RawURLEncoding.EncodeToString(randBytes)), nil
}
