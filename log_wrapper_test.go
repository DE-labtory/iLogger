/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package iLogger_test

import (
	"testing"

	"path/filepath"

	"os"

	"github.com/it-chain/iLogger"
	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	iLogger.SetToDebug()
	iLogger.EnableFileLogger(false, "")
	iLogger.Debug(nil, "debug test")
}

func TestError(t *testing.T) {
	iLogger.EnableFileLogger(false, "")
	iLogger.Errorf(&iLogger.Fields{"filedTest": "good"}, "%s good?", "testing is")
}

func TestEnableFileLogger(t *testing.T) {
	os.RemoveAll("./.tmp")
	absPath, _ := filepath.Abs("./.tmp/tmp.log")
	defer os.RemoveAll("./.tmp")
	err := iLogger.EnableFileLogger(true, absPath)
	assert.NoError(t, err)
	iLogger.Error(nil, "hahaha")
}
