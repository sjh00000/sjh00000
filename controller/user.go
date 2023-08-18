package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
	"tiktok/pjdata"
)

// 此map用于查找用户是否存在
var usersLoginInfo = map[string]pjdata.Author{
	"孙佳豪123456": {
		Id:            1,
		Name:          "孙佳豪",
		FollowCount:   0,
		FollowerCount: 99999999,
		IsFollow:      true,
		TotalFavorite: 99999999,
	},
}

// id生成器
var userIdSequence int64 = 1

// UserLoginResponse 返回登录，注册信息
type UserLoginResponse struct {
	pjdata.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// UserResponse 返回用户信息
type UserResponse struct {
	pjdata.Response
	User pjdata.Author `json:"user"`
}

// Register 注册
func Register(c *gin.Context) {

	//获取查询参数中名为username的值
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	/**
	检查账号是否已经注册：
	是：
	否：
	**/
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: pjdata.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		//对要增加的变量进行原子操作，避免并发问题
		atomic.AddInt64(&userIdSequence, 1)

		//创建新的账户，并添加token映射
		newUser := pjdata.Author{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser

		//返回用户数据
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: pjdata.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

// Login 登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	//登陆成功返回用户数据，失败返回不存在
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: pjdata.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: pjdata.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	//存在用户返回用户信息，否则返回用户不存在
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: pjdata.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: pjdata.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
