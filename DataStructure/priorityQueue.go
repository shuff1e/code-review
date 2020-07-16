package DataStructure

type Element interface {
	CompareTo(other interface{}) int
}

type PriorityQueue struct {
	arr []Element
	size int
}

func NewPriorityQueue(size int) *PriorityQueue {
	return &PriorityQueue{
		arr: make([]Element,size),
		size: 0,
	}
}

func (pq *PriorityQueue) Length() int {
	return pq.size
}

func (pq *PriorityQueue) Poll() Element {
	result := pq.arr[0]
	swap(pq.arr,0,pq.size-1)

	pq.size--

	siftDown(pq.arr,0,pq.size)
	return result
}

func siftDown(arr []Element,index,size int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index
	for left < size {
		if arr[left].CompareTo(arr[index]) >= 0 {
			largest = left
		}
		if right < size && arr[right].CompareTo(arr[largest]) >= 0 {
			largest = right
		}
		if index != largest {
			swap(arr,index,largest)
		} else {
			break
		}
		index = largest
		left = 2*index + 1
		right = 2*index + 2
	}
}
func (pq *PriorityQueue) Offer(e Element) {
	pq.size ++
	// grow
	if pq.size > len(pq.arr) {
		pq.arr = append(pq.arr,make([]Element,len(pq.arr))...)
	}
	pq.arr[pq.size-1] = e
	siftUP(pq.arr,pq.size-1)
}

func siftUP(arr []Element,index int) {
	for index > 0 {
		parent := (index-1)/2
		if arr[parent].CompareTo(arr[index]) < 0 {
			swap(arr,parent,index)
			index = parent
		} else {
			break
		}
	}
}

func swap(arr []Element,x,y int) {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}