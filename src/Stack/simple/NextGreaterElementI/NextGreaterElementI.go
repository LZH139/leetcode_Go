package NextGreaterElementI

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    var mSt Stack
    hashMap := make(map[int]int)
    ret := make([]int, 0, len(nums1))
    for i := 0; i < len(nums2); i++ {
        if mSt.Len() == 0 || nums2[mSt.Peek().(int)] >= nums2[i] {
            mSt.Push(i)
        } else {
            for mSt.Len() != 0 && nums2[mSt.Peek().(int)] < nums2[i] {
                topIndex := mSt.Pop().(int)
                hashMap[nums2[topIndex]] = nums2[i]
            }
            mSt.Push(i)
        }
    }
    for _, v := range nums1 {
        value, ok := hashMap[v]
        if ok {
            ret = append(ret, value)
        } else {
            ret = append(ret, -1)
        }
    }
    return ret
}

type (
    Stack struct {
        top    *node
        length int
    }
    node struct {
        value interface{}
        prev  *node
    }
)

// Create a new stack
func New() *Stack {
    return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
    return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
    if this.length == 0 {
        return nil
    }
    return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
    if this.length == 0 {
        return nil
    }

    n := this.top
    this.top = n.prev
    this.length--
    return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
    n := &node{value, this.top}
    this.top = n
    this.length++
}


