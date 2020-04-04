package DesignHashMap


type MyHashMap struct {
	hash_table []*Bucket
}


func Constructor() MyHashMap {
	var hashTable []*Bucket
	for i:=0;i<2069;i++ {
		hashTable = append(hashTable,&Bucket{})
	}
	return MyHashMap{hash_table:hashTable}
}


func (this *MyHashMap) Put(key int, value int)  {
	this.hash_table[key%2069].update(key,value)
}


func (this *MyHashMap) Get(key int) int {
	return this.hash_table[key%2069].get(key)
}


func (this *MyHashMap) Remove(key int)  {
	this.hash_table[key%2069].remove(key)
}

type Pair struct {
	first int
	second int
}

type Bucket struct {
	bucket []*Pair
}

func (b *Bucket) get(key int) int {
	for i:=0;i<len(b.bucket);i++ {
		if b.bucket[i].first == key {
			return b.bucket[i].second
		}
	}
	return -1
}

func (b *Bucket) update(key int, value int) {
	found := false
	for i:=0;i<len(b.bucket);i++ {
		if b.bucket[i].first == key {
			b.bucket[i].second = value
			found = true
			return
		}
	}
	if !found {
		b.bucket = append(b.bucket,&Pair{key,value})
	}
}

func (b *Bucket) remove(key int) {
	for i:=0;i<len(b.bucket);i++ {
		if b.bucket[i].first == key {
			b.bucket = append(b.bucket[:i],b.bucket[i+1:]...)
			break
		}
	}
}




