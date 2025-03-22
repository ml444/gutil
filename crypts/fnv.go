package crypts

import "hash/fnv"

func HashCode(s string) uint32 {
	f := fnv.New32a()
	_, _ = f.Write([]byte(s))
	return f.Sum32()
}
