package controllers

import (
	"net/http"
	"vote-backend/storage"

	"github.com/gin-gonic/gin"
)

func GetCandidates(c *gin.Context) {
	c.JSON(http.StatusOK, storage.GetCandidates())
}

func VoteCandidate(c *gin.Context) {
	name := c.Param("name")
	if candidate, ok := storage.Vote(name); ok {
		c.JSON(http.StatusOK, candidate)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "candidate not found"})
	}
}

func ResetVotes(c *gin.Context) {
	storage.ResetVotes()
	c.JSON(http.StatusOK, gin.H{"message": "reset ok"})
}
