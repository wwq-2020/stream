package commons

import (
	"math/rand"
	"sort"
)

type Uint8Stream struct {
	value         []uint8
	defaultReturn uint8
}

func StreamOfUint8(value []uint8) *Uint8Stream {
	return &Uint8Stream{value: value, defaultReturn: 0}
}
func (s *Uint8Stream) OrElase(defaultReturn uint8) *Uint8Stream {
	s.defaultReturn = defaultReturn
	return s
}
func (s *Uint8Stream) Concate(given []uint8) *Uint8Stream {
	value := make([]uint8, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}
func (s *Uint8Stream) Drop(n int) *Uint8Stream {
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
func (s *Uint8Stream) Filter(fn func(int, uint8) bool) *Uint8Stream {
	value := make([]uint8, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
func (s *Uint8Stream) First() uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[0]
}
func (s *Uint8Stream) Last() uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[len(s.value)-1]
}
func (s *Uint8Stream) Map(fn func(int, uint8)) *Uint8Stream {
	for i, each := range s.value {
		fn(i, each)
	}
	return s
}
func (s *Uint8Stream) Reduce(fn func(uint8, uint8, int) uint8, initial uint8) uint8 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}
func (s *Uint8Stream) Reverse() *Uint8Stream {
	value := make([]uint8, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func (s *Uint8Stream) Unique() *Uint8Stream {
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
func (s *Uint8Stream) Append(given uint8) *Uint8Stream {
	s.value = append(s.value, given)
	return s
}
func (s *Uint8Stream) Len() int {
	return len(s.value)
}
func (s *Uint8Stream) IsEmpty() bool {
	return len(s.value) == 0
}
func (s *Uint8Stream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func (s *Uint8Stream) Sort() *Uint8Stream {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s
}
func (s *Uint8Stream) All(fn func(int, uint8) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}
func (s *Uint8Stream) Any(fn func(int, uint8) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}
func (s *Uint8Stream) Paginate(size int) [][]uint8 {
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
func (s *Uint8Stream) Pop() uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value) - 1
	val := s.value[lastIdx]
	s.value = s.value[:lastIdx]
	return val
}
func (s *Uint8Stream) Prepend(given uint8) *Uint8Stream {
	s.value = append([]uint8{given}, s.value...)
	return s
}
func (s *Uint8Stream) Max() uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max uint8 = s.value[0]
	for _, each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func (s *Uint8Stream) Min() uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min uint8 = s.value[0]
	for _, each := range s.value {
		if each < min {
			min = each
		}
	}
	return min
}
func (s *Uint8Stream) Random() uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func (s *Uint8Stream) Shuffle() *Uint8Stream {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i]
	})

	return s
}
func (s *Uint8Stream) Collect() []uint8 {
	return s.value
}

type Uint8PStream struct {
	value         []*uint8
	defaultReturn *uint8
}

func PStreamOfUint8(value []*uint8) *Uint8PStream {
	return &Uint8PStream{value: value, defaultReturn: nil}
}
func (s *Uint8PStream) OrElse(defaultReturn *uint8) *Uint8PStream {
	s.defaultReturn = defaultReturn
	return s
}
func (s *Uint8PStream) Concate(given []*uint8) *Uint8PStream {
	value := make([]*uint8, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}
func (s *Uint8PStream) Drop(n int) *Uint8PStream {
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
func (s *Uint8PStream) Filter(fn func(int, *uint8) bool) *Uint8PStream {
	value := make([]*uint8, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
func (s *Uint8PStream) First() *uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[0]
}
func (s *Uint8PStream) Last() *uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[len(s.value)-1]
}
func (s *Uint8PStream) Map(fn func(int, *uint8)) *Uint8PStream {
	for i, each := range s.value {
		fn(i, each)
	}
	return s
}
func (s *Uint8PStream) Reduce(fn func(*uint8, *uint8, int) *uint8, initial *uint8) *uint8 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}
func (s *Uint8PStream) Reverse() *Uint8PStream {
	value := make([]*uint8, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func (s *Uint8PStream) Unique() *Uint8PStream {
	value := make([]*uint8, 0, len(s.value))
	seen := make(map[*uint8]struct{})
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
func (s *Uint8PStream) Append(given *uint8) *Uint8PStream {
	s.value = append(s.value, given)
	return s
}
func (s *Uint8PStream) Len() int {
	return len(s.value)
}
func (s *Uint8PStream) IsEmpty() bool {
	return len(s.value) == 0
}
func (s *Uint8PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func (s *Uint8PStream) Sort(less func(*uint8, *uint8) bool) *Uint8PStream {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i], s.value[j])
	})
	return s
}
func (s *Uint8PStream) All(fn func(int, *uint8) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

func (s *Uint8PStream) Any(fn func(int, *uint8) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

func (s *Uint8PStream) Paginate(size int) [][]*uint8 {
	var pages [][]*uint8
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
func (s *Uint8PStream) Pop() *uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value) - 1
	val := s.value[lastIdx]
	s.value = s.value[:lastIdx]
	return val
}
func (s *Uint8PStream) Prepend(given *uint8) *Uint8PStream {
	s.value = append([]*uint8{given}, s.value...)
	return s
}
func (s *Uint8PStream) Max() *uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *uint8 = s.value[0]
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
func (s *Uint8PStream) Min() *uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *uint8 = s.value[0]
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
func (s *Uint8PStream) Random() *uint8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func (s *Uint8PStream) Shuffle() *Uint8PStream {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i]
	})

	return s
}
func (s *Uint8PStream) Collect() []*uint8 {
	return s.value
}
