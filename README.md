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

	"github.com/ifrasoft/car_structure"
)

func main() {
    //carInformations คือข้อมูลที่เกี่ยวกับยางรถ เช่น id,ตำแหน่งบนรถ,หมายเลขยาง, ความลึกดอกยาง เป็นต้น
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
    
    //carCode คือ โค้ดล้อรถ
    carCode := "2S-4T-4T"
    
    cs := car_structure.NewCarStructureConvertor(carCode, carInformations)
    
    jsonResult, _ := cs.GetJsonResult()
    
	fmt.Println(jsonResult)
}
```
## Result
```json
{
   "axisQTY":3,
   "wheelQTY":10,
   "axis":[
      {
         "axisID":3,
         "left":[
            {
               "TireID":7,
               "position":"3-L1",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":0,
               "turnable":false
            },
            {
               "TireID":8,
               "position":"3-L2",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":0,
               "turnable":false
            }
         ],
         "right":[
            {
               "TireID":10,
               "position":"3-R2",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":1,
               "turnable":false
            },
            {
               "TireID":9,
               "position":"3-R1",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":1,
               "turnable":false
            }
         ],
         "spareWheels":null
      },
      {
         "axisID":2,
         "left":[
            {
               "TireID":3,
               "position":"2-L1",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":1,
               "turnable":false
            },
            {
               "TireID":4,
               "position":"2-L2",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":1,
               "turnable":false
            }
         ],
         "right":[
            {
               "TireID":6,
               "position":"2-R2",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":1,
               "turnable":false
            },
            {
               "TireID":5,
               "position":"2-R1",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":1,
               "turnable":false
            }
         ],
         "spareWheels":null
      },
      {
         "axisID":1,
         "left":[
            {
               "TireID":1,
               "position":"1-L1",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":1,
               "turnable":true
            }
         ],
         "right":[
            {
               "TireID":2,
               "position":"1-R1",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":1,
               "turnable":true
            }
         ],
         "spareWheels":null
      },
      {
         "axisID":0,
         "left":null,
         "right":null,
         "spareWheels":[
            {
               "TireID":11,
               "position":"0-B1",
               "tireSerialNumber":"TireSerialNumber",
               "tireDepth":1,
               "turnable":false
            }
         ]
      }
   ]
}
```