package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
import java.util.*;

public class Solution {
    public String PrintMinNumber(int [] numbers) {
        if(numbers.length == 0){return "";}
        String[] str1 = new String[numbers.length];
        for(int i = 0; i < numbers.length; i++){
            str1[i] = numbers[i] + "";
        }
        //贪心算法：相邻两者的字典序最小
        Arrays.sort(str1, new Comparator<String>() {
            @Override
            public int compare(String o1, String o2) {
                return (o1 + o2).compareTo(o2 + o1);
            }
        });
        String res = new String();
        for (String tmp : str1) {
            res += tmp;
        }
        return res;
    }
}
*/

type com []string

func (s com) Len() int {
	return len(s)
}
func (s com) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s com) Less(i, j int) bool {
	return s[i]+s[j] < s[j]+s[i]
}

func PrintMinNumber(numbers []int) string {
	if len(numbers) == 0 {
		return ""
	}
	str1 := make([]string, len(numbers))
	for i := 0; i < len(numbers); i++ {
		str1[i] = strconv.Itoa(numbers[i])
	}
	//贪心算法：相邻两者的字典序最小
	//sort.Sort(sort.StringSlice(str1))
	//sort.Sort(com(str1))
	sort.SliceStable(str1, func(i, j int) bool {
		return str1[i]+str1[j] < str1[j]+str1[i]
	})
	fmt.Println(str1)
	res := ""
	for _, tmp := range str1 {
		res += tmp
	}
	return res
}

func main() {
	fmt.Println(PrintMinNumber([]int{3, 32, 321})) // 332321
	//fruits := []string{"peach", "banana", "kiwi"}
	//sort.Sort(byLength(fruits))
	//fmt.Println(fruits)
}
