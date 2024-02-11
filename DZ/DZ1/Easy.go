package DZ1

import "fmt"

func Easy() {
	var number int
	fmt.Print("Введите целое число: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("Число [%d] - чётное", number)
	} else {
		fmt.Printf("Число [%d] - нечётное", number)
	}
}
