package order

import (
	"github.com/gin-gonic/gin"
	"shequn1/internal/entities"
	app2 "shequn1/internal/foundation/app"
	"shequn1/internal/foundation/database/orm"
	"shequn1/internal/foundation/validator"
)

// ListToVisitor get id
type ListToVisitor struct {
	ID string `form:"id" binding:"required,max=32"`
}

// List 示例接口
// @Summary 订单
// @Tags 订单列表
// @Produce  json
// @Param    id        query    string     true      "订单id"
// @Success  0         {object}  entities.Order
// @failure  404
// @Router  /order [get]
func List(ctx *gin.Context) {
	var listToVisitor ListToVisitor
	if validate := validator.Bind(ctx, &listToVisitor); !validate.IsValid() {
		app2.NewResponse(app2.Fail, validate.ErrorsInfo).End(ctx)
		return
	}

	var order entities.Order
	result := orm.Slave().Where("id = ?", listToVisitor.ID).Find(&order)
	if result.RowsAffected > 0 {
		app2.NewResponse(app2.Success, order).End(ctx)
		return
	}
	app2.NewResponse(app2.NotFound, nil).End(ctx)
	return
}
