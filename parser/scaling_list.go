package parser

import (
	"fmt"

	"github.com/ibbbpbbbp/H264analyzer/bits"
	"github.com/ibbbpbbbp/gobits"
)

func parseScalingList(bs *gobits.BitStream, sizeOfScalingList int, prefix string) {
	lastScale := 8
	nextScale := 8
	for i := 0; i < sizeOfScalingList; i++ {
		if nextScale != 0 {
			delta_scale := bits.SE(bs)
			fmt.Printf("%sdelta_scale: %d\n", prefix, delta_scale)
			nextScale = (lastScale + delta_scale + 256) % 256
		}
		if nextScale != 0 {
			lastScale = nextScale
		}
	}
}
