package base

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

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
// func authHandler(c *gin.Context) {
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
// 		tokenString, _ := GenToken(user.Username)
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

//------------

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println(file.Filename)

		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Printf("username:%s  password:%s \n", username, password)
	if username == "qiudaoyu" && password == "123" {
		fmt.Printf("返回数据1")
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			// "data": gin.H{
			// 	"code":    200,
			// 	"message": "登录成功",
			// },
			"message": "登录成功",
		})
	} else {
		fmt.Printf("返回数据")
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "帐号密码错误",
		})

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
			"code": 200,
			"data": data,
		})

	} else {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    50001,
			"message": "输入内容非法",
		})
	}

}

func Page404(c *gin.Context) {
	// c.HTML(http.StatusNotFound, "views/404.html", nil)
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/page404")
}
