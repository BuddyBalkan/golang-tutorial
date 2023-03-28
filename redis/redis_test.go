package redis

import (
	"kk.com/time_tool"
	"testing"
)

const (
	RewardCoinLimitCount = "RewardCoinLimitCount"
)

func TestHSetData(t *testing.T) {
	err := Init()
	if err != nil {
		t.Fatalf("redis init err %s \n", err)
	}
	HSetData(RewardCoinLimitCount, "779598809:1:4", "1")
}

func TestHGetData(t *testing.T) {
	err := Init()
	if err != nil {
		t.Fatalf("redis init err %s \n", err)
	}
	data, err := HGetData(RewardCoinLimitCount, "779598809:1:4")
	if err != nil {
		t.Fatalf("redis read HGet happens err: %s", err)
	}
	t.Logf("the data is %q", data)
}

func TestExpireData(t *testing.T) {
	e := Init()
	if e != nil {
		t.Fatalf("redis init err %s \n", e)
	}
	ExpireData(RewardCoinLimitCount, time_tool.Get2NextDayZeroTimeSecond())
}

// 测试获取的nil判断问题
func TestHGetIfNotExistThenHSet(t *testing.T) {
	e := Init()
	if e != nil {
		t.Fatalf("redis init err: %s \n", e)
	}
	data, _ := HGetData(RewardCoinLimitCount, "779598809:1:4")
	if data == "" { // condition confirmed
		t.Logf("HGet from redis is nil : '%v' \n", data)
		HSetData(RewardCoinLimitCount, "779598809:1:4", "0")
	} else {
		t.Logf("HGet from redis is : '%v' \n", data)
	}
}
