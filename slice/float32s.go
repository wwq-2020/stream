package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Float32Slice float32的Slice
type Float32Slice []float32

// Concat 拼接
func (s Float32Slice) Concat(given []float32) Float32Slice {
	value := make([]float32, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return Float32Slice(value)
}

// Drop 丢弃前n个
func (s Float32Slice) Drop(n int) Float32Slice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return Float32Slice(s[n:])
}

// Filter 过滤
func (s Float32Slice) Filter(fn func(int, float32) bool) Float32Slice {
	value := make([]float32, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return Float32Slice(value)
}

// First 获取第一个元素
func (s Float32Slice) First() (float32, error) {
	if len(s) <= 0 {
		var defaultReturn float32
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s Float32Slice) Last() (float32, error) {
	if len(s) <= 0 {
		var defaultReturn float32
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s Float32Slice) Map(fn func(int, float32) float32) Float32Slice {
	value := make([]float32, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return Float32Slice(value)
}

// Reduce reduce
func (s Float32Slice) Reduce(fn func(float32, float32, int) float32, initial float32) float32 {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s Float32Slice) Reverse() Float32Slice {
	value := make([]float32, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return Float32Slice(value)
}

// Unique 唯一
func (s Float32Slice) Unique() Float32Slice {
	value := make([]float32, 0, len(s))
	seen := make(map[float32]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return Float32Slice(value)
}

// Append 在尾部添加
func (s Float32Slice) Append(given float32) Float32Slice {
	return append(s, given)
}

// Len 获取长度
func (s Float32Slice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s Float32Slice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s Float32Slice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s Float32Slice) Sort() Float32Slice {
	value := make([]float32, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return Float32Slice(value)
}

// All 是否所有元素满足条件
func (s Float32Slice) All(fn func(int, float32) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s Float32Slice) Any(fn func(int, float32) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s Float32Slice) Paginate(size int) [][]float32 {
	if size <= 0 {
		size = 1
	}
	var pages [][]float32
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
func (s Float32Slice) Preappend(given float32) Float32Slice {
	value := make([]float32, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return Float32Slice(value)
}

// Max 获取最大元素
func (s Float32Slice) Max() (float32, error) {
	if len(s) <= 0 {
		var defaultReturn float32
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
func (s Float32Slice) Min() (float32, error) {
	if len(s) <= 0 {
		var defaultReturn float32
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
func (s Float32Slice) Random() (float32, error) {
	if len(s) <= 0 {
		var defaultReturn float32
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]float32
func (s Float32Slice) Shuffle() Float32Slice {
	if len(s) <= 0 {
		return s
	}
	value := make([]float32, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return Float32Slice(value)
}

// Collect 获取[]float32
func (s Float32Slice) Collect() []float32 {
	return s
}
