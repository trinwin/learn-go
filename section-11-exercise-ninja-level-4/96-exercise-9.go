package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`: {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`: {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)

	m[`nguyen_trinh`] = []string{`asian food`, `traveling`, `relaxing`}

	for k, v := range m {
		fmt.Println(k, v)
		for i, vv := range v {
			fmt.Println(i, vv)
		}
	}
}

