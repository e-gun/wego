// Copyright © 2020 wego authors
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

package model

import (
	"context"
	"io"

	"github.com/e-gun/wego/pkg/model/modelutil/matrix"
	"github.com/e-gun/wego/pkg/model/modelutil/vector"
)

type Model interface {
	Train(io.ReadSeeker) error
	Save(io.Writer, vector.Type) error
	WordVector(vector.Type) *matrix.Matrix
	Reporter(chan int, chan string)
}

type ModelWithCtx interface {
	Train(io.ReadSeeker) error
	Save(io.Writer, vector.Type) error
	WordVector(vector.Type) *matrix.Matrix
	Reporter(chan int, chan string)
	InsertContext(ctx context.Context)
}
