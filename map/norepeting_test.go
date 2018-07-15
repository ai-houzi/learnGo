package main

import "testing"

func TestSubString(t *testing.T) {
	tests := []struct {
		s   string
		num int
	}{
		{"abcd", 4},
		{"abcdasdqwe", 6},
		{"这是我的测试", 6},
		{"一二三三二一", 3},
	}

	for _, test := range tests {
		str := LengthOfNoRepetingSubStr(test.s)

		if str != test.num {
			t.Errorf("字符串 %s  的测试结果为 %d；期望结果为%d ", test.s, str, test.num)
		}
	}
}

func BenchmarkLengthOfNoRepetingSubStr(b *testing.B) {
	s, ans := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("字符串长度为：%d",len(s))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str := LengthOfNoRepetingSubStr(s)
		if str != ans {
			b.Errorf("字符串 %s  的测试结果为 %d；期望结果为%d ", s, str, ans)
		}
	}

}
