package outter

import (
	"math/rand"
)

// SomeSlice Some的Slice
type SomeSlice []Some

// AOfSomeMap AOfSomeMap
type AOfSomeMap map[string]SomeSlice

// FlatMap FlatMap
func (m AOfSomeMap) FlatMap(fn func([]Some)) {
	for _, list := range m {
		fn(list)
	}
}

// BOfSomeMap BOfSomeMap
type BOfSomeMap map[string]SomeSlice

// FlatMap FlatMap
func (m BOfSomeMap) FlatMap(fn func([]Some)) {
	for _, list := range m {
		fn(list)
	}
}

// COfSomeMap COfSomeMap
type COfSomeMap map[*Some]SomeSlice

// FlatMap FlatMap
func (m COfSomeMap) FlatMap(fn func([]Some)) {
	for _, list := range m {
		fn(list)
	}
}

// SomeResult SomeResult
type SomeResult struct {
	value     Some
	isPresent bool
}

// IsPresent 是否存在
func (r SomeResult) IsPresent() bool {
	return r.isPresent
}

// Get 获取值
func (r SomeResult) Get() Some {
	return r.value
}

// Concat 拼接
func (s SomeSlice) Concat(given []Some) SomeSlice {
	result := make([]Some, len(s)+len(given))
	copy(result, s)
	copy(result[len(s):], given)
	return result
}

// Limit 取前n个
func (s SomeSlice) Limit(n int) SomeSlice {
	result := make([]Some, 0, len(s))
	for idx, each := range s {
		if idx < n {
			result = append(result, each)
		}
	}
	return result
}

// Peek Peek
func (s SomeSlice) Peek(fn func(Some)) {
	for _, each := range s {
		fn(each)
	}
}

// Skip 丢弃前n个
func (s SomeSlice) Skip(n int) SomeSlice {
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
func (s SomeSlice) Filter(filters ...func(Some) bool) SomeSlice {
	result := make([]Some, 0, len(s))
	for _, each := range s {
		valid := true
		for _, filter := range filters {
			if !filter(each) {
				valid = false
			}
		}
		if valid {
			result = append(result, each)
		}
	}
	return result
}

// GroupByA 通过A分组
func (s SomeSlice) GroupByA(comparator func(string, string) bool) AOfSomeMap {
	result := make(map[string]SomeSlice, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip := skip[i]; skip {
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter.A, inner.A) {
				skip[j] = struct{}{}
			}
		}
		result[outter.A] = append(result[outter.A], outter)
	}
	return result
}

// GroupByB 通过B分组
func (s SomeSlice) GroupByB(comparator func(string, string) bool) BOfSomeMap {
	result := make(map[string]SomeSlice, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip := skip[i]; skip {
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter.B, inner.B) {
				skip[j] = struct{}{}
			}
		}
		result[outter.B] = append(result[outter.B], outter)
	}
	return result
}

// GroupByC 通过C分组
func (s SomeSlice) GroupByC(comparator func(*Some, *Some) bool) COfSomeMap {
	result := make(map[*Some]SomeSlice, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip := skip[i]; skip {
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter.C, inner.C) {
				skip[j] = struct{}{}
			}
		}
		result[outter.C] = append(result[outter.C], outter)
	}
	return result
}

// First 获取第一个元素
func (s SomeSlice) First() SomeResult {
	if len(s) <= 0 {
		var defaultReturn Some
		return SomeResult{value: defaultReturn, isPresent: false}
	}
	return SomeResult{value: s[0], isPresent: true}
}

// Last 获取最后一个元素
func (s SomeSlice) Last(value *Some) SomeResult {
	if len(s) <= 0 {
		var defaultReturn Some
		return SomeResult{value: defaultReturn, isPresent: false}
	}
	return SomeResult{value: s[len(s)-1], isPresent: true}
}

// Map 对每个元素进行操作
func (s SomeSlice) Map(fn func(Some) Some) SomeSlice {
	result := make([]Some, len(s))
	for i, each := range s {
		result[i] = fn(each)
	}
	return result
}

// Reduce reduce
func (s SomeSlice) Reduce(fn func(Some, Some) Some, initial Some) Some {
	final := initial
	for _, each := range s {
		final = fn(final, each)
	}
	return final
}

// Reverse 逆序
func (s SomeSlice) Reverse() SomeSlice {
	result := make([]Some, len(s))
	for i, each := range s {
		result[len(s)-1-i] = each
	}
	return result
}

// Distinct 去重
func (s SomeSlice) Distinct(comparator func(Some, Some) bool) SomeSlice {
	result := make(SomeSlice, 0, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip := skip[i]; skip {
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter, inner) {
				skip[j] = struct{}{}
			}
		}
		result = append(result, outter)
	}
	return result
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

// AllMatch 是否所有元素满足添加
func (s SomeSlice) AllMatch(matchFuncs ...func(Some) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if !matchFunc(each) {
				return false
			}
		}
	}
	return true
}

