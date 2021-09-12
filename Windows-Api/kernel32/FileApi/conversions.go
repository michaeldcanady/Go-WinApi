package fileapi

import (
	"math/bits"
	"strings"
	"syscall"
	"unsafe"
	"strconv"
)

func bitsToBits(data []byte) (st []int) {
	st = make([]int, len(data)*8) // Performance x 2 as no append occurs.
	for i, d := range data {
		for j := 0; j < 8; j++ {
			if bits.LeadingZeros8(d) == 0 {
				// No leading 0 means that it is a 1
				st[i*8+j] = 1
			} else {
				st[i*8+j] = 0
			}
			d = d << 1
		}
	}
	return
}

func bitsToDrives(bitMap uint32) (drives []string) {
	availableDrives := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for i := range availableDrives {
		if bitMap&1 == 1 {
			drives = append(drives, availableDrives[i])
		}
		bitMap >>= 1
	}
	return drives
}

func LPSTRsToStrings(in [][]uint16) []string {
	if len(in) == 0 {
		return nil
	}

	out := make([]string, len(in))
	for i, s := range in {
		out[i] = syscall.UTF16ToString(s)
	}

	return out
}

func UintptrFromString(s *string) uintptr {
	if *s == "" {
		return 0
	}
	var ret *uint16
	// Some Windows API functions like GetTextExtentPoint32() panic when given
	// a string containing NUL. This block checks & returns the part before NUL.
	zeroAt := strings.Index(*s, "\x00")
	if zeroAt == -1 {
		ret, _ = syscall.UTF16PtrFromString(*s)
		return uintptr(unsafe.Pointer(ret))
	}
	if zeroAt == 0 {
		return 0
	}
	ret, _ = syscall.UTF16PtrFromString((*s)[:zeroAt])
	return uintptr(unsafe.Pointer(ret))
}

//seperateFlags takes SystemFlags as hex, converts them to binary to determine each flag
//Then converts back into an int64 for later usage
func seperateFlags(SystemFlags uint32) (flags []int64) {
	binary := string(strconv.FormatInt(int64(SystemFlags), 2))
	bin := strings.Split(binary, "")

	endingZeros := strings.Repeat("0", int(len(bin)-1))
	endingZeros1 := strings.Split(endingZeros, "")
	for i, b := range bin {
		flag := b + strings.Join(endingZeros1, "")
		intFlag, _ := strconv.ParseInt(parseBinToHex(flag), 16, 64)
		flags = append(flags, intFlag)

		if i != int(len(bin)-1) {
			endingZeros1 = remove(endingZeros1, 0)
		}
	}
	return
}