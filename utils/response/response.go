package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(r *http.Request, w http.ResponseWriter, resp interface{}) {
	httpx.WriteJson(w, http.StatusOK, &Body{
		Code:    "SUCCESS",
		Message: "success",
		Data:    resp,
	})
}

func ParamErrorResponse(r *http.Request, w http.ResponseWriter, err error)  {
	httpx.WriteJson(w, http.StatusUnprocessableEntity, &Body{
		Code: "REQUEST_PARAM_ERROR",
		Message: err.Error(),
	})
}