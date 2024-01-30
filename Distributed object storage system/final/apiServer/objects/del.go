package objects

import (
	"../../../src/lib/es"
	"log"
	"net/http"
	"strings"
)

/**
 * @Description: 删除对象的操作	逻辑很简单
 * @param w
 * @param r
 */
//将该对象中存储在ES的元数据hash置空，version＋1
func del(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	version, e := es.SearchLatestVersion(name) //找出最近的版本
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	e = es.PutMetadata(name, version.Version+1, 0, "") //删除也是用的PutMetadata函数，按照规定版本+1只不过放进去的是空的
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
