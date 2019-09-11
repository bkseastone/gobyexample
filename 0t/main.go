package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(str string) string {
	rs := []rune(str)
	len := len(rs)
	var tt []rune

	tt = make([]rune, 0)
	for i := 0; i < len; i++ {
		tt = append(tt, rs[len-i-1])
	}
	return string(tt[0:])
}
func main() {
	const n = 9
	count := 0
	max, _ := strconv.Atoi(strings.Repeat("9", n))
	for i := 0; i < max; i++ {
		flipStr := reverse(strconv.Itoa(i))
		len := len(flipStr)
		flip, _ := strconv.Atoi(flipStr)
		if i/len == flip {
			count++
			fmt.Printf("%d / %d = %d\n", i, len, flip)
		}
	}
}

//$n     = 9;
//$max   = (int) str_repeat(9, $n);
//$count = 0;
//for ($i = 0; $i < $max; $i++) {
//$filpStr = strrev($i);
//$len     = strlen($filpStr);
//$filp    = (int) strrev($i);
//if ($i / $len === $filp) {
//$count++;
//echo "{$i} / {$len} = {$filp}\n";
//}
//}
