package objects

import (
	"../../../src/lib/utils"
	"../locate"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// 函数先校验一下ES中有没有重复的哈希值，有的话直接返回传输完成
func storeObject(r io.Reader, hash string, size int64) (int, error) {
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}

	///用putStream（）在一个数据服务器上创建了一个写入流
	stream, e := putStream(url.PathEscape(hash), size)
	if e != nil {
		return http.StatusInternalServerError, e
	}

	////创建一个reader写入流，把数据r同时写入到reader和stream流中，然后计算r的哈希，如果算出来的值与http请求给定的值一样，
	//就commit到数据服务器的持久化存储上，不一样就舍弃。
	reader := io.TeeReader(r, stream)
	d := utils.CalculateHash(reader)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Commit(true)
	return http.StatusOK, nil
}
