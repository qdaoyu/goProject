package achieve

import (
	"log"
	"qiudaoyu/models"
	"qiudaoyu/models/menuInfo"
)

// 获取塑颜业绩表
func GetSyAchieve(uid int, username string, currentPage int, size int) (map[string]interface{}, error) {
	var syAchieve []SyAchieve
	var user menuInfo.User
	var total int64
	//存储信息
	syMap := make(map[string]interface{})
	err := models.Conn.Table("t_admin").Where("id = ? ", uid).Find(&user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//分页的固定写法
	offsetVal := (currentPage - 1) * size

	rid := user.Roleid
	// log.Println("user:", user)
	// log.Println("roleid:", rid)
	if rid == 1 || rid == 2 {
		err := models.Conn.Table("t_syachieve").Count(&total).Limit(size).Offset(offsetVal).Find(&syAchieve).Error
		log.Println("条数：", total)
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
			syMap["total"] = total
			return syMap, nil
		}
	} else {
		//limit offset
		err := models.Conn.Table("t_syachieve").Where("depart = ? ", username).Count(&total).Limit(size).Offset(offsetVal).Find(&syAchieve).Error
		log.Println("条数：", total)
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
			syMap["total"] = total
			return syMap, nil
		}
	}

}
