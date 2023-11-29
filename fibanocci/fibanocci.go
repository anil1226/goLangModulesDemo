package fibanocci

import (
	"fmt"
)

func Sequence(limit int) string {
	j, k, l := 0, 1, 0
	var q string
	//var res strings.Builder
	for i := 0; i < limit; i++ {
		l = j + k
		//res.WriteString(strconv.FormatInt(int64(l), 10) + ", ")
		q += fmt.Sprint(l) + ","
		j = k
		k = l
	}
	return q
}
