package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaaa"
	n := int32(len(s))

	fmt.Println(substrCount(n, s))
}

func substrCount(n int32, s string) int64 {
    count := n

    // collect all string data want to be checked is palindrome or not and save to data
    for i := 2; i <= int(n); i++ {

        for j := 0; j < int(n)-i+1; j++ {

            var data string
            for k := j; k < j+i; k++ {
                data += string(s[k])
            }
            
            if data[0] == data[len(data)-1] {
                if isPalindrome(data) {
                    count++
                }
            }
        }
    }
    return int64(count)
}

func isPalindrome(s string) bool {
    var is string
    for i := len(s) - 1; i >= 0; i-- {
        is += string(s[i])
    }

    if is == s {
        // check jika panjang lebih dari 3 apakah karakter2 pembentuk hanya terdiri dari 2 atau 1 karakter jika lebih maka false
        if len(s) > 3 {
            slice := strings.Split(s, "")
            slice = unique(slice)
            if len(slice) > 2 {
                return false
            }
        }
        return true
    }

    return false
}

func unique(s []string) []string {
    key := make(map[string]bool)
    new := []string{}
    for _, v := range s {
        if key[v] == false {
            key[v] = true
            new = append(new, v)
        }
    }
    return new
}
