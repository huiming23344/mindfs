package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/mindfs/metaServer/meta"
	"github.com/huiming23344/mindfs/metaServer/server"
)

func AddUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "add user",
	})
	user := &meta.User{}
	err := c.ShouldBind(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}
	err = server.AddUser(user.Name, user.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete user",
	})
	name := c.Param("name")
	err := server.DeleteUser(name)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func ListUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list users",
	})
	users := server.ListUsers()
	c.JSON(200, gin.H{
		"users": users,
	})
}

func AddGroup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "add group",
	})
	group := &meta.UserGroup{}
	err := c.ShouldBind(group)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}
	err = server.AddGroup(group.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func DeleteGroup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete group",
	})
	name := c.Param("name")
	err := server.DeleteGroup(name)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func AddUserToGroup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "add user to group",
	})
	userName := c.Param("user")
	groupName := c.Param("group")
	err := server.AddUserToGroup(userName, groupName)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func RemoveUserFromGroup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "remove user from group",
	})
	userName := c.Param("user")
	groupName := c.Param("group")
	err := server.RemoveUserFromGroup(userName, groupName)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
}
