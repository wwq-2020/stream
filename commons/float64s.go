package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Float64Slice float64的Slice
type Float64Slice []float64

// Concat 拼接
func (s Float64Slice) Concat(given []float64) Float64Slice {
	value := make([]float64, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return Float64Slice(value)
}

// Drop 丢弃前n个
func (s Float64Slice) Drop(n int) Float64Slice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return Float64Slice(s[n:])
}

// Filter 过滤
func (s Float64Slice) Filter(fn func(int, float64) bool) Float64Slice {
	value := make([]float64, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return Float64Slice(value)
}

// First 获取第一个元素
func (s Float64Slice) First() (float64, error) {
	if len(s) <= 0 {
		var defaultReturn float64
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s Float64Slice) Last() (float64, error) {
	if len(s) <= 0 {
		var defaultReturn float64
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s Float64Slice) Map(fn func(int, float64) float64) Float64Slice {
	value := make([]float64, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return Float64Slice(value)
}

// Reduce reduce
func (s Float64Slice) Reduce(fn func(float64, float64, int) float64, initial float64) float64 {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s Float64Slice) Reverse() Float64Slice {
	value := make([]float64, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return Float64Slice(value)
}

// Unique 唯一
func (s Float64Slice) Unique() Float64Slice {
	value := make([]float64, 0, len(s))
	seen := make(map[float64]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return Float64Slice(value)
}

// Append 在尾部添加
func (s Float64Slice) Append(given float64) Float64Slice {
	return append(s, given)
}

// Len 获取长度
func (s Float64Slice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s Float64Slice) IsEmpty() bool {
	return len(s) == 0
}

// IsEmpty 判断是否非空
func (s Float64Slice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s Float64Slice) Sort() Float64Slice {
	value := make([]float64, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return Float64Slice(value)
}

// All 是否所有元素满足条件
func (s Float64Slice) All(fn func(int, float64) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s Float64Slice) Any(fn func(int, float64) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s Float64Slice) Paginate(size int) [][]float64 {
	if size <= 0 {
		size = 1
	}
	var pages [][]float64
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
func (s Float64Slice) Preappend(given float64) Float64Slice {
	value := make([]float64, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return s
}

// Max 获取最大元素
func (s Float64Slice) Max() (float64, error) {
	if len(s) <= 0 {
		var defaultReturn float64
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
func (s Float64Slice) Min() (float64, error) {
	if len(s) <= 0 {
		var defaultReturn float64
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
func (s Float64Slice) Random() (float64, error) {
	if len(s) <= 0 {
		var defaultReturn float64
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]float64
func (s Float64Slice) Shuffle() Float64Slice {
	if len(s) <= 0 {
		return s
	}
	value := make([]float64, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return Float64Slice(value)
}

// Collect 获取[]float64
func (s Float64Slice) Collect() []float64 {
	return s
}
