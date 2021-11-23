package v1

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/service"
)

type User struct {
	session *Session
	authorizationService service.Authorization
}

func NewUser(
	session *Session,
	authorizationService service.Authorization,
) *User {
	return &User{
		session: session,
		authorizationService: authorizationService,
	}
}

type UserInfo struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

func (u *User) PostSignup(c *gin.Context) {
	var userInfo UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := u.authorizationService.Signup(
		c.Request.Context(),
		values.NewUserName(userInfo.Name),
		values.NewUserPassword([]byte(userInfo.Password)),
	)
	if errors.Is(err, service.ErrUserAlreadyExists) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user already exists",
		})
		return
	}
	if err != nil {
		log.Printf("failed to signup: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to signup",
		})
		return
	}

	session := u.session.getSession(c)
	u.session.setUser(session, user)
	err = u.session.save(c, session)
	if err != nil {
		log.Printf("failed to save session: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save session",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (u *User) PostLogin(c *gin.Context) {
	var userInfo UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := u.authorizationService.Login(
		c.Request.Context(),
		values.NewUserName(userInfo.Name),
		values.NewUserPassword([]byte(userInfo.Password)),
	)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user or password",
		})
		return
	}
	if err != nil {
		log.Printf("failed to login: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to login",
		})
		return
	}

	session := u.session.getSession(c)
	u.session.setUser(session, user)
	err = u.session.save(c, session)
	if err != nil {
		log.Printf("failed to save session: %v\n", err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (u *User) GetMe(c *gin.Context) {
	session := u.session.getSession(c)
	user, err := u.session.getUser(session)
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user",
		})
		return
	}

	c.JSON(http.StatusOK, &UserInfo{
		Name: string(user.GetName()),
	})
}

func (u *User) PatchMe(c *gin.Context) {
	var userInfo UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	session := u.session.getSession(c)
	user, err := u.session.getUser(session)
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user",
		})
		return
	}

	user, err = u.authorizationService.UpdateUserInfo(
		c.Request.Context(),
		user,
		values.NewUserName(userInfo.Name),
		values.NewUserPassword([]byte(userInfo.Password)),
	)
	if err != nil {
		log.Printf("failed to update user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update user",
		})
	}

	u.session.setUser(session, user)
	err = u.session.save(c, session)
	if err != nil {
		log.Printf("failed to save session: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save session",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (u *User) DeleteMe(c *gin.Context) {
	session := u.session.getSession(c)
	user, err := u.session.getUser(session)
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user",
		})
		return
	}

	err = u.authorizationService.DeleteAccount(
		c.Request.Context(),
		user,
	)
	if err != nil {
		log.Printf("failed to delete user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete user",
		})
		return
	}

	u.session.revoke(session)
	err = u.session.save(c, session)
	if err != nil {
		log.Printf("failed to save session: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save session",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
