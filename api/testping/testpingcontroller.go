package testping

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"iCloudDisk/pkg/log"
)

func pingFunc(c *gin.Context) {
	reqIp := c.ClientIP()
	log.Info("ping success, resource ip: %s", reqIp)
	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("this is a test ping request, ip: %s", reqIp)})
}
