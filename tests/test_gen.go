package tests

import (
	"errors"
	"math/rand"
	"sort"
		
	commons "github.com/wwq1988/stream/commons"						
					
	"github.com/wwq1988/stream/outter"
)

// SomeSlice Some的Slice
type SomeSlice struct {
	value []Some
}

// ToSomeSlice Some列表转成SomeSlice
func ToSomeSlice(value []Some) *SomeSlice {
	return &SomeSlice{value: value}
}

// Concat 拼接
func (s *SomeSlice) Concat(given []Some) *SomeSlice {
	value := make([]Some, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *SomeSlice) Drop(n int) *SomeSlice {
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
func (s *SomeSlice) Filter(fn func(int, Some) bool) *SomeSlice {
	value := make([]Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}


// FilterByA 通过过滤器过滤
func (s *SomeSlice) FilterByA(fn func(int, string) bool) *SomeSlice {
	value := make([]Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.A) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// FilterByB 通过过滤器过滤
func (s *SomeSlice) FilterByB(fn func(int, string) bool) *SomeSlice {
	value := make([]Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.B) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// FilterByC 通过过滤器过滤
func (s *SomeSlice) FilterByC(fn func(int, *Some) bool) *SomeSlice {
	value := make([]Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.C) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// FilterByD 通过过滤器过滤
func (s *SomeSlice) FilterByD(fn func(int, *outter.Some) bool) *SomeSlice {
	value := make([]Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.D) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}


// First 获取第一个元素
func (s *SomeSlice) First(value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *SomeSlice) Last(value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *SomeSlice) Map(fn func(int, Some) Some) *SomeSlice {
	value := make([]Some, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *SomeSlice) Reduce(fn func(Some, Some, int) Some, initial Some) Some {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *SomeSlice) Reverse() *SomeSlice {
	value := make([]Some, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}



// UniqueByA 通过A唯一
func (s *SomeSlice) UniqueByA() *SomeSlice {
	value := make([]Some, 0, len(s.value))
	seen := make(map[string]struct{})
	for _, each := range s.value {
		if _, dup := seen[each.A]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.A] = struct{}{}	
	}
	s.value = value
	return s
}



// UniqueByB 通过B唯一
func (s *SomeSlice) UniqueByB() *SomeSlice {
	value := make([]Some, 0, len(s.value))
	seen := make(map[string]struct{})
	for _, each := range s.value {
		if _, dup := seen[each.B]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.B] = struct{}{}	
	}
	s.value = value
	return s
}




// UniqueByC 通过C唯一
func (s *SomeSlice) UniqueByC(compare func(*Some, *Some) bool) *SomeSlice {
	value := make([]Some, 0, len(s.value))
	seen := make(map[int]struct{})
	for i, outter := range s.value {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s.value {
			if i == j {
				continue
			}
			if compare(inner.C, outter.C) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value, outter)			
	}
	s.value = value
	return s
}





// UniqueByD 通过D唯一
func (s *SomeSlice) UniqueByD(compare func(*outter.Some, *outter.Some) bool) *SomeSlice {
	value := make([]Some, 0, len(s.value))
	seen := make(map[int]struct{})
	for i, outter := range s.value {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s.value {
			if i == j {
				continue
			}
			if compare(inner.D, outter.D) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value, outter)			
	}
	s.value = value
	return s
}




// Append 在尾部添加元素
func (s *SomeSlice) Append(given Some) *SomeSlice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *SomeSlice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *SomeSlice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsNotEmpty 判断是否非空
func (s *SomeSlice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// All 是否所有元素满足添加
func (s *SomeSlice) All(fn func(int, Some) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *SomeSlice) Any(fn func(int, Some) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *SomeSlice) Paginate(size int) [][]Some {
	if size <= 0 {
		size = 1
	}
	var pages [][]Some
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
func (s *SomeSlice) Preappend(given Some) *SomeSlice {
	value := make([]Some, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最后元素
func (s *SomeSlice) Max(bigger func(Some, Some) bool, value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if bigger(each, *value) {
			*value = each
		}
	}
	return nil
}

// Min 获取最小元素
func (s *SomeSlice) Min(less func(Some, Some) bool, value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if less(each, *value) {
			*value = each
		}
	}
	return nil
}

// Random 随机获取一个元素
func (s *SomeSlice) Random(value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *SomeSlice) Shuffle() *SomeSlice {
	if len(s.value) <= 0 {
		return s
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	return s
}



// SortByA 根据A排序
func (s *SomeSlice) SortByA() *SomeSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i].A < s.value[j].A
	})
	return s 
}



// SortByB 根据B排序
func (s *SomeSlice) SortByB() *SomeSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i].B < s.value[j].B
	})
	return s 
}



// SortByC 根据C排序
func (s *SomeSlice) SortByC(less func(*Some, *Some) bool) *SomeSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i].C, s.value[j].C)
	})
	return s 
}



