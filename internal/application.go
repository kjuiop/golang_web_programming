package internal

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	if app.checkEmptyValue(request.UserName) {
		return CreateResponse{}, errEmptyUsername
	}

	if app.checkEmptyValue(request.MembershipType) {
		return CreateResponse{}, errEmptyMemberShip
	}

	if app.notMemberShipType(request.MembershipType) {
		return CreateResponse{}, errNotApplyMemberShip
	}

	data := app.repository.data

	_, exists := data[request.UserName]
	if exists {
		return CreateResponse{}, errAlreadyExistUsername
	}

	data[request.UserName] = Membership{request.UserName, request.UserName, request.MembershipType}

	return CreateResponse{request.UserName, request.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	return UpdateResponse{}, nil
}

func (app *Application) Delete(id string) error {
	return nil
}

func (app *Application) Select(request SelectRequest) (SelectResponse, error) {
	return SelectResponse{}, nil
}

func (app *Application) checkEmptyValue(s string) bool {
	if len(s) > 0 {
		return false
	}
	return true
}

func (app *Application) notMemberShipType(s string) bool {
	switch s {
	case "naver", "toss", "payco":
		return false
	default:
		return true
	}
}

func (app *Application) isDuplicateUsername(s string) bool {
	return true
}
