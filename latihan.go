package main
import "fmt"
const NMAX int = 100
type tabInt[NMAX]int
func main(){
	var n, i, max, min, idx int
	var A tabInt
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	max = A[0]
	min = A[0]
	for i = 0; i < n; i++{
		if max < A[i]{
			max = A[i]
		}else if min > A[i]{
			min = A[i]
		}
	}
	for  i = min; i < max; i++{
		if i != A[i]{
			idx = i
		}
	}
	fmt.Print(idx)
}
