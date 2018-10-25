package main

import "fmt"

/** 二叉堆:本质上是一种完全二叉树，分为两种类型：
1.最大堆：最大堆任何一个父节点的值，都大于等于它左右孩子节点的值；
2.最小堆：最小堆任何一个父节点的值，都小于等于它左右孩子节点的值；
二叉树的根节点叫做堆顶，最大堆的最大值是堆顶，最小堆的最小值是堆顶。

二叉堆有如下几种操作：
1.插入节点
2.删除节点
3.构建二叉堆

以最小堆为例：
插入节点：插入位置是完全二叉树的最后一个位置，然后让新节点与其父节点比较，如果新节点更小，这让新节点上浮，即与父节点交换位置。继续用新节点和父节点比较，若小则同样上浮，直到父节点比新节点小为止；
删除节点：二叉堆的节点删除过程和插入过程正好相反，所删除的是处于堆顶的节点。删除堆顶节点后，为了维持完全二叉树的结构，把堆的最后一个节点补到原本堆顶的位置。接下来将新的堆顶节点和它的左右孩子进行比较，
	如果左右孩子中最小的一个比父节点小，则父节点下沉。继续让该节点与左右孩子比较大小，同样若左右孩子的最小节点比该节点小，则该节点下沉，直到比左右子节点都大为止。
构建二叉堆：就是把一个无序的完全二叉树调整为二叉堆，本质上就是让所有非叶子节点依次下沉。首先，从最后一个非叶子节点开始，若该节点大于其左右节点中最小一个，则下沉，
	直到该节点变成叶子节点或比其左右子节点都小为止。然后换一个非叶子节点继续进行比较和交换操作，直到检查完所有非叶子节点为止。


堆的代码实现：
二叉堆虽然是一颗完全二叉树，但它的存储方式并不是链式存储，而是顺序存储。换句话说，二叉堆的所有节点都存储在数组当中。

二叉堆：
				1
			  /   \
			 3	   2
			/ \   / \
		   6   5 7	 8
		  / \
		 9	 10

				↓

数组：1  3  2  6  5  7  8  9  10

在没有左右指针的情况下，如何定位到一个父节点的左右孩子呢？
假设父节点的下标是parent，那么它的左孩子下标就是2*parent+1，右孩子下标就是2*parent+2
*/
func main() {
	array := []int{1, 3, 2, 6, 5, 7, 8, 9, 10}
	upAdjust(&array)
	fmt.Println(array)

	array = []int{9, 4, 5, 3, 2, 7, 6, 10, 8}
	heapSort(&array)
	fmt.Println(array)
}

// 上浮调整
func upAdjust(array *[]int) {
	childIndex := len(*array) - 1
	parentIndex := (childIndex - 1) / 2
	// temp保存插入叶子节点的值，用于最后的赋值
	temp := (*array)[childIndex]
	for childIndex > 0 && temp < (*array)[parentIndex] {
		// 无需正真交换，单向赋值即可
		(*array)[childIndex] = (*array)[parentIndex]
		// 交换后子节点的下标变成父节点，新的父节点为原父节点的父节点
		childIndex = parentIndex
		parentIndex = (parentIndex - 1) / 2
	}
	(*array)[childIndex] = temp
}

// 下沉调整，将小的根节点下沉
// parentIndex 要下沉的父节点
// length 堆的有效大小
func downAdjust(array *[]int, parentIndex, length int) {
	// temp保存父节点的值，用于最后的赋值
	temp := (*array)[parentIndex]
	childIndex := 2*parentIndex + 1
	for childIndex < length {
		// 如果有右孩子，且右孩子大于左孩子的值，则定位到右孩子
		if childIndex+1 < length && (*array)[childIndex+1] > (*array)[childIndex] {
			childIndex++
		}
		// 如果父节点不小于任何一个孩子的值，直接跳出
		if temp >= (*array)[childIndex] {
			break
		}
		// 无需正真交换，单向赋值即可
		(*array)[parentIndex] = (*array)[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	(*array)[parentIndex] = temp
}

// 构建最大二叉堆
func buildHeap(array *[]int) {
	// 从最后一个非叶子节点开始，依次下沉调整
	// 完全二叉树的最后一个非叶子节点即：len(array) / 2
	for i := len(*array) / 2; i >= 0; i-- {
		downAdjust(array, i, len(*array))
	}
}

/**
二叉堆的用途：
1.堆排序

堆排序步骤：
1.把无序数组构建成最大二叉堆
2.循环删除堆顶元素，移到集合尾部，调节堆产生新的堆顶
*/

// 堆排序。假设array有n个元素，那么堆排序的时间复杂度为O(nlogn)，空间复杂度为O(1)
// 堆排序是不稳定排序
func heapSort(array *[]int) {
	// 1.构建最大二叉堆
	buildHeap(array)
	// 2.循环删除堆顶元素，移动到集合尾部。调节堆产生新的堆顶
	// i用来控制新堆的大小。通过构建最大堆，将堆顶元素[0]放到[i]的位置，然后对大小为i的数组进行重新下沉，获得新堆
	for i := len(*array) - 1; i > 0; i-- {
		max := (*array)[0]
		(*array)[0] = (*array)[i]
		(*array)[i] = max
		downAdjust(array, 0, i)
	}
}

/*
堆排序和快速排序的异同点：
1.堆排和快排的平均时间复杂度都是O(nlogn)，且都是不稳定排序
2.快排的最坏时间复杂度是O(n^2)，而堆排序的最坏时间复杂度稳定在O(nlogn)
3.快排的递归和非递归方法空间复杂度都是O(n)，而堆排序不需要额外空间，因此空间复杂度为O(1)
*/
