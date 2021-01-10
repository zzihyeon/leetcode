package Dec2021

import "log"

func CountArrangement(n int) int {
	result := 1
	duplicate := 0
	for i := 1; i <= n; i++ {
		sum := 0
		for j := 1; j <= n; j++ {
			sum++
			if i*j > n {
				sum--
				break
			}
			if j > 1 && j < i {
				log.Println(i * j)
				log.Println("i: ", i, ",j: ", j, ",j*i: ", i*j)
				duplicate++
			}
		}
		log.Println(sum)
		result = result * sum
	}
	return result - duplicate*n
}
