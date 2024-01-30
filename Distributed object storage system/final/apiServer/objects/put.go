package objects

import (
	"../../../src/lib/es"
	"../../../src/lib/utils"
	"log"
	"net/http"
	"strings"
)

/**
 * @Description: 核心函数PUT
 * @param w
 * @param r
 */
/////////////////////////////////////////////////////////////////////////////////////////
//utils.GetHashFromHeader(r.Header) 用于从请求头中获取哈希值。在这里，hash 保存了获取的哈希值。如果获取不到哈希值，则记录日志，并返回状态码
//http.StatusBadRequest 表示请求参数错误。
func put(w http.ResponseWriter, r *http.Request) {
	hash := utils.GetHashFromHeader(r.Header) //先获取hash值
	if hash == "" {                           //hash值空的记得返回问题
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	/////////////////////////////////////////////////////////////////////////////////////////
	//从头部获取size信息,然后调用 storeObject(r.Body, hash, size)----》跳转
	//来存储客户端传输的数据。返回的 c 表示存储结果的状态码
	//e 表示可能出现的错误。如果存储发生错误，则记录日志，并返回存储结果的状态码。
	//如果存储结果的状态码不是 http.StatusOK，则返回对应的状态码作为响应。
	size := utils.GetSizeFromHeader(r.Header)
	c, e := storeObject(r.Body, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return
	}
	if c != http.StatusOK {
		w.WriteHeader(c)
		return
	}

	/////////////////////////////////////////////////////////////////////////////////////////
	//已经存储成功，在ES中写入元数据
	name := strings.Split(r.URL.EscapedPath(), "/")[2] //组成名字
	e = es.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
