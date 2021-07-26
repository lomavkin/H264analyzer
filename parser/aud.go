package parser

import (
	"fmt"

	"github.com/ibbbpbbbp/H264analyzer/bits"
	"github.com/ibbbpbbbp/gobits"
)

func parseAUD(bs *gobits.BitStream, prefix string) {
	prefix = prefix + "aud."

	primary_pic_type := bits.U(bs, 3)
	fmt.Printf("%sprimary_pic_type: %d\n", prefix, primary_pic_type)
}
