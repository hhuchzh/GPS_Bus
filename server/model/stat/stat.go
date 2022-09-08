package stat

//"time"

type StatInfo struct {
	RouteNum        int           `json:"routeNum" form:"routeNum"`
	ShuttleNum      int           `json:"shuttleNum" form:"shuttleNum"`
	StationNum      int           `json:"stationNum" form:"StationNum"`
	ClassesNum      int           `json:"classesNum" form:"ClassesNum"`
	TotalMileage    int           `json:"totalMileage" form:"totalMileage"`
	ShuttleLineList []ShuttleLine `json:"shuttleLineList" form:"shuttleLineList"`
}

type ShuttleLine struct {
	RouteName  string `json:"routeName" form:"routeName"`
	ShuttleNum int    `josn:"shuttleNum" form:"shuttleNum"`
}
