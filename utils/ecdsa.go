package utils

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int //Public key's X coordinate
	S *big.Int //
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
