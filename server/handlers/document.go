package handlers

import (
	"net/http"
	"simple-qa/core"
	"simple-qa/util"

	"github.com/gin-gonic/gin"
)

type Document struct {
	Text string `json:"text"`
}

type DocumentPostPayload struct {
	Documents []Document `json:"documents"`
}

func (h *Handler) DocumentPost(c *gin.Context) {
	var payload DocumentPostPayload
	err := c.BindJSON(&payload)

	if err != nil || len(payload.Documents) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Malformed request body"})

		return
	}

	docs := util.Map(payload.Documents, func(document Document) string {
		return document.Text
	})

	_, err = core.ProcessDocuments(docs, h.VectorStore)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error processing documents"})

		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Documents added successfully"})
}
