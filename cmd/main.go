package main

import (
	"fmt"

	"github.com/ifrasoft/car_structure"
)

func main() {
	var carInformations = []*car_structure.TireInformation{
		&car_structure.TireInformation{
			TireID:           1,
			PositionCode:     "1-L1",
			TireSerialNumber: "TireSerialNumber",
			TireDepth:        1.0,
		},
		&car_structure.TireInformation{
			TireID:           2,
			PositionCode:     "1-R1",
			TireSerialNumber: "TireSerialNumber",
			TireDepth:        1.0,
		},
		&car_structure.TireInformation{
			TireID:           4,
			PositionCode:     "2-L2",
			TireSerialNumber: "TireSerialNumber",
			TireDepth:        1.0,
		},
		&car_structure.TireInformation{
			TireID:           3,
			PositionCode:     "2-L1",
			TireSerialNumber: "TireSerialNumber",
			TireDepth:        1.0,
		},

		&car_structure.TireInformation{
			TireID:           5,
			PositionCode:     "2-R1",
			TireSerialNumber: "TireSerialNumber",
			TireDepth:        1.0,
		},
		&car_structure.TireInformation{
			TireID:           6,
			PositionCode:     "2-R2",
			TireSerialNumber: "TireSerialNumber",
			TireDepth:        1.0,
		},
		&car_structure.TireInformation{
			TireID:           7,
			PositionCode:     "3-L1",
			TireSerialNumber: "TireSerialNumber",
		},
		&car_structure.TireInformation{
			TireID:           8,
			PositionCode:     "3-L2",
			TireSerialNumber: "TireSerialNumber",
		},
		&car_structure.TireInformation{
			TireID:           9,
			PositionCode:     "3-R1",
			TireSerialNumber: "TireSerialNumber",
			TireDepth:        1.0,
		},
		&car_structure.TireInformation{
			TireID:           10,
			PositionCode:     "3-R2",
			TireSerialNumber: "TireSerialNumber",
			TireDepth:        1.0,
		},
		&car_structure.TireInformation{
			TireID:           11,
			PositionCode:     "0-B1",
			TireSerialNumber: "TireSerialNumber",
			TireDepth:        1.0,
		},
	}

	carCode := "2S-4T-4T"
	cs := car_structure.NewCarStructureConvertor(carCode, carInformations)
	jsonResult, _ := cs.GetJsonResult()
	fmt.Println(jsonResult)
}
