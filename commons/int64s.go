package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Int64Slice int64的Slice
type Int64Slice struct {
	value []int64
}

// ToInt64Slice int64列表转为Int64Slice
func ToInt64Slice(value []int64) *Int64Slice {
	return &Int64Slice{value: value}
}

// Concat 拼接
func (s *Int64Slice) Concat(given []int64) *Int64Slice {
	value := make([]int64, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *Int64Slice) Drop(n int) *Int64Slice {
	if n < 0 {
		n = 0
	}
	l := len(s.value) - n
	if l < 0 {
		n = len(s.value)
	}
	s.value = s.value[n:]
	return s
}

// Filter 过滤
func (s *Int64Slice) Filter(fn func(int, int64) bool) *Int64Slice {
	value := make([]int64, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *Int64Slice) First(value *int64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *Int64Slice) Last(value *int64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *Int64Slice) Map(fn func(int, int64) int64) *Int64Slice {
	value := make([]int64, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *Int64Slice) Reduce(fn func(int64, int64, int) int64, initial int64) int64 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *Int64Slice) Reverse() *Int64Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] > s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *Int64Slice) Unique() *Int64Slice {
	value := make([]int64, 0, len(s.value))
	seen := make(map[int64]struct{})
	for _, each := range s.value {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	s.value = value
	return s
}

// Append 在尾部添加
func (s *Int64Slice) Append(given int64) *Int64Slice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *Int64Slice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *Int64Slice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *Int64Slice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *Int64Slice) Sort() *Int64Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *Int64Slice) All(fn func(int, int64) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *Int64Slice) Any(fn func(int, int64) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *Int64Slice) Paginate(size int) [][]int64 {
	if size <= 0 {
		size = 1
	}
	var pages [][]int64
	prev := -1
	for i := range s.value {
		if (i-prev) < size && i != (len(s.value)-1) {
			continue
		}
		pages = append(pages, s.value[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s *Int64Slice) Preappend(given int64) *Int64Slice {
	value := make([]int64, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *Int64Slice) Max(value *int64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if *value < each {
			*value  = each
		}
	}
	return nil 
}

// Min 获取最小元素
func (s *Int64Slice) Min(value *int64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if each < *value {
			*value = each
		}
	}
	return nil
}

// Random 随机获取一个元素
func (s *Int64Slice) Random(value *int64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *Int64Slice) Shuffle() *Int64Slice {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *Int64Slice) Collect() []int64 {
	return s.value
}
