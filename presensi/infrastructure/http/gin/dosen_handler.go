package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"its.id/akademik/presensi/application/query"
)

type DosenResponse struct {
	DosenId string `json:"dosen_id"`
	Nama    string `json:"nama"`
}

type DosenHandler struct {
	dosenQuery query.DosenQueryHandler
}

func NewDosenHandler(dosenQuery query.DosenQueryHandler) *DosenHandler {
	return &DosenHandler{dosenQuery: dosenQuery}
}

func (d *DosenHandler) GetDosen(c *gin.Context) {

	userId := c.Query("user_id")

	if len(userId) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	dosen, err := d.dosenQuery.GetByUserId(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if dosen == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	response := DosenResponse{
		DosenId: dosen.DosenId.String(),
		Nama:    dosen.Nama,
	}

	c.JSON(http.StatusOK, response)
}
