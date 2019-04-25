package countries

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// RegionCode - Region code (UN M.49 code standart)
type RegionCode int64 // int64 for database/sql/driver.Valuer compatibility

// Region - all info about region
type Region struct {
	Name string
	Code RegionCode
}

// TypeRegionCode for Typer interface
const TypeRegionCode = "countries.RegionCode"

// TypeRegion for Typer interface
const TypeRegion = "countries.Region"

const (
	RegionUnknown RegionCode = 0
	RegionAF      RegionCode = 2
	RegionNA      RegionCode = 3
	RegionSA      RegionCode = 5
	RegionOC      RegionCode = 9
	RegionAN      RegionCode = 999
	RegionAS      RegionCode = 142
	RegionEU      RegionCode = 150
)

const (
	RegionAfrica       RegionCode = 2
	RegionNorthAmerica RegionCode = 3
	RegionSouthAmerica RegionCode = 5
	RegionOceania      RegionCode = 9
	RegionAntarctica   RegionCode = 999
	RegionAsia         RegionCode = 142
	RegionEurope       RegionCode = 150
)

// Type implements Typer interface
func (c RegionCode) Type() string {
	return TypeRegionCode
}

// String - implements fmt.Stringer, returns a Region name in english
func (c RegionCode) String() string {
	switch c {
	case RegionAF:
		return "Africa"
	case RegionNA:
		return "North America"
	case RegionOC:
		return "Oceania"
	case RegionAN:
		return "Antarctica"
	case RegionAS:
		return "Asia"
	case RegionEU:
		return "Europe"
	case RegionSA:
		return "South America"
	}
	return UnknownMsg
}

// String - implements fmt.Stringer, returns a Region name in russian
func (c RegionCode) StringRus() string {
	switch c {
	case RegionAF:
		return "Африка"
	case RegionNA:
		return "Северная Америка"
	case RegionOC:
		return "Океания"
	case RegionAN:
		return "Антарктика"
	case RegionAS:
		return "Азия"
	case RegionEU:
		return "Европа"
	case RegionSA:
		return "Южная Америка"
	}
	return UnknownMsg
}

// TotalRegions - returns number of Regions codes in the package
func TotalRegions() int {
	return 7
}

func (c RegionCode) Info() *Region {
	return &Region{
		Name: c.String(),
		Code: c,
	}
}

// Type implements Typer interface
func (r *Region) Type() string {
	return TypeRegion
}

// Value implements database/sql/driver.Valuer
func (r Region) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan implements database/sql.Scanner
func (r *Region) Scan(src interface{}) error {
	if r == nil {
		return fmt.Errorf("countries::Scan: Region scan err: region == nil")
	}
	switch src := src.(type) {
	case *Region:
		*r = *src
	case Region:
		*r = src
	case nil:
		r = nil
	default:
		return fmt.Errorf("countries::Scan: Region scan err: unexpected value of type %T for %T", src, *r)
	}
	return nil
}

// AllRegions - returns all Regions
func AllRegions() []RegionCode {
	return []RegionCode{
		RegionAF,
		RegionNA,
		RegionOC,
		RegionAN,
		RegionAS,
		RegionEU,
		RegionSA,
	}
}

// AllRegions - return all currencies as []Region
func AllRegionsInfo() []*Region {
	all := AllRegions()
	regions := make([]*Region, 0, len(all))
	for _, v := range all {
		regions = append(regions, v.Info())
	}
	return regions
}
