// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

package tokenizer

// Token represents a piece of text extracted from a larger input.  It
// holds information regarding its beginning and ending offsets in the
// input text.  A token may span the entire input.
type Token interface {
	Text() string
	Begin() int
	End() int
	Type() TokenType
}
