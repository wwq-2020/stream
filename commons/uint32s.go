package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Uint32Slice uint32的Slice
type Uint32Slice struct {
	value []uint32
}

// ToUint32Slice uint32列表转为Uint32Slice
func ToUint32Slice(value []uint32) *Uint32Slice {
	return &Uint32Slice{value: value}
}

// Concat 拼接
func (s *Uint32Slice) Concat(given []uint32) *Uint32Slice {
	value := make([]uint32, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *Uint32Slice) Drop(n int) *Uint32Slice {
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
func (s *Uint32Slice) Filter(fn func(int, uint32) bool) *Uint32Slice {
	value := make([]uint32, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *Uint32Slice) First(value *uint32) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *Uint32Slice) Last(value *uint32) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *Uint32Slice) Map(fn func(int, uint32) uint32) *Uint32Slice {
	value := make([]uint32, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *Uint32Slice) Reduce(fn func(uint32, uint32, int) uint32, initial uint32) uint32 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *Uint32Slice) Reverse() *Uint32Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] > s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *Uint32Slice) Unique() *Uint32Slice {
	value := make([]uint32, 0, len(s.value))
	seen := make(map[uint32]struct{})
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
func (s *Uint32Slice) Append(given uint32) *Uint32Slice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *Uint32Slice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *Uint32Slice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *Uint32Slice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *Uint32Slice) Sort() *Uint32Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *Uint32Slice) All(fn func(int, uint32) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *Uint32Slice) Any(fn func(int, uint32) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *Uint32Slice) Paginate(size int) [][]uint32 {
	if size <= 0 {
		size = 1
	}
	var pages [][]uint32
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
func (s *Uint32Slice) Preappend(given uint32) *Uint32Slice {
	value := make([]uint32, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *Uint32Slice) Max(value *uint32) error {
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
func (s *Uint32Slice) Min(value *uint32) error {
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
func (s *Uint32Slice) Random(value *uint32) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *Uint32Slice) Shuffle() *Uint32Slice {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *Uint32Slice) Collect() []uint32 {
	return s.value
}
