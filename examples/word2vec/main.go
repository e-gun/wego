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

package main

import (
	"os"

	"github.com/e-gun/wego/pkg/model/modelutil/vector"
	"github.com/e-gun/wego/pkg/model/word2vec"
)

func main() {
	model, err := word2vec.New(
		word2vec.Window(5),
		word2vec.Model(word2vec.Cbow),
		word2vec.Optimizer(word2vec.NegativeSampling),
		word2vec.NegativeSampleSize(5),
		word2vec.Verbose(),
	)
	if err != nil {
		// failed to create word2vec.
	}

	input, _ := os.Open("text8")
	defer input.Close()
	if err = model.Train(input); err != nil {
		// failed to train.
	}

	// write word vector.
	model.Save(os.Stdin, vector.Agg)
}
