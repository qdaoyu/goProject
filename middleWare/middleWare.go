package middleWare

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"username"`
	Userid   int64  `json:"userid"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 3

var MySecret = []byte("夏天夏天悄悄过去")

// GenToken 生成JWT
func GenToken(username string, userid int64) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		userid,   //用户ID
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "qiudaoyu",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		//当请求登录接口时，跳过token验证:
		// fmt.Println(c.FullPath())
		if c.FullPath() != "/login" && c.FullPath() != "/admin/info" {
			// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
			// 这里假设Token放在Header的Authorization中，并使用Bearer开头
			// 这里的具体实现方式要依据你的实际业务情况决定
			authHeader := c.Request.Header.Get("Authorization")
			uid, _ := strconv.Atoi(c.Request.Header.Get("userID"))
			fmt.Printf("token:%s\n", authHeader)
			if authHeader == "" {
				// c.Redirect(http.StatusMovedPermanently, "http://localhost:8080")
				c.JSON(http.StatusOK, gin.H{
					"code":    2003,
					"message": "请求头中auth为空,请重新登录",
				})
				// c.Redirect(http.StatusMovedPermanently, "http://localhost:8080")
				c.Abort()
				return
			}
			// 按空格分割
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				c.JSON(http.StatusOK, gin.H{
					"code":    2004,
					"message": "请求头中auth格式有误",
				})
				c.Abort()
				return
			}
			// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
			// fmt.Println("----------------------")
			mc, err := ParseToken(parts[1])
			// fmt.Println("mc", mc)
			// fmt.Println("mc", mc.Username)
			// fmt.Println("mc", mc.Userid)
			//判断userid是否有误
			if mc.Userid != int64(uid) {
				c.JSON(http.StatusOK, gin.H{
					"code":    2006,
					"message": "userid错误,与token存储信息匹配失败",
				})
				c.Abort()
				return
			}

			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    2005,
					"message": "无效的Token",
				})
				c.Abort()
				return
			}
			// 将当前请求的username信息保存到请求的上下文c上
			c.Set("username", mc.Username)
			log.Println(c.Get("username"))
			c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		} else {
			c.Next()
		}

	}
}

//------------

//------------------------------------------

// StatCost 是一个统计耗时请求耗时的中间件：记录接口耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name1", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

// 处理跨域
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
