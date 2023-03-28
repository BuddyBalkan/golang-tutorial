package redis

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"runtime"
	"time"
)

const (
	WeekSeconds   = 7 * DaySeconds     // 一周的秒数
	DaySeconds    = 24 * HourSeconds   // 1天的秒数
	HourSeconds   = 60 * MinuteSeconds // 一小时的秒数
	MinuteSeconds = 60                 // 一分钟的秒数
)

const (
	Day   = 24 * time.Hour // 一天
	Month = 30 * Day       // 一个月
)

var (
	Client *redis.Client
)

// 初始化redis
func Init() error {
	// 初始化连接
	Client = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		PoolSize:     runtime.NumCPU() * 4,
		MinIdleConns: runtime.NumCPU(),
		OnConnect: func(conn *redis.Conn) error {
			glog.Infof("current redis PoolState:%v\n", Client.PoolStats())
			return nil
		},
	})
	result, err := Client.Ping().Result()
	if err != nil {
		return err
	}
	glog.Infof("redis client ping return:%s\n", result)
	glog.Infof("redis init after PoolState:%v\n", Client.PoolStats())
	return nil
}

// 关闭redis
func Close() {
	//closePubSub = true
	//if pubSub != nil {
	//	err := pubSub.Close()
	//	if err != nil {
	//		glog.Errorf("pubSub close err:%s\n", err.Error())
	//	}
	//}
	if Client != nil {
		return
	}
	if err := Client.Close(); err != nil {
		glog.Errorf("redis client close err:%s\n", err.Error())
	} else {
		glog.Info("redis client closed")
	}
}

// 初始化发布订阅
//func InitPubSub(pubSubProc func(msg *redis.Message), subChannels ...string) {
//	if len(subChannels) < 1 {
//		return
//	}
//	if pubSubProc == nil {
//		panic("no pubSubHandler")
//	}
//	pubSubHandler = pubSubProc
//	pubSub = Client.Subscribe(subChannels...)
//	go func() {
//		for {
//			if closePubSub {
//				break
//			}
//			msg, err := pubSub.ReceiveMessage()
//			if err != nil {
//				glog.Errorf("redis pubSub err:%s\n", err.Error())
//			} else {
//				if pubSubHandler != nil {
//					go pubSubHandler(msg)
//				} else {
//					glog.Errorf("no pubSubHandler, msg:%v\n", msg)
//				}
//			}
//		}
//	}()
//}

// set 命令
func Save(key string, data interface{}) {
	if Client == nil {
		glog.Error("redis client is nil")
		return
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		glog.Errorf("db set key %s marshal err %s", key, err.Error())
		return
	}
	s, err := Client.Set(key, bytes, Month).Result()
	if err != nil {
		glog.Errorf("redis set key %s return %s err %v", key, s, err.Error())
	}
}

// get 命令
func Load(key string) []byte {
	if Client == nil {
		glog.Error("redis client is nil")
		return nil
	}
	reply, err := Client.Get(key).Bytes()
	if err != nil {
		glog.Errorf("redis get key %s err %s", key, err.Error())
		return nil
	}
	return reply
}

// set command
func SetData(key string, data interface{}, ttl int64) {
	if Client == nil {
		glog.Error("redis client is nil")
		return
	}
	s, err := Client.Set(key, data, time.Duration(ttl)).Result()
	if err != nil {
		glog.Errorf("redis set key %s return %s err %v", key, s, err.Error())
	}
}

// get command
func GetData(key string) string {
	if Client == nil {
		glog.Error("redis client is nil")
		return ""
	}
	reply := Client.Get(key).Val()
	return reply
}

//func ZAddData(key string, scores int64, member string) {
//	if Client == nil {
//		glog.Error("redis client is nil")
//		return
//	}
//	zaddArgMember := redis.Z{Score: float64(scores), Member: member}
//	add := Client.ZAdd(key, zaddArgMember)
//}

// HSet command
func HSetData(key, filed, value string) {
	if Client == nil {
		glog.Error("redis client is nil")
		return
	}
	s, e := Client.HSet(key, filed, value).Result()
	if e != nil {
		glog.Errorf("redis set key %s return %s err %v", key, s, e.Error())
	}
}

// HGet command redis返回nil时 取到的result为空字符串("")
func HGetData(key, filed string) (string, error) {
	if Client == nil {
		glog.Error("redis client is nil")
		return "", errors.New("redis is nil")
	}
	result := Client.HGet(key, filed).Val()
	return result, nil
}

// HIncrby command
func HIncrByData(key, filed string, countValue int64) {
	if Client == nil {
		glog.Error("redis client is nil")
		return
	}
	result, err := Client.HIncrBy(key, filed, countValue).Result()
	if err != nil {
		glog.Errorf("redis HIncrBy key %s filed %s return %s err %v", key, filed, result, err.Error())
	}
}

// Exist command
func ExistsDatum(key string) (bool, error) {
	if Client == nil {
		glog.Error("redis client is nil")
		return false, errors.New("redis client is nil")
	}
	result, _ := Client.Exists(key).Result()
	if result == int64(1) {
		return true, nil
	} else {
		return false, nil
	}
}

// Expire command 设定key的ttl命令
func ExpireData(key string, timeSeconds time.Duration) bool {
	if Client == nil {
		glog.Error("redis client is nil")
		return false
	}
	result, _ := Client.Expire(key, timeSeconds).Result()
	return result
}
