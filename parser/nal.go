package parser

import (
	"fmt"

	"github.com/ibbbpbbbp/H264analyzer/bits"
	"github.com/ibbbpbbbp/gobits"
)

func findNAL(ba gobits.ByteAccessor, pos int64) (*nalindex, int64) {
	index := nalindex{
		offset: 0,
		size:   0,
	}
	for i := pos; ; {
		b2, ok := ba.At(i + 2)
		if !ok {
			index.size = ba.Length() - index.offset
			break
		}
		if b2 > 1 {
			i += 3
		} else if b2 == 1 {
			b1, _ := ba.At(i + 1)
			b0, _ := ba.At(i)
			if b1 == 0 && b0 == 0 {
				// start code is 0x000001
				start := i
				if start > 0 {
					bs, _ := ba.At(start - 1)
					if bs == 0 {
						// start code is 0x00000001
						start--
					}
				}
				if index.offset > 0 {
					index.size = start - index.offset
					return &index, i
				}
				index.offset = i + 3
			}
			i += 3
		} else {
			i++
		}
	}

	return &index, 0
}

func parseNAL(h264 *H264Stream, rbsp []byte, prefix string) {
	h264.nal = &nal{}
	bs := gobits.NewBitStream(gobits.NewSliceByteAccessor(rbsp))

	h264.nal.forbidden_zero_bit = bits.U(bs, 1)
	fmt.Printf("%snal.forbidden_zero_bit: %d\n", prefix, h264.nal.forbidden_zero_bit)
	h264.nal.nal_ref_idc = bits.U(bs, 2)
	fmt.Printf("%snal.nal_ref_idc: %d\n", prefix, h264.nal.nal_ref_idc)
	h264.nal.nal_unit_type = bits.U(bs, 5)
	fmt.Printf("%snal.nal_unit_type: %d\n", prefix, h264.nal.nal_unit_type)

	switch h264.nal.nal_unit_type {
	case naluTypeOfSlice:
		fallthrough
	case naluTypeOfIDR:
		fallthrough
	case naluTypeOfAUX:
		parseSliceHeader(h264, bs, prefix)
	case naluTypeOfSEI:
	case naluTypeOfSPS:
		parseSPS(h264, bs, prefix)
	case naluTypeOfPPS:
		parsePPS(h264, bs, prefix)
	case naluTypeOfAUD:
		parseAUD(bs, prefix)
	case naluTypeOfEndOfSequence:
	case naluTypeOfEndOfStream:
	case naluTypeOfFiller:
	default:
	}
}
