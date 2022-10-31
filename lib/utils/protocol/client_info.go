/*
*

	@author: yaoqiang
	@date: 2021/11/9
	@note:

*
*/
package protocol

import (
	"match/lib/result"
)

func (c *ClientInfo) IsValid() (code int32, msg string) {
	/*	if c.AppName == "" {
			log.Error("AppName is empty string")
			return result.Param, "the AppName of request invalid"
		}
		if c.AppId == 0 {
			log.Error("AppId is empty string")
			return result.Param, "the AppId of request invalid"
		}
		if c.VersionName == "" {
			log.Error("VersionName is empty string")
			return result.Param, "the VersionName of request invalid"
		}
		if c.VersionCode == 0 {
			log.Error("VersionCode is empty string")
			return result.Param, "the VersionCode of request invalid"
		}*/
	//if !version.IsValid(c.VersionName, int(c.VersionCode)) {
	//	log.WithField("VersionCode", c.VersionCode).WithField("VersionName", c.VersionName).Error("version.IsValid return false")
	//	return result.Param, "VersionName or VersionCode invalid"
	//}
	return result.Success, result.SuccessMsg
}
