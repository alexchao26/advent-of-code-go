package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int64 {
	bin := parseInput(input)

	versionTotal, _, _ := handlePacket(bin)

	return versionTotal
}

// bitsRead is often just called n in Go. See io.Read() for example
func handlePacket(pack string) (versionTotal int64, expressionValue int, bitsRead int) {
	version, err := strconv.ParseInt(pack[0:3], 2, 64)
	versionTotal += version
	if err != nil {
		panic("version strconv.ParseInt(): " + err.Error())
	}
	typeID, err := strconv.ParseInt(pack[3:6], 2, 64)
	if err != nil {
		panic("typeID strconv.ParseInt(): " + err.Error())
	}

	read := 6 // version and typeID

	switch typeID {
	case 4: // literal value
		// parse 5 bits at a time
		var bits string
		for i := 6; i < len(pack); i += 5 {
			fiveBits := pack[i : i+5]
			read += 5

			bits += fiveBits[1:]

			if fiveBits[0] == '0' {
				break
			}
		}

		decimalVal, err := strconv.ParseInt(bits, 2, 64)
		if err != nil {
			panic("parsing packet bits: " + err.Error())
		}
		return version, int(decimalVal), read
	default: // operator types?
		// contains one or more packets
		lengthTypeID := pack[6:7]
		read++
		var bitsToRead int
		switch lengthTypeID {
		case "0":
			// next 15 bits == total length in bits for REST of subpackets
			bitsToRead = 15
		case "1":
			// next 11 bits == NUMBER of subpackets
			bitsToRead = 11
		}

		rawLength := pack[7 : 7+bitsToRead]
		read += bitsToRead
		length, err := strconv.ParseInt(rawLength, 2, 64)
		if err != nil {
			panic("parsing 0 lengthTypeID: " + err.Error())
		}

		// followed by the subpackets themselves
		var subPacketExpressionValues []int
		switch lengthTypeID {
		case "0":
			// next 15 bits == total length in bits for REST of subpackets
			for length > 0 {
				// continue reading until length of bits are read
				ver, expVal, n := handlePacket(pack[read:])
				read += n
				length -= int64(n)
				subPacketExpressionValues = append(subPacketExpressionValues, expVal)
				versionTotal += ver
			}
		case "1":
			// next 11 bits == NUMBER of subpackets
			for length > 0 {
				// continue reading until number of packets are read
				ver, expVal, n := handlePacket(pack[read:])
				read += n
				length-- // reduce length by 1 (ie one packet read)
				subPacketExpressionValues = append(subPacketExpressionValues, expVal)
				versionTotal += ver
			}
		}

		switch typeID {
		// note: case 0 already handled above, literal value
		case 0: // sum
			return versionTotal, mathy.SumIntSlice(subPacketExpressionValues), read
		case 1: // product
			return versionTotal, mathy.MultiplyIntSlice(subPacketExpressionValues), read
		case 2: // min
			return versionTotal, mathy.MinInt(subPacketExpressionValues...), read
		case 3: // max
			return versionTotal, mathy.MaxInt(subPacketExpressionValues...), read
			// 4 is literal...
		case 5: // greater than (first subpacket > second, will always have exactly 2)
			var ans int
			if subPacketExpressionValues[0] > subPacketExpressionValues[1] {
				ans = 1 // otherwise int zero val works
			}
			return versionTotal, ans, read
		case 6: // less than (opposite)
			var ans int
			if subPacketExpressionValues[0] < subPacketExpressionValues[1] {
				ans = 1 // otherwise int zero val works
			}
			return versionTotal, ans, read
		case 7: // equal to
			var ans int
			if subPacketExpressionValues[0] == subPacketExpressionValues[1] {
				ans = 1 // otherwise int zero val works
			}
			return versionTotal, ans, read
		default:
			panic(fmt.Sprintf("unknown typeID: %d", typeID))
		}
	}
}

func part2(input string) int {
	bin := parseInput(input)

	_, expVal, _ := handlePacket(bin)

	return expVal
}

var hexToBin = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func parseInput(input string) string {
	var binarySb strings.Builder
	for _, char := range strings.Split(input, "") {
		binarySb.WriteString(hexToBin[char])
	}
	return binarySb.String()
}
