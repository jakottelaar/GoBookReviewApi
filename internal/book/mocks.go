package book

import "github.com/stretchr/testify/mock"

type MockBookRepository struct {
	mock.Mock
}

type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) Create(req *CreateBookRequest) (*Book, error) {
	args := m.Called(req)
	return args.Get(0).(*Book), args.Error(1)
}

func (m *MockBookService) GetBookById(id string) (*Book, error) {
	args := m.Called(id)
	return args.Get(0).(*Book), args.Error(1)
}

func (m *MockBookService) Update(id string, req *UpdateBookRequest) (*Book, error) {
	args := m.Called(id, req)
	return args.Get(0).(*Book), args.Error(1)
}

func (m *MockBookService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookRepository) Save(book *Book) (*Book, error) {
	args := m.Called(book)
	return args.Get(0).(*Book), args.Error(1)
}

func (m *MockBookRepository) FindById(id string) (*Book, error) {
	args := m.Called(id)
	return args.Get(0).(*Book), args.Error(1)
}

func (m *MockBookRepository) Update(book *Book) (*Book, error) {
	args := m.Called(book)
	return args.Get(0).(*Book), args.Error(1)
}

func (m *MockBookRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
