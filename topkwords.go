package main

import "fmt"
import "strings"
import "sort"

//time: n * 

func main() {

	myarr := "It is a far far better thing that I do than I have ever done it is a far far better rest I go to than I have ever known"
	
	split_slice := strings.Split(myarr, " ")

	glob_map := make(map[string]int)

	for _, val := range split_slice {

		if count, ok := glob_map[val]; ok {
			glob_map[val] = count+1
		} else {
			glob_map[val] = 1
		}
	}

    type kv struct {
        Key   string
        Value int
    }

    var ss []kv
    for k, v := range glob_map {
        ss = append(ss, kv{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })

    for _, kv := range ss {
        fmt.Printf("%s \t: %d\n", kv.Key, kv.Value)
    }

	
}
