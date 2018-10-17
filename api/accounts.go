package api

import (
	"net/http"
	"strconv"
	"errors"

	"columbus/utils"
	"columbus/models"
	. "columbus/database"
	. "columbus/log"

	"github.com/gin-gonic/gin"
)

// TODO 参考swag生成 修改校验逻辑 类似swagger g.json
type AccountArg struct {
	Name string `form:"name" json:"name"`
}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @Tags accounts
// @Security Bearer
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} models.Account
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	Info.Println(id, "query id....")
	aid, err := strconv.Atoi(id)

	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := new(models.Account)
	account.Uid = aid
	Info.Println(account)
	has, err := DB.Get(account)
	if err != nil{
		Error.Println(err)
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	if has == false {
		utils.NewError(ctx, http.StatusNotFound, errors.New("account not found"))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Tags accounts
// @Security Bearer
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q" Format(email)
// @Success 200 {array} models.Account
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /accounts [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	//q := ctx.Request.URL.Query().Get("q")
	accounts := make([]models.Account, 0)
	err := DB.Asc("uid").Find(&accounts)
	if err != nil{
		Error.Println(err)
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

// AddAccount godoc
// @Summary Add a account
// @Description add by json account
// @Tags accounts
// @Security Bearer
// @Accept  json
// @Produce  json
// @Param account body AccountArg true "Add account"
// @Success 200 {object} models.Account
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /accounts [post]
func (c *Controller) AddAccount(ctx *gin.Context) {
	var accountArg AccountArg
	ctx.ShouldBindJSON(&accountArg)
	account := new(models.Account)
	account.Name = accountArg.Name
	res, err := DB.Insert(account)
	if err != nil {
		Error.Println(err)
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// UpdateAccount godoc
// @Summary Update a account
// @Description Update by json account
// @Tags accounts
// @Accept  json
// @Security Bearer
// @Produce  json
// @Param  id path int true "Account ID"
// @Param  account body AccountArg true "Update account"
// @Success 200 {object} models.Account
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /accounts/{id} [patch]
func (c *Controller) UpdateAccount(ctx *gin.Context) {
	var accountArg AccountArg
	ctx.ShouldBindJSON(&accountArg)
	account := new(models.Account)
	account.Name = accountArg.Name
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		Error.Println(err)
		return
	}
	res, err := DB.Id(aid).Update(account)
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteAccount godoc
// @Summary 修改用户
// @Description Delete by account ID
// @Tags accounts
// @Accept  json
// @Security Bearer
// @Produce  json
// @Param  id path int true "Account ID" Format(int64)
// @Success 204 {object} models.Account
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /accounts/{id} [delete]
func (c *Controller) DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := new(models.Account)
	res, err := DB.Id(aid).Delete(account)
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, res)
}

// AccountToken godoc
// @Summary 微信授权获取token接口
// @Description get token
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body AccountArg true "Add account"
// @Success 200 {object} string
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /accounts/token [post]
func (c *Controller) AccountToken(ctx *gin.Context) {
	// 根据code获取用户信息
	code := ctx.Request.URL.Query().Get("code")
    user_info := utils.GetUserInfoByCode(code)
	account := new(models.Account)
	has, _ := DB.Where("open_id = ?", user_info.OpenID).Get(account)
	if has == false {
		account.Name = user_info.Nickname
		account.Avatar = user_info.HeadImgURL
		account.City = user_info.City
		account.OpenId = user_info.OpenID
		account.UnionId = user_info.Unionid
		// 建立用户信息
		_, err := DB.Insert(account)
		if err != nil {
			Error.Println(err)
			utils.NewError(ctx, http.StatusBadRequest, err)
			return
		}
	}
	// 返回token
	token := utils.GenerateToken(*account)
	ctx.JSON(http.StatusOK, token)
}