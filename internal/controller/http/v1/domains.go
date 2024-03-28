package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rsdmike/rps/internal/entity"
	"github.com/rsdmike/rps/internal/usecase"
	"github.com/rsdmike/rps/pkg/logger"
)

type domainRoutes struct {
	t usecase.Domain
	l logger.Interface
}

func newDomainRoutes(handler *gin.RouterGroup, t usecase.Domain, l logger.Interface) {
	r := &domainRoutes{t, l}

	h := handler.Group("/domains")
	{
		h.GET("/", r.get)
		h.GET("/:domainName", r.getByName)
		h.POST("/", r.insert)
		h.PATCH("/", r.update)
		h.DELETE("/:domainName", r.delete)
	}
}

type OData struct {
	Top   int  `form:"$top"`
	Skip  int  `form:"$skip"`
	Count bool `form:"$count"`
}
type CountResponse struct {
	Count int             `json:"totalAccount"`
	Data  []entity.Domain `json:"data"`
}

// @Summary     Show Domains
// @Description Show all domains
// @ID          history
// @Tags  	    domains
// @Accept      json
// @Produce     json
// @Success     200 {object} domainresponse
// @Failure     500 {object} response
// @Router      /api/v1/admin/domains [get]
func (r *domainRoutes) get(c *gin.Context) {
	var odata OData
	if err := c.ShouldBindQuery(&odata); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	domains, err := r.t.Get(c.Request.Context(), odata.Top, odata.Skip, "")
	if err != nil {
		r.l.Error(err, "http - v1 - getCount")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	if odata.Count {
		count, err := r.t.GetCount(c.Request.Context(), "")
		if err != nil {
			r.l.Error(err, "http - v1 - getCount")
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}

		countResponse := CountResponse{
			Count: count,
			Data:  domains,
		}

		c.JSON(http.StatusOK, countResponse)
	} else {
		c.JSON(http.StatusOK, domains)
	}
}

func (r *domainRoutes) getByName(c *gin.Context) {
	var domain entity.Domain
	if err := c.ShouldBindUri(&domain); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	translations, err := r.t.GetByName(c.Request.Context(), domain.ProfileName, "")
	if err != nil {
		r.l.Error(err, "http - v1 - getByName")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, translations)
}

func (r *domainRoutes) insert(c *gin.Context) {
	var domain entity.Domain
	if err := c.ShouldBindJSON(&domain); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	translations, err := r.t.Insert(c.Request.Context(), &domain)
	if err != nil {
		r.l.Error(err, "http - v1 - insert")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, translations)
}

func (r *domainRoutes) update(c *gin.Context) {
	var domain entity.Domain
	if err := c.ShouldBindJSON(&domain); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	translations, err := r.t.Update(c.Request.Context(), &domain)
	if err != nil {
		r.l.Error(err, "http - v1 - update")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, translations)
}

func (r *domainRoutes) delete(c *gin.Context) {
	var domain entity.Domain
	if err := c.ShouldBindUri(&domain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	translations, err := r.t.Delete(c.Request.Context(), domain.ProfileName, "")
	if err != nil {
		r.l.Error(err, "http - v1 - delete")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, translations)
}
