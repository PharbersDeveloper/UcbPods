package UcbDataStorage

import (
	"fmt"
	"errors"
	"Ucb/UcbModel"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
)

// UcbLevelConfigStorage stores all of the tasty modelleaf, needs to be injected into
// LevelConfig and LevelConfig Resource. In the real world, you would use a database for that.
type UcbLevelConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbLevelConfigStorage) NewLevelConfigStorage(args []BmDaemons.BmDaemon) *UcbLevelConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbLevelConfigStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbLevelConfigStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.LevelConfig {
	in := UcbModel.LevelConfig{}
	var out []UcbModel.LevelConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(&iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s UcbLevelConfigStorage) GetOne(id string) (UcbModel.LevelConfig, error) {
	in := UcbModel.LevelConfig{ID: id}
	out := UcbModel.LevelConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("LevelConfig for id %s not found", id)
	return UcbModel.LevelConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbLevelConfigStorage) Insert(c UcbModel.LevelConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbLevelConfigStorage) Delete(id string) error {
	in := UcbModel.LevelConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("LevelConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbLevelConfigStorage) Update(c UcbModel.LevelConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("LevelConfig with id does not exist")
	}

	return nil
}