// AnyMatch 是否有元素满足条件
func (s SomeSlice) AnyMatch(matchFuncs ...func(Some) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if matchFunc(each) {
				return true
			}
		}
	}
	return false
}

// NoneMatch 是否没有元素满足条件
func (s SomeSlice) NoneMatch(matchFuncs ...func(Some) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if matchFunc(each) {
				return false
			}
		}
	}
	return true
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
	result := make([]Some, len(s)+1)
	result = append(result, given)
	result[0] = given
	copy(result[1:], s)
	return result
}

// Max 获取最大元素
func (s SomeSlice) Max(comparator func(Some, Some) bool) SomeResult {
	if len(s) <= 0 {
		var defaultReturn Some
		return SomeResult{value: defaultReturn, isPresent: false}
	}
	max := s[0]
	for _, each := range s {
		if comparator(each, max) {
			max = each
		}
	}
	return SomeResult{value: max, isPresent: true}
}

// Min 获取最小元素
func (s SomeSlice) Min(less func(Some, Some) bool) SomeResult {
	if len(s) <= 0 {
		var defaultReturn Some
		return SomeResult{value: defaultReturn, isPresent: true}
	}
	min := s[0]
	for _, each := range s {
		if less(each, min) {
			min = each
		}
	}
	return SomeResult{value: min, isPresent: true}
}

// Random 随机获取一个元素
func (s SomeSlice) Random() SomeResult {
	if len(s) <= 0 {
		var defaultReturn Some
		return SomeResult{value: defaultReturn, isPresent: true}
	}
	n := rand.Intn(len(s))
	return SomeResult{value: s[n], isPresent: true}
}

// Shuffle 打乱列表
func (s SomeSlice) Shuffle() SomeSlice {
	if len(s) <= 0 {
		return s
	}

	result := make([]Some, len(s))
	copy(result, s)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
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

// Collect 获取最终的列表
func (s SomeSlice) Collect() []Some {
	return s
}

// SomePSlice Some的PSlice
type SomePSlice []*Some

// AOfSomePMap AOfSomePMap
type AOfSomePMap map[string]SomePSlice

// FlatMap FlatMap
func (m AOfSomePMap) FlatMap(fn func([]*Some)) {
	for _, list := range m {
		fn(list)
	}
}

// BOfSomePMap BOfSomePMap
type BOfSomePMap map[string]SomePSlice

// FlatMap FlatMap
func (m BOfSomePMap) FlatMap(fn func([]*Some)) {
	for _, list := range m {
		fn(list)
	}
}

// COfSomePMap COfSomePMap
type COfSomePMap map[*Some]SomePSlice

// FlatMap FlatMap
func (m COfSomePMap) FlatMap(fn func([]*Some)) {
	for _, list := range m {
		fn(list)
	}
}

// SomePResult SomePResult
type SomePResult struct {
	value     *Some
	isPresent bool
}

// IsPresent 是否存在
func (r SomePResult) IsPresent() bool {
	return r.isPresent
}

// Get 获取值
func (r SomePResult) Get() *Some {
	return r.value
}

// Concat 拼接
func (s SomePSlice) Concat(given []*Some) SomePSlice {
	result := make([]*Some, len(s)+len(given))
	copy(result, s)
	copy(result[len(s):], given)
	return result
}

// Limit 取前n个
func (s SomePSlice) Limit(n int) SomePSlice {
	result := make([]*Some, 0, len(s))
	for idx, each := range s {
		if idx < n {
			result = append(result, each)
		}
	}
	return result
}

// Peek Peek
func (s SomePSlice) Peek(fn func(*Some)) {
	for _, each := range s {
		fn(each)
	}
}

// Skip 丢弃前n个
func (s SomePSlice) Skip(n int) SomePSlice {
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
func (s SomePSlice) Filter(filters ...func(*Some) bool) SomePSlice {
	result := make([]*Some, 0, len(s))
	for _, each := range s {
		valid := true
		for _, filter := range filters {
			if !filter(each) {
				valid = false
			}
		}
		if valid {
			result = append(result, each)
		}
	}
	return result
}

// GroupByA 通过A分组
func (s SomePSlice) GroupByA(comparator func(string, string) bool) AOfSomePMap {
	result := make(map[string]SomePSlice, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip := skip[i]; skip {
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter.A, inner.A) {
				skip[j] = struct{}{}
			}
		}
		result[outter.A] = append(result[outter.A], outter)
	}
	return result
}

// GroupByB 通过B分组
func (s SomePSlice) GroupByB(comparator func(string, string) bool) BOfSomePMap {
	result := make(map[string]SomePSlice, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip := skip[i]; skip {
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter.B, inner.B) {
				skip[j] = struct{}{}
			}
		}
		result[outter.B] = append(result[outter.B], outter)
	}
	return result
}

