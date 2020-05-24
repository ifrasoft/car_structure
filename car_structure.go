package car_structure

import (
	"fmt"
	"strconv"
	"strings"
)

type Axis struct {
	AxisID int64  `json:"axisID"`
	Left   []Tire `json:"left"`
	Right  []Tire `json:"right"`
}

type Wheel struct {
	Tires []Tire `json:"tires"`
}

type Tire struct {
	TireID           int64   `json:"TireID"`
	Position         string  `json:"position"`
	TireSerialNumber string  `json:"tireSerialNumber"`
	TireDepth        float64 `json:"tireDepth"`
	Turnable         bool    `json:"turnable"`
}

type Summanry struct {
	AxisQTY  int    `json:"axisQTY"`
	WheelQTY int    `json:"wheelQTY"`
	Axles    []Axis `json:"axis"`
}

type carStructure struct {
	TextIntput       string
	TireInformations []*TireInformation
}

type TireInformation struct {
	TireID           int64   `json:"TireID"`
	PositionCode     string  `json:"position"`
	TireSerialNumber string  `json:"tireSerialNumber"`
	TireDepth        float64 `json:"tireDepth"`
}

func NewCarStructureConvertor(input string, tireInformations []*TireInformation) *carStructure {
	return &carStructure{
		TextIntput:       input,
		TireInformations: tireInformations,
	}
}

func (cs *carStructure) GetAxisQTY() int {
	return len(strings.Split(cs.TextIntput, "-"))
}

func (cs *carStructure) GetWheelQTY() int {

	count := 0
	axisLists := strings.Split(cs.TextIntput, "-")
	for _, axis := range axisLists {
		numberStr := axis[0:1]
		numberInt, _ := strconv.Atoi(numberStr)
		count = count + numberInt

	}

	return count
}

func (cs *carStructure) filTireInformation() {

}

func extectCode(positionCode string) (side string, axisNo, WheelNo int) {
	data := strings.Split(positionCode, "-")
	axisNo, _ = strconv.Atoi(data[0])
	WheelNo, _ = strconv.Atoi(data[1][1:2])

	side = "L"
	if strings.Contains(data[1], "R") {
		side = "R"
	}
	return
}

func (cs *carStructure) GetSummary() {

	a := cs.GetAxisQTY()
	fmt.Println("Axis QTY:", a)

	var summary Summanry
	summary.AxisQTY = cs.GetAxisQTY()
	summary.WheelQTY = cs.GetWheelQTY()

	for i := summary.AxisQTY; i >= 0; i-- {

		summary.Axles = append(summary.Axles, Axis{
			AxisID: int64(i),
		})
		for _, tireInformation := range cs.TireInformations {

			side, axisNo, WheelNo := extectCode(tireInformation.PositionCode)
			if axisNo == i {
				fmt.Println(tireInformation.PositionCode, side, WheelNo)

				var tire = Tire{
					TireID:           tireInformation.TireID,
					Position:         tireInformation.PositionCode,
					TireSerialNumber: tireInformation.TireSerialNumber,
					TireDepth:        tireInformation.TireDepth,
				}
				if side == "R" {
					for a := 0; a < len(summary.Axles); a++ {
						if summary.Axles[a].AxisID == int64(i) {
							summary.Axles[a].Right = append(summary.Axles[a].Right, tire)
							break
						}
					}
				} else {
					for a := 0; a < len(summary.Axles); a++ {
						if summary.Axles[a].AxisID == int64(i) {
							summary.Axles[a].Right = append(summary.Axles[a].Right, tire)
							break
						}
					}
				}
			}
		}
		// 	summary.Left = append(summary.Left, Wheel{
		// 		AxisID: int64(i),
		// 	})
		// 	summary.Right = append(summary.Left, Wheel{
		// 		AxisID: int64(i),
		// 	})
	}

	// for _, tireInformation := range cs.TireInformations {
	// 	side, axisNo, WheelNo := extectCode(tireInformation.PositionCode)
	// 	if side == "R" {
	// 		summary.Left[]
	// 	} else {

	// 	}
	// }

	fmt.Printf("%+v\n", summary)
}
