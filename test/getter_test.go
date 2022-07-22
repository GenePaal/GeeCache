package test

import (
	"fmt"
	gc "prac/geecache"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {

	var f gc.Getter = gc.GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")

	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("callback failed")
	} else {
		fmt.Println("callback succeed")
	}

}
