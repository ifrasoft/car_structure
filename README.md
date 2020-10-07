# Carcode to JSON API Converter
## แปลง โค้ดล้อรถ เช่น  "2S-4T-4T"


### Install with go get

```
go get github.com/ifrasoft/car_structure
```
### Install with go dep
```
dep ensure -add github.com/ifrasoft/car_structure
```

### ตัวอย่างการใช้งาน
```go
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

	cs := car_structure.NewCarStructureConvertor(carCode, carInformations)

	cs.ApplyPolicies(policies)

		imageFile, err := os.Open("../image/tractor.jpg")
	if err != nil {
		fmt.Println("img.jpg file not found!")
	}
	defer imageFile.Close()
	img, _, err := image.Decode(imageFile)
	if err != nil {
		fmt.Println(err)
	}

	cs.InjectImageCarType(img)

	jsonResult, _ := cs.GetJsonResult()

	fmt.Println(jsonResult)
}

```