
package commons

const EmptyString string =""

type StringChain struct{
	value	[]string
}

func NewStringChain(value []string) *StringChain {
	return &StringChain{value:value}
}

func(c *StringChain) Concate(given []string)  *StringChain {
	value := make([]string, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *StringChain) Drop(n int)  *StringChain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *StringChain) Filter(fn func(int, string)bool)  *StringChain {
	value := make([]string, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *StringChain) First() string {
	if len(c.value) <= 0 {
		return EmptyString
	} 
	return c.value[0]
}

func(c *StringChain) Last() string {
	if len(c.value) <= 0 {
		return EmptyString
	} 
	return c.value[len(c.value)-1]
}

func(c *StringChain) Map(fn func(int, string)) *StringChain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *StringChain) Reduce(fn func(string, string, int) string,initial string) string   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *StringChain) Reverse()  *StringChain {
	value := make([]string, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *StringChain) Unique()  *StringChain{
	value := make([]string, 0, len(c.value))
	seen:=make(map[string]struct{})
	for _, each := range c.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	c.value = value
	return c
}

func(c *StringChain) Collect() []string{
	return c.value
}
