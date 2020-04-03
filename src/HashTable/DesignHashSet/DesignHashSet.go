package DesignHashSet

type MyHashSet struct {
	bucketArray [769]*bucket
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	bucketlist := [769]*bucket{}
	for i:=0;i<769;i++ {
		bucketlist[i] = new(bucket)
	}
	return MyHashSet{
		bucketArray: bucketlist,
	}
}

func (this *MyHashSet) _hash(key int) int {
	return key%769
}

func (this *MyHashSet) Add(key int)  {
	this.bucketArray[this._hash(key)].insert(key)
}


func (this *MyHashSet) Remove(key int)  {
	this.bucketArray[this._hash(key)].delete(key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.bucketArray[this._hash(key)].exists(key)
}

type bucket struct {
	container []int
}

func (b *bucket) insert(key int) {
	for _,value := range b.container {
		if value == key {
			return
		}
	}
	b.container = append(b.container,key)
}

func (b *bucket) delete(key int) {
	for index,value := range b.container {
		if value == key {
			b.container = append(b.container[0:index],b.container[index+1:]...)
			return
		}
	}
}

func (b *bucket) exists(key int) bool {
	for _,value := range b.container {
		if value == key {
			return true
		}
	}
	return false
}

