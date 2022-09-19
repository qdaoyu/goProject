package base

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"qiudaoyu/middleWare"
	"qiudaoyu/models/achieve"
	"qiudaoyu/models/menuInfo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Situ01 []string
	Situ02 []string
	Situ03 []string
	Situ04 []string
	Situ05 []string
}

// ----token部分----
type UserInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func UploadSyAchieveTb(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println(file.Filename)

		dst := "./assets/fileRec/" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}

func Login(c *gin.Context) {
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	var user UserInfo
	err := c.ShouldBind(&user)
	fmt.Println("userInfo:", user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    2001,
			"message": "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确,数据库取用户信息
	//调用数据库查询，根据返回值，判定是否为t_admin表里的用户
	res, err2 := menuInfo.LoginConfirm(user.Username, user.Password)
	if err2 == nil {
		// 生成Token
		fmt.Println("res:", res)
		// var userStruct menuInfo.User
		// err3 := json.Unmarshal(res, &userStruct)
		// resByte, _ := reflect.TypeOf(res["userInfo"]).FieldByName("ID")
		// fmt.Println("res:", reflect.TypeOf(res))
		// fmt.Println(reflect.TypeOf(res["userInfo"]).FieldByName("ID"))
		// fmt.Println(reflect.ValueOf(res["userInfo"]).FieldByName("ID").Int())  res["userInfo"].ID
		fmt.Println(res["userInfo"])
		// res.
		// fmt.Println(res["userInfo"].ID)

		//类型断言
		userId, ok := res["userInfo"].(menuInfo.User)
		if ok {
			fmt.Println(userId.ID)
			// tokenString, _ := middleWare.GenToken(user.Username, reflect.ValueOf(res["userInfo"]).FieldByName("ID").Int())
			tokenString, _ := middleWare.GenToken(user.Username, int64(userId.ID))
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "登录成功",
				"data":    gin.H{"token": tokenString, "userInfo": res["userInfo"]},
			})
			return

		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    2002,
				"message": "用户名或密码错误",
			})
		}
		// tokenString, _ := middleWare.GenToken(user.Username, reflect.ValueOf(res["userInfo"]).FieldByName("ID").Int())

		return
	}
}

func HomeMenuHandler(c *gin.Context) {
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	userID, err1 := strconv.Atoi(c.Request.Header.Get("userID"))
	if err1 != nil {
		c.JSON(5001, gin.H{
			"code":    5001,
			"message": "获取菜单失败",
			"data":    gin.H{},
		})
		return
	}
	res, err := menuInfo.GetMenuDb(userID)
	fmt.Println(res)
	if err != nil {
		c.JSON(5001, gin.H{
			"code":    5001,
			"message": "获取菜单失败",
			"data":    gin.H{},
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取菜单成功",
			"data":    res,
		})
		return
	}
}

func GetUserInfoHandler(c *gin.Context) {
	//获取menu菜单
	res, err := menuInfo.GetUserInfo(c)
	fmt.Println(res)
	if err != nil {
		c.JSON(5001, gin.H{
			"code":    5001,
			"message": "获取用户信息失败",
			"data":    gin.H{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取用户信息成功",
			"data":    res,
		})
		return
	}
}

func GetUserBasicInfoHandler(c *gin.Context) {
	//获取menu菜单
	userID, _ := strconv.Atoi(c.Request.Header.Get("userID"))
	res, err := menuInfo.GetUserBasicInfo(c, userID)
	fmt.Println(res)
	if err != nil {
		c.JSON(5001, gin.H{
			"code":    5001,
			"message": "获取用户基本信息失败",
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取用户基本信息成功",
			"data":    res,
		})
		return
	}
}
func AddSyAchieveInfoHandler(c *gin.Context) {

	// userID, _ := strconv.Atoi(c.Request.Header.Get("userID"))
	err := achieve.SyAchieveExcelize(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    5001,
			"message": "插入塑颜业绩数据库失败,原因:" + err.Error(),
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "插入塑颜业绩数据库成功",
		})
		return
	}
}

// 获取塑颜业绩表信息(管理员_id1和测试角色_id2默认可以返回所有数据)
func GetSyAchieveInfoHandler(c *gin.Context) {
	var syMap = make(map[string]interface{})

	userID, _ := strconv.Atoi(c.Request.Header.Get("userID"))
	//类型断言
	userName, _ := c.Get("username")
	userNameAssert, ok := userName.(string)
	// userName, _ := c.Get("username")
	if !ok {
		c.JSON(200, gin.H{
			"code":    5003,
			"message": "用户名断言失败",
			"data":    nil,
		})
		return
	}
	syMap, err := achieve.GetSyAchieve(userID, userNameAssert)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    5002,
			"message": syMap["message"],
			"data":    nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": syMap["message"],
			"data":    syMap["data"],
		})
		return
	}
}
func CalDupName(c *gin.Context) {
	// 注意：下面为了举例子方便，暂时忽略了错误处理
	b, err := c.GetRawData() // 从c.Request.Body读取请求数据
	fmt.Println(b)
	if err == nil {
		// 定义map或结构体
		var m map[string]string
		// 反序列化
		// fmt.Printf()
		_ = json.Unmarshal(b, &m)
		arg2 := m["textArea"]
		arg1 := m["label"]
		fmt.Println("----以下为接收到的参数-----")
		fmt.Println(arg1)
		fmt.Println(arg2)
		// fmt.Println("---------")
		//执行python脚本-------
		cmd := exec.Command("python", "D:/tempdownload/重名判断.py", arg1, arg2)
		out, _ := cmd.CombinedOutput()
		fmt.Println("concatenation: ", string(out))
		msg := string(out)
		var data Data
		_ = json.Unmarshal([]byte(msg), &data)
		fmt.Println(string(out))
		//-----------------------
		// fmt.Println(m)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "计算成功",
			"data":    data,
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    50001,
			"message": "输入内容非法",
			"data":    gin.H{},
		})
	}

}

func Page404(c *gin.Context) {
	// c.HTML(http.StatusNotFound, "views/404.html", nil)
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/page404")
}

// func AuthHandler(c *gin.Context) {
// 	// 用户发送用户名和密码过来
// 	var user UserInfo
// 	err := c.ShouldBind(&user)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": 2001,
// 			"msg":  "无效的参数",
// 		})
// 		return
// 	}
// 	// 校验用户名和密码是否正确
// 	if user.Username == "q1mi" && user.Password == "q1mi123" {
// 		// 生成Token
// 		tokenString, _ := middleWare.GenToken(user.Username)
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": 2000,
// 			"msg":  "success",
// 			"data": gin.H{"token": tokenString},
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 2002,
// 		"msg":  "鉴权失败",
// 	})
// 	return
// }

//---------------
