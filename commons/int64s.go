package commons

import (
	"math/rand"
	"sort"
)

type Int64Stream struct {
	value         []int64
	defaultReturn int64
}

func StreamOfInt64(value []int64) *Int64Stream {
	return &Int64Stream{value: value, defaultReturn: 0}
}
func (s *Int64Stream) OrElase(defaultReturn int64) *Int64Stream {
	s.defaultReturn = defaultReturn
	return s
}
func (s *Int64Stream) Concate(given []int64) *Int64Stream {
	value := make([]int64, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}
func (s *Int64Stream) Drop(n int) *Int64Stream {
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
func (s *Int64Stream) Filter(fn func(int, int64) bool) *Int64Stream {
	value := make([]int64, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
func (s *Int64Stream) First() int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[0]
}
func (s *Int64Stream) Last() int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[len(s.value)-1]
}
func (s *Int64Stream) Map(fn func(int, int64)) *Int64Stream {
	for i, each := range s.value {
		fn(i, each)
	}
	return s
}
func (s *Int64Stream) Reduce(fn func(int64, int64, int) int64, initial int64) int64 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}
func (s *Int64Stream) Reverse() *Int64Stream {
	value := make([]int64, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func (s *Int64Stream) Unique() *Int64Stream {
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
func (s *Int64Stream) Append(given int64) *Int64Stream {
	s.value = append(s.value, given)
	return s
}
func (s *Int64Stream) Len() int {
	return len(s.value)
}
func (s *Int64Stream) IsEmpty() bool {
	return len(s.value) == 0
}
func (s *Int64Stream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func (s *Int64Stream) Sort() *Int64Stream {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s
}
func (s *Int64Stream) All(fn func(int, int64) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}
func (s *Int64Stream) Any(fn func(int, int64) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}
func (s *Int64Stream) Paginate(size int) [][]int64 {
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
func (s *Int64Stream) Pop() int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value) - 1
	val := s.value[lastIdx]
	s.value = s.value[:lastIdx]
	return val
}
func (s *Int64Stream) Prepend(given int64) *Int64Stream {
	s.value = append([]int64{given}, s.value...)
	return s
}
func (s *Int64Stream) Max() int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max int64 = s.value[0]
	for _, each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func (s *Int64Stream) Min() int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min int64 = s.value[0]
	for _, each := range s.value {
		if each < min {
			min = each
		}
	}
	return min
}
func (s *Int64Stream) Random() int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func (s *Int64Stream) Shuffle() *Int64Stream {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i]
	})

	return s
}
func (s *Int64Stream) Collect() []int64 {
	return s.value
}

type Int64PStream struct {
	value         []*int64
	defaultReturn *int64
}

func PStreamOfInt64(value []*int64) *Int64PStream {
	return &Int64PStream{value: value, defaultReturn: nil}
}
func (s *Int64PStream) OrElse(defaultReturn *int64) *Int64PStream {
	s.defaultReturn = defaultReturn
	return s
}
func (s *Int64PStream) Concate(given []*int64) *Int64PStream {
	value := make([]*int64, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}
func (s *Int64PStream) Drop(n int) *Int64PStream {
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
func (s *Int64PStream) Filter(fn func(int, *int64) bool) *Int64PStream {
	value := make([]*int64, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
func (s *Int64PStream) First() *int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[0]
}
func (s *Int64PStream) Last() *int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[len(s.value)-1]
}
func (s *Int64PStream) Map(fn func(int, *int64)) *Int64PStream {
	for i, each := range s.value {
		fn(i, each)
	}
	return s
}
func (s *Int64PStream) Reduce(fn func(*int64, *int64, int) *int64, initial *int64) *int64 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}
func (s *Int64PStream) Reverse() *Int64PStream {
	value := make([]*int64, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func (s *Int64PStream) Unique() *Int64PStream {
	value := make([]*int64, 0, len(s.value))
	seen := make(map[*int64]struct{})
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
func (s *Int64PStream) Append(given *int64) *Int64PStream {
	s.value = append(s.value, given)
	return s
}
func (s *Int64PStream) Len() int {
	return len(s.value)
}
func (s *Int64PStream) IsEmpty() bool {
	return len(s.value) == 0
}
func (s *Int64PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func (s *Int64PStream) Sort(less func(*int64, *int64) bool) *Int64PStream {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i], s.value[j])
	})
	return s
}
func (s *Int64PStream) All(fn func(int, *int64) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

func (s *Int64PStream) Any(fn func(int, *int64) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

func (s *Int64PStream) Paginate(size int) [][]*int64 {
	var pages [][]*int64
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
func (s *Int64PStream) Pop() *int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value) - 1
	val := s.value[lastIdx]
	s.value = s.value[:lastIdx]
	return val
}
func (s *Int64PStream) Prepend(given *int64) *Int64PStream {
	s.value = append([]*int64{given}, s.value...)
	return s
}
func (s *Int64PStream) Max() *int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *int64 = s.value[0]
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
func (s *Int64PStream) Min() *int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *int64 = s.value[0]
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
func (s *Int64PStream) Random() *int64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func (s *Int64PStream) Shuffle() *Int64PStream {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i]
	})

	return s
}
func (s *Int64PStream) Collect() []*int64 {
	return s.value
}
