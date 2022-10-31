/*
*

	@author: yaoqiang
	@date: 2021/11/9
	@note:

*
*/
package protocol

import (
	"github.com/mediocregopher/radix/v3"
	"match/lib/redis"
	"match/lib/result"
	"strings"
)

/**
 * @Author: yaoqiang
 * @Description: 验证token并返回account
 * @Date: 2021/11/10 下午2:37
 * @Param:
 * @return:
 **/
func VerifyToken(token string) (string, int32) {
	cli := redis.GetClient("match")
	var num int
	err := cli.Do(radix.FlatCmd(&num, "SISMEMBER", "match_token", token))
	if err != nil || num == 0 {
		return "", result.Unknown
	}

	tokenArray := strings.Split(token, "-")
	//userId, err := strconv.ParseInt(tokenArray[len(tokenArray)-1], 10, 64)
	//if err != nil {
	//	log.Println("VerifyToken", err)
	//	return "", result.Unknown
	//}
	return tokenArray[len(tokenArray)-1], result.Success
}
