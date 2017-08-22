package karatsuba

import "testing"

func TestKaratsubaAdd(t *testing.T) {
	cases := []struct {
		a, b, want string
	}{
		{"1", "1", "2"},
		{"10009", "1", "10010"},
		{"1", "10009", "10010"},
		{"1", "99", "100"},
	}
	for _, c := range cases {
		got := KaratsubaAdd(c.a, c.b)
		if got != c.want {
			t.Errorf("karatsuba_add(%q, %q) == %q, want %q", c.a, c.b, got, c.want)
		}
	}
}

func TestKaratsubaMulSimple(t *testing.T) {
	cases := []struct {
		a, b, want string
	}{
		{"1", "1", "1"},
		{"1", "2", "2"},
		{"2", "2", "4"},
		{"10", "3", "30"},
		{"111", "9", "999"},
		{"122", "9", "1098"},
	}
	for _, c := range cases {
		got := KaratsubaMulSimple(c.a, c.b)
		if got != c.want {
			t.Errorf("karatsuba_mul_simple(%q, %q) == %q, want %q", c.a, c.b, got, c.want)
		}
	}
}

func TestKaratsuba(t *testing.T) {
	cases := []struct {
		a, b, want string
	}{
		{"3141592653589793238462643383279502884197169399375105820974944592",
			"2718281828459045235360287471352662497757247093699959574966967627",
			"8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184"},
	}
	for _, c := range cases {
		got := KaratsubaMul(c.a, c.b)
		if got != c.want {
			t.Errorf("KaratsubaMul(%q, %q) == %q, want %q", c.a, c.b, got, c.want)
		}
	}
}
