package day16

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
	"strings"
)

const (
	SUM_ID           = iota
	PRODUCT_ID       = iota
	MINIMUM_ID       = iota
	MAXIMUM_ID       = iota
	LITERAL_VALUE_ID = iota
	GREATER_THAN_ID  = iota
	LESS_THAN_ID     = iota
	EQUAL_TO_ID      = iota
)

func getTypeIdDesc(id uint8) string {
	switch id {
	case SUM_ID:
		return fmt.Sprintf("SUM (%v)", id)
	case PRODUCT_ID:
		return fmt.Sprintf("PRODUCT (%v)", id)
	case MINIMUM_ID:
		return fmt.Sprintf("MINIMUM (%v)", id)
	case MAXIMUM_ID:
		return fmt.Sprintf("MAXIMUM (%v)", id)
	case LITERAL_VALUE_ID:
		return fmt.Sprintf("LITERAL (%v)", id)
	case GREATER_THAN_ID:
		return fmt.Sprintf("GREATER THAN (%v)", id)
	case LESS_THAN_ID:
		return fmt.Sprintf("LESS THAN (%v)", id)
	case EQUAL_TO_ID:
		return fmt.Sprintf("EQUAL TO (%v)", id)
	}

	panic("incorrect type id")
}

type Packet struct {
	parent  *Packet
	version uint8
	typeId  uint8

	operatorMode   uint8
	subpacketBits  int16
	subpacketCount int16

	literalValue []uint8

	subpackets []*Packet
}

func getLiteralValue(val []uint8) int64 {
	initial := int64(0)

	for _, v := range val {
		initial = initial << 8
		initial = initial | int64(v)
	}

	return initial
}

func readVersion(index int, bytes []uint8) (uint8, int) {
	version := libs.GetNextNBits(3, index, bytes)

	return version, index + 3
}

func readTypeId(index int, bytes []uint8) (uint8, int) {
	typeId := libs.GetNextNBits(3, index, bytes)

	return typeId, index + 3
}

func readOperatorMode(index int, bytes []uint8) (uint8, int) {
	mode := libs.GetNextNBits(1, index, bytes)

	return mode, index + 1
}

func readSubpacketsLengthBits(index int, bytes []uint8) (int16, int) {
	firstHalf := libs.GetNextNBits(8, index, bytes)
	secondHalf := libs.GetNextNBits(7, index+8, bytes)

	return int16((int16(firstHalf) << 7) | int16(secondHalf)), index + 15
}

func readSubpacketCount(index int, bytes []uint8) (int16, int) {
	firstHalf := libs.GetNextNBits(8, index, bytes)
	secondHalf := libs.GetNextNBits(3, index+8, bytes)

	return int16((int16(firstHalf) << 3) | int16(secondHalf)), index + 11
}

func readLiteral(index int, bytes []uint8) ([]uint8, int) {
	newIndex := index
	res := []uint8(nil)

	firstHalf := true

	for {
		val := libs.GetNextNBits(5, newIndex, bytes)
		newIndex += 5

		lastGroupMask := uint8(1 << 4)
		lastGroupBit := lastGroupMask & val

		// clear last group bit if set
		if lastGroupBit != 0 {
			val = val & 15
		}

		if firstHalf {
			res = append(res, val<<4)
			firstHalf = false
		} else {
			lastBytes := res[len(res)-1]
			lastBytes = lastBytes | val
			res[len(res)-1] = lastBytes
			firstHalf = true
		}

		if lastGroupBit == 0 {
			break
		}
	}

	if !firstHalf {
		// need to shift everything for 4
		fromPrevious := uint8(0)
		for i := 0; i < len(res); i++ {
			val := res[i]
			newFromPrevious := (val & 15) << 4
			val = fromPrevious | (val >> 4)
			res[i] = val
			fromPrevious = newFromPrevious
		}
	}

	return res, newIndex
}

func countSubpackets(packet *Packet) int {
	if len(packet.subpackets) == 0 {
		return 0
	}

	count := len(packet.subpackets)
	for i := 0; i < len(packet.subpackets); i++ {
		count += countSubpackets(packet.subpackets[i])
	}

	return count
}

