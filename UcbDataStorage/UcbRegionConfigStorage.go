package UcbDataStorage

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"Ucb/UcbModel"
	"github.com/manyminds/api2go"
	"net/http"
)

// UcbRegionConfigStorage stores all of the tasty chocolate, needs to be injected into
// RegionConfig Resource. In the real world, you would use a database for that.
type UcbRegionConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbRegionConfigStorage) NewRegionConfigStorage(args []BmDaemons.BmDaemon) *UcbRegionConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbRegionConfigStorage{mdb}
}

// GetAll of the chocolate
func (s UcbRegionConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.RegionConfig {
	in := UcbModel.RegionConfig{}
	var out []UcbModel.RegionConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.RegionConfig
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Student)
	}
}

// GetOne
func (s UcbRegionConfigStorage) GetOne(id string) (UcbModel.RegionConfig, error) {
	in := UcbModel.RegionConfig{ID: id}
	out := UcbModel.RegionConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		// TODO: 双重绑定没明白啥意思
		//双重绑定
		//if out.RegionID != "" {
		//	item, err := UcbRegionStorage{db: s.db}.GetOne(out.RegionID)
		//	if err != nil {
		//		return UcbModel.RegionConfig{}, err
		//	}
		//	out.Region = item
		//}

		return out, nil
	}
	errMessage := fmt.Sprintf("RegionConfig for id %s not found", id)
	return UcbModel.RegionConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbRegionConfigStorage) Insert(c UcbModel.RegionConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbRegionConfigStorage) Delete(id string) error {
	in := UcbModel.RegionConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RegionConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *UcbRegionConfigStorage) Update(c UcbModel.RegionConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RegionConfig with id does not exist")
	}

	return nil
}

func (s *UcbRegionConfigStorage) Count(req api2go.Request, c UcbModel.RegionConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
