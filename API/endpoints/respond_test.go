package endpoints

import (
	"eisandbar/anbox/typing"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Data struct {
}

type RespondTestSuite struct {
	suite.Suite
}

func (suite *RespondTestSuite) TestRespBodyType() {
	w := httptest.NewRecorder()
	respond(w, "", Data{})
	resp := w.Result()
	respBody, _ := io.ReadAll(resp.Body)
	suite.IsType([]byte{}, respBody)
}

func (suite *RespondTestSuite) TestResponseGood() {
	w := httptest.NewRecorder()
	respond(w, "", Data{})

	resp := w.Result()
	respBody, _ := io.ReadAll(resp.Body)

	var response typing.Response
	json.Unmarshal(respBody, &response)

	suite.Equal("Success", response.Status)
	suite.Equal(200, response.StatusCode)
}

func (suite *RespondTestSuite) TestResponseBad() {
	w := httptest.NewRecorder()
	respond(w, "Error", Data{})

	resp := w.Result()
	respBody, _ := io.ReadAll(resp.Body)

	var response typing.Response
	json.Unmarshal(respBody, &response)

	suite.Equal("Error", response.ErrorMsg)
	suite.Equal(400, response.ErrorCode)

}

func TestRespondTestSuite(t *testing.T) {
	suite.Run(t, new(RespondTestSuite))
}
