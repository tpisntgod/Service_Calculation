package client

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	min     = 1000000
	max     = 9999999
	urlhead = "http://api.fanyi.baidu.com/api/trans/vip/translate"
	from    = "en"
	to      = "zh"
	appid   = "20180103000111810"
	key     = "7N1Csy2YSMhU0JcOjCIf"
)

type resultJSON struct {
	From        string              `json:"from"`
	To          string              `json:"to"`
	TransResult []map[string]string `json:"trans_result"`
}

func getSalt() string {
	rand.Seed(time.Now().Unix())
	t := min + rand.Intn(max-min)
	return strconv.Itoa(t)
}

func getMD5(s string) string {
	md5Obj := md5.New()
	md5Obj.Write([]byte(s))
	digest := md5Obj.Sum(nil)
	return hex.EncodeToString(digest)
}

func getURL(q string) string {
	salt := getSalt()
	url := urlhead + "?q=" + q + "&from=" + from + "&to=" + to +
		"&appid=" + appid + "&salt=" + salt + "&sign=" + getMD5(appid+q+salt+key)
	return url
}

//HTTPGetSync 分别使用同步和异步方法向百度翻译发送5个翻译请求
func HTTPGetSync(q string) string {
	res, err := http.Get(getURL(q))
	CheckErr(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	CheckErr(err)
	var result resultJSON
	if err = json.Unmarshal(body, &result); err == nil {
		return fmt.Sprintln(result.TransResult[0]["src"], ":", result.TransResult[0]["dst"])
	}
	return q + "failed when parsing json response!\n"
}

//HTTPGetAsync 使用 go HTTPClient 实现图 6-2 的 Naive Approach
func HTTPGetAsync(q string, ch chan string) {
	res, err := http.Get(getURL(q))
	CheckErr(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	CheckErr(err)
	var result resultJSON
	err = json.Unmarshal(body, &result)
	CheckErr(err)
	ch <- fmt.Sprintln(result.TransResult[0]["src"], ":", result.TransResult[0]["dst"])
}

//CheckErr check err
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
