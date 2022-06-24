package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	ss := []string{
		"aabbccddeefghi",    // NO
		"abcdefghhgfedecba", // YES
		"aabbcd",            // NO
		"aaaabbcc",          // NO
		"aaaaabc",           // NO
		"ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd", // YES
	}

	for _, v := range ss {
		fmt.Println(isValid(v))
	}
}

func isValid(s string) string {
	// Write your code here

	// split s
	slice := strings.Split(s, "")
	sort.Strings(slice)
	fmt.Println(slice)

	// make s be unique
	slice = Unique(slice).([]string)

	// count each letter in s
	counts := []int{}
	for _, v := range slice {
		counts = append(counts, strings.Count(s, string(v)))
	}

	// collect unique of counts and sort it
	countsUnique := Unique(counts).([]int)
	sort.Ints(countsUnique)

	// if countsUnique just contain one element it means all of letter has same amount so that s is valid
	if len(countsUnique) == 1 {
		return "YES"
	}

	// if countsUnique contain two element it is valid if one of them just 1 times
	if len(countsUnique) == 2 {
		// count how many times each of counts element come up
		key := make(map[int]int)
		for _, v := range counts {
			key[v]++
		}

		// Hitung masing2 dari countsUnique dari berapa region
		// Pengecekan selanjutnya akan berdasarkan yang lebih besar

		for k, v := range key {
			if v == 1 {
				if v == 1 {
					if k-1 == countsUnique[0] || k-1 == 0 {
						return "YES"
					}
				}
			}
		}
	}
	return "NO"
}

func Unique(list interface{}) interface{} {
	key := make(map[interface{}]bool)
	switch list.(type) {
	case []string:
		new := []string{}
		for _, v := range list.([]string) {
			if key[v] == false {
				key[v] = true
				new = append(new, v)
			}
		}
		return new
	case []int:
		new := []int{}
		for _, v := range list.([]int) {
			if key[v] == false {
				key[v] = true
				new = append(new, v)
			}
		}
		return new
	}

	return key
}
