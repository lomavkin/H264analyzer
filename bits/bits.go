package bits

import "github.com/ibbbpbbbp/gobits"

func U(bs *gobits.BitStream, bitCount byte) int {
	b, _ := bs.ReadBits(bitCount)
	return int(b)
}

func UE(bs *gobits.BitStream) int {
	b, _ := bs.ReadExponentialGolomb()
	return int(b)
}

func SE(bs *gobits.BitStream) int {
	b, _ := bs.ReadSignedExponentialGolomb()
	return int(b)
}

func RBSP(data []byte) []byte {
	length := len(data)
	rbsp := make([]byte, 0, length)
	for i := 0; i < length; {
		if length-i >= 3 && data[i] == 0 && data[i+1] == 0 && data[i+2] == 3 {
			rbsp = append(rbsp, data[i])
			rbsp = append(rbsp, data[i+1])
			i += 3 // skip emulation_prevention_three_byte
		} else {
			rbsp = append(rbsp, data[i])
			i++
		}
	}
	return rbsp
}

func MoreRBSP(bs *gobits.BitStream) bool {
	pos := bs.SavePos()
	defer bs.RestorePos(pos)

	b, ok := bs.ReadBits(1)
	if !ok {
		return false
	}
	if b == 0 {
		return true
	}

	for ok {
		b, ok = bs.ReadBits(1)
		if b == 1 {
			return true
		}
	}

	return false
}

func CeilLog2(val int) int {
	if val < 0 {
		return 0
	}
	log := 0
	for ; val > 0; val >>= 1 {
		log++
	}
	if log > 0 && val&(val-1) == 0 {
		log--
	}
	return log
}
