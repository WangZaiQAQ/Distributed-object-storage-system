package temp

import "net/http"

/**
 * @Description: 处理handler请求，去各个文件中找实现
 * @param w
 * @param r
 */
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodHead {
		head(w, r)
		return
	}
	if m == http.MethodPut {
		put(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
