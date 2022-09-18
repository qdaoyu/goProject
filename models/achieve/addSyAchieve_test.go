package achieve

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Achieve struct {
	Id              int
	Area            string
	Depart          string
	Name            string
	Post            string
	Hireday         time.Time
	St_attendance   float64
	Ac_attendance   float64
	Restday         float64
	Leaveday        float64
	Absenteeism     float64
	Yq_close        float64
	Overtime        float64
	Business_travel float64
	Basesalary      float64
	Wx_achieve      float64
	Pf_achieve      float64
	Zc_achieve      float64
	Gd_achieve      float64
	Bs_achieve      float64
	Shop_achieve    float64
	Wxpf_commission float64
	Zc_commission   float64
	Gd_commission   float64
	Bs_commission   float64
	Shop_commission float64
	Achievement     float64
	Subsidy         float64
	Handcost        float64
	Ac_basesalary   float64
	Deduction       float64
	Month_salary    float64
}

func StrToInt(str string) int {
	if str == "" {
		return 0
	} else {
		res, err := strconv.Atoi(str)
		if err != nil {
			log.Println(err)
			return -1111
		}
		return res

	}

}

func StrToFloat(str string) float64 {
	if str == "" {
		return 0
	} else {
		res, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Println(err)
			return -1111.11
		}
		return res
	}

}

func TestAddSyAchieve(t *testing.T) {
	//读取excel
	dst := "d:/achieve.xlsx"
	xlsx, err := excelize.OpenFile(dst)
	if err != nil {
		log.Println(err)
		return
	}
	rows, err := xlsx.Rows("Sheet1")
	if err != nil {
		log.Println(err)
		return
	}
	j := 0
	//声明一个数组
	var datas []Achieve
	var data Achieve

	for rows.Next() {
		row := rows.Columns()
		if j == 0 {
			//判断表抬头是否正确
			log.Println("以后验证抬头")
			j++
			continue
		}
		log.Println(row)
		data.Id = StrToInt(row[0])
		data.Area = row[1]
		data.Depart = row[2]
		data.Name = row[3]
		data.Post = row[4]

		currentDate, err := ExcelDateToDate(row[5])
		if err != nil {
			log.Println("塑颜业绩表入职日期转化失败:", err)
		}
		data.Hireday = currentDate
		data.St_attendance = StrToFloat(row[6])
		data.Ac_attendance = StrToFloat(row[7])
		data.Restday = StrToFloat(row[8])
		data.Leaveday = StrToFloat(row[9])
		data.Absenteeism = StrToFloat(row[10])
		data.Yq_close = StrToFloat(row[11])
		data.Overtime = StrToFloat(row[12])
		data.Business_travel = StrToFloat(row[13])
		data.Basesalary = StrToFloat(row[14])
		data.Wx_achieve = StrToFloat(row[15])
		data.Pf_achieve = StrToFloat(row[16])
		data.Zc_achieve = StrToFloat(row[17])
		data.Gd_achieve = StrToFloat(row[18])
		data.Bs_achieve = StrToFloat(row[19])
		data.Shop_achieve = StrToFloat(row[20])
		data.Wxpf_commission = StrToFloat(row[21])
		data.Zc_commission = StrToFloat(row[22])
		data.Gd_commission = StrToFloat(row[23])
		data.Bs_commission = StrToFloat(row[24])
		data.Shop_commission = StrToFloat(row[25])
		data.Achievement = StrToFloat(row[26])
		data.Subsidy = StrToFloat(row[27])
		data.Handcost = StrToFloat(row[28])
		data.Ac_basesalary = StrToFloat(row[29])
		data.Deduction = StrToFloat(row[30])
		data.Month_salary = StrToFloat(row[31])

		datas = append(datas, data)
		j++
	}
	// log.Println(datas)
	AddAchieve(datas)
	// return
}

// 批量新增塑颜业绩表
func AddAchieve(Achieve []Achieve) {
	// 数据库连接
	// 定义数据地址
	var Dsn string = "root:qiudaoyu@tcp(127.0.0.1:3306)/qiudaoyu?charset=utf8mb4&parseTime=True&loc=Local"
	// var Conn *gorm.DB = nil

	Conn, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB, _ := Conn.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(200)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	err = Conn.Table("t_syachieve").Create(&Achieve).Error
	if err == nil {
		log.Println(err)
		return
	}

}

// for i, row := range rows {
// 	// 去掉第一行是execl表头部分
// 	if i == 0 {
// 		continue
// 	}
// 	var data SyAchieve
// 	for k, v := range row {
// 		// 第一列是员工姓名
// 		if k == 0 {
// 			data.Name = v
// 		}
// 		//
// 		if k == 1 {
// 			log.Println("日期起始:", v)
// 			// currentDate := row[1]
// 			currentDate, err := ExcelDateToDate(v)
// 			if err != nil {
// 				log.Println(err)
// 				return

// 			}
// 			// log.Println("current_date is " + currentDate.Format("2006-01-02 15:04:05"))
// 			data.Date = currentDate
// 			log.Println("日期结束:", data.Date)
// 		}
// 		//
// 		if k == 2 {
// 			log.Println("业绩起始:", v)
// 			vInt, _ := strconv.Atoi(v)
// 			data.Achieve = vInt
// 			log.Println("业绩结束:", data.Achieve)
// 		}
// 	}
// 	//将数据追加到datas数组中
// 	datas = append(datas, data)
