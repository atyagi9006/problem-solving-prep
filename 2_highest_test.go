package main

import "testing"

// i Think i am done

func FindSecHigestIndex(input []int) int {
	if input == nil {
		return -1
	}

	if len(input) < 2 {
		return -1
	}

	if len(input) == 2 {
		if input[0] > input[1] {
			return 1
		} else {
			return 0
		}
	}

	min := -12015545
	first := min
	second := min
	index := 0
	for _, v := range input {
		if first < v {
			first = v
		}
	}

	for i, v := range input {
		if second < v && v != first {
			second = v
			index = i
		}
	}

	if second == min {
		return -1
	}

	return index
}

func FindSecondHighest(input []int) int {
	min := -12015545
	first := min
	second := min
	index := 0
	for i, v := range input {
		if first < v {
			second = first
			first = v
		} else if second < v && v != first {
			second = v
			index = i
		}
	}

	if second == min {
		return -1
	}

	return index
}

func TestfindSceondHighest(t *testing.T) {
	type testlBuilder func()
	testcases := map[string]testlBuilder{
		"test [1,4,3]": func() {
			test := []int{1, 4, 3}
			res := FindSecondHighest(test)

			if res == 2 {
				t.Log("test [1,4,3]", res)
			} else {
				t.Log("test Fail")
			}
		},
	}
	t.Log("FindSecHigestIndex Tests ...")
	for testname, test := range testcases {
		t.Run(testname, func(t *testing.T) {
			test()
		})
	}

}
/* 
func TestFindSecHigestIndex(t *testing.T) {
	type testlBuilder func()
	testcases := map[string]testlBuilder{
		"When input is nil": func() {
			res := FindSecHigestIndex(nil)
			if res == -1 {
				t.Log(" nil test Pass")
			} else {
				t.Log(" nil test fail")
			}
		},
		"When input is empty": func() {
			test := make([]int, 0, 0)
			res := FindSecHigestIndex(test)
			if res == -1 {
				t.Log(" empty test Pass")
			} else {
				t.Log(" empty test fail")
			}
		},
		"When input has only one element": func() {
			test := make([]int, 0, 0)
			test = append(test, 1)
			res := FindSecHigestIndex(test)
			if res == -1 {
				t.Log(" only one elment test Pass")
			} else {
				t.Log(" only one elment test fail")
			}
		},
		"test [1,4,3]": func() {
			test := []int{1, 4, 3}
			res := FindSecHigestIndex(test)

			if res == 2 {
				t.Log("test [1,4,3]", res)
			} else {
				t.Log("test Fail")
			}
		},
		"test [1,2,3]": func() {
			test := []int{1, 2, 3}
			res := FindSecHigestIndex(test)
			if res == 1 {
				t.Log("test [1,2,3]", res)
			} else {
				t.Log("test Fail")
			}
		},
		"test [3,2,1]": func() {
			test := []int{3, 2, 1}
			res := FindSecHigestIndex(test)
			if res == 1 {
				t.Log("[3,2,1]", res)
			} else {
				t.Log("test Fail")
			}
		},
		"test [2,1]": func() {
			test := []int{2, 1}
			res := FindSecHigestIndex(test)
			if res == 1 {
				t.Log("[2,1]", res)
			} else {
				t.Log("test Fail")
			}
		},
		"when input has repeat smaller element ": func() {
			test := []int{2, 1, 1}
			res := FindSecHigestIndex(test)
			if res == 1 {
				t.Log("[2,1,1]", res)
			} else {
				t.Log("test Fail")
			}
		},
		"when input has repeat bigger element ": func() {
			test := []int{2, 3, 3}
			res := FindSecHigestIndex(test)
			if res == 0 {
				t.Log("[2,3,3]", res)
			} else {
				t.Log("test Fail")
			}
		},
		"when all the elements are same": func() {
			test := []int{1, 1, 1}
			res := FindSecHigestIndex(test)
			if res == -1 {
				t.Log("[1,1,1]", res)
			} else {
				t.Log("test Fail")
			}
		},
	}
	t.Log("FindSecHigestIndex Tests ...")
	for testname, test := range testcases {
		t.Run(testname, func(t *testing.T) {
			test()
		})
	}
} */
