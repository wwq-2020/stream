package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Float32Slice float32的Slice
type Float32Slice struct {
	value []float32
}

// ToFloat32Slice float32列表转为Float32Slice
func ToFloat32Slice(value []float32) *Float32Slice {
	return &Float32Slice{value: value}
}

// Concat 拼接
func (s *Float32Slice) Concat(given []float32) *Float32Slice {
	value := make([]float32, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *Float32Slice) Drop(n int) *Float32Slice {
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
func (s *Float32Slice) Filter(fn func(int, float32) bool) *Float32Slice {
	value := make([]float32, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *Float32Slice) First(value *float32) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *Float32Slice) Last(value *float32) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *Float32Slice) Map(fn func(int, float32) float32) *Float32Slice {
	value := make([]float32, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *Float32Slice) Reduce(fn func(float32, float32, int) float32, initial float32) float32 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *Float32Slice) Reverse() *Float32Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] > s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *Float32Slice) Unique() *Float32Slice {
	value := make([]float32, 0, len(s.value))
	seen := make(map[float32]struct{})
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
func (s *Float32Slice) Append(given float32) *Float32Slice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *Float32Slice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *Float32Slice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *Float32Slice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *Float32Slice) Sort() *Float32Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *Float32Slice) All(fn func(int, float32) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *Float32Slice) Any(fn func(int, float32) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *Float32Slice) Paginate(size int) [][]float32 {
	if size <= 0 {
		size = 1
	}
	var pages [][]float32
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
func (s *Float32Slice) Preappend(given float32) *Float32Slice {
	value := make([]float32, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *Float32Slice) Max(value *float32) error {
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
func (s *Float32Slice) Min(value *float32) error {
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
func (s *Float32Slice) Random(value *float32) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *Float32Slice) Shuffle() *Float32Slice {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *Float32Slice) Collect() []float32 {
	return s.value
}
