package achieve

import (
	"log"
	"qiudaoyu/models"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

type SyAchieve struct {
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

// 批量新增塑颜业绩表
func AddSyAchieve(syAchieve []SyAchieve) error {
	var err error = models.Conn.Table("t_syachieve").Create(&syAchieve).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func ExcelDateToDate(excelDate string) (time.Time, error) {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, err = strconv.ParseFloat(excelDate, 64)
	if err != nil {
		log.Println(err)
		return time.Date(1900, 1, 30, 0, 0, 0, 0, time.UTC), err
	}
	return excelTime.Add(time.Second * time.Duration(days*86400)), nil
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

// 获取上传的文件，存储后进行读取并存入数据库
func SyAchieveExcelize(c *gin.Context) error {
	file, err := c.FormFile("file")
	userID := c.Request.Header.Get("userID")
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println(file.Filename)
		dst := "./assets/upFile/" + userID + "_" + time.Now().GoString() + "_" + file.Filename
		// log.Println("dst:", dst)
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)

		// c.String(http.StatusOK, log.Sprintf("'%s' uploaded!", file.Filename))
		//读取excel
		xlsx, err := excelize.OpenFile(dst)
		if err != nil {
			log.Printf("open excel error:[%s]", err.Error())
			return err
		}
		//解析excel

		if err := ReadExcel(xlsx); err != nil {
			log.Println(err)
			return err
		}
		return nil

	}

}

// ReadExcel
func ReadExcel(xlsx *excelize.File) error {
	//根据名字获取cells的内容，返回的是一个[][]string
	// rows := xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex()))
	rows, err := xlsx.Rows("Sheet1")
	if err != nil {
		log.Println(err)
		return err
	}
	j := 0
	//声明一个数组
	var datas []SyAchieve
	var data SyAchieve

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
	err = AddSyAchieve(datas)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
