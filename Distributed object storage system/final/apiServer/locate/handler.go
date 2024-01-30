package locate

import (
	"encoding/json"
	"net/http"
	"strings"
)

/**
 * @Description: 处理handler请求
 * @param w
 * @param r
 */
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {		//首先判断方法不是get就返回405错误
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])		//然后读取地址
	if len(info) == 0 {			//地址如果是空的就返回"没找到"的错误
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	w.Write(b)
}
