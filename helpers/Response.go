package helpers

import "github.com/labstack/echo"

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Err struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func FetchData(c echo.Context, data interface{}) error {
	var res Res
	res.Code = 200
	res.Message = "success retrieve data"
	res.Data = data
	return c.JSON(res.Code, res)
}

func UserAuthenticated(c echo.Context, data interface{}) error {
	var res Res
	res.Code = 200
	res.Message = "success authenticating user"
	res.Data = data
	return c.JSON(res.Code, res)
}
func UnAuthorizedUser(c echo.Context) error {
	var err Err
	err.Code = 401
	err.Message = "user not found"
	err.Errors = nil
	return c.JSON(err.Code, err)
}
func CreateData(c echo.Context, data interface{}) error {
	var res Res
	res.Code = 200
	res.Message = "success store data"
	res.Data = data
	return c.JSON(res.Code, res)
}
func UpdateData(c echo.Context, data interface{}) error {
	var res Res
	res.Code = 202
	res.Message = "success update data"
	res.Data = data
	return c.JSON(res.Code, res)
}
func DeleteData(c echo.Context) error {
	var res Res
	res.Code = 200
	res.Message = "success delete data"
	res.Data = nil
	return c.JSON(res.Code, res)
}
func DataNotFound(c echo.Context) error {
	var err Err
	err.Code = 200
	err.Message = "data not found"
	err.Errors = nil
	return c.JSON(err.Code, err)
}
func ValidationErrors(c echo.Context, Error interface{}) error {
	var err Err
	err.Code = 406
	err.Message = "validation error"
	err.Errors = Error
	return c.JSON(err.Code, err)
}
func ClientErrors(c echo.Context, Error interface{}) error {
	var err Err
	err.Code = 400
	err.Message = "bad request error"
	err.Errors = Error
	return c.JSON(err.Code, err)
}
func ServerErrors(c echo.Context, Error interface{}) error {
	var err Err
	err.Code = 500
	err.Message = "server error"
	err.Errors = Error
	return c.JSON(err.Code, err)
}
