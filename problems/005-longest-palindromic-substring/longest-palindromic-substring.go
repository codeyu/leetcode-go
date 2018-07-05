package problem005
import "strings"
func longestPalindrome(s string) string {
    ss := []string{}
    for i := 0; i < len(s); i++ {
        ss = append(ss, string(s[i]))
    }
    t := strings.Join(ss, "#")
	t = "$#" + t + "#"
    p, id, mx, resId, resMx := make([]int,len(t)), 0, 0, 0, 0
    for i := 1; i < len(t); i++ {
        if mx > i {
        p[i] = min(p[2 * id - i], mx - i)
        }else{
            p[i] = 1
		} 
        for{
            if i+p[i]<len(t) && i-p[i]>=0 && t[i+p[i]]==t[i-p[i]]{
                p[i]++
            }else{
				break
			} 
		} 
        if mx < i + p[i] {
            mx = i + p[i]
            id = i
        }
        if resMx < p[i] {
            resMx = p[i]
            resId = i
        }
	}
    return substr(s, (resId - resMx) / 2, resMx - 1)
}

func min(first int, args... int) int {
    for _ , v := range args{
        if first > v {
            first = v
        }
    }
    return first
}

func substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
 
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
 
	if start > end {
		start, end = end, start
	}
 
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
