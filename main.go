package main

import (
	"fmt"
	"strings"
)

func main() {
	var x []string
	letter := [3]string{"a", "б", "в"}

	for i := 0; i < len(letter); i++ {
		x = append(x, strings.Join([]string{"a", "b"}, "-"))
	}
	fmt.Println(x)
}

// func main() {
// 	benchmarking.Run(benchmarking.BenchmarkConfig{
// 		BenchFunc: BenchmarkConcat,
// 	})
// }

// func reverseWords(s string) string {
//     words := strings.Split(s," ");
//     res := "";
//     for i:=0; i < len(words);i++{
//         curWord := strings.Split(words[i],"");
//         for j:=len(curWord)-1; j >= 0;j--{
//             res += curWord[j];
//         }
//         if(i!=len(words)-1){
//             res += " ";
//         }
//     }
//     return res;
// }
