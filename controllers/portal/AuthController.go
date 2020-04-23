package portal

type AuthController struct {
	BaseController
}

func (c *AuthController) login() {
	//c.BaseController.BaseSuccess
	data := new(Data)
	data.Error("cuo_wu")
	data.Success(make(map[string]interface{}))
}