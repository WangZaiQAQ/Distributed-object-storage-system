package locate

import (
	"../../../src/lib/rabbitmq"
	"../../../src/lib/rs"
	"../../../src/lib/types"
	"encoding/json"
	"os"
	"time"
)

/*
*
  - @Description: 用来定位一个object对象的地址;  这两个函数用于获取文件的位置信息和判断文件是否存在
    新建一个消息队列，群发广播去找这个对象，如果一秒后没有响应就关闭，返回没找到
  - @param name
  - @return locateInfo	对象所在的地址
*/
func Locate(name string) (locateInfo map[int]string) {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	locateInfo = make(map[int]string)
	for i := 0; i < rs.ALL_SHARDS; i++ {
		msg := <-c
		if len(msg.Body) == 0 {
			return
		}
		var info types.LocateMessage
		json.Unmarshal(msg.Body, &info)
		locateInfo[info.Id] = info.Addr
	}
	return
}

func Exist(name string) bool {
	return len(Locate(name)) >= rs.DATA_SHARDS
}

/*
这两个函数用于获取文件的位置信息和判断文件是否存在。Locate 函数通过向 RabbitMQ 发布消息通知数据服务器进行定位操作，并接收定位信息返回，
最终将定位信息保存在字典中返回。Exist 函数则通过调用 Locate 函数获取数据服务器的位置信息，并根据返回结果判断文件是否存在于至少 rs.DATA_SHARDS 个数据服务器上。
*/