// SortByD 根据D排序
func (s *SomeSlice) SortByD(less func(*outter.Some, *outter.Some) bool) *SomeSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i].D, s.value[j].D)
	})
	return s 
}








// APSlice 获取A的Slice
func (s *SomeSlice) ASlice() *commons.StringSlice {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.A)
	}
	newSlice := commons.ToStringSlice(value)
	return newSlice
}





// BPSlice 获取B的Slice
func (s *SomeSlice) BSlice() *commons.StringSlice {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.B)
	}
	newSlice := commons.ToStringSlice(value)
	return newSlice
}





// CPSlice 获取C的PSlice
func (s *SomeSlice) CPSlice() *SomePSlice {	
	value := make([]*Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.C)
	}
	newSlice := ToSomePSlice(value)
	return newSlice
}





// DPSlice 获取D的PSlice
func (s *SomeSlice) DPSlice() *outter.SomePSlice {	
	value := make([]*outter.Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.D)
	}
	newSlice := outter.ToSomePSlice(value)
	return newSlice
}





// As 获取A的列表
func (s *SomeSlice) As() []string {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.A)
	}
	return value
}

// Bs 获取B的列表
func (s *SomeSlice) Bs() []string {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.B)
	}
	return value
}

// Cs 获取C的列表
func (s *SomeSlice) Cs() []*Some {	
	value := make([]*Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.C)
	}
	return value
}

// Ds 获取D的列表
func (s *SomeSlice) Ds() []*outter.Some {	
	value := make([]*outter.Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.D)
	}
	return value
}


// Collect 获取最终的列表
func (s *SomeSlice) Collect() []Some {
	return s.value
}
	
// SomePSlice	Some的PSlice		
type SomePSlice struct {
	value []*Some
}

// ToSomePSlice Some的指针列表转成SomePSlice 
func ToSomePSlice(value []*Some) *SomePSlice {
	return &SomePSlice{value: value}
}

// Concat 拼接
func (s *SomePSlice) Concat(given []*Some) *SomePSlice {
	value := make([]*Some, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *SomePSlice) Drop(n int) *SomePSlice {
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
func (s *SomePSlice) Filter(fn func(int, *Some) bool) *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}


// FilterByA 通过过滤器过滤
func (s *SomePSlice) FilterByA(fn func(int, string) bool) *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.A) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// FilterByB 通过过滤器过滤
func (s *SomePSlice) FilterByB(fn func(int, string) bool) *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.B) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// FilterByC 通过过滤器过滤
func (s *SomePSlice) FilterByC(fn func(int, *Some) bool) *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.C) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// FilterByD 通过过滤器过滤
func (s *SomePSlice) FilterByD(fn func(int, *outter.Some) bool) *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.D) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}


// First 获取第一个元素
func (s *SomePSlice) First(value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = *s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *SomePSlice) Last(value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = *s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *SomePSlice) Map(fn func(int, *Some) *Some) *SomePSlice {
	value := make([]*Some, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *SomePSlice) Reduce(fn func(*Some, *Some, int) *Some, initial *Some) *Some {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *SomePSlice) Reverse() *SomePSlice {
	value := make([]*Some, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

// UniqueBy 通过比较器唯一
func (s *SomePSlice) UniqueBy(compare func(*Some, *Some)bool) *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	seen := make(map[int]struct{})
	for i, outter := range s.value {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s.value {
			if i == j {
				continue
			}
			if compare(inner, outter) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value, outter)			
	}
	s.value = value
	return s
}

// Append 在尾部添加
func (s *SomePSlice) Append(given *Some) *SomePSlice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *SomePSlice) Len() int {
	return len(s.value)
}

// IsEmpty 是否为空
func (s *SomePSlice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsNotEmpty 是否非空
func (s *SomePSlice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// SortBy 根据比较器排序
func (s *SomePSlice) SortBy(less func(*Some, *Some) bool) *SomePSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i], s.value[j])
	})
	
	return s 
}

