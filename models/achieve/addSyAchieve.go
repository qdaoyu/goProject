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
	// ID           int       `json:"id" `
	Name    string    `json:"name" `
	Date    time.Time `json:"date" `
	Achieve int       `json:"achieve" `
}

// 批量新增塑颜业绩表
func AddSyAchieve(syAchieve []SyAchieve) {
	var err error = models.Conn.Table("t_syAchieve").Create(&syAchieve).Error
	if err == nil {
		log.Println(err)
		return
	}

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
	rows := xlsx.GetRows("Sheet1")
	//声明一个数组
	var datas []SyAchieve
	for i, row := range rows {
		// 去掉第一行是execl表头部分
		if i == 0 {
			log.Println(row)

		}
		var data SyAchieve
		for k, v := range row {
			// 第一列是员工姓名
			if k == 0 {

				log.Println("姓名起始:", v)
				data.Name = v
				log.Println("姓名结束:", data.Name)
			}
			//
			if k == 1 {
				log.Println("日期起始:", v)
				// currentDate := row[1]
				currentDate, err1 := ExcelDateToDate(v)
				if err1 != nil {
					log.Println(err1)
					return err1

				}
				// log.Println("current_date is " + currentDate.Format("2006-01-02 15:04:05"))
				data.Date = currentDate
				log.Println("日期结束:", data.Date)
			}
			//
			if k == 2 {
				log.Println("业绩起始:", v)
				vInt, _ := strconv.Atoi(v)
				data.Achieve = vInt
				log.Println("业绩结束:", data.Achieve)
			}
		}
		//将数据追加到datas数组中
		datas = append(datas, data)

	}

	AddSyAchieve(datas)
	return nil
}
