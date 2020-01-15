package main

import (
	"fmt"
	"sort"
)

type Developer struct {
	Name string
	Age  int
}

type Developers []Developer

func (devs Developers) Len() int {
	return len(devs)
}

func (devs Developers) Less(i, j int) bool {
	return devs[i].Age < devs[j].Age
}

func (devs Developers) Swap(i, j int) {
	devs[i], devs[j] = devs[j], devs[i]
}

func main() {
	developers := Developers{
		{Name: "Koo", Age: 33},
		{Name: "Sanghyeon Lee", Age: 38},
		{Name: "Seungkyu", Age: 25},
	}

	sort.Sort(developers)
	fmt.Println(developers)

	sort.Sort(sort.Reverse(developers))
	fmt.Println(developers)
}
