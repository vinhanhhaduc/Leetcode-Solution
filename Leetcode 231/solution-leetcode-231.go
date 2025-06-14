func isPowerOfTwo(num int) bool {
	return (num > 0 && ((num & (num - 1)) == 0))
}

func isPowerOfTwo1(num int) bool {
	return num > 0 && (1073741824%num == 0)
}

func isPowerOfTwo2(num int) bool {
	allPowerOfTwoMap := map[int]int{1: 1, 2: 2, 4: 4, 8: 8, 16: 16, 32: 32, 64: 64, 128: 128, 256: 256, 512: 512, 1024: 1024, 2048: 2048, 4096: 4096, 8192: 8192, 16384: 16384, 32768: 32768, 65536: 65536, 131072: 131072, 262144: 262144, 524288: 524288, 1048576: 1048576, 2097152: 2097152, 4194304: 4194304, 8388608: 8388608, 16777216: 16777216, 33554432: 33554432, 67108864: 67108864, 134217728: 134217728, 268435456: 268435456, 536870912: 536870912, 1073741824: 1073741824}
	_, ok := allPowerOfTwoMap[num]
	return ok
}
func isPowerOfTwo3(num int) bool {
	for num >= 2 {
		if num%2 == 0 {
			num = num / 2
		} else {
			return false
		}
	}
	return num == 1
}