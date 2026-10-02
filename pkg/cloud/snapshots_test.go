//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package cloud

import (
	"errors"
	"testing"
)

func TestIsSnapshotGoneError(t *testing.T) {
	cases := []struct {
		err  error
		want bool
	}{
		{nil, false},
		{errors.New(`CloudStack API error 431 (CSExceptionErrorCode: 4350): Invalid parameter id`), true},
		{errors.New(`Undefined error: {"errorcode":431,"errortext":"Snapshot [Snapshot {\"id\":21,\"state\":\"Destroyed\"}] is already destroyed"}`), true},
		{errors.New(`CloudStack API error 530: Failed to delete snapshot`), false},
	}
	for _, c := range cases {
		if got := isSnapshotGoneError(c.err); got != c.want {
			t.Errorf("isSnapshotGoneError(%v) = %v, want %v", c.err, got, c.want)
		}
	}
}
