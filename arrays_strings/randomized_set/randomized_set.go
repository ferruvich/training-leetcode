package randomized_set

import "math/rand"

type RandomizedSet struct {
	elemList []int
	elemMap  map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		elemList: []int{},
		elemMap:  map[int]int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (set *RandomizedSet) Insert(val int) bool {
	if _, ok := set.elemMap[val]; ok {
		return false
	}

	set.elemList = append(set.elemList, val)
	set.elemMap[val] = len(set.elemList) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (set *RandomizedSet) Remove(val int) bool {
	el, ok := set.elemMap[val]
	if !ok {
		return false
	}

	latestElement := set.elemList[len(set.elemList)-1]
	set.elemList[el] = latestElement
	set.elemMap[latestElement] = el

	set.elemList = set.elemList[:len(set.elemList)-1]
	delete(set.elemMap, val)

	return true
}

/** Get a random element from the set. */
func (set *RandomizedSet) GetRandom() int {
	r := rand.Intn(len(set.elemList))
	return set.elemList[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
