package internal

import (
	"errors"
)

var ErrNotFoundMembership = errors.New("not found membership")

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) Create(membership Membership) {
	r.data[membership.ID] = membership
}

func (r *Repository) Update(membership Membership) {
	r.data[membership.ID] = membership
}

func (r *Repository) Delete(id string) {
	delete(r.data, id)
}

func (r *Repository) checkExistId(id string) bool {
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

func (r *Repository) GetAll() map[string]Membership {
	return r.data
}