// All 是否所有元素满足条件
func (s *SomePSlice) All(fn func(int, *Some) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}


// AllByA 是否所有元素的A满足条件
func (s *SomePSlice) AllByA(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.A){
			return false
		}
	}
	return true
}

// AllByB 是否所有元素的B满足条件
func (s *SomePSlice) AllByB(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.B){
			return false
		}
	}
	return true
}

// AllByC 是否所有元素的C满足条件
func (s *SomePSlice) AllByC(fn func(int, *Some) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.C){
			return false
		}
	}
	return true
}

// AllByD 是否所有元素的D满足条件
func (s *SomePSlice) AllByD(fn func(int, *outter.Some) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.D){
			return false
		}
	}
	return true
}




// AllByA 是否所有元素的A满足条件
func (s *SomeSlice) AllByA(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.A){
			return false
		}
	}
	return true
}

// AllByB 是否所有元素的B满足条件
func (s *SomeSlice) AllByB(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.B){
			return false
		}
	}
	return true
}

// AllByC 是否所有元素的C满足条件
func (s *SomeSlice) AllByC(fn func(int, *Some) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.C){
			return false
		}
	}
	return true
}

// AllByD 是否所有元素的D满足条件
func (s *SomeSlice) AllByD(fn func(int, *outter.Some) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.D){
			return false
		}
	}
	return true
}


// Any 是否有元素满足条件
func (s *SomePSlice) Any(fn func(int, *Some) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}



// AnyByA 是否有元素的A满足条件
func (s *SomePSlice) AnyByA(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if fn(i, each.A) {
			return true
		}
	}
	return false
}

// AnyByB 是否有元素的B满足条件
func (s *SomePSlice) AnyByB(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if fn(i, each.B) {
			return true
		}
	}
	return false
}

// AnyByC 是否有元素的C满足条件
func (s *SomePSlice) AnyByC(fn func(int, *Some) bool) bool {
	for i, each := range s.value {
		if fn(i, each.C) {
			return true
		}
	}
	return false
}

// AnyByD 是否有元素的D满足条件
func (s *SomePSlice) AnyByD(fn func(int, *outter.Some) bool) bool {
	for i, each := range s.value {
		if fn(i, each.D) {
			return true
		}
	}
	return false
}



// AnyByA 是否有元素的A满足条件
func (s *SomeSlice) AnyByA(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if fn(i, each.A) {
			return true
		}
	}
	return false
}

// AnyByB 是否有元素的B满足条件
func (s *SomeSlice) AnyByB(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if fn(i, each.B) {
			return true
		}
	}
	return false
}

// AnyByC 是否有元素的C满足条件
func (s *SomeSlice) AnyByC(fn func(int, *Some) bool) bool {
	for i, each := range s.value {
		if fn(i, each.C) {
			return true
		}
	}
	return false
}

// AnyByD 是否有元素的D满足条件
func (s *SomeSlice) AnyByD(fn func(int, *outter.Some) bool) bool {
	for i, each := range s.value {
		if fn(i, each.D) {
			return true
		}
	}
	return false
}


// Paginate 分页
func (s *SomePSlice) Paginate(size int) [][]*Some {
	if size <= 0 {
		size = 1
	}
	var pages [][]*Some
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
func (s *SomePSlice) Preappend(given *Some) *SomePSlice {
	value := make([]*Some, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *SomePSlice) Max(bigger func(*Some, *Some) bool, value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = *s.value[0]
	for _, each := range s.value {
		if bigger(each, value) {
			*value = *each
		}
	}
	return nil
}

// Min 获取最小元素
func (s *SomePSlice) Min(less func(*Some, *Some) bool, value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = *s.value[0]
	for _, each := range s.value {
		if less(each, value) {
			*value = *each
		}
	}
	return nil
}

// Random 随机获取元素
func (s *SomePSlice) Random(value *Some) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = *s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *SomePSlice) Shuffle() *SomePSlice {
	if len(s.value) <= 0 {
		return s
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}



// SortByA 根据元素的A排序
func (s *SomePSlice) SortByA() *SomePSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i].A < s.value[j].A
	})
	return s 
}



