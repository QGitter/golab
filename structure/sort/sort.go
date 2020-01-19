package sort

//冒泡排序：两两比较相邻记录的 关键字，如果反序则交换，直到没有反序的记录为止
func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//选择排序：在长度为N的无序数组中，第一次遍历n-1个数，找到最小的数值与第一个元素交换，第二次遍历n-2个数，
// 找到最小的数值与第二个元素交换。。。第n-1次遍历，找到最小的数值与第n-1个元素交换，排序完成
func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var minIndex int
	for i := 0; i < len(arr)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

//插入排序：在要排序的一组数中，假定前n-1个数已经排好序，现在将第n个数插到前面的有序数列中，
// 使得这n个数也是排好顺序的。如此反复循环，直到全部排号顺序
func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func ShellSort(arr []int) []int {

	return arr
}
