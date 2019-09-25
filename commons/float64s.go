package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Float64Slice float64的Slice
type Float64Slice struct {
	value []float64
}

// ToFloat64Slice float64列表转为Float64Slice
func ToFloat64Slice(value []float64) *Float64Slice {
	return &Float64Slice{value: value}
}

// Concat 拼接
func (s *Float64Slice) Concat(given []float64) *Float64Slice {
	value := make([]float64, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *Float64Slice) Drop(n int) *Float64Slice {
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
func (s *Float64Slice) Filter(fn func(int, float64) bool) *Float64Slice {
	value := make([]float64, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *Float64Slice) First(value *float64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *Float64Slice) Last(value *float64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *Float64Slice) Map(fn func(int, float64) float64) *Float64Slice {
	value := make([]float64, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *Float64Slice) Reduce(fn func(float64, float64, int) float64, initial float64) float64 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *Float64Slice) Reverse() *Float64Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] > s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *Float64Slice) Unique() *Float64Slice {
	value := make([]float64, 0, len(s.value))
	seen := make(map[float64]struct{})
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
func (s *Float64Slice) Append(given float64) *Float64Slice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *Float64Slice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *Float64Slice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *Float64Slice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *Float64Slice) Sort() *Float64Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *Float64Slice) All(fn func(int, float64) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *Float64Slice) Any(fn func(int, float64) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *Float64Slice) Paginate(size int) [][]float64 {
	if size <= 0 {
		size = 1
	}
	var pages [][]float64
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
func (s *Float64Slice) Preappend(given float64) *Float64Slice {
	value := make([]float64, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *Float64Slice) Max(value *float64) error {
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
func (s *Float64Slice) Min(value *float64) error {
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
func (s *Float64Slice) Random(value *float64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *Float64Slice) Shuffle() *Float64Slice {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *Float64Slice) Collect() []float64 {
	return s.value
}
