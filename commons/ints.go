
package commons

import (
	"sort"
	"math/rand"
)

type IntStream struct{
	value	[]int
	defaultReturn int
}

func StreamOfInt(value []int) *IntStream {
	return &IntStream{value:value,defaultReturn:0}
}

func(s *IntStream) OrElase(defaultReturn int)  *IntStream {
	s.defaultReturn = defaultReturn
	return s
}


func(s *IntStream) Concate(given []int)  *IntStream {
	value := make([]int, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}

func(s *IntStream) Drop(n int)  *IntStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}

func(s *IntStream) Filter(fn func(int, int)bool)  *IntStream {
	value := make([]int, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}

func(s *IntStream) First() int {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}

func(s *IntStream) Last() int {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}

func(s *IntStream) Map(fn func(int, int)) *IntStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}

func(s *IntStream) Reduce(fn func(int, int, int) int,initial int) int   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}

func(s *IntStream) Reverse()  *IntStream {
	value := make([]int, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

func(s *IntStream) Unique()  *IntStream{
	value := make([]int, 0, len(s.value))
	seen:=make(map[int]struct{})
	for _, each := range s.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	s.value = value
	return s
}

func(s *IntStream) Append(given int) *IntStream {
	s.value=append(s.value,given)
	return s
}

func(s *IntStream) Len() int {
	return len(s.value)
}

func(s *IntStream) IsEmpty() bool {
	return len(s.value) == 0
}

func(s *IntStream) IsNotEmpty() bool {
	return len(s.value) != 0
}

func(s *IntStream)  Sort()  *IntStream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i] < s.value[j]
	})
	return s 
}

func(s *IntStream) All(fn func(int, int)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(s *IntStream) Any(fn func(int, int)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(s *IntStream) Paginate(size int)  [][]int {
	var pages  [][]int
	prev := -1
	for i := range s.value {
		if (i-prev) < size-1 && i != (len(s.value)-1) {
			continue
		}
		pages=append(pages,s.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(s *IntStream) Pop() int{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}

func(s *IntStream) Prepend(given int) *IntStream {
	s.value = append([]int{given},s.value...)
	return s
}

func(s *IntStream) Max() int{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max int = s.value[0]
	for _,each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}


func(s *IntStream) Min() int{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min int = s.value[0]
	for _,each := range s.value {
		if each  < min {
			min = each
		}
	}
	return min
}

func(s *IntStream) Random() int{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}

func(s *IntStream) Shuffle() *IntStream {
	if len(s.value) <= 0 {
		return s
	}
	indexes := make([]int, len(s.value))
	for i := range s.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = 	s.value[j], s.value[i] 
	})
	
	return s
}

func(s *IntStream) Collect() []int{
	return s.value
}


type IntPStream struct{
	value	[]*int
	defaultReturn *int
}

func PStreamOfInt(value []*int) *IntPStream {
	return &IntPStream{value:value,defaultReturn:nil}
}

func(s *IntPStream) OrElse(defaultReturn *int)  *IntPStream {
	s.defaultReturn = defaultReturn
	return s
}

func(s *IntPStream) Concate(given []*int)  *IntPStream {
	value := make([]*int, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}

func(s *IntPStream) Drop(n int)  *IntPStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}

func(s *IntPStream) Filter(fn func(int, *int)bool)  *IntPStream {
	value := make([]*int, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}

func(s *IntPStream) First() *int {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}

func(s *IntPStream) Last() *int {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}

func(s *IntPStream) Map(fn func(int, *int)) *IntPStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}

func(s *IntPStream) Reduce(fn func(*int, *int, int) *int,initial *int) *int   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}

func(s *IntPStream) Reverse()  *IntPStream {
	value := make([]*int, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

func(s *IntPStream) Unique()  *IntPStream{
	value := make([]*int, 0, len(s.value))
	seen:=make(map[*int]struct{})
	for _, each := range s.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	s.value = value
	return s
}

func(s *IntPStream) Append(given *int) *IntPStream {
	s.value=append(s.value,given)
	return s
}

func(s *IntPStream) Len() int {
	return len(s.value)
}

func(s *IntPStream) IsEmpty() bool {
	return len(s.value) == 0
}

func(s *IntPStream) IsNotEmpty() bool {
	return len(s.value) != 0
}

func(s *IntPStream)  Sort(less func(*int,*int) bool )  *IntPStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}

func(s *IntPStream) All(fn func(int, *int)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(s *IntPStream) Any(fn func(int, *int)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(s *IntPStream) Paginate(size int)  [][]*int {
	var pages  [][]*int
	prev := -1
	for i := range s.value {
		if (i-prev) < size-1 && i != (len(s.value)-1) {
			continue
		}
		pages=append(pages,s.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(s *IntPStream) Pop() *int{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}

func(s *IntPStream) Prepend(given *int) *IntPStream {
	s.value = append([]*int{given},s.value...)
	return s
}

func(s *IntPStream) Max() *int{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *int = s.value[0]
	for _,each := range s.value {
		if max == nil{
			max = each
			continue
		}
		if each != nil && *max <= *each {
			max = each
		}
	}
	return max
}


func(s *IntPStream) Min() *int{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *int = s.value[0]
	for _,each := range s.value {
		if min == nil{
			min = each
			continue
		}
		if  each != nil && *each  <= *min {
			min = each
		}
	}
	return min
}

func(s *IntPStream) Random() *int{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}

func(s *IntPStream) Shuffle() *IntPStream {
	if len(s.value) <= 0 {
		return s
	}
	indexes := make([]int, len(s.value))
	for i := range s.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = 	s.value[j], s.value[i] 
	})
	
	return s
}

func(s *IntPStream) Collect() []*int{
	return s.value
}
