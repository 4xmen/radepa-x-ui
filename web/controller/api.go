package controller

import (
	"x-ui/logger"
	"x-ui/web/service"
	"x-ui/web/session"

	"github.com/gin-gonic/gin"
)

type ClientForm struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type APIController struct {
	BaseController
	inboundController *InboundController
	userService       service.UserService
	settingService    service.SettingService
	tgBot             service.Tgbot
}

func NewAPIController(g *gin.RouterGroup) *APIController {
	a := &APIController{}
	a.initRouter(g)
	return a
}

func (a *APIController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/cpsess7945419007/frontend/jupiter/api/v1")

	g.POST("/login", a.login)
	g.POST("/getData", a.getClientTraffic)

	g.Use(a.checkLogin)

	g.GET("/list", a.getAllInbounds)
	g.GET("/get/:id", a.getSingleInbound)
	g.GET("/getClientTraffics/:email", a.getClientTraffics)
	g.POST("/add", a.addInbound)
	g.POST("/del/:id", a.delInbound)
	g.POST("/update/:id", a.updateInbound)
	g.POST("/clientIps/:email", a.getClientIps)
	g.POST("/clearClientIps/:email", a.clearClientIps)
	g.POST("/addClient", a.addInboundClient)
	g.POST("/:id/delClient/:clientId", a.delInboundClient)
	g.POST("/updateClient/:clientId", a.updateInboundClient)
	g.POST("/:id/resetClientTraffic/:email", a.resetClientTraffic)
	g.POST("/resetAllTraffics", a.resetAllTraffics)
	g.POST("/resetAllClientTraffics/:id", a.resetAllClientTraffics)
	g.POST("/delDepletedClients/:id", a.delDepletedClients)
	g.GET("/createbackup", a.createBackup)
	a.inboundController = NewInboundController(g)
}

func (a *APIController) getAllInbounds(c *gin.Context) {
	a.inboundController.getInbounds(c)
}

func (a *APIController) getSingleInbound(c *gin.Context) {
	a.inboundController.getInbound(c)
}

func (a *APIController) getClientTraffics(c *gin.Context) {
	a.inboundController.getClientTraffics(c)
}

func (a *APIController) getClientTraffic(c *gin.Context) {
	var form ClientForm
	err := c.ShouldBind(&form)
	if err != nil {
		pureJsonMsg(c, false, I18nWeb(c, "pages.login.toasts.invalidFormData"))
		return
	}
	if form.Email == "" {
		pureJsonMsg(c, false, I18nWeb(c, "pages.login.toasts.emptyUsername"))
		return
	}
	if form.Password == "" {
		pureJsonMsg(c, false, I18nWeb(c, "pages.login.toasts.emptyPassword"))
		return
	}
	a.inboundController.getClientTraffic(form.Email, form.Password, c)
}

func (a *APIController) addInbound(c *gin.Context) {
	a.inboundController.addInbound(c)
}

func (a *APIController) delInbound(c *gin.Context) {
	a.inboundController.delInbound(c)
}

func (a *APIController) updateInbound(c *gin.Context) {
	a.inboundController.updateInbound(c)
}

func (a *APIController) getClientIps(c *gin.Context) {
	a.inboundController.getClientIps(c)
}

func (a *APIController) clearClientIps(c *gin.Context) {
	a.inboundController.clearClientIps(c)
}

func (a *APIController) addInboundClient(c *gin.Context) {
	a.inboundController.addInboundClient(c)
}

func (a *APIController) delInboundClient(c *gin.Context) {
	a.inboundController.delInboundClient(c)
}

func (a *APIController) updateInboundClient(c *gin.Context) {
	a.inboundController.updateInboundClient(c)
}

func (a *APIController) resetClientTraffic(c *gin.Context) {
	a.inboundController.resetClientTraffic(c)
}

func (a *APIController) resetAllTraffics(c *gin.Context) {
	a.inboundController.resetAllTraffics(c)
}

func (a *APIController) resetAllClientTraffics(c *gin.Context) {
	a.inboundController.resetAllClientTraffics(c)
}

func (a *APIController) delDepletedClients(c *gin.Context) {
	a.inboundController.delDepletedClients(c)
}

func (a *APIController) createBackup(c *gin.Context) {
	a.tgBot.SendBackupToAdmins()
}

func (a *APIController) login(c *gin.Context) {
	var form LoginForm
	err := c.ShouldBind(&form)
	if err != nil {
		pureJsonMsg(c, false, I18nWeb(c, "pages.login.toasts.invalidFormData"))
		return
	}
	if form.Username == "" {
		pureJsonMsg(c, false, I18nWeb(c, "pages.login.toasts.emptyUsername"))
		return
	}
	if form.Password == "" {
		pureJsonMsg(c, false, I18nWeb(c, "pages.login.toasts.emptyPassword"))
		return
	}

	user := a.userService.CheckUser(form.Username, form.Password, form.LoginSecret)
	//timeStr := time.Now().Format("2006-01-02 15:04:05")
	//if user == nil {
	//	logger.Infof("wrong username or password: \"%s\" \"%s\"", form.Username, form.Password)
	//	a.tgbot.UserLoginNotify(form.Username, getRemoteIp(c), timeStr, 0)
	//	pureJsonMsg(c, false, I18nWeb(c, "pages.login.toasts.wrongUsernameOrPassword"))
	//	return
	//} else {
	//	logger.Infof("%s login success, Ip Address: %s\n", form.Username, getRemoteIp(c))
	//	a.tgbot.UserLoginNotify(form.Username, getRemoteIp(c), timeStr, 1)
	//}

	sessionMaxAge, err := a.settingService.GetSessionMaxAge()
	if err != nil {
		logger.Infof("Unable to get session's max age from DB")
	}

	if sessionMaxAge > 0 {
		err = session.SetMaxAge(c, sessionMaxAge*60)
		if err != nil {
			logger.Infof("Unable to set session's max age")
		}
	}

	err = session.SetLoginUser(c, user)
	logger.Info("user", user.Id, "login success")
	jsonMsg(c, I18nWeb(c, "pages.login.toasts.successLogin"), err)
}
