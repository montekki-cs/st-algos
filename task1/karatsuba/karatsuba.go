package karatsuba

import "strings"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func KaratsubaAdd(a, b string) string {
	carryover := byte(0)

	if len(b) > len(a) {
		a, b = b, a
	}

	j := len(b) - 1

	for i := len(a) - 1; i >= 0; i-- {
		res := a[i] - '0' + carryover

		if j >= 0 {
			res += (b[j] - '0')
			j--
		}

		if res > 9 {
			carryover = 1
			res -= 10
		} else {
			carryover = 0
		}

		a = a[:i] + string('0'+res) + a[i+1:]
	}

	if carryover != 0 {
		a = string('1') + a
	}

	return a
}

func KaratsubaMulSimple(a, b string) string {
	carryover := byte(0)

	if len(b) > len(a) {
		a, b = b, a
	}

	for i := len(a) - 1; i >= 0; i-- {
		res := ('0'-a[i])*('0'-b[0]) + carryover

		if res > 9 {
			carryover = res / 10
			res = res % 10
		} else {
			carryover = 0
		}

		a = a[:i] + string('0'+res) + a[i+1:]
	}

	if carryover != 0 {
		a = string(carryover+'0') + a
	}

	return a
}

func KaratsubaMul(a, b string) string {
	if len(a) == 0 || len(b) == 0 {
		return ""
	}
	if len(a) == 1 || len(b) == 1 {
		return KaratsubaMulSimple(a, b)
	}

	m := max(len(a), len(b))
	m2 := m / 2

	high1 := a[0:m2]
	low1 := a[m2:]

	high2 := b[0:m2]
	low2 := b[m2:]

	z0 := KaratsubaMul(low1, low2)
	z1 := KaratsubaAdd(KaratsubaMul(low2, high1), KaratsubaMul(low1, high2))
	z2 := KaratsubaMul(high1, high2)

	res1 := KaratsubaAdd((z2 + strings.Repeat("0", m)), z0)
	res1 = KaratsubaAdd(res1, (z1 + strings.Repeat("0", m2)))

	return res1
}
