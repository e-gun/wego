```
NB the vectorbot *can* trigger a panic when calling this code, but it does not *always* do so. Very frustrating...
```
```
[HGS] [AV: 98.384s][Δ: 98.384s] (0.5%) checking need to model Isocrates (gr0010)
[HGS] [AV: 167.781s][Δ: 69.396s] (0.9%) checking need to model Hesiodus (gr0020)
[HGS] [AV: 173.646s][Δ: 5.865s] (1.4%) checking need to model Xenophon (gr0032)
[HGS] [AV: 187.382s][Δ: 13.736s] (1.8%) checking need to model Alexandri Magni Epistulae (gr0042)
[HGS] [AV: 187.461s][Δ: 0.079s] (2.3%) checking need to model Menippus (gr0052)
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x28 pc=0x102e9a674]

goroutine 7537 [running]:
github.com/e-gun/wego/pkg/model/word2vec.(*hierarchicalSoftmax).optim(0x1401ade0000, 0x102964180?, 0x3f9999999999999a, {0x1404027dd80, 0x7d, 0x102931150?}, {0x1401945a400, 0x7d, 0x14020a8bf1f?})
	/Users/erik/Development/go/pkg/mod/github.com/e-gun/wego@v0.0.11/pkg/model/word2vec/optimizer.go:115 +0x84
github.com/e-gun/wego/pkg/model/word2vec.(*skipGram).trainOne(0x14017b30000, {0x140102a7f78, 0x2bef, 0x1?}, 0x8, 0x829de59?, 0x1401b958000, {0x103c06920, 0x1401ade0000})
	/Users/erik/Development/go/pkg/mod/github.com/e-gun/wego@v0.0.11/pkg/model/word2vec/model.go:73 +0x1c4
github.com/e-gun/wego/pkg/model/word2vec.(*word2vec).trainPerThread(0x1401adf4000, {0x140102a7f78?, 0x2bef, 0x40411}, 0x14012f6f7a0?, 0x10303f368?, 0x1401470c000?)
	/Users/erik/Development/go/pkg/mod/github.com/e-gun/wego@v0.0.11/pkg/model/word2vec/word2vec.go:203 +0x160
created by github.com/e-gun/wego/pkg/model/word2vec.(*word2vec).modifiedtrain
	/Users/erik/Development/go/pkg/mod/github.com/e-gun/wego@v0.0.11/pkg/model/word2vec/word2vec.go:379 +0x274

```

```
another:

panic: runtime error: slice bounds out of range [:1] with capacity 0

goroutine 9721 [running]:
github.com/e-gun/wego/pkg/corpus/dictionary/node.(*Node).GetPath(0x140087a2e88?, 0x10249bfb8?)
	/Users/erik/Development/go/pkg/mod/github.com/e-gun/wego@v0.0.11/pkg/corpus/dictionary/node/node.go:42 +0x84
github.com/e-gun/wego/pkg/model/word2vec.(*hierarchicalSoftmax).optim(...)
	/Users/erik/Development/go/pkg/mod/github.com/e-gun/wego@v0.0.11/pkg/model/word2vec/optimizer.go:112
github.com/e-gun/wego/pkg/model/word2vec.(*skipGram).trainOne(0x140086ac480, {0x1400a2fc000, 0x8cb, 0x1?}, 0x1, 0x3f9999999999999a, 0x14003229680, {0x1040a43a0, 0x14003229800?})
	/Users/erik/Development/go/pkg/mod/github.com/e-gun/wego@v0.0.11/pkg/model/word2vec/model.go:73 +0x1f0
github.com/e-gun/wego/pkg/model/word2vec.(*word2vec).trainPerThread(0x14006785d40, {0x1400a2fc000?, 0x8cb, 0xd800}, 0x0?, 0x1025091e4?, 0x0?)
	/Users/erik/Development/go/pkg/mod/github.com/e-gun/wego@v0.0.11/pkg/model/word2vec/word2vec.go:203 +0x164
created by github.com/e-gun/wego/pkg/model/word2vec.(*word2vec).modifiedtrain in goroutine 9717
	/Users/erik/Development/go/pkg/mod/github.com/e-gun/wego@v0.0.11/pkg/model/word2vec/word2vec.go:379 +0x254


```


# Word Embeddings in Go