// GroupByC 通过C分组
func (s SomePSlice) GroupByC(comparator func(*Some, *Some) bool) COfSomePMap {
	result := make(map[*Some]SomePSlice, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip := skip[i]; skip {
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter.C, inner.C) {
				skip[j] = struct{}{}
			}
		}
		result[outter.C] = append(result[outter.C], outter)
	}
	return result
}

// First 获取第一个元素
func (s SomePSlice) First() SomePResult {
	if len(s) <= 0 {
		var defaultReturn *Some
		return SomePResult{value: defaultReturn, isPresent: false}
	}
	return SomePResult{value: s[0], isPresent: true}
}

// Last 获取最后一个元素
func (s SomePSlice) Last(value *Some) SomePResult {
	if len(s) <= 0 {
		var defaultReturn *Some
		return SomePResult{value: defaultReturn, isPresent: false}
	}
	return SomePResult{value: s[len(s)-1], isPresent: true}
}

// Map 对每个元素进行操作
func (s SomePSlice) Map(fn func(*Some) *Some) SomePSlice {
	result := make([]*Some, len(s))
	for i, each := range s {
		result[i] = fn(each)
	}
	return result
}

// Reduce reduce
func (s SomePSlice) Reduce(fn func(*Some, *Some) *Some, initial *Some) *Some {
	final := initial
	for _, each := range s {
		final = fn(final, each)
	}
	return final
}

// Reverse 逆序
func (s SomePSlice) Reverse() SomePSlice {
	result := make([]*Some, len(s))
	for i, each := range s {
		result[len(s)-1-i] = each
	}
	return result
}

// Distinct 去重
func (s SomePSlice) Distinct(comparator func(*Some, *Some) bool) SomePSlice {
	result := make([]*Some, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip := skip[i]; skip {
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter, inner) {
				skip[j] = struct{}{}
			}
		}
		result = append(result, outter)
	}
	return result
}

// Append 在尾部添加元素
func (s SomePSlice) Append(given *Some) SomePSlice {
	return append(s, given)
}

// Len 获取长度
func (s SomePSlice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s SomePSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s SomePSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// AllMatch 是否所有元素满足添加
func (s SomePSlice) AllMatch(matchFuncs ...func(*Some) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if !matchFunc(each) {
				return false
			}
		}
	}
	return true
}

// AnyMatch 是否有元素满足条件
func (s SomePSlice) AnyMatch(matchFuncs ...func(*Some) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if matchFunc(each) {
				return true
			}
		}
	}
	return false
}

// NoneMatch 是否没有元素满足条件
func (s SomePSlice) NoneMatch(matchFuncs ...func(*Some) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if matchFunc(each) {
				return false
			}
		}
	}
	return true
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
	result := make([]*Some, len(s)+1)
	result[0] = given
	copy(result[1:], s)
	return result
}

// Max 获取最大元素
func (s SomePSlice) Max(comparator func(*Some, *Some) bool) SomePResult {
	if len(s) <= 0 {
		var defaultReturn *Some
		return SomePResult{value: defaultReturn, isPresent: false}
	}
	max := s[0]
	for _, each := range s {
		if comparator(each, max) {
			max = each
		}
	}
	return SomePResult{value: max, isPresent: true}
}

// Min 获取最小元素
func (s SomePSlice) Min(comparator func(*Some, *Some) bool) SomePResult {
	if len(s) <= 0 {
		var defaultReturn *Some
		return SomePResult{value: defaultReturn, isPresent: false}
	}
	min := s[0]
	for _, each := range s {
		if comparator(each, min) {
			min = each
		}
	}
	return SomePResult{value: min, isPresent: true}
}

// Random 随机获取一个元素
func (s SomePSlice) Random() SomePResult {
	if len(s) <= 0 {
		var defaultReturn *Some
		return SomePResult{value: defaultReturn, isPresent: false}
	}
	n := rand.Intn(len(s))
	return SomePResult{value: s[n], isPresent: true}
}

// Shuffle 打乱列表
func (s SomePSlice) Shuffle() SomePSlice {
	if len(s) <= 0 {
		return s
	}

	result := make([]*Some, len(s))
	copy(result, s)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

// As 获取A的列表
func (s SomePSlice) As() []string {
	value := make([]string, 0, len(s))
	for _, each := range s {
		value = append(value, each.A)
	}
	return value
}

// Bs 获取B的列表
func (s SomePSlice) Bs() []string {
	value := make([]string, 0, len(s))
	for _, each := range s {
		value = append(value, each.B)
	}
	return value
}

// Cs 获取C的列表
func (s SomePSlice) Cs() []*Some {
	value := make([]*Some, 0, len(s))
	for _, each := range s {
		value = append(value, each.C)
	}
	return value
}

// Collect 获取最终的列表
func (s SomePSlice) Collect() []*Some {
	return s
}
