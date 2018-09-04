package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"errors"

	"columbus/httputil"
	"columbus/model"
	"columbus/db"

	"github.com/gin-gonic/gin"
)

// TODO 参考swag生成 修改校验逻辑 类似swagger g.json
type AccountArg struct {
	Name string `form:"name" json:"name"`
}

func (c *Controller) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := new(model.Account)
	account.Id = aid
	engine, err:= db.SqliteEngine()
	engine.ShowSQL(true)
	fmt.Println(account)
	if err != nil {
		fmt.Println(err)
		return
	}
	has, err := engine.Get(account)
	if err != nil{
		fmt.Println(err)
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	if has == false {
		httputil.NewError(ctx, http.StatusNotFound, errors.New("account not found"))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (c *Controller) ListAccounts(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")
	fmt.Sprintf(q)
	accounts := make([]model.Account, 0)
	engine, err:= db.SqliteEngine()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = engine.Where("name = ?", q).Find(&accounts)
	if err != nil{
		fmt.Println(err)
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

func (c *Controller) AddAccount(ctx *gin.Context) {
	var accountArg AccountArg
	ctx.ShouldBindJSON(&accountArg)
	account := new(model.Account)
	account.Name = accountArg.Name
	engine, err:= db.SqliteEngine()
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := engine.Insert(account)
	if err != nil {
		fmt.Println(err)
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) UpdateAccount(ctx *gin.Context) {
	var accountArg AccountArg
	ctx.ShouldBindJSON(&accountArg)
	account := new(model.Account)
	account.Name = accountArg.Name
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	engine, err:= db.SqliteEngine()
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := engine.Id(aid).Update(account)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := new(model.Account)
	engine, err:= db.SqliteEngine()
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := engine.Id(aid).Delete(account)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, res)
}
