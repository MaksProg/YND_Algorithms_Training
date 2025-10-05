//НЕ ЯВЛЯЕТСЯ РЕШЕНИЕМ ТАК КАК НЕ ПРОХОДИТ 1 из ТЕСТОВ
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	topics := make([]int, 0, n)
	uniqueTopics := make(map[int]struct{})
	results := make([]int, 0, k)

	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	parts := strings.Fields(data)

	for _, s := range parts {
		num, _ := strconv.Atoi(s)
		uniqueTopics[num] = struct{}{}
		topics = append(topics, num)
	}

	if k <= len(uniqueTopics) {
		keys := make([]int, 0, len(uniqueTopics))
		for key := range uniqueTopics {
			keys = append(keys, key)
		}
		for i := 0; i < k; i++ {
			fmt.Print(keys[i], " ")
		}
		fmt.Println()
		return
	}


	for key := range uniqueTopics {
		results = append(results, key)
	}

	for _, t := range topics {
		if len(results) == k {
			break
		}
		results = append(results, t)
	}

	for _, val := range results {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
