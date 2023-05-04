package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
		h.GET("/", r.getCount)
		// h.POST("/domains", r.doTranslate)
	}
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *domainRoutes) getCount(c *gin.Context) {
	translations, err := r.t.GetCount(c.Request.Context(), "")
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, translations)
}

// type doTranslateRequest struct {
// 	Source      string `json:"source"       binding:"required"  example:"auto"`
// 	Destination string `json:"destination"  binding:"required"  example:"en"`
// 	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
// }

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /translation/do-translate [post]
// func (r *domainRoutes) doTranslate(c *gin.Context) {
// 	var request doTranslateRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}

// 	translation, err := r.t.Translate(
// 		c.Request.Context(),
// 		entity.Translation{
// 			Source:      request.Source,
// 			Destination: request.Destination,
// 			Original:    request.Original,
// 		},
// 	)
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")
// 		errorResponse(c, http.StatusInternalServerError, "translation service problems")

// 		return
// 	}

// 	c.JSON(http.StatusOK, translation)
// }
