package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// 定义函数类型 Hash， 采取依赖注入的方式，允许用于替换成自定义的Hash函数
// 也方便测试时替换， 默认为 crc32.ChecksumIEEE 算法
// 通俗来讲就是将 数据转成哈希码的算法
type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int   // 虚拟节点倍数
	keys     []int // 哈希环  这是那个真正用于一致性哈希的数据结构(被当做环来使用)
	// 虚拟节点与真实节点的映射表 hashMap
	// 键是虚拟节点的哈希值，值是真实节点的名称
	hashMap map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 添加真实节点， 往一致性哈希map中添加
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			// 生成虚拟节点(i 拼接 key), 然后用hash转换成32位hash码
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			// 将该虚拟节点的hash码存入哈希环中
			m.keys = append(m.keys, hash)
			// 然后将改虚拟节点的hash码 存入到真实节点的映射中
			m.hashMap[hash] = key
		}
	}
	// 对环上的哈希值排序
	sort.Ints(m.keys)
}

// 选择节点
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	// 从 [0, n) 中用二分搜索的方式找到第一个使得
	// 传入的方法为true的下标
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// 在这里取一下余 考虑到 idx=len(m.keys) 时需要置为 0， 因此在这里取余
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
