package controller

func initarray() [20]int {
	arr := [20]int{}
	for i := 1; i <= 20; i++ {
		arr[i] = i
	}
	return arr
}

func Displayaray() [20]int {
	array := initarray()
	return array
}
