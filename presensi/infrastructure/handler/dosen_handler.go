package handler

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
	dosenQuery query.DosenQuery
}

func NewDosenHandler(dosenQuery query.DosenQuery) *DosenHandler {
	return &DosenHandler{dosenQuery: dosenQuery}
}

func (d *DosenHandler) GetDosen(c *gin.Context) {

	userId := c.Query("user_id")

	if len(userId) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	dosen, err := d.dosenQuery.Execute(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if dosen == nil {
		c.JSON(http.StatusNotFound, nil)
	}

	response := DosenResponse{
		DosenId: dosen.DosenId.String(),
		Nama:    dosen.Nama,
	}

	c.JSON(http.StatusOK, response)
}
