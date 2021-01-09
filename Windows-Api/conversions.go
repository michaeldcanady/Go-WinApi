package winapi

import(
  "syscall"
)

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
