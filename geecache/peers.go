package geecache

type PeerPicker interface {
	// 用于根据传入的key 选择响应节点
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 每个节点必须实现的接口
type PeerGetter interface {
	// 用于从对应group查找缓存值
	// 该流程在HTTP客户端中完成
	Get(group string, key string) ([]byte, error)
}
