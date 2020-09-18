package main

func intToRoman(num int) string {
	levelMap := map[int][]string {
		0:[]string{"I","V","IV","IX"},
		1:[]string{"X","L","XL","XC"},
		2:[]string{"C","D","CD","CM"},
		3:[]string{"M"},
	}

	level := 0
	result := ""

	for num != 0 {
		pop := num%10
		temp := ""
		if pop == 4 {
			temp = levelMap[level][2]
		} else if pop == 9 {
			temp = levelMap[level][3]
		} else {
			if pop/5 == 1 {
				temp += levelMap[level][1]
			}

			for i:=0;i<pop%5;i++ {
				temp += levelMap[level][0]
			}
		}
		result = temp + result

		num/=10
		level ++
	}
	return result
}
