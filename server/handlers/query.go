package handlers

import (
	"fmt"
	"net/http"
	"simple-qa/core"

	"github.com/gin-gonic/gin"
)

type QueryPostPayload struct {
	Query string `json:"query"`
}

func (h *Handler) QueryPost(c *gin.Context) {
	var payload QueryPostPayload
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Malformed request body"})

		return
	}

	response, err := core.QueryDocument(payload.Query, h.VectorStore, h.LanguageModel)

	fmt.Printf("response: %s\n", err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error processing query"})

		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Query processed successfully", Data: *response})
}
