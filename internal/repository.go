package internal

import "errors"

var ErrNotFoundMembership = errors.New("not found membership")

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) Create(membership Membership) {
	r.data[membership.UserName] = membership
}

func (r *Repository) checkDuplicateId(id string) bool {
	_, existsId := r.data[id]
	return existsId
}

func (r *Repository) GetById(id string) (Membership, error) {
	for _, membership := range r.data {
		if membership.ID == id {
			return membership, nil
		}
	}
	return Membership{}, ErrNotFoundMembership
}
