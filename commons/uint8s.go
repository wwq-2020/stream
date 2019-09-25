package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Uint8Slice uint8的Slice
type Uint8Slice struct {
	value []uint8
}

// ToUint8Slice uint8列表转为Uint8Slice
func ToUint8Slice(value []uint8) *Uint8Slice {
	return &Uint8Slice{value: value}
}

// Concat 拼接
func (s *Uint8Slice) Concat(given []uint8) *Uint8Slice {
	value := make([]uint8, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *Uint8Slice) Drop(n int) *Uint8Slice {
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
func (s *Uint8Slice) Filter(fn func(int, uint8) bool) *Uint8Slice {
	value := make([]uint8, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *Uint8Slice) First(value *uint8) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *Uint8Slice) Last(value *uint8) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *Uint8Slice) Map(fn func(int, uint8) uint8) *Uint8Slice {
	value := make([]uint8, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *Uint8Slice) Reduce(fn func(uint8, uint8, int) uint8, initial uint8) uint8 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *Uint8Slice) Reverse() *Uint8Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] > s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *Uint8Slice) Unique() *Uint8Slice {
	value := make([]uint8, 0, len(s.value))
	seen := make(map[uint8]struct{})
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
func (s *Uint8Slice) Append(given uint8) *Uint8Slice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *Uint8Slice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *Uint8Slice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *Uint8Slice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *Uint8Slice) Sort() *Uint8Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *Uint8Slice) All(fn func(int, uint8) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *Uint8Slice) Any(fn func(int, uint8) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *Uint8Slice) Paginate(size int) [][]uint8 {
	if size <= 0 {
		size = 1
	}
	var pages [][]uint8
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
func (s *Uint8Slice) Preappend(given uint8) *Uint8Slice {
	value := make([]uint8, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *Uint8Slice) Max(value *uint8) error {
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
func (s *Uint8Slice) Min(value *uint8) error {
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
func (s *Uint8Slice) Random(value *uint8) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *Uint8Slice) Shuffle() *Uint8Slice {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *Uint8Slice) Collect() []uint8 {
	return s.value
}
