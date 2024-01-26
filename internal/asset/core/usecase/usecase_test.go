package usecase

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/lavinas/keel/internal/asset/adapter/tools"
	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/dto"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

type CreateTestCase struct {
	Name           string
	Desc           string
	Instr          string
	DtoIn          port.CreateDtoIn
	DtoOut         port.CreateDtoOut
	ExpectedDtoOut port.CreateDtoOut
	ExpectedError  *kerror.KError
	Messages       []string
}

var CreateTestCases = []CreateTestCase{
	{
		Name:  "create tax - ok",
		Desc:  "should create a new tax",
		Instr: "create tax ok",
		DtoIn: &dto.TaxCreateIn{ID: "tax1", Name: "tax1", Period: "Y",
			TaxItens: []dto.TaxItemCreate{{ID: "tax1", Until: 1, Value: 0.1}}},
		DtoOut: &dto.TaxCreateOut{},
		ExpectedDtoOut: &dto.TaxCreateOut{ID: "tax1", Name: "tax1", Period: "Y",
			TaxItens: []dto.TaxItemCreate{{ID: "tax1", Until: 1, Value: 0.1}}},
		ExpectedError: nil,
		Messages:      []string{},
	},
	{
		Name:           "create class - ok",
		Desc:           "should create a new class",
		Instr:          "create class ok",
		DtoIn:          &dto.ClassCreateIn{ID: "class1", Name: "class1", TaxID: "tax1"},
		DtoOut:         &dto.ClassCreateOut{},
		ExpectedDtoOut: &dto.ClassCreateOut{ID: "class1", Name: "class1", TaxID: "tax1"},
		ExpectedError:  nil,
		Messages:       []string{},
	},
}

func TestCreate(t *testing.T) {
	repo := &MockRepository{}
	logger := &MockLogger{}
	config := tools.NewConfig()
	usecase := NewUseCase(repo, logger, config)

	for _, tc := range CreateTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo.Instr = tc.Instr
			err := usecase.Create(tc.DtoIn, tc.DtoOut)
			if !reflect.DeepEqual(tc.ExpectedError, err) {
				t.Errorf("%s - Expected error: '%v', got: '%v'", tc.Name, tc.ExpectedError, err)
			}
			if !reflect.DeepEqual(tc.DtoOut, tc.ExpectedDtoOut) {
				t.Errorf("%s - Expected dto out: '%v', got: '%v'", tc.Name, tc.ExpectedDtoOut, tc.DtoOut)
			}
			if !reflect.DeepEqual(tc.Messages, logger.Message) && (len(tc.Messages) > 0 || len(logger.Message) > 0) {
				t.Errorf("%s - Expected messages: '%v', got: '%v'", tc.Name, tc.Messages, logger.Message)
			}
		})
	}
}

// MockConfig is a mock implementation of the Repository interface
type MockRepository struct {
	Instr string
}

func (r *MockRepository) Exists(obj interface{}, id string) (bool, error) {
	if r.Instr == "create class ok" && reflect.TypeOf(obj) == reflect.TypeOf(&domain.Tax{}) {
		return true, nil
	}
	return false, nil
}
func (r *MockRepository) GetByID(obj interface{}, id string) (bool, error) {
	println(3, obj)
	if r.Instr == "create class ok" && reflect.TypeOf(obj) == reflect.TypeOf(&domain.Tax{}) {
		obj = &domain.Tax{ID: id, Name: id, Period: "Y", TaxItems: []*domain.TaxItem{{ID: id, Until: 1, Value: 0.1}}}
		println(1, obj)
		return true, nil
	}
	return false, nil
}
func (r *MockRepository) Add(obj interface{}) error {
	return nil
}
func (r *MockRepository) Close() {
}

// MockLogger is a mock implementation of the Repository interface
type MockLogger struct {
	Message []string
}

func (r *MockLogger) GetFile() *os.File {
	return nil
}
func (r *MockLogger) GetName() string {
	return ""
}
func (r *MockLogger) Info(message string) {
	r.Message = append(r.Message, message)
}
func (r *MockLogger) Infof(format string, a ...any) {
	r.Message = append(r.Message, fmt.Sprintf(format, a...))
}
func (r *MockLogger) Error(err error) {
	r.Message = append(r.Message, err.Error())
}
func (r *MockLogger) Fatal(err error) {
	r.Message = append(r.Message, err.Error())
}
func (r *MockLogger) Errorf(format string, a ...any) {
	r.Message = append(r.Message, fmt.Sprintf(format, a...))
}
func (r *MockLogger) Close() {
}
