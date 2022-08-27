package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)

		req2 := CreateRequest{"jenny", "naver"}
		res2, err2 := app.Create(req2)
		assert.ErrorIs(t, err2, errAlreadyExistUsername)
		assert.Empty(t, res2.ID)

	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}
		res, err := app.Create(req)
		assert.ErrorIs(t, err, errEmptyUsername)
		assert.Empty(t, res.ID)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", ""}
		res, err := app.Create(req)
		assert.ErrorIs(t, err, errEmptyMemberShip)
		assert.Empty(t, res.ID)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "kakao"}
		res, err := app.Create(req)
		assert.ErrorIs(t, err, errNotApplyMemberShip)
		assert.Empty(t, res.ID)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("internal 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)

		updateReq := UpdateRequest{"jenny", "jake", "toss"}
		updateRes, updateErr := app.Update(updateReq)
		assert.Nil(t, updateErr)
		assert.NotEmpty(t, updateRes.ID)
		assert.Equal(t, updateReq.UserName, updateRes.UserName)
		assert.Equal(t, updateReq.MembershipType, updateRes.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)

		req2 := CreateRequest{"jake", "naver"}
		res2, err2 := app.Create(req2)
		assert.Nil(t, err2)
		assert.NotEmpty(t, res2.ID)

		updateReq := UpdateRequest{"jenny", "jake", "naver"}
		updateRes, updateErr := app.Update(updateReq)
		assert.ErrorIs(t, updateErr, errAlreadyExistUsername)
		assert.Empty(t, updateRes.ID)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)

		updateReq := UpdateRequest{"", "jenny", "toss"}
		updateRes, updateErr := app.Update(updateReq)
		assert.ErrorIs(t, updateErr, errEmptyId)
		assert.Empty(t, updateRes.ID)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)

		updateReq := UpdateRequest{"jenny", "", "toss"}
		updateRes, updateErr := app.Update(updateReq)
		assert.ErrorIs(t, updateErr, errEmptyUsername)
		assert.Empty(t, updateRes.ID)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)

		updateReq := UpdateRequest{"jenny", "jenny", ""}
		updateRes, updateErr := app.Update(updateReq)
		assert.ErrorIs(t, updateErr, errEmptyMemberShip)
		assert.Empty(t, updateRes.ID)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)

		updateReq := UpdateRequest{"jenny", "jenny", "kakao"}
		updateRes, updateErr := app.Update(updateReq)
		assert.ErrorIs(t, updateErr, errNotApplyMemberShip)
		assert.Empty(t, updateRes.ID)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {

	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {

	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {

	})
}
