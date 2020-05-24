package car_structure

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

type Axis struct {
	AxisID     int64  `json:"axisID"`
	Left       []Tire `json:"left"`
	Right      []Tire `json:"right"`
	SpareWheel []Tire `json:"spareWheels"`
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

func (sm *Summanry) Sort() {

	for index, axle := range sm.Axles {
		sort.SliceStable(axle.Left, func(i, j int) bool {
			_, _, WheelNoi := extectCode(axle.Left[i].Position)
			_, _, WheelNoj := extectCode(axle.Left[j].Position)
			return WheelNoi < WheelNoj
		})

		sort.SliceStable(axle.Right, func(i, j int) bool {
			_, _, WheelNoi := extectCode(axle.Right[i].Position)
			_, _, WheelNoj := extectCode(axle.Right[j].Position)
			return WheelNoi > WheelNoj
		})

		sm.Axles[index].Left = axle.Left
		sm.Axles[index].Right = axle.Right
	}

}

func (cs *carStructure) Turnable(axisNo int) bool {
	axisLists := strings.Split(cs.TextIntput, "-")
	for no, axis := range axisLists {
		if axisNo == no+1 {
			str := axis[1:2]
			if str == "S" {
				return true
			}
		}

	}
	return false
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

func (cs *carStructure) GetJsonResult() (error, string) {

	var summary Summanry
	summary.AxisQTY = cs.GetAxisQTY()
	summary.WheelQTY = cs.GetWheelQTY()

	for i := summary.AxisQTY; i >= 0; i-- {

		//Add initial struct
		summary.Axles = append(summary.Axles, Axis{
			AxisID: int64(i),
		})

		for _, tireInformation := range cs.TireInformations {

			side, axisNo, _ := extectCode(tireInformation.PositionCode)
			var tire = Tire{
				TireID:           tireInformation.TireID,
				Position:         tireInformation.PositionCode,
				TireSerialNumber: tireInformation.TireSerialNumber,
				TireDepth:        tireInformation.TireDepth,
				Turnable:         cs.Turnable(axisNo),
			}

			if axisNo == 0 { //ยางอะไหล่
				for a := 0; a < len(summary.Axles); a++ {
					if summary.Axles[a].AxisID == 0 {
						summary.Axles[a].SpareWheel = append(summary.Axles[a].SpareWheel, tire)
						break
					}
				}
			}

			if axisNo == i && i != 0 { //Exclude axis 0

				if side == "R" { //ด้านขวา
					for a := 0; a < len(summary.Axles); a++ {
						if summary.Axles[a].AxisID == int64(i) {
							summary.Axles[a].Right = append(summary.Axles[a].Right, tire)
							break
						}
					}
				} else { //ด้านซ้าย
					for a := 0; a < len(summary.Axles); a++ {
						if summary.Axles[a].AxisID == int64(i) {
							summary.Axles[a].Left = append(summary.Axles[a].Left, tire)
							break
						}
					}
				}
			}
		}

	}

	//เรียงข้อมูล
	summary.Sort()

	//แปลง struct to json format
	b, err := json.Marshal(summary)
	if err != nil {
		return err, ""
	}

	// fmt.Println(string(b))

	return nil, string(b)
}