// SortByB 根据元素的B排序
func (s *SomePSlice) SortByB() *SomePSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i].B < s.value[j].B
	})
	return s 
}



// SortByC 根据元素的C和比较器排序
func (s *SomePSlice) SortByC(less func(*Some, *Some) bool) *SomePSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i].C, s.value[j].C)
	})
	return s 
}



// SortByD 根据元素的D和比较器排序
func (s *SomePSlice) SortByD(less func(*outter.Some, *outter.Some) bool) *SomePSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i].D, s.value[j].D)
	})
	return s 
}





// UniqueByA 根据元素的A唯一
func (s *SomePSlice) UniqueByA() *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[string]struct{})
	for _, each := range s.value {
		if _, dup := seen[each.A]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.A] = struct{}{}	
	}
	s.value = value
	return s
}



// UniqueByB 根据元素的B唯一
func (s *SomePSlice) UniqueByB() *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[string]struct{})
	for _, each := range s.value {
		if _, dup := seen[each.B]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.B] = struct{}{}	
	}
	s.value = value
	return s
}




// UniqueByC 根据元素的C和比较器唯一
func (s *SomePSlice) UniqueByC(compare func (*Some, *Some) bool) *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j,inner :=range s.value {
			if i == j {
				continue
			}
			if compare(inner.C, outter.C) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value, outter)			
	}
	s.value = value
	
	return s
}





// UniqueByD 根据元素的D和比较器唯一
func (s *SomePSlice) UniqueByD(compare func (*outter.Some, *outter.Some) bool) *SomePSlice {
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j,inner :=range s.value {
			if i == j {
				continue
			}
			if compare(inner.D, outter.D) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value, outter)			
	}
	s.value = value
	
	return s
}







// ASlice 获取A的Slice
func (s *SomePSlice) ASlice() *commons.StringSlice {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.A)
	}
	newSlice := commons.ToStringSlice(value)
	return newSlice
}





// BSlice 获取B的Slice
func (s *SomePSlice) BSlice() *commons.StringSlice {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.B)
	}
	newSlice := commons.ToStringSlice(value)
	return newSlice
}





// CPSlice 获取C的PSlice
func (s *SomePSlice) CPSlice() *SomePSlice {	
	value := make([]*Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.C)
	}
	newSlice := ToSomePSlice(value)
	return newSlice
}





// DPSlice 获取D的PSlice
func (s *SomePSlice) DPSlice() *outter.SomePSlice {	
	value := make([]*outter.Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.D)
	}
	newSlice := outter.ToSomePSlice(value)
	return newSlice
}





// As 获取A列表
func (s *SomePSlice) As() []string {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.A)
	}
	return value
}

// Bs 获取B列表
func (s *SomePSlice) Bs() []string {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.B)
	}
	return value
}

// Cs 获取C列表
func (s *SomePSlice) Cs() []*Some {	
	value := make([]*Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.C)
	}
	return value
}

// Ds 获取D列表
func (s *SomePSlice) Ds() []*outter.Some {	
	value := make([]*outter.Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.D)
	}
	return value
}


// Collect 获取列表
func (s *SomePSlice) Collect() []*Some {
	return s.value
}


// BSlice B的Slice
type BSlice struct {
	value []B
}

// ToBSlice B列表转成BSlice
func ToBSlice(value []B) *BSlice {
	return &BSlice{value: value}
}

