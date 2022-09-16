package menuInfo

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 结构体
type T_menu struct {
	// gorm.Model         //自动添加主键
	ID          int      `gorm:"type:int(0) " json:"id"  binding:"required"`
	Url         string   `gorm:"type:varchar(64)" json:"url"  binding:"required"`
	Path        string   `gorm:"type:varchar(64)" json:"path" binding:"required"`
	Component   string   `gorm:"type:varchar(64) " json:"component"  binding:"required"`
	Name        string   `gorm:"type:varchar(64) " json:"name"  binding:" required"`
	IconCls     string   `gorm:"type:varchar(64) " json:"iconCls"  binding:"required"`
	KeepAlive   int      `gorm:"type:int(0) " json:"keepAlive"   binding:"required"`
	RequireAuth int      `gorm:"type:int(0) " json:"requireAuth"  binding:"required"`
	ParentId    int      `gorm:"type:int(0) " column:"parentId"  binding:"required"`
	Enabled     int      `gorm:"type:int(0) " json:"enabled"  binding:"required"`
	Children    []T_menu `gorm:"type:varchar(64) " json:"children" binding:"required"`
}

func GetMenuDb() (gin.H, error) {
	dsn := "root:qiudaoyu@tcp(127.0.0.1:3306)/qiudaoyu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	var dataList []T_menu
	// p_dataList := &dataList
	// db.Model(&dataList).Limit(-1).Offset(-1).Find(&dataList)
	db.Table("t_menu").Find(&dataList)
	//查询数据库  返回总数  limit 跟offset 参数如果是-1，就是无限制
	if len(dataList) == 0 {
		return gin.H{"message": "菜单获取失败"}, errors.New("菜单获取失败")
	} else {
		// fmt.Println(dataList)
		//对json进行处理
		// 先循环遍历看下
		for i := range dataList {
			idTemp := dataList[i].ID
			for j := range dataList {
				if idTemp != dataList[j].ID {
					if dataList[j].ParentId == idTemp {
						// dataList[i].Name = dataList[i].Name + fmt.Sprint(i)
						// fmt.Println("------------")
						// fmt.Println("原datalist", dataList[i])
						dataList[i].Children = append(dataList[i].Children, dataList[j])
						// fmt.Println("后datalist", dataList[i])
					}

				}
			}
		}

		//剔除掉子菜单
		for i := 0; i < len(dataList); i++ {
			if dataList[i].Component != "Home" {
				dataList = append(dataList[:i], dataList[i+1:]...)
				i--
			}
		}
		// fmt.Println(dataList)
		//查询到数据
		return gin.H{
			"data":    dataList,
			"message": "查询成功",
		}, nil
	}

}
