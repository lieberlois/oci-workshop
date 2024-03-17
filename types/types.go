package types

import "io"

type DecoderFunc func(io.Reader) io.Reader
