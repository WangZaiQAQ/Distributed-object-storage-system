package objects

import (
	"../../../src/lib/es"
	"../../../src/lib/utils"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/**
 * @Description:	取出对象，注意按照版本取，URL里没指定就取最新的版本
 * @param w
 * @param r
 */
func get(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	versionId := r.URL.Query()["version"]
	version := 0
	var e error
	if len(versionId) != 0 {
		version, e = strconv.Atoi(versionId[0])		//拿到版本号
		if e != nil {		//有问题就报错
			log.Println(e)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	meta, e := es.GetMetadata(name, version)	//	从ES中取出来
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if meta.Hash == "" {			//空的就是没找到咯，del里面删除不就是置空吗
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hash := url.PathEscape(meta.Hash)
	stream, e := GetStream(hash, meta.Size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	offset := utils.GetOffsetFromHeader(r.Header)
	if offset != 0 {
		stream.Seek(offset, io.SeekCurrent)
		w.Header().Set("content-range", fmt.Sprintf("bytes %d-%d/%d", offset, meta.Size-1, meta.Size))
		w.WriteHeader(http.StatusPartialContent)
	}
	acceptGzip := false
	encoding := r.Header["Accept-Encoding"]
	for i := range encoding {
		if encoding[i] == "gzip" {
			acceptGzip = true
			break
		}
	}
	if acceptGzip {
		w.Header().Set("content-encoding", "gzip")
		w2 := gzip.NewWriter(w)
		io.Copy(w2, stream)
		w2.Close()
	} else {
		io.Copy(w, stream)
	}
	stream.Close()
}
