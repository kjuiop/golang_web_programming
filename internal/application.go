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

	exists := app.repository.checkExistId(request.UserName)
	if exists {
		return CreateResponse{}, errAlreadyExistUsername
	}

	data := app.repository.data

	data[request.UserName] = Membership{request.UserName, request.UserName, request.MembershipType}

	return CreateResponse{request.UserName, request.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {

	if app.checkEmptyValue(request.ID) {
		return UpdateResponse{}, errEmptyId
	}

	if app.checkEmptyValue(request.UserName) {
		return UpdateResponse{}, errEmptyUsername
	}

	if app.checkEmptyValue(request.MembershipType) {
		return UpdateResponse{}, errEmptyMemberShip
	}

	if app.notMemberShipType(request.MembershipType) {
		return UpdateResponse{}, errNotApplyMemberShip
	}

	existsId := app.repository.checkExistId(request.ID)
	if !existsId {
		return UpdateResponse{}, errNotFoundId
	}

	data := app.repository.data

	_, existsUsername := data[request.UserName]
	if existsUsername {
		return UpdateResponse{}, errAlreadyExistUsername
	}

	member := data[request.ID]
	member.UserName = request.UserName
	member.MembershipType = request.MembershipType

	data[request.ID] = member

	return UpdateResponse{member.ID, member.UserName, member.MembershipType}, nil
}

func (app *Application) Delete(id string) error {

	if app.checkEmptyValue(id) {
		return errEmptyId
	}

	existsId := app.repository.checkExistId(id)
	if !existsId {
		return errNotFoundId
	}

	data := app.repository.data

	delete(data, id)

	return nil
}

func (app *Application) SelectAll() (SelectAllResponse, error) {
	data := app.repository.data
	return SelectAllResponse{data}, nil
}

func (app *Application) SelectById(request SelectRequest) (SelectOneResponse, error) {

	if app.checkEmptyValue(request.ID) {
		return SelectOneResponse{}, errEmptyId
	}

	data := app.repository.data

	val, exists := data[request.ID]
	if !exists {
		return SelectOneResponse{}, errNotFoundException
	}

	return SelectOneResponse{val}, nil
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
