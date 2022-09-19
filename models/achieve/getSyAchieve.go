package achieve

import (
	"log"
	"qiudaoyu/models"
	"qiudaoyu/models/menuInfo"

	"github.com/gin-gonic/gin"
)

// 获取塑颜业绩表
func GetSyAchieve(uid int, username string) (map[string]interface{}, error) {
	var syAchieve []SyAchieve
	var user menuInfo.User
	//存储信息
	syMap := make(map[string]interface{})
	err := models.Conn.Table("t_admin").Where("id = ? ", uid).Find(&user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rid := user.Roleid
	// log.Println("user:", user)
	// log.Println("roleid:", rid)
	if rid == 1 || rid == 2 {
		err := models.Conn.Table("t_syachieve").Find(&syAchieve).Error
		if err != nil {
			log.Println(err)
			syMap["message"] = "查询失败" + err.Error()
			return syMap, err
		}
		// models.Conn.Raw(sqlString, userID).Scan(&user)
		// db.Raw(sqlString, userID).Create(&user)
		if len(syAchieve) == 0 {
			log.Println("数据库无数据")
			syMap["message"] = "数据库无数据"
			return syMap, err
		} else {
			syMap["data"] = syAchieve
			syMap["message"] = "查询成功"
			return syMap, nil
		}
	} else {
		err := models.Conn.Table("t_syachieve").Where("depart = ? ", username).Find(&syAchieve).Error
		if err != nil {
			log.Println(err)
			syMap["message"] = "查询失败" + err.Error()
			return syMap, err
		}
		// models.Conn.Raw(sqlString, userID).Scan(&user)
		// db.Raw(sqlString, userID).Create(&user)
		if len(syAchieve) == 0 {
			log.Println("无此人数据")
			syMap["message"] = "查无数据"
			return syMap, err
		} else {
			syMap["data"] = syAchieve
			syMap["message"] = "查询成功"
			return gin.H{"data": syAchieve}, nil
		}
	}

}
