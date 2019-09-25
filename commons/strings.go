package commons

import (
	"errors"
	"math/rand"
	"sort"
)

// StringSlice string的Slice
type StringSlice []string

// Concat 拼接
func (s StringSlice) Concat(given []string) StringSlice {
	value := make([]string, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return StringSlice(value)
}

// Drop 丢弃前n个
func (s StringSlice) Drop(n int) StringSlice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return StringSlice(s[n:])
}

// Filter 过滤
func (s StringSlice) Filter(fn func(int, string) bool) StringSlice {
	value := make([]string, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return StringSlice(value)
}

// First 获取第一个元素
func (s StringSlice) First() (string, error) {
	if len(s) <= 0 {
		var defaultReturn string
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s StringSlice) Last() (string, error) {
	if len(s) <= 0 {
		var defaultReturn string
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s StringSlice) Map(fn func(int, string) string) StringSlice {
	value := make([]string, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return StringSlice(value)
}

// Reduce reduce
func (s StringSlice) Reduce(fn func(string, string, int) string, initial string) string {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s StringSlice) Reverse() StringSlice {
	value := make([]string, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return StringSlice(value)
}

// Unique 唯一
func (s StringSlice) Unique() StringSlice {
	value := make([]string, 0, len(s))
	seen := make(map[string]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}
		seen[each] = struct{}{}
		value = append(value, each)
	}
	return StringSlice(value)
}

// Append 在尾部添加
func (s StringSlice) Append(given string) StringSlice {
	return append(s, given)
}

// Len 获取长度
func (s StringSlice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s StringSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsEmpty 判断是否非空
func (s StringSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s StringSlice) Sort() StringSlice {
	value := make([]string, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return StringSlice(value)
}

// All 是否所有元素满足条件
func (s StringSlice) All(fn func(int, string) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s StringSlice) Any(fn func(int, string) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s StringSlice) Paginate(size int) [][]string {
	if size <= 0 {
		size = 1
	}
	var pages [][]string
	prev := -1
	for i := range s {
		if (i-prev) < size && i != (len(s)-1) {
			continue
		}
		pages = append(pages, s[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s StringSlice) Preappend(given string) StringSlice {
	value := make([]string, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return s
}

// Max 获取最大元素
func (s StringSlice) Max() (string, error) {
	if len(s) <= 0 {
		var defaultReturn string
		return defaultReturn, errors.New("empty")
	}
	max := s[0]
	for _, each := range s {
		if max < each {
			max = each
		}
	}
	return max, nil
}

// Min 获取最小元素
func (s StringSlice) Min() (string, error) {
	if len(s) <= 0 {
		var defaultReturn string
		return defaultReturn, errors.New("empty")
	}
	min := s[0]
	for _, each := range s {
		if each < min {
			min = each
		}
	}
	return min, nil
}

// Random 随机获取一个元素
func (s StringSlice) Random() (string, error) {
	if len(s) <= 0 {
		var defaultReturn string
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]string
func (s StringSlice) Shuffle() StringSlice {
	if len(s) <= 0 {
		return s
	}
	value := make([]string, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i]
	})
	return StringSlice(value)
}

// Collect 获取[]string
func (s StringSlice) Collect() []string {
	return s
}
