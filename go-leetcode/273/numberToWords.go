package main

import "fmt"

/*

273. 整数转换英文表示
将非负整数 num 转换为其对应的英文表示。



示例 1：

输入：num = 123
输出："One Hundred Twenty Three"
示例 2：

输入：num = 12345
输出："Twelve Thousand Three Hundred Forty Five"
示例 3：

输入：num = 1234567
输出："One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
示例 4：

输入：num = 1234567891
输出："One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"


提示：

0 <= num <= 231 - 1

 */

func main() {
	x := 1234567891
	fmt.Println(numberToWords(x))
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	billion := num/1000000000
	million := (num - billion * 1000000000) / 1000000
	thousand := (num - billion * 1000000000 - million * 1000000) / 1000
	rest := num - billion * 1000000000 - million * 1000000 - thousand * 1000

	result := ""
	if billion != 0 {
		result += three(billion) + " Billion"
	}
	if million != 0 {
		if len(result) > 0 {
			result += " "
		}
		result += three(million) + " Million"
	}
	if thousand != 0 {
		if len(result) > 0 {
			result += " "
		}
		result += three(thousand) + " Thousand"
	}
	if rest != 0 {
		if len(result) > 0 {
			result += " "
		}
		result += three(rest)
	}
	return result
}

func one(x int) string {
	switch x {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	}
	return ""
}

func ten(x int) string {
	switch x {
	case 2:
		return "Twenty"
	case 3:
		return "Thirty"
	case 4:
		return "Forty"
	case 5:
		return "Fifty"
	case 6:
		return "Sixty"
	case 7:
		return "Seventy"
	case 8:
		return "Eighty"
	case 9:
		return "Ninety"
	}
	return ""
}

func twoLessThan20(x int) string {
	switch x {
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 14:
		return "Fourteen"
	case 15:
		return "Fifteen"
	case 16:
		return "Sixteen"
	case 17:
		return "Seventeen"
	case 18:
		return "Eighteen"
	case 19:
		return "Nineteen"
	}
	return ""
}

func two(x int) string {
	if x == 0 {
		return ""
	} else if x < 10 {
		return one(x)
	} else if x < 20 {
		return twoLessThan20(x)
	} else {
		tenner := x / 10
		rest := x - 10*tenner
		if rest != 0 {
			return ten(tenner) + " " + one(rest)
		} else {
			return ten(tenner)
		}
	}
}

func three(x int) string {
	hundred := x / 100
	rest := x - hundred * 100
	if hundred != 0 && rest != 0 {
		return one(hundred) + " Hundred " + two(rest)
	} else if hundred == 0 && rest != 0 {
		return two(rest)
	} else if hundred != 0 && rest == 0 {
		return one(hundred) + " Hundred"
	}
	return ""
}

/*

方法一：分治
我们将这个问题分解成一系列子问题。例如，对于数字 1234567890，我们将它从低位开始每三个分成一组，
得到 1,234,567,890，它的英文表示为 1 Billion 234 Million 567 Thousand 890。
这样我们就将原问题分解成若干个三位整数转换为英文表示的问题了。

接下来，我们可以继续将三位整数分解，例如数字 234 可以分别成百位 2 和十位个位 34，
它的英文表示为 2 Hundred 34。这样我们继续将原问题分解成一位整数和两位整数的英文表示。
其中一位整数的表示是很容易的，而两位整数中除了 10 到 19 以外，其余整数的的表示可以分解成两个一位整数的表示，
这样问题就被圆满地解决了。





class Solution {
    public String one(int num) {
        switch(num) {
            case 1: return "One";
            case 2: return "Two";
            case 3: return "Three";
            case 4: return "Four";
            case 5: return "Five";
            case 6: return "Six";
            case 7: return "Seven";
            case 8: return "Eight";
            case 9: return "Nine";
        }
        return "";
    }

    public String twoLessThan20(int num) {
        switch(num) {
            case 10: return "Ten";
            case 11: return "Eleven";
            case 12: return "Twelve";
            case 13: return "Thirteen";
            case 14: return "Fourteen";
            case 15: return "Fifteen";
            case 16: return "Sixteen";
            case 17: return "Seventeen";
            case 18: return "Eighteen";
            case 19: return "Nineteen";
        }
        return "";
    }

    public String ten(int num) {
        switch(num) {
            case 2: return "Twenty";
            case 3: return "Thirty";
            case 4: return "Forty";
            case 5: return "Fifty";
            case 6: return "Sixty";
            case 7: return "Seventy";
            case 8: return "Eighty";
            case 9: return "Ninety";
        }
        return "";
    }

    public String two(int num) {
        if (num == 0)
            return "";
        else if (num < 10)
            return one(num);
        else if (num < 20)
            return twoLessThan20(num);
        else {
            int tenner = num / 10;
            int rest = num - tenner * 10;
            if (rest != 0)
              return ten(tenner) + " " + one(rest);
            else
              return ten(tenner);
        }
    }

    public String three(int num) {
        int hundred = num / 100;
        int rest = num - hundred * 100;
        String res = "";
        if (hundred * rest != 0)
            res = one(hundred) + " Hundred " + two(rest);
        else if ((hundred == 0) && (rest != 0))
            res = two(rest);
        else if ((hundred != 0) && (rest == 0))
            res = one(hundred) + " Hundred";
        return res;
    }

    public String numberToWords(int num) {
        if (num == 0)
            return "Zero";

        int billion = num / 1000000000;
        int million = (num - billion * 1000000000) / 1000000;
        int thousand = (num - billion * 1000000000 - million * 1000000) / 1000;
        int rest = num - billion * 1000000000 - million * 1000000 - thousand * 1000;

        String result = "";
        if (billion != 0)
            result = three(billion) + " Billion";
        if (million != 0) {
            if (! result.isEmpty())
                result += " ";
            result += three(million) + " Million";
        }
        if (thousand != 0) {
            if (! result.isEmpty())
                result += " ";
            result += three(thousand) + " Thousand";
        }
        if (rest != 0) {
            if (! result.isEmpty())
                result += " ";
            result += three(rest);
        }
        return result;
    }
}

 */