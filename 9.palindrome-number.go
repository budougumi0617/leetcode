package main

func isPalindrome(x int) bool {
	// エッジケースは早期リターンする
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	if x < 10 {
		return true
	}

	var rn int
	for x > rn {
		// いい感じに2つも数字列になるまで分けていく。
		rn = rn*10 + x%10
		x /= 10
	}

	// 1桁余っている場合も考慮する
	return x == rn || x == rn/10
}
