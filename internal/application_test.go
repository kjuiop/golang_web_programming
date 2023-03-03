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
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.ErrorIs(t, err, errAlreadyExistUsername)
		assert.Empty(t, res.ID)

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
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		req := UpdateRequest{"jenny", "jake", "toss"}
		res, err := app.Update(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.UserName, res.UserName)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
			"jake":  {"jake", "jake", "naver"},
		}))

		req := UpdateRequest{"jenny", "jake", "naver"}
		res, err := app.Update(req)
		assert.ErrorIs(t, err, errAlreadyExistUsername)
		assert.Empty(t, res.ID)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		req := UpdateRequest{"", "jenny", "toss"}
		res, err := app.Update(req)
		assert.ErrorIs(t, err, errEmptyId)
		assert.Empty(t, res.ID)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		req := UpdateRequest{"jenny", "", "toss"}
		res, err := app.Update(req)
		assert.ErrorIs(t, err, errEmptyUsername)
		assert.Empty(t, res.ID)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		req := UpdateRequest{"jenny", "jenny", ""}
		res, err := app.Update(req)
		assert.ErrorIs(t, err, errEmptyMemberShip)
		assert.Empty(t, res.ID)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		req := UpdateRequest{"jenny", "jenny", "kakao"}
		res, err := app.Update(req)
		assert.ErrorIs(t, err, errNotApplyMemberShip)
		assert.Empty(t, res.ID)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		err := app.Delete("jenny")
		assert.Nil(t, err)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		err := app.Delete("")
		assert.ErrorIs(t, err, errEmptyId)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		err := app.Delete("jake")
		assert.ErrorIs(t, err, errNotFoundId)
	})
}

func TestSelect(t *testing.T) {
	t.Run("전체 멤버십 정보를 조회한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, err := app.SelectAll()
		assert.Nil(t, err)
	})

	t.Run("존재하는 특정 사용자이름을 조회한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		req := SelectRequest{"jenny"}
		res, err := app.SelectById(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.data)
	})

	t.Run("존재하지 않는 사용자이름을 조회한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := SelectRequest{"jenny"}
		res, err := app.SelectById(req)
		assert.ErrorIs(t, err, errNotFoundException)
		assert.Empty(t, res.data)
	})

	t.Run("특정 아이디를 조회할 때 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{
			"jenny": {"jenny", "jenny", "naver"},
		}))

		req := SelectRequest{""}
		res, err := app.SelectById(req)
		assert.ErrorIs(t, err, errEmptyId)
		assert.Empty(t, res.data)
	})

}
