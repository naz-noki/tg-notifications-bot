package botController

type botController struct {
	service botService
}

func New(service botService) *botController {
	bc := new(botController)
	bc.service = service

	return bc
}
