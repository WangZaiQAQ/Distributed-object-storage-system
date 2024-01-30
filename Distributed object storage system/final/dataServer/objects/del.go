package objects

import (
	"../locate"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func del(w http.ResponseWriter, r *http.Request) {
	//获取要删除的对象的哈希值和路径
	hash := strings.Split(r.URL.EscapedPath(), "/")[2]
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/" + hash + ".*")
	if len(files) != 1 {
		return
	}
	//调用delete删除索引
	locate.Del(hash)
	//
	os.Rename(files[0], os.Getenv("STORAGE_ROOT")+"/garbage/"+filepath.Base(files[0]))
}
