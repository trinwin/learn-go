package main

import "fmt"

func main()  {
	// maps are key value store - fast lookup - unordered list
	m := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
		"d": 40,
	} 
		fmt.Println(m)
		fmt.Println(m["a"])
		fmt.Println(m["x"])

		// Check to make sure the key exists
		v, ok := m["x"]
		fmt.Println(v)
		fmt.Println(ok)

		if v, ok := m["a"]; ok {
			fmt.Println("this is the if print", v)
		}

		m["f"] = 99

		for k, v := range m {
			fmt.Println(k, v)
		}

		xi := []int{1,2,3,4,5}

		for i, v := range xi {
			fmt.Println(i, v)
		}

		// delete an entry
		delete(m, "a")
		delete(m, "hello")
		fmt.Println(m)

		if v, ok := m["b"]; ok {
			fmt.Println("value", v)
			delete(m, "b")
			fmt.Println(m)	
		}
}	