// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	user "ToDo/internal/domain/user"

	uuid "github.com/google/uuid"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryInterface struct {
	mock.Mock
}

type UserRepositoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepositoryInterface) EXPECT() *UserRepositoryInterface_Expecter {
	return &UserRepositoryInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: u
func (_m *UserRepositoryInterface) Create(u user.User) (user.User, error) {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(user.User) (user.User, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(user.User) user.User); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(user.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepositoryInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UserRepositoryInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - u user.User
func (_e *UserRepositoryInterface_Expecter) Create(u interface{}) *UserRepositoryInterface_Create_Call {
	return &UserRepositoryInterface_Create_Call{Call: _e.mock.On("Create", u)}
}

func (_c *UserRepositoryInterface_Create_Call) Run(run func(u user.User)) *UserRepositoryInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(user.User))
	})
	return _c
}

func (_c *UserRepositoryInterface_Create_Call) Return(_a0 user.User, _a1 error) *UserRepositoryInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepositoryInterface_Create_Call) RunAndReturn(run func(user.User) (user.User, error)) *UserRepositoryInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *UserRepositoryInterface) Delete(id uuid.UUID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepositoryInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type UserRepositoryInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *UserRepositoryInterface_Expecter) Delete(id interface{}) *UserRepositoryInterface_Delete_Call {
	return &UserRepositoryInterface_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *UserRepositoryInterface_Delete_Call) Run(run func(id uuid.UUID)) *UserRepositoryInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *UserRepositoryInterface_Delete_Call) Return(_a0 error) *UserRepositoryInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepositoryInterface_Delete_Call) RunAndReturn(run func(uuid.UUID) error) *UserRepositoryInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with no fields
func (_m *UserRepositoryInterface) GetAll() ([]user.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]user.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []user.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepositoryInterface_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type UserRepositoryInterface_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *UserRepositoryInterface_Expecter) GetAll() *UserRepositoryInterface_GetAll_Call {
	return &UserRepositoryInterface_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *UserRepositoryInterface_GetAll_Call) Run(run func()) *UserRepositoryInterface_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UserRepositoryInterface_GetAll_Call) Return(_a0 []user.User, _a1 error) *UserRepositoryInterface_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepositoryInterface_GetAll_Call) RunAndReturn(run func() ([]user.User, error)) *UserRepositoryInterface_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserRepositoryInterface) GetByEmail(email string) (*user.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*user.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepositoryInterface_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type UserRepositoryInterface_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//   - email string
func (_e *UserRepositoryInterface_Expecter) GetByEmail(email interface{}) *UserRepositoryInterface_GetByEmail_Call {
	return &UserRepositoryInterface_GetByEmail_Call{Call: _e.mock.On("GetByEmail", email)}
}

func (_c *UserRepositoryInterface_GetByEmail_Call) Run(run func(email string)) *UserRepositoryInterface_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserRepositoryInterface_GetByEmail_Call) Return(_a0 *user.User, _a1 error) *UserRepositoryInterface_GetByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepositoryInterface_GetByEmail_Call) RunAndReturn(run func(string) (*user.User, error)) *UserRepositoryInterface_GetByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: id
func (_m *UserRepositoryInterface) GetById(id uuid.UUID) (*user.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*user.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *user.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepositoryInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type UserRepositoryInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *UserRepositoryInterface_Expecter) GetById(id interface{}) *UserRepositoryInterface_GetById_Call {
	return &UserRepositoryInterface_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *UserRepositoryInterface_GetById_Call) Run(run func(id uuid.UUID)) *UserRepositoryInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *UserRepositoryInterface_GetById_Call) Return(_a0 *user.User, _a1 error) *UserRepositoryInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepositoryInterface_GetById_Call) RunAndReturn(run func(uuid.UUID) (*user.User, error)) *UserRepositoryInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: id, updated
func (_m *UserRepositoryInterface) Update(id uuid.UUID, updated user.User) (*user.User, error) {
	ret := _m.Called(id, updated)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, user.User) (*user.User, error)); ok {
		return rf(id, updated)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, user.User) *user.User); ok {
		r0 = rf(id, updated)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, user.User) error); ok {
		r1 = rf(id, updated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepositoryInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type UserRepositoryInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - id uuid.UUID
//   - updated user.User
func (_e *UserRepositoryInterface_Expecter) Update(id interface{}, updated interface{}) *UserRepositoryInterface_Update_Call {
	return &UserRepositoryInterface_Update_Call{Call: _e.mock.On("Update", id, updated)}
}

func (_c *UserRepositoryInterface_Update_Call) Run(run func(id uuid.UUID, updated user.User)) *UserRepositoryInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(user.User))
	})
	return _c
}

func (_c *UserRepositoryInterface_Update_Call) Return(_a0 *user.User, _a1 error) *UserRepositoryInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepositoryInterface_Update_Call) RunAndReturn(run func(uuid.UUID, user.User) (*user.User, error)) *UserRepositoryInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRepositoryInterface creates a new instance of UserRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryInterface {
	mock := &UserRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
