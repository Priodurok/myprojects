package main

import (
	"fmt"
	"strconv"
)

func implContains(sl map[string]int, name string) bool {

	for key, _ := range sl {
		if key == name {
			return true
		}
	}
	return false
}

func entryFull(sl []string, sd []string, name1 string, name2 string) bool {

	for _, value := range sl {
		if value == name1 {
			for _, value := range sl {
				if value == name2 {
					return true
				}
			}
		}
	}
	for _, value := range sd {
		if value == name1 {
			for _, value := range sd {
				if value == name2 {
					return true
				}
			}
		}
	}
	return false
}

func mapContains(sl map[string]int, name string) bool {

	for key, _ := range sl {
		if key == name {
			return true
		}
	}
	return false
}

func operContains(sl []string, name string) bool {

	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func main() {

	var firstV string
	var secondV string
	var oper string
	var firstValue int
	var secondValue int

	fmt.Scan(&firstV, &oper, &secondV)

	var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var numbers_rim = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var operations = []string{"+", "-", "*", "/"}

	var rim = make(map[string]int)
	rim["I"] = 1
	rim["II"] = 2
	rim["III"] = 3
	rim["IV"] = 4
	rim["V"] = 5
	rim["VI"] = 6
	rim["VII"] = 7
	rim["VIII"] = 8
	rim["IX"] = 9
	rim["X"] = 10
	rim["XI"] = 11
	rim["XII"] = 12
	rim["XIII"] = 13
	rim["XIV"] = 14
	rim["XV"] = 15
	rim["XVI"] = 16
	rim["XVII"] = 17
	rim["XVIII"] = 18
	rim["XIX"] = 19
	rim["XX"] = 20
	rim["XXI"] = 21
	rim["XXII"] = 22
	rim["XXIII"] = 23
	rim["XXIV"] = 24
	rim["XXV"] = 25
	rim["XXVI"] = 26
	rim["XXVII"] = 27
	rim["XXVIII"] = 28
	rim["XXIX"] = 29
	rim["XXX"] = 30
	rim["XXXI"] = 31
	rim["XXXII"] = 32
	rim["XXXIII"] = 33
	rim["XXXIV"] = 34
	rim["XXXV"] = 35
	rim["XXXVI"] = 36
	rim["XXXVII"] = 37
	rim["XXXVIII"] = 38
	rim["XXXIX"] = 39
	rim["XL"] = 40
	rim["XLI"] = 41
	rim["XLII"] = 42
	rim["XLIII"] = 43
	rim["XLIV"] = 44
	rim["XLV"] = 45
	rim["XLVI"] = 46
	rim["XLVII"] = 47
	rim["XLVIII"] = 48
	rim["XLIX"] = 49
	rim["L"] = 50
	rim["LI"] = 51
	rim["LII"] = 52
	rim["LIII"] = 53
	rim["LIV"] = 54
	rim["LV"] = 55
	rim["LVI"] = 56
	rim["LVII"] = 57
	rim["LVIII"] = 58
	rim["LIX"] = 59
	rim["LX"] = 60
	rim["LXI"] = 61
	rim["LXII"] = 62
	rim["LXIII"] = 63
	rim["LXIV"] = 64
	rim["LXV"] = 65
	rim["LXVI"] = 66
	rim["LXVII"] = 67
	rim["LXVIII"] = 68
	rim["LXIX"] = 69
	rim["LXX"] = 70
	rim["LXXI"] = 71
	rim["LXXII"] = 72
	rim["LXXIII"] = 73
	rim["LXXIV"] = 74
	rim["LXXV"] = 75
	rim["LXXVI"] = 76
	rim["LXXVII"] = 77
	rim["LXXVIII"] = 78
	rim["LXXIX"] = 79
	rim["LXXX"] = 80
	rim["LXXXI"] = 81
	rim["LXXXII"] = 82
	rim["LXXXIII"] = 83
	rim["LXXXIV"] = 84
	rim["LXXXV"] = 85
	rim["LXXXVI"] = 86
	rim["LXXXVII"] = 87
	rim["LXXXVIII"] = 88
	rim["LXXXIX"] = 89
	rim["XC"] = 90
	rim["XCI"] = 91
	rim["XCII"] = 92
	rim["XCIII"] = 93
	rim["XCIV"] = 94
	rim["XCV"] = 95
	rim["XCVI"] = 96
	rim["XCVII"] = 97
	rim["XCVIII"] = 98
	rim["XCIX"] = 99
	rim["C"] = 100

	var arab = make(map[string]int)
	arab["1"] = 1
	arab["2"] = 2
	arab["3"] = 3
	arab["4"] = 4
	arab["5"] = 5
	arab["6"] = 6
	arab["7"] = 7
	arab["8"] = 8
	arab["9"] = 9
	arab["10"] = 10
	arab["-1"] = -1
	arab["-2"] = -2
	arab["-3"] = -3
	arab["-4"] = -4
	arab["-5"] = -5
	arab["-6"] = -6
	arab["-7"] = -7
	arab["-8"] = -8
	arab["-9"] = -9
	arab["-10"] = -10

	presentRim1 := implContains(rim, firstV)
	presentRim2 := implContains(rim, secondV)
	//presentArab1 := implContains(arab, firstV)
	//presentArab2 := implContains(arab, secondV)

	entryArab := entryFull(numbers, numbers_rim, firstV, secondV)

	operEntry := operContains(operations, oper)

	var newArab1 int
	var newArab2 int

	s := firstV
	if n, err := strconv.Atoi(s); err == nil {
		newArab1 = n
	}

	c := secondV
	if n, err := strconv.Atoi(c); err == nil {
		newArab2 = n
	}

	if presentRim1 {
		for key, value := range rim {

			if key == firstV {
				firstValue = value
				if firstValue <= 10 {
					break
				} else {
					fmt.Println("Ошибка! Числа должны быть от I до X!")
				}

			}
		}

		if presentRim2 {
			for key, value := range rim {

				if key == secondV {
					secondValue = value
					if secondValue <= 10 {
						break
					} else {
						fmt.Println("Ошибка! Числа должны быть от I до X!")
					}
				}
			}
		}
	}

	if newArab1 <= 10 && newArab1 > 0 {

		if newArab2 <= 10 && newArab2 > 0 {

			for key, value := range arab {
				if key == firstV {
					firstValue = value
					if firstValue <= 10 && firstValue > 0 {
						break
					} else {
						fmt.Println("Ошибка! Числа должны быть от 1 до 10!")
					}
				}
			}

			for key, value := range arab {
				if key == secondV {
					secondValue = value
					if secondValue <= 10 && secondValue > 0 {
						break
					} else {
						fmt.Println("Ошибка! Числа должны быть от 1 до 10!")
					}
				}
			}
		} else if newArab2 > 10 || newArab2 < 0 {
			fmt.Println("Ошибка! Числа должны быть от 1 до 10!")
		}
	} else if newArab1 > 10 || newArab1 < 0 {
		fmt.Println("Ошибка! Числа должны быть от 1 до 10!")
	}

	if entryArab == false {
		fmt.Println("Ошибка! Оба числа должны быть арабскими, либо римскими!")
	}

	if operEntry {
		for _, i := range numbers {
			if i == firstV {
				for _, i := range numbers {
					if i == secondV {

						if oper == "+" {
							fmt.Println(firstValue + secondValue)
						} else if oper == "-" {
							fmt.Println(firstValue - secondValue)
						} else if oper == "*" {
							fmt.Println(firstValue * secondValue)
						} else if oper == "/" {
							fmt.Println(firstValue / secondValue)
						}

					}
				}
			}
		}
	} else {
		fmt.Println("Ошибка! Нет такой операции!")
	}

	if operEntry {
		for _, i := range numbers_rim {
			if i == firstV {
				for _, i := range numbers_rim {
					if i == secondV {

						var answer_rim int

						if oper == "+" {

							answer_rim = firstValue + secondValue

							for key, value := range rim {
								if value == answer_rim {
									fmt.Println(key)
								}
							}

						} else if oper == "-" {
							answer_rim = firstValue - secondValue

							if answer_rim > 0 {
								for key, value := range rim {
									if value == answer_rim {
										fmt.Println(key)
									}
								}
							} else {
								fmt.Println("Ошибка! В римской системе нет отрицательных чисел!")
							}
						} else if oper == "*" {
							answer_rim = firstValue * secondValue
							for key, value := range rim {
								if value == answer_rim {
									fmt.Println(key)
								}
							}
						} else if oper == "/" {
							answer_rim = firstValue / secondValue
							for key, value := range rim {
								if value == answer_rim {
									fmt.Println(key)
								}
							}
						}
					}
				}
			}
		}
	}
}
