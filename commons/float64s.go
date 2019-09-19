package commons

import (
	"math/rand"
	"sort"
)

type Float64Stream struct {
	value         []float64
	defaultReturn float64
}

func StreamOfFloat64(value []float64) *Float64Stream {
	return &Float64Stream{value: value, defaultReturn: 0.0}
}
func (s *Float64Stream) OrElase(defaultReturn float64) *Float64Stream {
	s.defaultReturn = defaultReturn
	return s
}
func (s *Float64Stream) Concate(given []float64) *Float64Stream {
	value := make([]float64, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}
func (s *Float64Stream) Drop(n int) *Float64Stream {
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
func (s *Float64Stream) Filter(fn func(int, float64) bool) *Float64Stream {
	value := make([]float64, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
func (s *Float64Stream) First() float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[0]
}
func (s *Float64Stream) Last() float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[len(s.value)-1]
}
func (s *Float64Stream) Map(fn func(int, float64)) *Float64Stream {
	for i, each := range s.value {
		fn(i, each)
	}
	return s
}
func (s *Float64Stream) Reduce(fn func(float64, float64, int) float64, initial float64) float64 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}
func (s *Float64Stream) Reverse() *Float64Stream {
	value := make([]float64, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func (s *Float64Stream) Unique() *Float64Stream {
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
func (s *Float64Stream) Append(given float64) *Float64Stream {
	s.value = append(s.value, given)
	return s
}
func (s *Float64Stream) Len() int {
	return len(s.value)
}
func (s *Float64Stream) IsEmpty() bool {
	return len(s.value) == 0
}
func (s *Float64Stream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func (s *Float64Stream) Sort() *Float64Stream {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s
}
func (s *Float64Stream) All(fn func(int, float64) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}
func (s *Float64Stream) Any(fn func(int, float64) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}
func (s *Float64Stream) Paginate(size int) [][]float64 {
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
func (s *Float64Stream) Pop() float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value) - 1
	val := s.value[lastIdx]
	s.value = s.value[:lastIdx]
	return val
}
func (s *Float64Stream) Prepend(given float64) *Float64Stream {
	s.value = append([]float64{given}, s.value...)
	return s
}
func (s *Float64Stream) Max() float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max float64 = s.value[0]
	for _, each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func (s *Float64Stream) Min() float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min float64 = s.value[0]
	for _, each := range s.value {
		if each < min {
			min = each
		}
	}
	return min
}
func (s *Float64Stream) Random() float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func (s *Float64Stream) Shuffle() *Float64Stream {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i]
	})

	return s
}
func (s *Float64Stream) Collect() []float64 {
	return s.value
}

type Float64PStream struct {
	value         []*float64
	defaultReturn *float64
}

func PStreamOfFloat64(value []*float64) *Float64PStream {
	return &Float64PStream{value: value, defaultReturn: nil}
}
func (s *Float64PStream) OrElse(defaultReturn *float64) *Float64PStream {
	s.defaultReturn = defaultReturn
	return s
}
func (s *Float64PStream) Concate(given []*float64) *Float64PStream {
	value := make([]*float64, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}
func (s *Float64PStream) Drop(n int) *Float64PStream {
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
func (s *Float64PStream) Filter(fn func(int, *float64) bool) *Float64PStream {
	value := make([]*float64, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
func (s *Float64PStream) First() *float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[0]
}
func (s *Float64PStream) Last() *float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	return s.value[len(s.value)-1]
}
func (s *Float64PStream) Map(fn func(int, *float64)) *Float64PStream {
	for i, each := range s.value {
		fn(i, each)
	}
	return s
}
func (s *Float64PStream) Reduce(fn func(*float64, *float64, int) *float64, initial *float64) *float64 {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}
func (s *Float64PStream) Reverse() *Float64PStream {
	value := make([]*float64, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func (s *Float64PStream) Unique() *Float64PStream {
	value := make([]*float64, 0, len(s.value))
	seen := make(map[*float64]struct{})
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
func (s *Float64PStream) Append(given *float64) *Float64PStream {
	s.value = append(s.value, given)
	return s
}
func (s *Float64PStream) Len() int {
	return len(s.value)
}
func (s *Float64PStream) IsEmpty() bool {
	return len(s.value) == 0
}
func (s *Float64PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func (s *Float64PStream) Sort(less func(*float64, *float64) bool) *Float64PStream {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i], s.value[j])
	})
	return s
}
func (s *Float64PStream) All(fn func(int, *float64) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

func (s *Float64PStream) Any(fn func(int, *float64) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

func (s *Float64PStream) Paginate(size int) [][]*float64 {
	var pages [][]*float64
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
func (s *Float64PStream) Pop() *float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value) - 1
	val := s.value[lastIdx]
	s.value = s.value[:lastIdx]
	return val
}
func (s *Float64PStream) Prepend(given *float64) *Float64PStream {
	s.value = append([]*float64{given}, s.value...)
	return s
}
func (s *Float64PStream) Max() *float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *float64 = s.value[0]
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
func (s *Float64PStream) Min() *float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *float64 = s.value[0]
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
func (s *Float64PStream) Random() *float64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func (s *Float64PStream) Shuffle() *Float64PStream {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i]
	})

	return s
}
func (s *Float64PStream) Collect() []*float64 {
	return s.value
}
