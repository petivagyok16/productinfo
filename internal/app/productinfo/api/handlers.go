package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *RouteHandler) getServices(c *gin.Context) {

	pathParams := getPathParams(c)
	log.Debug("Path params: [%s]", pathParams)

	services := r.prod.GetServices(pathParams.Provider, pathParams.Region)
	c.JSON(http.StatusOK, services)
}

func (r *RouteHandler) getService(c *gin.Context) {

	pathParams := getPathParams(c)
	log.Debug("Path params: [%s]", pathParams)

	log.Error("not yet implemented")
	c.JSON(http.StatusInternalServerError, "not yet implemented")
}

func (r *RouteHandler) getServiceImages(c *gin.Context) {

	pathParams := getPathParams(c)
	log.Debug("Path params: [%s]", pathParams)

	log.Error("not yet implemented")
	c.JSON(http.StatusInternalServerError, "not yet implemented")

}

func (r *RouteHandler) getServiceProducts(c *gin.Context) {
	pathParams := getPathParams(c)
	log.Debug("Path params: [%s]", pathParams)

	log.Error("not yet implemented")
	c.JSON(http.StatusInternalServerError, "not yet implemented")

}

func (r *RouteHandler) getServiceAttributes(c *gin.Context) {
	pathParams := getPathParams(c)
	log.Debug("Path params: [%s]", pathParams)

	log.Error("not yet implemented")
	c.JSON(http.StatusInternalServerError, "not yet implemented")

}

// getPathParam handles path params retrieval in case tha path param is not set a warning is logged
func getPathParam(pathParam string, c *gin.Context) string {
	param := c.Param(pathParam)
	if param == "" {
		log.Warnf("path param %s is not set", pathParam)
	}
	return param
}

// getPathParams parses the path information from the gin Contexts
func getPathParams(c *gin.Context) PathParams {
	provider := getPathParam(providerParam, c)
	region := getPathParam(regionParam, c)
	service := getPathParam(serviceParam, c)
	return newPathParams(provider, region, service)
}