// returns decoded packet and how many bits was read
func decodePacket(startIndex int, bytes []uint8, parent *Packet) (*Packet, int) {
	packet := new(Packet)
	packet.parent = parent

	var value uint8
	index := startIndex
	value, index = readVersion(index, bytes)
	packet.version = value

	value, index = readTypeId(index, bytes)
	packet.typeId = value

	if value == 4 {
		val, newIndex := readLiteral(index, bytes)
		packet.literalValue = val

		return packet, newIndex - startIndex
	} else {
		value, index = readOperatorMode(index, bytes)
		packet.operatorMode = value

		if value == 0 {
			subpacketBitLength, newIndex := readSubpacketsLengthBits(index, bytes)
			packet.subpacketBits = subpacketBitLength

			bitCount := 0
			for bitCount < int(subpacketBitLength) {
				subpacket, childBitsRead := decodePacket(newIndex, bytes, packet)
				packet.subpackets = append(packet.subpackets, subpacket)
				newIndex += childBitsRead
				bitCount += childBitsRead
			}

			return packet, newIndex - startIndex
		} else if value == 1 {
			subpacketCount, newIndex := readSubpacketCount(index, bytes)
			packet.subpacketCount = subpacketCount

			count := 0

			for count < int(subpacketCount) {
				subpacket, childBitsRead := decodePacket(newIndex, bytes, packet)
				packet.subpackets = append(packet.subpackets, subpacket)
				newIndex += childBitsRead

				count++
			}

			return packet, newIndex - startIndex
		} else {
			panic("unknown operator mode")
		}
	}
}

func part1(bytes []uint8) *Packet {
	packet, _ := decodePacket(0, bytes, nil)
	sum := sumVersions(packet)
	fmt.Println("Total version sum:", sum)

	return packet
}

func sumVersions(packet *Packet) int {
	versionSum := int(packet.version)

	for _, sub := range packet.subpackets {
		versionSum += sumVersions(sub)
	}

	return versionSum
}

func bytesToString(bytes []uint8) string {
	res := ""

	for _, bits := range bytes {
		tmp := fmt.Sprintf("%08b", bits)
		res += tmp
	}

	return res
}

func printPacket(packet *Packet, indent int) {
	padding := strings.Repeat(" ", indent)
	fmt.Printf("%vVer: %v, TypeID: %v, OpMode: %v, SubBits: %v, SubCount: %v (%v), Literal: %v, Value: %v\n",
		padding, packet.version, getTypeIdDesc(packet.typeId), packet.operatorMode, packet.subpacketBits, packet.subpacketCount, len(packet.subpackets), bytesToString(packet.literalValue), packet.GetValue())

	for _, sub := range packet.subpackets {
		printPacket(sub, indent+2)
	}
}

func (p *Packet) GetValue() int64 {
	res := int64(0)

	switch p.typeId {
	case SUM_ID:
		for _, sub := range p.subpackets {
			res += sub.GetValue()
		}
	case PRODUCT_ID:
		if len(p.subpackets) == 1 {
			res = p.subpackets[0].GetValue()
		} else {
			res = 1
			for _, sub := range p.subpackets {
				res = res * sub.GetValue()
			}
		}
	case MINIMUM_ID:
		res = p.subpackets[0].GetValue()

		for _, sub := range p.subpackets {
			val := sub.GetValue()
			if res > val {
				res = val
			}
		}
	case MAXIMUM_ID:
		res = p.subpackets[0].GetValue()

		for _, sub := range p.subpackets {
			val := sub.GetValue()
			if res < val {
				res = val
			}
		}
	case LITERAL_VALUE_ID:
		res = getLiteralValue(p.literalValue)
	case GREATER_THAN_ID:
		first := p.subpackets[0].GetValue()
		second := p.subpackets[1].GetValue()

		if first > second {
			res = 1
		} else {
			res = 0
		}
	case LESS_THAN_ID:
		first := p.subpackets[0].GetValue()
		second := p.subpackets[1].GetValue()

		if first < second {
			res = 1
		} else {
			res = 0
		}
	case EQUAL_TO_ID:
		first := p.subpackets[0].GetValue()
		second := p.subpackets[1].GetValue()

		if first == second {
			res = 1
		} else {
			res = 0
		}
	default:
		panic("unknown type id")
	}

	return res
}

func part2(rootPacket *Packet) {
	// print hierarchy
	printPacket(rootPacket, 0)
	fmt.Println("Value: ", rootPacket.GetValue())
}

func Run() {
	input := libs.ReadTxtFileLines("days/day16/input.txt")

	bytes := []uint8(nil)

	firstHalf := true
	for _, char := range input[0] {
		num, err := strconv.ParseUint(string(char), 16, 8)
		if err != nil {
			panic(err)
		}
		bits := uint8(num)

		if firstHalf {
			bytes = append(bytes, bits<<4)
			firstHalf = false
		} else {
			lastBytes := bytes[len(bytes)-1]
			lastBytes = lastBytes | bits
			bits = lastBytes
			bytes[len(bytes)-1] = lastBytes
			firstHalf = true
		}
	}

	rootPacket := part1(bytes)
	part2(rootPacket)
}

/*

0 = 0000
1 = 0001
2 = 0010
3 = 0011
4 = 0100
5 = 0101
6 = 0110
7 = 0111
8 = 1000
9 = 1001
A = 1010
B = 1011
C = 1100
D = 1101
E = 1110
F = 1111

A = 10
B = 11
C = 12
D = 13
E = 14
F = 15

*/
