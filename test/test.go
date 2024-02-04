package test

// test func
// func InsertSlice(list []int, key int, pointer int) {
// 	for i := 0; i < len(list); i++ {
// 		if list[i] < key {
// 			continue
// 		}

// 		copy(list[i:], list[i-1:])
// 		list[i] = key

// 		break
// 	}

// 	list[pointer] = key
// }

func InsertSlice2(list []int, key int, pointer int) {
	for i := 0; i < len(list); i++ {
		if list[i] > key {
			pointer = i
			break
		}
	}

	copy(list[pointer:], list[pointer-1:])
	list[pointer] = key
}

func InsertSlice3(list []int, key int, pointer int) {
	for i := 0; i < len(list); i++ {
		if list[i] > key {
			list = append(list[:i], list[i-1:cap(list)-1]...)
			list[i] = key

			return
		}
	}

	list[pointer] = key
}

func Memory(list []int) (int, int, int) {
	pero := list[0]
	marko := list[1]
	fja := list[2]

	return pero, marko, fja
}
