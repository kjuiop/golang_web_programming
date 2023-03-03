package practice

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var errDivisorZero = errors.New("0으로 나눌 수 없습니다")

func sum(num1, num2 int) int {
	return num1 + num2
}

func divide(dividend, divisor int) (float32, error) {
	if divisor == 0 {
		return 0, errDivisorZero
	}
	return float32(dividend / divisor), nil
}

func generateRandomID() string {
	return uuid.New().String()
}

func TestPractice(t *testing.T) {
	t.Run("두 숫자를 더하면 합이 나온다", func(t *testing.T) {
		actual := sum(1, 2)
		expected := 3
		assert.Equal(t, expected, actual)
	})

	t.Run("두 숫자를 더하면 합이 나온다", func(t *testing.T) {
		actual := sum(1, 2)
		expected := float32(3)
		assert.EqualValues(t, expected, actual)
	})

	t.Run("두 숫자를 나눗셈 할 수 있다.", func(t *testing.T) {
		actual, err := divide(10, 2)
		assert.Nil(t, err)
		assert.EqualValues(t, 5, actual)
	})

	t.Run("0으로 나누기를 할 수 없다.", func(t *testing.T) {
		actual, err := divide(10, 0)
		assert.ErrorIs(t, err, errDivisorZero)
		assert.Zero(t, actual)
	})

	t.Run("uuid가 생성된다.", func(t *testing.T) {
		uuid := generateRandomID()
		assert.NotEmpty(t, uuid)
	})
}
