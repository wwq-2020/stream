
package commons
import (
	"sort"
	"math/rand"
)
type Float32Stream struct{
	value	[]float32
	defaultReturn float32
}
func StreamOfFloat32(value []float32) *Float32Stream {
	return &Float32Stream{value:value,defaultReturn:0.0}
}
func(s *Float32Stream) OrElase(defaultReturn float32)  *Float32Stream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *Float32Stream) Concate(given []float32)  *Float32Stream {
	value := make([]float32, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *Float32Stream) Drop(n int)  *Float32Stream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *Float32Stream) Filter(fn func(int, float32)bool)  *Float32Stream {
	value := make([]float32, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *Float32Stream) First() float32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *Float32Stream) Last() float32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *Float32Stream) Map(fn func(int, float32)) *Float32Stream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *Float32Stream) Reduce(fn func(float32, float32, int) float32,initial float32) float32   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *Float32Stream) Reverse()  *Float32Stream {
	value := make([]float32, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *Float32Stream) Unique()  *Float32Stream{
	value := make([]float32, 0, len(s.value))
	seen:=make(map[float32]struct{})
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
func(s *Float32Stream) Append(given float32) *Float32Stream {
	s.value=append(s.value,given)
	return s
}
func(s *Float32Stream) Len() int {
	return len(s.value)
}
func(s *Float32Stream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *Float32Stream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *Float32Stream)  Sort()  *Float32Stream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i] < s.value[j]
	})
	return s 
}
func(s *Float32Stream) All(fn func(int, float32)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}
func(s *Float32Stream) Any(fn func(int, float32)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}
func(s *Float32Stream) Paginate(size int)  [][]float32 {
	var pages  [][]float32
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
func(s *Float32Stream) Pop() float32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *Float32Stream) Prepend(given float32) *Float32Stream {
	s.value = append([]float32{given},s.value...)
	return s
}
func(s *Float32Stream) Max() float32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max float32 = s.value[0]
	for _,each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func(s *Float32Stream) Min() float32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min float32 = s.value[0]
	for _,each := range s.value {
		if each  < min {
			min = each
		}
	}
	return min
}
func(s *Float32Stream) Random() float32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *Float32Stream) Shuffle() *Float32Stream {
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
func(s *Float32Stream) Collect() []float32{
	return s.value
}
type Float32PStream struct{
	value	[]*float32
	defaultReturn *float32
}
func PStreamOfFloat32(value []*float32) *Float32PStream {
	return &Float32PStream{value:value,defaultReturn:nil}
}
func(s *Float32PStream) OrElse(defaultReturn *float32)  *Float32PStream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *Float32PStream) Concate(given []*float32)  *Float32PStream {
	value := make([]*float32, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *Float32PStream) Drop(n int)  *Float32PStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *Float32PStream) Filter(fn func(int, *float32)bool)  *Float32PStream {
	value := make([]*float32, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *Float32PStream) First() *float32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *Float32PStream) Last() *float32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *Float32PStream) Map(fn func(int, *float32)) *Float32PStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *Float32PStream) Reduce(fn func(*float32, *float32, int) *float32,initial *float32) *float32   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *Float32PStream) Reverse()  *Float32PStream {
	value := make([]*float32, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *Float32PStream) Unique()  *Float32PStream{
	value := make([]*float32, 0, len(s.value))
	seen:=make(map[*float32]struct{})
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
func(s *Float32PStream) Append(given *float32) *Float32PStream {
	s.value=append(s.value,given)
	return s
}
func(s *Float32PStream) Len() int {
	return len(s.value)
}
func(s *Float32PStream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *Float32PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *Float32PStream)  Sort(less func(*float32,*float32) bool )  *Float32PStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}
func(s *Float32PStream) All(fn func(int, *float32)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}



func(s *Float32PStream) Any(fn func(int, *float32)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}



func(s *Float32PStream) Paginate(size int)  [][]*float32 {
	var pages  [][]*float32
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
func(s *Float32PStream) Pop() *float32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *Float32PStream) Prepend(given *float32) *Float32PStream {
	s.value = append([]*float32{given},s.value...)
	return s
}
func(s *Float32PStream) Max() *float32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *float32 = s.value[0]
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
func(s *Float32PStream) Min() *float32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *float32 = s.value[0]
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
func(s *Float32PStream) Random() *float32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *Float32PStream) Shuffle() *Float32PStream {
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
func(s *Float32PStream) Collect() []*float32{
	return s.value
}
