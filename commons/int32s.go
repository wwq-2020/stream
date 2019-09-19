package commons

import (
	"math/rand"
	"sort"
)

type Int32Stream struct {
	value         []int32
	defaultReturn int32
}

func StreamOfInt32(value []int32) *Int32Stream {
	return &Int32Stream{value: value, defaultReturn: 0}
}
func (s *Int32Stream) OrElase(defaultReturn int32) *Int32Stream {
	s.defaultReturn = defaultReturn
	return s
}
func (s *Int32Stream) Concate(given []int32) *Int32Stream {
	value := make([]int32, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}
func (s *Int32Stream) Drop(n int) *Int32Stream {
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
func (s *Int32Stream) Filter(fn func(int, int32) bool) *Int32Stream {
	value := make([]int32, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
func (s *Int32Stream) First() int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[0]
}
func (s *Int32Stream) Last() int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[len(s.value)-1]
}
func (s *Int32Stream) Map(fn func(int, int32)) *Int32Stream {
	for i, each := range s.value {
		fn(i, each)
	}
	return s
}
func (s *Int32Stream) Reduce(fn func(int32, int32, int) int32, initial int32) int32 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}
func (s *Int32Stream) Reverse() *Int32Stream {
	value := make([]int32, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func (s *Int32Stream) Unique() *Int32Stream {
	value := make([]int32, 0, len(s.value))
	seen := make(map[int32]struct{})
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
func (s *Int32Stream) Append(given int32) *Int32Stream {
	s.value = append(s.value, given)
	return s
}
func (s *Int32Stream) Len() int {
	return len(s.value)
}
func (s *Int32Stream) IsEmpty() bool {
	return len(s.value) == 0
}
func (s *Int32Stream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func (s *Int32Stream) Sort() *Int32Stream {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s
}
func (s *Int32Stream) All(fn func(int, int32) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}
func (s *Int32Stream) Any(fn func(int, int32) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}
func (s *Int32Stream) Paginate(size int) [][]int32 {
	var pages [][]int32
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
func (s *Int32Stream) Pop() int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value) - 1
	val := s.value[lastIdx]
	s.value = s.value[:lastIdx]
	return val
}
func (s *Int32Stream) Prepend(given int32) *Int32Stream {
	s.value = append([]int32{given}, s.value...)
	return s
}
func (s *Int32Stream) Max() int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max int32 = s.value[0]
	for _, each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func (s *Int32Stream) Min() int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min int32 = s.value[0]
	for _, each := range s.value {
		if each < min {
			min = each
		}
	}
	return min
}
func (s *Int32Stream) Random() int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func (s *Int32Stream) Shuffle() *Int32Stream {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i]
	})

	return s
}
func (s *Int32Stream) Collect() []int32 {
	return s.value
}

type Int32PStream struct {
	value         []*int32
	defaultReturn *int32
}

func PStreamOfInt32(value []*int32) *Int32PStream {
	return &Int32PStream{value: value, defaultReturn: nil}
}
func (s *Int32PStream) OrElse(defaultReturn *int32) *Int32PStream {
	s.defaultReturn = defaultReturn
	return s
}
func (s *Int32PStream) Concate(given []*int32) *Int32PStream {
	value := make([]*int32, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}
func (s *Int32PStream) Drop(n int) *Int32PStream {
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
func (s *Int32PStream) Filter(fn func(int, *int32) bool) *Int32PStream {
	value := make([]*int32, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
func (s *Int32PStream) First() *int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[0]
}
func (s *Int32PStream) Last() *int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[len(s.value)-1]
}
func (s *Int32PStream) Map(fn func(int, *int32)) *Int32PStream {
	for i, each := range s.value {
		fn(i, each)
	}
	return s
}
func (s *Int32PStream) Reduce(fn func(*int32, *int32, int) *int32, initial *int32) *int32 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}
func (s *Int32PStream) Reverse() *Int32PStream {
	value := make([]*int32, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func (s *Int32PStream) Unique() *Int32PStream {
	value := make([]*int32, 0, len(s.value))
	seen := make(map[*int32]struct{})
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
func (s *Int32PStream) Append(given *int32) *Int32PStream {
	s.value = append(s.value, given)
	return s
}
func (s *Int32PStream) Len() int {
	return len(s.value)
}
func (s *Int32PStream) IsEmpty() bool {
	return len(s.value) == 0
}
func (s *Int32PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func (s *Int32PStream) Sort(less func(*int32, *int32) bool) *Int32PStream {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i], s.value[j])
	})
	return s
}
func (s *Int32PStream) All(fn func(int, *int32) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

func (s *Int32PStream) Any(fn func(int, *int32) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

func (s *Int32PStream) Paginate(size int) [][]*int32 {
	var pages [][]*int32
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
func (s *Int32PStream) Pop() *int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value) - 1
	val := s.value[lastIdx]
	s.value = s.value[:lastIdx]
	return val
}
func (s *Int32PStream) Prepend(given *int32) *Int32PStream {
	s.value = append([]*int32{given}, s.value...)
	return s
}
func (s *Int32PStream) Max() *int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *int32 = s.value[0]
	for _, each := range s.value {
		if max == nil {
			max = each
			continue
		}
		if each != nil && *max <= *each {
			max = each
		}
	}
	return max
}
func (s *Int32PStream) Min() *int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *int32 = s.value[0]
	for _, each := range s.value {
		if min == nil {
			min = each
			continue
		}
		if each != nil && *each <= *min {
			min = each
		}
	}
	return min
}
func (s *Int32PStream) Random() *int32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func (s *Int32PStream) Shuffle() *Int32PStream {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i]
	})

	return s
}
func (s *Int32PStream) Collect() []*int32 {
	return s.value
}