// Concat 拼接
func (s *BSlice) Concat(given []B) *BSlice {
	value := make([]B, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *BSlice) Drop(n int) *BSlice {
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
func (s *BSlice) Filter(fn func(int, B) bool) *BSlice {
	value := make([]B, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}



// First 获取第一个元素
func (s *BSlice) First(value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *BSlice) Last(value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *BSlice) Map(fn func(int, B) B) *BSlice {
	value := make([]B, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *BSlice) Reduce(fn func(B, B, int) B, initial B) B {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *BSlice) Reverse() *BSlice {
	value := make([]B, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}



// Append 在尾部添加元素
func (s *BSlice) Append(given B) *BSlice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *BSlice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *BSlice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsNotEmpty 判断是否非空
func (s *BSlice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// All 是否所有元素满足添加
func (s *BSlice) All(fn func(int, B) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *BSlice) Any(fn func(int, B) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *BSlice) Paginate(size int) [][]B {
	if size <= 0 {
		size = 1
	}
	var pages [][]B
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
func (s *BSlice) Preappend(given B) *BSlice {
	value := make([]B, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最后元素
func (s *BSlice) Max(bigger func(B, B) bool, value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if bigger(each, *value) {
			*value = each
		}
	}
	return nil
}

// Min 获取最小元素
func (s *BSlice) Min(less func(B, B) bool, value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if less(each, *value) {
			*value = each
		}
	}
	return nil
}

// Random 随机获取一个元素
func (s *BSlice) Random(value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *BSlice) Shuffle() *BSlice {
	if len(s.value) <= 0 {
		return s
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	return s
}









// Collect 获取最终的列表
func (s *BSlice) Collect() []B {
	return s.value
}
	
// BPSlice	B的PSlice		
type BPSlice struct {
	value []*B
}

// ToBPSlice B的指针列表转成BPSlice 
func ToBPSlice(value []*B) *BPSlice {
	return &BPSlice{value: value}
}

// Concat 拼接
func (s *BPSlice) Concat(given []*B) *BPSlice {
	value := make([]*B, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *BPSlice) Drop(n int) *BPSlice {
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
func (s *BPSlice) Filter(fn func(int, *B) bool) *BPSlice {
	value := make([]*B, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}



// First 获取第一个元素
func (s *BPSlice) First(value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = *s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *BPSlice) Last(value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = *s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *BPSlice) Map(fn func(int, *B) *B) *BPSlice {
	value := make([]*B, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *BPSlice) Reduce(fn func(*B, *B, int) *B, initial *B) *B {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *BPSlice) Reverse() *BPSlice {
	value := make([]*B, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

// UniqueBy 通过比较器唯一
func (s *BPSlice) UniqueBy(compare func(*B, *B)bool) *BPSlice {
	value := make([]*B, 0, len(s.value))
	seen := make(map[int]struct{})
	for i, outter := range s.value {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s.value {
			if i == j {
				continue
			}
			if compare(inner, outter) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value, outter)			
	}
	s.value = value
	return s
}

// Append 在尾部添加
func (s *BPSlice) Append(given *B) *BPSlice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *BPSlice) Len() int {
	return len(s.value)
}

// IsEmpty 是否为空
func (s *BPSlice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsNotEmpty 是否非空
func (s *BPSlice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// SortBy 根据比较器排序
func (s *BPSlice) SortBy(less func(*B, *B) bool) *BPSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i], s.value[j])
	})
	
	return s 
}

// All 是否所有元素满足条件
func (s *BPSlice) All(fn func(int, *B) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}






// Any 是否有元素满足条件
func (s *BPSlice) Any(fn func(int, *B) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}






// Paginate 分页
func (s *BPSlice) Paginate(size int) [][]*B {
	if size <= 0 {
		size = 1
	}
	var pages [][]*B
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
func (s *BPSlice) Preappend(given *B) *BPSlice {
	value := make([]*B, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *BPSlice) Max(bigger func(*B, *B) bool, value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = *s.value[0]
	for _, each := range s.value {
		if bigger(each, value) {
			*value = *each
		}
	}
	return nil
}

// Min 获取最小元素
func (s *BPSlice) Min(less func(*B, *B) bool, value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = *s.value[0]
	for _, each := range s.value {
		if less(each, value) {
			*value = *each
		}
	}
	return nil
}

// Random 随机获取元素
func (s *BPSlice) Random(value *B) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = *s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *BPSlice) Shuffle() *BPSlice {
	if len(s.value) <= 0 {
		return s
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}









// Collect 获取列表
func (s *BPSlice) Collect() []*B {
	return s.value
}
