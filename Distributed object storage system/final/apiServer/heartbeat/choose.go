package heartbeat

import (
	"math/rand"
)

/**
 * @Description: 随机选取一个数据节点服务器返回出去
 * @param n
 * @param exclude
 * @return ds
 */
func ChooseRandomDataServers(n int, exclude map[int]string) (ds []string) {
	candidates := make([]string, 0)
	reverseExcludeMap := make(map[string]int)
	for id, addr := range exclude {
		reverseExcludeMap[addr] = id
	}
	servers := GetDataServers()		//首先我得知道现在有哪些节点存在
	for i := range servers {
		s := servers[i]
		_, excluded := reverseExcludeMap[s]
		if !excluded {
			candidates = append(candidates, s)
		}
	}
	length := len(candidates)
	if length < n {
		return
	}
	p := rand.Perm(length)
	for i := 0; i < n; i++ {
		ds = append(ds, candidates[p[i]])
	}
	return
}
