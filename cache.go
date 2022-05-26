package cache

// import (
// 	"fmt"
// 	"time"
// )

import "time"

type cacheValue struct {
	value string
	deadline time.Time
	isTimeout bool
}

type Cache struct {
	values map[string]cacheValue
}

func NewCache() Cache {
	values := map[string]cacheValue{}
	return Cache{values: values}
}

func newValue(value string, isTimeout bool, deadline time.Time) cacheValue {
	return cacheValue{value: value, isTimeout: isTimeout, deadline: deadline}
}

func (c Cache) Get(key string) (string, bool) {
	if val, ok := c.values[key]; ok {
		if val.isTimeout && val.deadline.Before(time.Now())  {
			return "", false
		}
		return val.value, true
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.values[key] = newValue(value, false, time.Now())
}

func (c Cache) Keys() []string {
	ret := make([]string, len(c.values))
	for k := range c.values {
		ret = append(ret, k)
	  }
	  return ret
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.values[key] = newValue(value, true, deadline)
}

// func main() {
// 	c := NewCache()
// 	fmt.Println(c)
// 	c.Put("one", "111111")
// 	c.Put("two", "222222")
// 	keys := c.Keys()
// 	fmt.Println(keys)
// 	t := time.Now().Add(24 * time.Hour)
// 	c.PutTill("four", "444444", t)
// 	newval, err := c.Get("four")
// 	fmt.Println(newval, err)
// 	_, err = c.Get("five")
// 	fmt.Println(err)
// 	fmt.Println(c)
// }