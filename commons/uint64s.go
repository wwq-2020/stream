package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Uint64Slice uint64的Slice
type Uint64Slice struct {
	value []uint64
}

// ToUint64Slice uint64列表转为Uint64Slice
func ToUint64Slice(value []uint64) *Uint64Slice {
	return &Uint64Slice{value: value}
}

// Concat 拼接
func (s *Uint64Slice) Concat(given []uint64) *Uint64Slice {
	value := make([]uint64, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *Uint64Slice) Drop(n int) *Uint64Slice {
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
func (s *Uint64Slice) Filter(fn func(int, uint64) bool) *Uint64Slice {
	value := make([]uint64, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *Uint64Slice) First(value *uint64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *Uint64Slice) Last(value *uint64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *Uint64Slice) Map(fn func(int, uint64) uint64) *Uint64Slice {
	value := make([]uint64, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *Uint64Slice) Reduce(fn func(uint64, uint64, int) uint64, initial uint64) uint64 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *Uint64Slice) Reverse() *Uint64Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] > s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *Uint64Slice) Unique() *Uint64Slice {
	value := make([]uint64, 0, len(s.value))
	seen := make(map[uint64]struct{})
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
func (s *Uint64Slice) Append(given uint64) *Uint64Slice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *Uint64Slice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *Uint64Slice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *Uint64Slice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *Uint64Slice) Sort() *Uint64Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *Uint64Slice) All(fn func(int, uint64) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *Uint64Slice) Any(fn func(int, uint64) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *Uint64Slice) Paginate(size int) [][]uint64 {
	if size <= 0 {
		size = 1
	}
	var pages [][]uint64
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
func (s *Uint64Slice) Preappend(given uint64) *Uint64Slice {
	value := make([]uint64, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *Uint64Slice) Max(value *uint64) error {
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
func (s *Uint64Slice) Min(value *uint64) error {
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
func (s *Uint64Slice) Random(value *uint64) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *Uint64Slice) Shuffle() *Uint64Slice {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *Uint64Slice) Collect() []uint64 {
	return s.value
}
