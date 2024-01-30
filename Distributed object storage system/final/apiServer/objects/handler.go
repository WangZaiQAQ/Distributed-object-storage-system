package objects

import "net/http"

/**
 * @Description: 就是调用各种函数，去各个文件看具体的实现
 * @param w
 * @param r
 */
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPost {
		post(w, r)
		return
	}
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodDelete {
		del(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
