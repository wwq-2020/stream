package tests

import (

	"errors"						

	"math/rand"						

	"sort"						

		
	commons "github.com/wwq1988/stream/commons"						

	"github.com/wwq1988/stream/outter"
)

// SomeSlice Some的Slice
type SomeSlice []Some

// Concat 拼接
func (s SomeSlice) Concat(given []Some) SomeSlice {
	value := make([]Some, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return SomeSlice(value)
}

// Drop 丢弃前n个
func (s SomeSlice) Drop(n int) SomeSlice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return s[n:]
}

// Filter 过滤
func (s SomeSlice) Filter(fn func(int, Some) bool) SomeSlice {
	value := make([]Some, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return SomeSlice(value)
}


// FilterByA 通过过滤器过滤
func (s SomeSlice) FilterByA(fn func(int, string) bool) SomeSlice {
	value := make([]Some, 0, len(s))
	for i, each := range s {
		if fn(i, each.A) {
			value = append(value, each)
		}
	}
	return SomeSlice(value)
}

// FilterByB 通过过滤器过滤
func (s SomeSlice) FilterByB(fn func(int, string) bool) SomeSlice {
	value := make([]Some, 0, len(s))
	for i, each := range s {
		if fn(i, each.B) {
			value = append(value, each)
		}
	}
	return SomeSlice(value)
}

// FilterByC 通过过滤器过滤
func (s SomeSlice) FilterByC(fn func(int, *Some) bool) SomeSlice {
	value := make([]Some, 0, len(s))
	for i, each := range s {
		if fn(i, each.C) {
			value = append(value, each)
		}
	}
	return SomeSlice(value)
}

// FilterByD 通过过滤器过滤
func (s SomeSlice) FilterByD(fn func(int, *outter.Some) bool) SomeSlice {
	value := make([]Some, 0, len(s))
	for i, each := range s {
		if fn(i, each.D) {
			value = append(value, each)
		}
	}
	return SomeSlice(value)
}


// First 获取第一个元素
func (s SomeSlice) First() (Some, error) {
	if len(s) <= 0 {
		var defaultReturn Some
		return defaultReturn, errors.New("empty")
	} 
	return s[0], nil
}

// Last 获取最后一个元素
func (s SomeSlice) Last(value *Some) (Some, error) {
	if len(s) <= 0 {
		var defaultReturn Some
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s SomeSlice) Map(fn func(int, Some) Some) SomeSlice {
	value := make([]Some, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return SomeSlice(value)
}

// Reduce reduce
func (s SomeSlice) Reduce(fn func(Some, Some, int) Some, initial Some) Some {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s SomeSlice) Reverse() SomeSlice {
	value := make([]Some, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return SomeSlice(value)
}



// UniqueByA 通过A唯一
func (s SomeSlice) UniqueByA() SomeSlice {
	value := make([]Some, 0, len(s))
	seen := make(map[string]struct{})
	for _, each := range s {
		if _, dup := seen[each.A]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.A] = struct{}{}	
	}
	return SomeSlice(value)
}



// UniqueByB 通过B唯一
func (s SomeSlice) UniqueByB() SomeSlice {
	value := make([]Some, 0, len(s))
	seen := make(map[string]struct{})
	for _, each := range s {
		if _, dup := seen[each.B]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.B] = struct{}{}	
	}
	return SomeSlice(value)
}




// UniqueByC 通过C唯一
func (s SomeSlice) UniqueByC(compare func(*Some, *Some) bool) SomeSlice {
	value := make([]Some, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s {
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
	return SomeSlice(value)
}





// UniqueByD 通过D唯一
func (s SomeSlice) UniqueByD(compare func(*outter.Some, *outter.Some) bool) SomeSlice {
	value := make([]Some, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s {
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
	return SomeSlice(value)
}




// Append 在尾部添加元素
func (s SomeSlice) Append(given Some) SomeSlice {
	return append(s, given)
}

// Len 获取长度
func (s SomeSlice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s SomeSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s SomeSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// All 是否所有元素满足添加
func (s SomeSlice) All(fn func(int, Some) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s SomeSlice) Any(fn func(int, Some) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s SomeSlice) Paginate(size int) [][]Some {
	if size <= 0 {
		size = 1
	}
	var pages [][]Some
	prev := -1
	for i := range s {
		if (i-prev) < size && i != (len(s)-1) {
			continue
		}
		pages = append(pages, s[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s SomeSlice) Preappend(given Some) SomeSlice {
	value := make([]Some, len(s)+1)
	value = append(value, given)
	value[0] = given
	copy(value[1:], s)
	return SomeSlice(value)
}

// Max 获取最后元素
func (s SomeSlice) Max(bigger func(Some, Some) bool) (Some, error) {
	if len(s) <= 0 {
		var defaultReturn Some
		return defaultReturn, errors.New("empty")
	}
	max := s[0]
	for _, each := range s {
		if bigger(each, max) {
			max = each
		}
	}
	return max, nil
}

// Min 获取最小元素
func (s SomeSlice) Min(less func(Some, Some) bool) (Some, error) {
	if len(s) <= 0 {
		var defaultReturn Some
		return defaultReturn, errors.New("empty")
	}
	min := s[0]
	for _, each := range s {
		if less(each, min) {
			min = each
		}
	}
	return min, nil
}

// Random 随机获取一个元素
func (s SomeSlice) Random() (Some, error) {
	if len(s) <= 0 {
		var defaultReturn Some
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱列表
func (s SomeSlice) Shuffle() SomeSlice {
	if len(s) <= 0 {
		return s
	}
	
	value := make([]Some, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return s
}



// SortByA 根据A排序
func (s SomeSlice) SortByA() SomeSlice {
	value := make([]Some, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i].A < value[j].A
	})
	return s 
}



// SortByB 根据B排序
func (s SomeSlice) SortByB() SomeSlice {
	value := make([]Some, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i].B < value[j].B
	})
	return s 
}



// SortByC 根据C排序
func (s SomeSlice) SortByC(less func(*Some, *Some) bool) SomeSlice {
	value := make([]Some, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return less(value[i].C, value[j].C)
	})
	return s 
}



// SortByD 根据D排序
func (s SomeSlice) SortByD(less func(*outter.Some, *outter.Some) bool) SomeSlice {
	value := make([]Some, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return less(value[i].D, value[j].D)
	})
	return s 
}








// APSlice 获取A的Slice
func (s SomeSlice) ASlice() commons.StringSlice {	
	value := make([]string, 0, len(s))	
	for _, each := range s {
		value = append(value, each.A)
	}
	newSlice := commons.StringSlice(value)
	return newSlice
}





// BPSlice 获取B的Slice
func (s SomeSlice) BSlice() commons.StringSlice {	
	value := make([]string, 0, len(s))	
	for _, each := range s {
		value = append(value, each.B)
	}
	newSlice := commons.StringSlice(value)
	return newSlice
}





// CPSlice 获取C的PSlice
func (s SomeSlice) CPSlice() SomePSlice {	
	value := make([]*Some, 0, len(s))	
	for _, each := range s {
		value = append(value, each.C)
	}
	newSlice := SomePSlice(value)
	return newSlice
}





// DPSlice 获取D的PSlice
func (s SomeSlice) DPSlice() outter.SomePSlice {	
	value := make([]*outter.Some, 0, len(s))	
	for _, each := range s {
		value = append(value, each.D)
	}
	newSlice := outter.SomePSlice(value)
	return newSlice
}





// As 获取A的列表
func (s SomeSlice) As() []string {	
	value := make([]string, 0, len(s))	
	for _, each := range s {
		value = append(value, each.A)
	}
	return value

}

// Bs 获取B的列表
func (s SomeSlice) Bs() []string {	
	value := make([]string, 0, len(s))	
	for _, each := range s {
		value = append(value, each.B)
	}
	return value

}

// Cs 获取C的列表
func (s SomeSlice) Cs() []*Some {	
	value := make([]*Some, 0, len(s))	
	for _, each := range s {
		value = append(value, each.C)
	}
	return value

}

// Ds 获取D的列表
func (s SomeSlice) Ds() []*outter.Some {	
	value := make([]*outter.Some, 0, len(s))	
	for _, each := range s {
		value = append(value, each.D)
	}
	return value

}


// Collect 获取最终的列表
func (s SomeSlice) Collect() []Some {
	return s
}
	
// SomePSlice	Some的PSlice		
type SomePSlice []*Some

// Concat 拼接
func (s SomePSlice) Concat(given []*Some) SomePSlice {
	value := make([]*Some, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return SomePSlice(value)
}

// Drop 丢弃前n个
func (s SomePSlice) Drop(n int) SomePSlice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return s[n:]
}

// Filter 过滤
func (s SomePSlice) Filter(fn func(int, *Some) bool) SomePSlice {
	value := make([]*Some, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return SomePSlice(value)
}


// FilterByA 通过过滤器过滤
func (s SomePSlice) FilterByA(fn func(int, string) bool) SomePSlice {
	value := make([]*Some, 0, len(s))
	for i, each := range s {
		if fn(i, each.A) {
			value = append(value, each)
		}
	}
	return SomePSlice(value)
}

// FilterByB 通过过滤器过滤
func (s SomePSlice) FilterByB(fn func(int, string) bool) SomePSlice {
	value := make([]*Some, 0, len(s))
	for i, each := range s {
		if fn(i, each.B) {
			value = append(value, each)
		}
	}
	return SomePSlice(value)
}

// FilterByC 通过过滤器过滤
func (s SomePSlice) FilterByC(fn func(int, *Some) bool) SomePSlice {
	value := make([]*Some, 0, len(s))
	for i, each := range s {
		if fn(i, each.C) {
			value = append(value, each)
		}
	}
	return SomePSlice(value)
}

// FilterByD 通过过滤器过滤
func (s SomePSlice) FilterByD(fn func(int, *outter.Some) bool) SomePSlice {
	value := make([]*Some, 0, len(s))
	for i, each := range s {
		if fn(i, each.D) {
			value = append(value, each)
		}
	}
	return SomePSlice(value)
}


// First 获取第一个元素
func (s SomePSlice) First() (*Some, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s SomePSlice) Last() (*Some, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	} 
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s SomePSlice) Map(fn func(int, *Some) *Some) SomePSlice {
	value := make([]*Some, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return SomePSlice(value)
}

// Reduce reduce
func (s SomePSlice) Reduce(fn func(*Some, *Some, int) *Some, initial *Some) *Some {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s SomePSlice) Reverse() SomePSlice {
	value := make([]*Some, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return SomePSlice(value)
}

// UniqueBy 通过比较器唯一
func (s SomePSlice) UniqueBy(compare func(*Some, *Some)bool) SomePSlice {
	value := make([]*Some, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s {
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
	return SomePSlice(value)
}

// Append 在尾部添加
func (s SomePSlice) Append(given *Some) SomePSlice {
	return append(s, given)
}

// Len 获取长度
func (s SomePSlice) Len() int {
	return len(s)
}

// IsEmpty 是否为空
func (s SomePSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 是否非空
func (s SomePSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// SortBy 根据比较器排序
func (s SomePSlice) SortBy(less func(*Some, *Some) bool) SomePSlice {
	value := make([]*Some, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return less(value[i], value[j])
	})
	
	return s 
}

// All 是否所有元素满足条件
func (s SomePSlice) All(fn func(int, *Some) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}


// AllByA 是否所有元素的A满足条件
func (s SomePSlice) AllByA(fn func(int, string) bool) bool {
	for i, each := range s {
		if !fn(i, each.A){
			return false
		}
	}
	return true
}

// AllByB 是否所有元素的B满足条件
func (s SomePSlice) AllByB(fn func(int, string) bool) bool {
	for i, each := range s {
		if !fn(i, each.B){
			return false
		}
	}
	return true
}

// AllByC 是否所有元素的C满足条件
func (s SomePSlice) AllByC(fn func(int, *Some) bool) bool {
	for i, each := range s {
		if !fn(i, each.C){
			return false
		}
	}
	return true
}

// AllByD 是否所有元素的D满足条件
func (s SomePSlice) AllByD(fn func(int, *outter.Some) bool) bool {
	for i, each := range s {
		if !fn(i, each.D){
			return false
		}
	}
	return true
}




// AllByA 是否所有元素的A满足条件
func (s SomeSlice) AllByA(fn func(int, string) bool) bool {
	for i, each := range s {
		if !fn(i, each.A){
			return false
		}
	}
	return true
}

// AllByB 是否所有元素的B满足条件
func (s SomeSlice) AllByB(fn func(int, string) bool) bool {
	for i, each := range s {
		if !fn(i, each.B){
			return false
		}
	}
	return true
}

// AllByC 是否所有元素的C满足条件
func (s SomeSlice) AllByC(fn func(int, *Some) bool) bool {
	for i, each := range s {
		if !fn(i, each.C){
			return false
		}
	}
	return true
}

// AllByD 是否所有元素的D满足条件
func (s SomeSlice) AllByD(fn func(int, *outter.Some) bool) bool {
	for i, each := range s {
		if !fn(i, each.D){
			return false
		}
	}
	return true
}


// Any 是否有元素满足条件
func (s SomePSlice) Any(fn func(int, *Some) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}



// AnyByA 是否有元素的A满足条件
func (s SomePSlice) AnyByA(fn func(int, string) bool) bool {
	for i, each := range s {
		if fn(i, each.A) {
			return true
		}
	}
	return false
}

// AnyByB 是否有元素的B满足条件
func (s SomePSlice) AnyByB(fn func(int, string) bool) bool {
	for i, each := range s {
		if fn(i, each.B) {
			return true
		}
	}
	return false
}

// AnyByC 是否有元素的C满足条件
func (s SomePSlice) AnyByC(fn func(int, *Some) bool) bool {
	for i, each := range s {
		if fn(i, each.C) {
			return true
		}
	}
	return false
}

// AnyByD 是否有元素的D满足条件
func (s SomePSlice) AnyByD(fn func(int, *outter.Some) bool) bool {
	for i, each := range s {
		if fn(i, each.D) {
			return true
		}
	}
	return false
}



// AnyByA 是否有元素的A满足条件
func (s SomeSlice) AnyByA(fn func(int, string) bool) bool {
	for i, each := range s {
		if fn(i, each.A) {
			return true
		}
	}
	return false
}

// AnyByB 是否有元素的B满足条件
func (s SomeSlice) AnyByB(fn func(int, string) bool) bool {
	for i, each := range s {
		if fn(i, each.B) {
			return true
		}
	}
	return false
}

// AnyByC 是否有元素的C满足条件
func (s SomeSlice) AnyByC(fn func(int, *Some) bool) bool {
	for i, each := range s {
		if fn(i, each.C) {
			return true
		}
	}
	return false
}

// AnyByD 是否有元素的D满足条件
func (s SomeSlice) AnyByD(fn func(int, *outter.Some) bool) bool {
	for i, each := range s {
		if fn(i, each.D) {
			return true
		}
	}
	return false
}


// Paginate 分页
func (s SomePSlice) Paginate(size int) [][]*Some {
	if size <= 0 {
		size = 1
	}
	var pages [][]*Some
	prev := -1
	for i := range s {
		if (i-prev) < size && i != (len(s)-1) {
			continue
		}
		pages = append(pages, s[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s SomePSlice) Preappend(given *Some) SomePSlice {
	value := make([]*Some, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return SomePSlice(value)
}

// Max 获取最大元素
func (s SomePSlice) Max(bigger func(*Some, *Some) bool) (*Some, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	}
	max := s[0]
	for _, each := range s {
		if bigger(each, max) {
			max = max
		}
	}
	return max, nil
}

// Min 获取最小元素
func (s SomePSlice) Min(less func(*Some, *Some) bool) (*Some, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	}
	min := s[0]
	for _, each := range s {
		if less(each, min) {
			min = each
		}
	}
	return min, nil
}

// Random 随机获取元素
func (s SomePSlice) Random() (*Some, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱列表
func (s SomePSlice) Shuffle() SomePSlice {
	if len(s) <= 0 {
		return s
	}
	value := make([]*Some, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	
	return SomePSlice(value)
}



// SortByA 根据元素的A排序
func (s SomePSlice) SortByA() SomePSlice {
	value := make([]*Some, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i].A < value[j].A
	})
	return SomePSlice(value)
}



// SortByB 根据元素的B排序
func (s SomePSlice) SortByB() SomePSlice {
	value := make([]*Some, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i].B < value[j].B
	})
	return SomePSlice(value)
}



// SortByC 根据元素的C和比较器排序
func (s SomePSlice) SortByC(less func(*Some, *Some) bool) SomePSlice {
	value := make([]*Some, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return less(value[i].C, value[j].C)
	})
	return SomePSlice(value)
}



// SortByD 根据元素的D和比较器排序
func (s SomePSlice) SortByD(less func(*outter.Some, *outter.Some) bool) SomePSlice {
	value := make([]*Some, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return less(value[i].D, value[j].D)
	})
	return SomePSlice(value)
}





// UniqueByA 根据元素的A唯一
func (s SomePSlice) UniqueByA() SomePSlice {
	value := make([]*Some, 0, len(s))
	seen := make(map[string]struct{})
	for _, each := range s {
		if _, dup := seen[each.A]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.A] = struct{}{}	
	}
	return SomePSlice(value)
}



// UniqueByB 根据元素的B唯一
func (s SomePSlice) UniqueByB() SomePSlice {
	value := make([]*Some, 0, len(s))
	seen := make(map[string]struct{})
	for _, each := range s {
		if _, dup := seen[each.B]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.B] = struct{}{}	
	}
	return SomePSlice(value)
}




// UniqueByC 根据元素的C和比较器唯一
func (s SomePSlice) UniqueByC(compare func (*Some, *Some) bool) SomePSlice {
	value := make([]*Some, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j,inner :=range s {
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
	return SomePSlice(value)
}





// UniqueByD 根据元素的D和比较器唯一
func (s SomePSlice) UniqueByD(compare func (*outter.Some, *outter.Some) bool) SomePSlice {
	value := make([]*Some, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j,inner :=range s {
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
	return SomePSlice(value)
}







// ASlice 获取A的Slice
func (s SomePSlice) ASlice() commons.StringSlice {	
	value := make([]string, 0, len(s))	
	for _, each := range s {
		value = append(value, each.A)
	}
	newSlice := commons.StringSlice(value)
	return newSlice
}





// BSlice 获取B的Slice
func (s SomePSlice) BSlice() commons.StringSlice {	
	value := make([]string, 0, len(s))	
	for _, each := range s {
		value = append(value, each.B)
	}
	newSlice := commons.StringSlice(value)
	return newSlice
}





// CPSlice 获取C的PSlice
func (s SomePSlice) CPSlice() SomePSlice {	
	value := make([]*Some, 0, len(s))	
	for _, each := range s {
		value = append(value, each.C)
	}
	newSlice := SomePSlice(value)
	return newSlice
}





// DPSlice 获取D的PSlice
func (s SomePSlice) DPSlice() outter.SomePSlice {	
	value := make([]*outter.Some, 0, len(s))	
	for _, each := range s {
		value = append(value, each.D)
	}
	newSlice := outter.SomePSlice(value)
	return newSlice
}





// As 获取A列表
func (s SomePSlice) As() []string {	
	value := make([]string, 0, len(s))	
	for _, each := range s {
		value = append(value, each.A)
	}
	return value

}

// Bs 获取B列表
func (s SomePSlice) Bs() []string {	
	value := make([]string, 0, len(s))	
	for _, each := range s {
		value = append(value, each.B)
	}
	return value

}

// Cs 获取C列表
func (s SomePSlice) Cs() []*Some {	
	value := make([]*Some, 0, len(s))	
	for _, each := range s {
		value = append(value, each.C)
	}
	return value

}

// Ds 获取D列表
func (s SomePSlice) Ds() []*outter.Some {	
	value := make([]*outter.Some, 0, len(s))	
	for _, each := range s {
		value = append(value, each.D)
	}
	return value

}




// A2Some A到Some的map
func (s SomePSlice) A2Some() map[string]*Some {
	result := make(map[string]*Some, len(s))
	for _, each := range s {
		result[each.A] = each
	}
	return result
}



// B2Some B到Some的map
func (s SomePSlice) B2Some() map[string]*Some {
	result := make(map[string]*Some, len(s))
	for _, each := range s {
		result[each.B] = each
	}
	return result
}







// Collect 获取列表
func (s SomePSlice) Collect() []*Some {
	return s
}


// BSlice B的Slice
type BSlice []B

// Concat 拼接
func (s BSlice) Concat(given []B) BSlice {
	value := make([]B, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return BSlice(value)
}

// Drop 丢弃前n个
func (s BSlice) Drop(n int) BSlice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return s[n:]
}

// Filter 过滤
func (s BSlice) Filter(fn func(int, B) bool) BSlice {
	value := make([]B, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return BSlice(value)
}



// First 获取第一个元素
func (s BSlice) First() (B, error) {
	if len(s) <= 0 {
		var defaultReturn B
		return defaultReturn, errors.New("empty")
	} 
	return s[0], nil
}

// Last 获取最后一个元素
func (s BSlice) Last(value *B) (B, error) {
	if len(s) <= 0 {
		var defaultReturn B
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s BSlice) Map(fn func(int, B) B) BSlice {
	value := make([]B, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return BSlice(value)
}

// Reduce reduce
func (s BSlice) Reduce(fn func(B, B, int) B, initial B) B {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s BSlice) Reverse() BSlice {
	value := make([]B, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return BSlice(value)
}



// Append 在尾部添加元素
func (s BSlice) Append(given B) BSlice {
	return append(s, given)
}

// Len 获取长度
func (s BSlice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s BSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s BSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// All 是否所有元素满足添加
func (s BSlice) All(fn func(int, B) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s BSlice) Any(fn func(int, B) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s BSlice) Paginate(size int) [][]B {
	if size <= 0 {
		size = 1
	}
	var pages [][]B
	prev := -1
	for i := range s {
		if (i-prev) < size && i != (len(s)-1) {
			continue
		}
		pages = append(pages, s[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s BSlice) Preappend(given B) BSlice {
	value := make([]B, len(s)+1)
	value = append(value, given)
	value[0] = given
	copy(value[1:], s)
	return BSlice(value)
}

// Max 获取最后元素
func (s BSlice) Max(bigger func(B, B) bool) (B, error) {
	if len(s) <= 0 {
		var defaultReturn B
		return defaultReturn, errors.New("empty")
	}
	max := s[0]
	for _, each := range s {
		if bigger(each, max) {
			max = each
		}
	}
	return max, nil
}

// Min 获取最小元素
func (s BSlice) Min(less func(B, B) bool) (B, error) {
	if len(s) <= 0 {
		var defaultReturn B
		return defaultReturn, errors.New("empty")
	}
	min := s[0]
	for _, each := range s {
		if less(each, min) {
			min = each
		}
	}
	return min, nil
}

// Random 随机获取一个元素
func (s BSlice) Random() (B, error) {
	if len(s) <= 0 {
		var defaultReturn B
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱列表
func (s BSlice) Shuffle() BSlice {
	if len(s) <= 0 {
		return s
	}
	
	value := make([]B, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return s
}









// Collect 获取最终的列表
func (s BSlice) Collect() []B {
	return s
}
	
// BPSlice	B的PSlice		
type BPSlice []*B

// Concat 拼接
func (s BPSlice) Concat(given []*B) BPSlice {
	value := make([]*B, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return BPSlice(value)
}

// Drop 丢弃前n个
func (s BPSlice) Drop(n int) BPSlice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return s[n:]
}

// Filter 过滤
func (s BPSlice) Filter(fn func(int, *B) bool) BPSlice {
	value := make([]*B, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return BPSlice(value)
}



// First 获取第一个元素
func (s BPSlice) First() (*B, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s BPSlice) Last() (*B, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	} 
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s BPSlice) Map(fn func(int, *B) *B) BPSlice {
	value := make([]*B, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return BPSlice(value)
}

// Reduce reduce
func (s BPSlice) Reduce(fn func(*B, *B, int) *B, initial *B) *B {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s BPSlice) Reverse() BPSlice {
	value := make([]*B, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return BPSlice(value)
}

// UniqueBy 通过比较器唯一
func (s BPSlice) UniqueBy(compare func(*B, *B)bool) BPSlice {
	value := make([]*B, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s {
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
	return BPSlice(value)
}

// Append 在尾部添加
func (s BPSlice) Append(given *B) BPSlice {
	return append(s, given)
}

// Len 获取长度
func (s BPSlice) Len() int {
	return len(s)
}

// IsEmpty 是否为空
func (s BPSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 是否非空
func (s BPSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// SortBy 根据比较器排序
func (s BPSlice) SortBy(less func(*B, *B) bool) BPSlice {
	value := make([]*B, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return less(value[i], value[j])
	})
	
	return s 
}

// All 是否所有元素满足条件
func (s BPSlice) All(fn func(int, *B) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}






// Any 是否有元素满足条件
func (s BPSlice) Any(fn func(int, *B) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}






// Paginate 分页
func (s BPSlice) Paginate(size int) [][]*B {
	if size <= 0 {
		size = 1
	}
	var pages [][]*B
	prev := -1
	for i := range s {
		if (i-prev) < size && i != (len(s)-1) {
			continue
		}
		pages = append(pages, s[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s BPSlice) Preappend(given *B) BPSlice {
	value := make([]*B, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return BPSlice(value)
}

// Max 获取最大元素
func (s BPSlice) Max(bigger func(*B, *B) bool) (*B, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	}
	max := s[0]
	for _, each := range s {
		if bigger(each, max) {
			max = max
		}
	}
	return max, nil
}

// Min 获取最小元素
func (s BPSlice) Min(less func(*B, *B) bool) (*B, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	}
	min := s[0]
	for _, each := range s {
		if less(each, min) {
			min = each
		}
	}
	return min, nil
}

// Random 随机获取元素
func (s BPSlice) Random() (*B, error) {
	if len(s) <= 0 {
		return nil, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱列表
func (s BPSlice) Shuffle() BPSlice {
	if len(s) <= 0 {
		return s
	}
	value := make([]*B, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	
	return BPSlice(value)
}











// Collect 获取列表
func (s BPSlice) Collect() []*B {
	return s
}
