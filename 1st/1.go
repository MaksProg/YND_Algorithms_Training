package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
 var n int
 fmt.Scan(&n)
  
 reader := bufio.NewReader(os.Stdin)
 line,_ := reader.ReadString('\n')
 fields := strings.Fields(line)

 nums := make([]int,0,n)
 for i:=0;i<n;i++{
	num,err:= strconv.Atoi(fields[i])
	if err == nil{
		nums = append(nums, num)
	}
 }

 min :=math.MaxInt32
 max := math.MinInt32
 for i:=0;i<len(nums);i++{
	if i % 2 == 0{
		if nums[i]<min{
			min = nums[i]
		}
	}else{
	if nums[i]>max{
		max = nums[i]
	}
	}
 }

sum:=0
for i:=0;i<len(nums);i++{
	if i % 2 == 0{
		sum+=nums[i]
	}else{
		sum-=nums[i]
	}

}

delta:= 2* (max-min)

if delta<0{
	delta = 0
}
fmt.Println(sum+delta)
}