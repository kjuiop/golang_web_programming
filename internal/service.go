package internal

import (
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(request CreateRequest) (CreateResponse, error) {

	exists := s.repository.checkExistId(request.UserName)
	if exists {
		return CreateResponse{}, errAlreadyExistUsername
	}

	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	s.repository.Create(membership)
	return CreateResponse{
		ID:             membership.ID,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) Update(request UpdateRequest) (UpdateResponse, error) {

	exists := s.repository.checkExistId(request.ID)
	if !exists {
		return UpdateResponse{}, errAlreadyExistUsername
	}

	membership := Membership{request.ID, request.UserName, request.MembershipType}
	s.repository.Update(membership)
	return UpdateResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) Delete(id string) error {

	exists := s.repository.checkExistId(id)
	if !exists {
		return errAlreadyExistUsername
	}

	s.repository.Delete(id)

	return nil
}

func (s *Service) GetByID(id string) (GetResponse, error) {
	membership, err := s.repository.GetById(id)
	if err != nil {
		return GetResponse{}, nil
	}
	return GetResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) GetAll() (map[string]Membership, error) {
	memberships := s.repository.GetAll()
	return memberships, nil
}

func (s *Service) checkEmptyValue(str string) bool {
	return len(str) == 0
}

func (s *Service) notMemberShipType(str string) bool {
	switch str {
	case "naver", "toss", "payco":
		return false
	default:
		return true
	}
}
