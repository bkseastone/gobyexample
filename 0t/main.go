package main

import "fmt"

func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
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
