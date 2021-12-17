package day16

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
)

const (
	LITERAL_VALUE_ID = 4
)

type Packet struct {
	parent *Packet

	id int

	version uint8
	typeId  uint8

	operatorMode   uint8
	subpacketBits  int16
	subpacketCount int16

	literalValue []uint8

	subpackets []*Packet
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

var totalPackets int = 0

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

	totalPackets++
	packet.id = totalPackets
	fmt.Print("Processing packet with id ", packet.id)
	if parent != nil {
		fmt.Print(" (", parent.id, ")")
	}
	fmt.Println()

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

				count += 1 + countSubpackets(subpacket)
				fmt.Println("packet limit", subpacketCount, "count", count)
			}

			return packet, newIndex - startIndex
		} else {
			panic("unknown operator mode")
		}
	}
}

func part1(chars []rune) {
	bytes := []uint8(nil)

	firstHalf := true
	for _, char := range chars {
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

	rootPackets := make([]*Packet, 0)

	totalBits := len(bytes) * 8

	index := 0
	for index < totalBits {
		// check if we only have zeros left
		justZeros := true
		testIndex := index
		for {
			bitsLeft := totalBits - testIndex
			if bitsLeft > 8 {
				res := libs.GetNextNBits(8, testIndex, bytes)
				testIndex += 8
				if res != 0 {
					justZeros = false
					break
				}
			} else {
				res := libs.GetNextNBits(bitsLeft, testIndex, bytes)
				testIndex += bitsLeft
				if res != 0 {
					justZeros = false
					break
				}
			}

			if bitsLeft <= 8 {
				break
			}
		}

		if justZeros {
			break
		}
		packet, bits := decodePacket(index, bytes, nil)
		index += bits
		rootPackets = append(rootPackets, packet)
	}

	sum := sumVersions(rootPackets)
	fmt.Println("Total sum:", sum)
}

func sumVersions(packets []*Packet) int {
	versionSum := int(0)

	for _, packet := range packets {
		versionSum += int(packet.version)
		if len(packet.subpackets) > 0 {
			versionSum += sumVersions(packet.subpackets)
		}
	}

	return versionSum
}

func Run() {
	input := libs.ReadTxtFileLines("days/day16/input.txt")

	chars := []rune(input[0])

	part1(chars)
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
