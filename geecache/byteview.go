package geecache

// 缓存中的item元素
type ByteView struct {
	b []byte
}

// 被缓存对象必须实现 Value接口
func (v ByteView) Len() int {
	return len(v.b)
}

// 返回一份切片的拷贝
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// 克隆一份切片
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
