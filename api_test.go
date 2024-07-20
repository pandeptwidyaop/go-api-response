package api_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/pandeptwidyaop/go-api-response"
	"github.com/stretchr/testify/assert"
)

func Test_CreateResponse(t *testing.T) {

	response := api.New().Status(http.StatusOK)

	j, err := json.Marshal(response)

	assert.Nil(t, err)
	assert.Equal(t, `{"status":200,"data":null,"message":null,"meta":null}`, string(j))
}

func Test_FillMessages(t *testing.T) {
	var msg []interface{}

	msg = append(msg, map[string]interface{}{
		"errors": map[string]interface{}{
			"email": []string{"Not Valid", "Is Required"},
		},
	})

	r := api.New().Message(msg)

	j, err := json.Marshal(r)

	assert.Nil(t, err)

	assert.Equal(t, `{"status":200,"data":null,"message":[{"errors":{"email":["Not Valid","Is Required"]}}],"meta":null}`, string(j))
}
