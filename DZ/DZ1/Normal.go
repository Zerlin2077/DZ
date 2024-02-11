package DZ1

import "fmt"

func Normal() {
	var arr [12]int
	fmt.Println("Введите одинадцать чисел через пробел или через enter:")
	fmt.Scan(&arr[1], &arr[2], &arr[3], &arr[4], &arr[5], &arr[6], &arr[7], &arr[8], &arr[9], &arr[10], &arr[11])

	var min, max = arr[1], arr[1]
	avg := 1

	for i := 1; i < 12; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
		avg += arr[i]
	}

	fmt.Println("\n===========\nМинимальное число:", min)
	fmt.Println("Максимальное число:", max)
	fmt.Println("Среднее арифметическое:", avg/11)

}
