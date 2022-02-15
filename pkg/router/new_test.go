package router_test

import (
	"testing"

	"github.com/ownfitness/template-go/pkg/gcp"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"

	"github.com/ownfitness/template-go/pkg/router"
)

func TestNew(t *testing.T) {
	c, _ := gcp.FirebaseClient("1234")
	r := router.New(false, c)

	routes := r.Routes()

	assert.Equal(t, "/health", routes[0].Path)
	assert.Equal(t, "GET", routes[0].Method)
	assert.Equal(t, "/users", routes[1].Path)
	assert.Equal(t, "POST", routes[1].Method)
	assert.IsType(t, &gin.Engine{}, r)
}
