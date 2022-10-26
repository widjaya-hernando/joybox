package middleware

import (
	"errors"
	"net/http"
	"sleekflow/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Authorization(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		c.JSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, errors.New("the request is unauthorized")))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	bearerToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	jwt, err := helper.ParseJwt(bearerToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, err))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user_id", jwt.UserID)
	c.Next()
}

// func (m *Middleware) AuthorizationAdmin(c *gin.Context) {
// 	role := c.MustGet("aggregator_name").(string)
// 	if role != constants.AdminRole {
// 		c.JSON(http.StatusForbidden, helper.NewResponse(http.StatusForbidden, nil, errors.New("you do not have authorization to access")))
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}
// 	log.Println(role)
// }

// func (m *Middleware) AuthorizationUser(c *gin.Context) {
// 	role := c.MustGet("role").(string)
// 	if role != constants.MemberRole {
// 		c.JSON(http.StatusForbidden, helper.NewResponse(http.StatusForbidden, nil, errors.New("your role is not allowed")))
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}
// 	log.Println(role)
// }
