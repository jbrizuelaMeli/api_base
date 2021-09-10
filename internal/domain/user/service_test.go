package user

import (
	"context"
	"github.com/api_base/internal/domain"
	"github.com/api_base/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type fakeContainer struct {
	domain.Container
	UserRepoMock *userRepositoryMock
}

func newContainerMock() *fakeContainer {
	fc := &fakeContainer{
		UserRepoMock: &userRepositoryMock{},
	}
	fc.Container = domain.Container{
		UserRepo: fc.UserRepoMock,
	}
	return fc
}

type userRepositoryMock struct {
	mock.Mock
}

func (mr *userRepositoryMock) Get(ctx context.Context, id string) (*model.User, error) {
	args := mr.Called(ctx, id)
	res := args.Get(0).(*model.User)
	return res, args.Error(1)
}

func initTest() (context.Context, *fakeContainer, Service) {
	ctn := newContainerMock()
	srv := NewService(ctn.Container)
	return context.Background(), ctn, srv
}

func TestService_Get(t *testing.T) {
	ctx, cnt, srv := initTest()

	resDb := &model.User{Id: "1"}
	cnt.UserRepoMock.On("Get", ctx, "1").Return(resDb, nil)

	res, err := srv.Get(ctx, "1")

	assert.Nil(t, err)
	assert.Equal(t, "1", res.Id)
}
