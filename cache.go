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

func isExpired(deadline time.Time) bool {
	return deadline.Before(time.Now())
}

func (c *Cache) Get(key string) (string, bool) {
	if val, ok := c.values[key]; ok {
		if val.isTimeout && isExpired(val.deadline)  {
			delete(c.values, key)
			return "", false
		}
		return val.value, true
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.values[key] = newValue(value, false, time.Now())
}

func (c *Cache) Keys() []string {
	ret := make([]string, 0, len(c.values))
	for k, val := range c.values {
		if !val.isTimeout || (val.isTimeout && !isExpired(val.deadline)) {
			ret = append(ret, k)
		} else {
			delete(c.values, k)
		}
		
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
// 	t := time.Now().Add(-24 * time.Hour)
// 	c.PutTill("four", "444444", t)
// 	newval, err := c.Get("four")
// 	fmt.Println(newval, err)
// 	_, err = c.Get("five")
// 	fmt.Println(err)
// 	keys := c.Keys()
// 	fmt.Println(keys)
// 	fmt.Println(c)
// }