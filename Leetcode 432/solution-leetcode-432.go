package main

import (
	"sort"
)

type AllOne struct {
	count map[string]int  
	set   map[int]map[string]bool
}

func Constructor() AllOne {
	return AllOne{
		count: make(map[string]int),
		set:   make(map[int]map[string]bool),
	}
}

func (this *AllOne) Inc(key string) {
	n := this.count[key]
	this.count[key]++

	if _, exists := this.set[n]; exists {
		delete(this.set[n], key)
		if len(this.set[n]) == 0 {
			delete(this.set, n)
		}
	}

	newCount := n + 1
	if _, exists := this.set[newCount]; !exists {
		this.set[newCount] = make(map[string]bool)
	}
	this.set[newCount][key] = true
}

func (this *AllOne) Dec(key string) {
	n := this.count[key]
	this.count[key]--

	delete(this.set[n], key)
	if len(this.set[n]) == 0 {
		delete(this.set, n)
	}

	if this.count[key] == 0 {
		delete(this.count, key)
	} else {
		newCount := n - 1
		if _, exists := this.set[newCount]; !exists {
			this.set[newCount] = make(map[string]bool)
		}
		this.set[newCount][key] = true
	}
}

func (this *AllOne) GetMaxKey() string {
	if len(this.set) == 0 {
		return ""
	}
	maxCount := maxKey(this.set)
	for key := range this.set[maxCount] {
		return key
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if len(this.set) == 0 {
		return ""
	}
	minCount := minKey(this.set)
	for key := range this.set[minCount] {
		return key
	}
	return ""
}

func maxKey(m map[int]map[string]bool) int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys[len(keys)-1]
}

func minKey(m map[int]map[string]bool) int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys[0]
}
