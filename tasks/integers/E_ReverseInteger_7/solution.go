package E_ReverseInteger_7

import "math"

/* Given a 32-bit signed integer, reverse digits of an integer. */

func reverse(x int) int {
	num, sum := x, 0
	for num != 0 {
		digit := num % 10    // Остаток от деления в переменную 321 -> 1
		sum = sum*10 + digit // Сумма * на 10 и добавлем переменную ; 0 * 10 + 1 = 1, 1 * 10 + 2 = 12, 12 * 10 + 3 = 123
		num = num / 10       // Делим на 10 оставшееся число.
	}
	// проверям что бы получившееся число не привысило int32
	// math.MaxInt32 = int(math.Pow(2,31)) - 1 и math.MinInt32 = int(math.Pow(-2,31))
	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}
	return sum
}