[![Go](https://github.com/ynqa/wego/actions/workflows/go.yml/badge.svg)](https://github.com/ynqa/wego/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/ynqa/wego?status.svg)](https://godoc.org/github.com/ynqa/wego)
[![Go Report Card](https://goreportcard.com/badge/github.com/ynqa/wego)](https://goreportcard.com/report/github.com/ynqa/wego)

*wego* is the implementations **from scratch** for word embeddings (a.k.a word representation) models in Go.

## What's word embeddings?

[Word embeddings](https://en.wikipedia.org/wiki/Word_embeddings) make words' meaning, structure, and concept mapping into vector space with a low dimension. For representative instance:
```
Vector("King") - Vector("Man") + Vector("Woman") = Vector("Queen")
```
Like this example, the models generate word vectors that could calculate word meaning by arithmetic operations for other vectors.

## Features

The following models to capture the word vectors are supported in *wego*:

- Word2Vec: Distributed Representations of Words and Phrases and their Compositionality [[pdf]](https://papers.nips.cc/paper/5021-distributed-representations-of-words-and-phrases-and-their-compositionality.pdf)

- GloVe: Global Vectors for Word Representation [[pdf]](http://nlp.stanford.edu/pubs/glove.pdf)

- LexVec: Matrix Factorization using Window Sampling and Negative Sampling for Improved Word Representations [[pdf]](http://anthology.aclweb.org/P16-2068)

Also, wego provides nearest neighbor search tools that calculate the distances between word vectors and find the nearest words for the target word. "near" for word vectors means "similar" for words.

Please see the [Usage](#Usage) section if you want to know how to use these for more details.

## Why Go?

Inspired by [Data Science in Go](https://speakerdeck.com/chewxy/data-science-in-go) @chewxy

## Installation

Use `go` command to get this pkg.

```
$ go get -u github.com/ynqa/wego
$ bin/wego -h
```

## Usage

*wego* provides CLI and Go SDK for word embeddings.

### CLI

```
Usage:
  wego [flags]
  wego [command]

Available Commands:
  console     Console to investigate word vectors
  glove       GloVe: Global Vectors for Word Representation
  help        Help about any command
  lexvec      Lexvec: Matrix Factorization using Window Sampling and Negative Sampling for Improved Word Representations
  query       Query similar words
  word2vec    Word2Vec: Continuous Bag-of-Words and Skip-gram model
```

`word2vec`, `glove` and `lexvec` executes the workflow to generate word vectors:
1. Build a dictionary for vocabularies and count word frequencies by scanning a given corpus.
2. Start training. The execution time depends on the size of the corpus, the hyperparameters (flags), and so on.
3. Save the words and their vectors as a text file.

`query` and `console` are the commands which are related to nearest neighbor searching for the trained word vectors.

`query` outputs similar words against a given word using sing word vectors which are generated by the above models.

e.g. `wego query -i word_vector.txt microsoft`:
```
  RANK |   WORD    | SIMILARITY
-------+-----------+-------------
     1 | hypercard |   0.791492
     2 | xp        |   0.768939
     3 | software  |   0.763369
     4 | freebsd   |   0.761084
     5 | unix      |   0.749563
     6 | linux     |   0.747327
     7 | ibm       |   0.742115
     8 | windows   |   0.731136
     9 | desktop   |   0.715790
    10 | linspire  |   0.711171
```

*wego* does not reproduce word vectors between each trial because it adopts HogWild! algorithm which updates the parameters (in this case word vector) async.

`console` is for REPL mode to calculate the basic arithmetic operations (`+` and `-`) for word vectors.

### Go SDK

It can define the hyper parameters for models by functional options.

```go
model, err := word2vec.New(
	word2vec.Window(5),
	word2vec.Model(word2vec.Cbow),
	word2vec.Optimizer(word2vec.NegativeSampling),
	word2vec.NegativeSampleSize(5),
	word2vec.Verbose(),
)
```

The models have some methods:

```go
type Model interface {
	Train(io.ReadSeeker) error
	Save(io.Writer, vector.Type) error
	WordVector(vector.Type) *matrix.Matrix
}
```

### Formats

As training word vectors wego requires the following file formats for inputs/outputs.

#### Input

Input corpus must be subject to the formats to be divided by space between words like [text8](http://mattmahoney.net/dc/textdata.html).

```
word1 word2 word3 ...
```

#### Output

After training *wego* save the word vectors into a txt file with the following format (`N` is the dimension for word vectors you given):

```
<word> <value_1> <value_2> ... <value_N>
```
