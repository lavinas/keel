package usecase

import (
	"fmt"
	"os"
	"testing"
	"time"
	"reflect"

	"github.com/lavinas/keel/internal/asset/adapter/tools"
	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/dto"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

type CreateTestCase struct {
	ItsOn          bool
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
		ItsOn:          false,
		Name:           "create tax - ok",
		Desc:           "should create a new tax",
		Instr:          "create tax ok",
		DtoIn:          &dto.TaxCreateIn{ID: "tax1", Name: "tax1", Period: "Y", TaxItens: []dto.TaxItemCreate{{ID: "tax1", Until: 1, Value: 0.1}}},
		DtoOut:         &dto.TaxCreateOut{},
		ExpectedDtoOut: &dto.TaxCreateOut{ID: "tax1", Name: "tax1", Period: "Y", TaxItens: []dto.TaxItemCreate{{ID: "tax1", Until: 1, Value: 0.1}}},
		ExpectedError:  nil,
		Messages:       []string{},
	},
	{
		ItsOn:          true,
		Name:           "create tax - no id",
		Desc:           "should create a new tax",
		Instr:          "create tax ok",
		DtoIn:          &dto.TaxCreateIn{ID: "", Name: "tax1", Period: "Y", TaxItens: []dto.TaxItemCreate{{ID: "tax1", Until: 1, Value: 0.1}}},
		DtoOut:         &dto.TaxCreateOut{},
		ExpectedDtoOut: &dto.TaxCreateOut{},
		ExpectedError:  kerror.NewKError(kerror.BadRequest, dto.ErrorTaxIDRequired),
		Messages:       []string{},
	},
	{
		ItsOn:          false,
		Name:           "create class - ok",
		Desc:           "should create a new class",
		Instr:          "create class ok",
		DtoIn:          &dto.ClassCreateIn{ID: "class1", Name: "class1", TaxID: "tax1"},
		DtoOut:         &dto.ClassCreateOut{},
		ExpectedDtoOut: &dto.ClassCreateOut{ID: "class1", Name: "class1", TaxID: "tax1", TaxName: "tax1"},
		ExpectedError:  nil,
		Messages:       []string{},
	},
	{
		ItsOn:          false,
		Name:           "create asset - ok",
		Desc:           "should create a new asset",
		Instr:          "create asset ok",
		DtoIn:          &dto.AssetCreateIn{ID: "asset1", Name: "asset1", ClassID: "class1", StartDate: "2020-01-01"},
		DtoOut:         &dto.AssetCreateOut{},
		ExpectedDtoOut: &dto.AssetCreateOut{ID: "asset1", Name: "asset1", ClassID: "class1", ClassName: "class1", StartDate: "2020-01-01"},
		ExpectedError:  nil,
		Messages:       []string{},
	},
	{
		ItsOn:          false,
		Name:           "create statement - ok",
		Desc:           "should create a new statement",
		Instr:          "create statement ok",
		DtoIn:          &dto.StatementCreateIn{ID: "statement1", AssetID: "asset1", Date: "2020-01-01", History: "FW", Value: 100, Comment: "comment1"},
		DtoOut:         &dto.StatementCreateOut{},
		ExpectedDtoOut: &dto.StatementCreateOut{ID: "statement1", AssetID: "asset1", AssetName: "asset1", Date: "2020-01-01", History: "FW", Value: 100, Comment: "comment1"},
		ExpectedError:  nil,
		Messages:       []string{},
	},
	{
		ItsOn:          false,
		Name:           "create portfolio - ok",
		Desc:           "should create a new portfolio",
		Instr:          "create portfolio ok",
		DtoIn:          &dto.PortfolioCreateIn{ID: "portfolio1", Name: "portfolio1", PortfolioItems: []*dto.PortfolioItemCreate{{AssetID: "asset1"}}},
		DtoOut:         &dto.PortfolioCreateOut{},
		ExpectedDtoOut: &dto.PortfolioCreateOut{ID: "portfolio1", Name: "portfolio1", PortfolioItems: []*dto.PortfolioItemCreate{{AssetID: "asset1"}}},
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
			if !tc.ItsOn {
				t.Skip()
			}
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
	if r.Instr == "create asset ok" && reflect.TypeOf(obj) == reflect.TypeOf(&domain.Class{}) {
		return true, nil
	}
	if r.Instr == "create statement ok" && reflect.TypeOf(obj) == reflect.TypeOf(&domain.Asset{}) {
		return true, nil
	}
	if r.Instr == "create portfolio ok" && reflect.TypeOf(obj) == reflect.TypeOf(&domain.Asset{}) {
		return true, nil
	}
	return false, nil
}
func (r *MockRepository) GetByID(obj interface{}, id string) (bool, error) {
	if r.Instr == "create class ok" && reflect.TypeOf(obj) == reflect.TypeOf(&domain.Tax{}) {
		obj1 := obj.(*domain.Tax)
		obj1.ID = "tax1"
		obj1.Name = "tax1"
		obj1.Period = "Y"
		obj1.TaxItems = []*domain.TaxItem{{ID: "tax1", Until: 1, Value: 0.1}}
		return true, nil
	}
	if r.Instr == "create asset ok" && reflect.TypeOf(obj) == reflect.TypeOf(&domain.Class{}) {
		obj1 := obj.(*domain.Class)
		obj1.ID = "class1"
		obj1.Name = "class1"
		obj1.TaxID = "tax1"
		return true, nil
	}
	if r.Instr == "create statement ok" && reflect.TypeOf(obj) == reflect.TypeOf(&domain.Asset{}) {
		obj1 := obj.(*domain.Asset)
		obj1.ID = "asset1"
		obj1.Name = "asset1"
		obj1.ClassID = "class1"
		obj1.StartDate = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		return true, nil
	}
	if r.Instr == "create portfolio ok" && reflect.TypeOf(obj) == reflect.TypeOf(&domain.Asset{}) {
		obj1 := obj.(*domain.Asset)
		obj1.ID = "asset1"
		obj1.Name = "asset1"
		obj1.ClassID = "class1"
		obj1.StartDate = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
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
