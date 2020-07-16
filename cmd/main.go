package main

import (
	"fmt"
	"image"
	"os"

	"github.com/ifrasoft/car_structure"
)

func main() {
	var carInformations = []*car_structure.TireInformation{
		&car_structure.TireInformation{
			TireID:              1,
			PositionCode:        "1-L1",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    10,
			TireDepthMinimum:    10,
			TirePressureMaximum: 10,
			TirePressureMinimum: 10,
		},
		&car_structure.TireInformation{
			TireID:              2,
			PositionCode:        "1-R1",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    4,
			TireDepthMinimum:    4,
			TirePressureMaximum: 30,
			TirePressureMinimum: 30,
		},
		&car_structure.TireInformation{
			TireID:              4,
			PositionCode:        "2-L2",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 1.0,
			TirePressureMinimum: 1.0,
		},
		&car_structure.TireInformation{
			TireID:              3,
			PositionCode:        "2-L1",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 25.0,
			TirePressureMinimum: 25.0,
		},

		&car_structure.TireInformation{
			TireID:              5,
			PositionCode:        "2-R1",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 1.0,
			TirePressureMinimum: 1.0,
		},
		&car_structure.TireInformation{
			TireID:              6,
			PositionCode:        "2-R2",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 1.0,
			TirePressureMinimum: 1.0,
		},
		&car_structure.TireInformation{
			TireID:              7,
			PositionCode:        "3-L1",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 1.0,
			TirePressureMinimum: 1.0,
		},
		&car_structure.TireInformation{
			TireID:              8,
			PositionCode:        "3-L2",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 1.0,
			TirePressureMinimum: 1.0,
		},
		&car_structure.TireInformation{
			TireID:              9,
			PositionCode:        "3-R1",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 1.0,
			TirePressureMinimum: 1.0,
		},
		&car_structure.TireInformation{
			TireID:              10,
			PositionCode:        "3-R2",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 1.0,
			TirePressureMinimum: 1.0,
		},
		&car_structure.TireInformation{
			TireID:              11,
			PositionCode:        "0-B1",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 1.0,
			TirePressureMinimum: 1.0,
		},
		&car_structure.TireInformation{
			TireID:              1,
			PositionCode:        "4-R1",
			TireSerialNumber:    "TireSerialNumber",
			TireDepthMaximum:    1.0,
			TireDepthMinimum:    1.0,
			TirePressureMaximum: 1.0,
			TirePressureMinimum: 1.0,
		},
	}

	policies := []*car_structure.Policy{
		&car_structure.Policy{
			AxlesNo:              1,
			StandardTireDepth:    1.0,
			WarningTireDepth:     10,
			CriticalTireDepth:    5,
			StandardTirePressure: 30,
			WarningTirePressure:  10,
			CriticalTirePressure: 20,
		}, &car_structure.Policy{
			AxlesNo:              2,
			StandardTireDepth:    1.0,
			WarningTireDepth:     1.0,
			CriticalTireDepth:    1.0,
			StandardTirePressure: 30,
			WarningTirePressure:  10,
			CriticalTirePressure: 20,
		},
		&car_structure.Policy{
			AxlesNo:              3,
			StandardTireDepth:    1.0,
			WarningTireDepth:     1.0,
			CriticalTireDepth:    1.0,
			StandardTirePressure: 1.0,
			WarningTirePressure:  1.0,
			CriticalTirePressure: 1.0,
		},
		&car_structure.Policy{
			AxlesNo:              4,
			StandardTireDepth:    1.0,
			WarningTireDepth:     1.0,
			CriticalTireDepth:    1.0,
			StandardTirePressure: 1.0,
			WarningTirePressure:  1.0,
			CriticalTirePressure: 1.0,
		},
	}

	carCode := "2S-4T-4T-2T"
	carType := "tractor"

	cs := car_structure.NewCarStructureConvertor(carCode, carType, carInformations)

	cs.ApplyPolicies(policies)

	Tractor, err := os.Open("../image/tractor-head.png")
	if err != nil {
		fmt.Println("img.png file not found!")
	}
	defer Tractor.Close()
	TractorH, _, err := image.Decode(Tractor)
	if err != nil {
		fmt.Println(err)
	}
	Trailer, err := os.Open("../image/trailer-head.png")
	if err != nil {
		fmt.Println("img.png file not found!")
	}
	defer Trailer.Close()
	TrailerH, _, err := image.Decode(Trailer)
	if err != nil {
		fmt.Println(err)
	}

	Center, err := os.Open("../image/center.png")
	if err != nil {
		fmt.Println("img.png file not found!")
	}
	defer Center.Close()
	Body, _, err := image.Decode(Center)
	if err != nil {
		fmt.Println(err)
	}

	Foot, err := os.Open("../image/footer.png")
	if err != nil {
		fmt.Println("img.png file not found!")
	}
	defer Foot.Close()
	Footer, _, err := image.Decode(Foot)
	if err != nil {
		fmt.Println(err)
	}

	cs.InjectImageCarType(TractorH, TrailerH, Body, Footer)

	jsonResult, _ := cs.GetJsonResult()

	fmt.Println(jsonResult)

}
