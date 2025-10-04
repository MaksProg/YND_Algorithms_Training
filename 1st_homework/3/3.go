package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    s, _ := reader.ReadString('\n')
    s = s[:len(s)-1] 

    n := len(s)
    count := make([]int, 26)
    for _, c := range s {
        count[c-'a']++
    }

    totalPairs := n * (n - 1) / 2
    samePairs := 0
    for _, x := range count {
        samePairs += x * (x - 1) / 2
    }

    fmt.Println(1 + totalPairs - samePairs)
}


/*
//Жрёт память, но имхо более простая для понимания версия
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	base, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil{
		fmt.Println("Ошибка при чтении строки:",err)
	}
	base = strings.TrimSpace(base)

	unique:=make(map[string]struct{})
	unique[base] = struct{}{}
	passymb := []rune(base)
	for i := 0; i <len(passymb);i++{
		for j:=i+1;j < len(passymb);j++{
			tmp := make([]rune,len(passymb))
			copy(tmp,passymb)
			tmp[i],tmp[j]= tmp[j],tmp[i]

			unique[string(tmp)] = struct{}{}
		}
	}
	fmt.Println(len(unique))
}
*/