package helpers

import (
	"errors"

	"github.com/bipuldan/MeAndYou/auth-service/common"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("unauthhorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == common.USER && uid != userId {
		err = errors.New("unauthhorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
