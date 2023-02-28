package service

import (
	"context"
	"testing"

	"github.com/manuonda/go-projects/inventario/internal/entitys"
	"github.com/manuonda/go-projects/inventario/internal/repository"
	"github.com/stretchr/testify/mock"
)

func testRegisteredUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisteredUser_Success",
			Email:         "test@test.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: nil,
		}, {
			Name:          "RegisteredUser_UserAlreadyExists",
			Email:         "test@test.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		}, {},
	}

	ctx := context.Background()
	repo := &repository.MockRepository{}
	repo.On("GetUserByEmail", ctx, mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", ctx, mock.Anything, "test@exists.com").Return(&entitys.User{Email: "test@exists.com"}, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			// configuration repo
			repo.Mock.Test(t)

			// s := New(repo)
			// err := s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)
			// if err != nil {
			// 	t.Error("Expected error %v, got %v", tc.ExpectedError, err)
			// }

		})
	}
}
