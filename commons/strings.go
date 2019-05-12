
package commons
import (
	"sort"
	"math/rand"
)
type StringStream struct{
	value	[]string
	defaultReturn string
}
func StreamOfString(value []string) *StringStream {
	return &StringStream{value:value,defaultReturn:""}
}
func(s *StringStream) OrElase(defaultReturn string)  *StringStream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *StringStream) Concate(given []string)  *StringStream {
	value := make([]string, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *StringStream) Drop(n int)  *StringStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *StringStream) Filter(fn func(int, string)bool)  *StringStream {
	value := make([]string, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *StringStream) First() string {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *StringStream) Last() string {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *StringStream) Map(fn func(int, string)) *StringStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *StringStream) Reduce(fn func(string, string, int) string,initial string) string   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *StringStream) Reverse()  *StringStream {
	value := make([]string, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *StringStream) Unique()  *StringStream{
	value := make([]string, 0, len(s.value))
	seen:=make(map[string]struct{})
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
func(s *StringStream) Append(given string) *StringStream {
	s.value=append(s.value,given)
	return s
}
func(s *StringStream) Len() int {
	return len(s.value)
}
func(s *StringStream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *StringStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *StringStream)  Sort()  *StringStream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i] < s.value[j]
	})
	return s 
}
func(s *StringStream) All(fn func(int, string)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}
func(s *StringStream) Any(fn func(int, string)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}
func(s *StringStream) Paginate(size int)  [][]string {
	var pages  [][]string
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
func(s *StringStream) Pop() string{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *StringStream) Prepend(given string) *StringStream {
	s.value = append([]string{given},s.value...)
	return s
}
func(s *StringStream) Max() string{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max string = s.value[0]
	for _,each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func(s *StringStream) Min() string{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min string = s.value[0]
	for _,each := range s.value {
		if each  < min {
			min = each
		}
	}
	return min
}
func(s *StringStream) Random() string{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *StringStream) Shuffle() *StringStream {
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
func(s *StringStream) Collect() []string{
	return s.value
}
type StringPStream struct{
	value	[]*string
	defaultReturn *string
}
func PStreamOfString(value []*string) *StringPStream {
	return &StringPStream{value:value,defaultReturn:nil}
}
func(s *StringPStream) OrElse(defaultReturn *string)  *StringPStream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *StringPStream) Concate(given []*string)  *StringPStream {
	value := make([]*string, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *StringPStream) Drop(n int)  *StringPStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *StringPStream) Filter(fn func(int, *string)bool)  *StringPStream {
	value := make([]*string, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *StringPStream) First() *string {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *StringPStream) Last() *string {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *StringPStream) Map(fn func(int, *string)) *StringPStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *StringPStream) Reduce(fn func(*string, *string, int) *string,initial *string) *string   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *StringPStream) Reverse()  *StringPStream {
	value := make([]*string, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *StringPStream) Unique()  *StringPStream{
	value := make([]*string, 0, len(s.value))
	seen:=make(map[*string]struct{})
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
func(s *StringPStream) Append(given *string) *StringPStream {
	s.value=append(s.value,given)
	return s
}
func(s *StringPStream) Len() int {
	return len(s.value)
}
func(s *StringPStream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *StringPStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *StringPStream)  Sort(less func(*string,*string) bool )  *StringPStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}
func(s *StringPStream) All(fn func(int, *string)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}



func(s *StringPStream) Any(fn func(int, *string)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}



func(s *StringPStream) Paginate(size int)  [][]*string {
	var pages  [][]*string
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
func(s *StringPStream) Pop() *string{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *StringPStream) Prepend(given *string) *StringPStream {
	s.value = append([]*string{given},s.value...)
	return s
}
func(s *StringPStream) Max() *string{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *string = s.value[0]
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
func(s *StringPStream) Min() *string{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *string = s.value[0]
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
func(s *StringPStream) Random() *string{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *StringPStream) Shuffle() *StringPStream {
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
func(s *StringPStream) Collect() []*string{
	return s.value
}
