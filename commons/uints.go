package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// UintSlice uint的Slice
type UintSlice struct {
	value []uint
}

// ToUintSlice uint列表转为UintSlice
func ToUintSlice(value []uint) *UintSlice {
	return &UintSlice{value: value}
}

// Concat 拼接
func (s *UintSlice) Concat(given []uint) *UintSlice {
	value := make([]uint, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *UintSlice) Drop(n int) *UintSlice {
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
func (s *UintSlice) Filter(fn func(int, uint) bool) *UintSlice {
	value := make([]uint, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *UintSlice) First(value *uint) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *UintSlice) Last(value *uint) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *UintSlice) Map(fn func(int, uint) uint) *UintSlice {
	value := make([]uint, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *UintSlice) Reduce(fn func(uint, uint, int) uint, initial uint) uint {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *UintSlice) Reverse() *UintSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] > s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *UintSlice) Unique() *UintSlice {
	value := make([]uint, 0, len(s.value))
	seen := make(map[uint]struct{})
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
func (s *UintSlice) Append(given uint) *UintSlice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *UintSlice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *UintSlice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *UintSlice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *UintSlice) Sort() *UintSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *UintSlice) All(fn func(int, uint) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *UintSlice) Any(fn func(int, uint) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *UintSlice) Paginate(size int) [][]uint {
	if size <= 0 {
		size = 1
	}
	var pages [][]uint
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
func (s *UintSlice) Preappend(given uint) *UintSlice {
	value := make([]uint, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *UintSlice) Max(value *uint) error {
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
func (s *UintSlice) Min(value *uint) error {
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
func (s *UintSlice) Random(value *uint) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *UintSlice) Shuffle() *UintSlice {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *UintSlice) Collect() []uint {
	return s.value
}
