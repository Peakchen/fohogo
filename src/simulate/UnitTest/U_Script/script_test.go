package U_Script

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/RedisConn"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/gomodule/redigo/redis"
)

var (
	// These variables are declared at package level to remove distracting
	// details from the examples.
	c     redis.Conn
	reply interface{}
	err   error
)

func ExampleScript() {
	// Initialize a package-level variable with a script.
	var getScript = redis.NewScript(1, `return redis.call('get', KEYS[1])`)

	// In a function, use the script Do method to evaluate the script. The Do
	// method optimistically uses the EVALSHA command. If the script is not
	// loaded, then the Do method falls back to the EVAL command.
	reply, err = getScript.Do(c, "foo")
}

func TestScript(t *testing.T) {
	rediscfg := serverConfig.GRedisconfigConfig.Get(0)
	redisconn = RedisConn.NewRedisConn(rediscfg.Connaddr, rediscfg.DBIndex, rediscfg.Passwd)

	c := redisconn.RedPool.Get()
	// To test fall back in Do, we make script unique by adding comment with current time.
	script := fmt.Sprintf("--%d\nreturn {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}", aktime.Now().UnixNano())
	s := redis.NewScript(2, script)
	reply := []interface{}{[]byte("key1"), []byte("key2"), []byte("arg1"), []byte("arg2")}

	v, err := s.Do(c, "key1", "key2", "arg1", "arg2")
	if err != nil {
		t.Errorf("s.Do(c, ...) returned %v", err)
	}

	if !reflect.DeepEqual(v, reply) {
		t.Errorf("s.Do(c, ..); = %v, want %v", v, reply)
	}

	var (
		params = []string{}
	)
	for _, param := range v.([]interface{}) {
		if reflect.TypeOf(param) == reflect.TypeOf([]byte{}) {
			params = append(params, string(param.([]byte)))
		}
	}

	if len(params) > 0 {
		akLog.FmtPrintf("run script params: %v.", params)
	}

}

func TestScript1(t *testing.T) {
	rediscfg := serverConfig.GRedisconfigConfig.Get(0)
	redisconn = RedisConn.NewRedisConn(rediscfg.Connaddr, rediscfg.DBIndex, rediscfg.Passwd)

	c := redisconn.RedPool.Get()

	srckeys := []string{"aa", "bb"}
	commonkey := "bigkey"
	for _, k := range srckeys {
		c.Do("HSET", commonkey, k, 1)
	}

	s := redis.NewScript(1, `local k1 = KEYS[1] 
							local a1 = ARGV[1] 
							redis.call('set', k1, a1)`)

	_, err := s.Do(c, "aa", "123")
	if err != nil {
		t.Errorf("s.Do(c, ...) returned %v", err)
	}

	_, err = s.Do(c, "bb", "456")
	if err != nil {
		t.Errorf("s.Do(c, ...) returned %v", err)
	}

	s2 := redis.NewScript(1, `local k1 = KEYS[1] 
							 return redis.call('get', k1)`)
	v2, err := s2.Do(c, "aa")
	if err != nil {
		t.Errorf("s.Do(c, ...) returned %v", err)
	}

	var (
		params = []string{}
	)

	if reflect.TypeOf(v2) == reflect.TypeOf([]byte{}) {
		params = append(params, string(v2.([]byte)))
	}

	if len(params) > 0 {
		akLog.FmtPrintf("run script params: %v.", params)
	}

	getkeys, err := c.Do("HKEYS", commonkey)
	if err != nil {
		return
	}

	for _, k := range getkeys.([]interface{}) {
		dstkey := string(k.([]byte))
		dstval, err := c.Do("GET", dstkey)
		if err != nil {
			akLog.FmtPrintln("get fail, err: ", err)
			continue
		}

		akLog.FmtPrintf("save val: %v. ", string(dstval.([]byte)))
	}
}
