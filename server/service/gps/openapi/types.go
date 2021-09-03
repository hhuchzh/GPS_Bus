package openapi

type AccessTokenResp struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Result AccessToken `json:"result"`
}

type AccessToken struct {
    AccessToken string `json:"accessToken"`
    ExpiresIn int `json:"expiresIn"`
    Account string `json:"account"`
    AppKey string `json:"appKey"`
    RefreshToken string `json:"refreshToken"`
    Time string `json:"time"`
}

type DeviceResp struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Result []*Device `json:"result"`
}

type Device struct {
    Imei string `json:"imei"`
    DeviceName string `json:"deviceName"`
    McType string `json:"mcType"`
    McTypeUseScope string `json:"mcTypeUseScope"`
    EquipType string `json:"equipType"`
    Sim string `json:"sim"`
    Expiration string `json:"expiration"`
    ActivationTime string `json:"activationTime"`
    ReMark string `json:"reMark"`
    VehicleName string `json:"vehicleName"`
    VehicleIcon string `json:"vehicleIcon"`
    VehicleNumber string `json:"vehicleNumber"`
    VehicleModels string `json:"vehicleModles"`
    CarFrame string `json:"carFrame"`
    DriverName string `json:"driverName"`
    DriverPhone string `json:"driverPhone"`
    EnabledFlag int `json:"enabledFlag"`
    EngineNumber string `json:"engineNumber"`
}

type LocationResp struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Result []*Location `json:"result"`
}

type Location struct {
    Imei string `json:"imei"`
    DeviceName string `json:"deviceName"`
    Icon string `json:"icon"`
    Status string `json:"status"`
    Lng float64 `json:"lng"`
    Lat float64 `json:"lat"`
    ExpireFlag string `json:"expireFlag"`
    ActivationFlag string `json:"activationFlag"`
    PosType string `json:"posType"`
    GpsTime string `json:"gpsTime"}`
    HbTime string `json:"hbTime"`
    Speed string `json:"speed"`
    AccStatus string `json:"accStatus"`
    ElectQuantity string `json:"electQuantity"`
    PowerValue string `json:"powerValue"`
    GpsNum string `json:"gpsNum"`
    Direction string `json:"direction"`
    Mileage string `json:"mileage"`
}

type TrackResp struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Result []*Track `json:"result"`
}

type Track struct {
    Lng float64 `json:"lng"`
    Lat float64 `json:"lat"`
    GpsTime string `json:"gpsTime"`
    Direction int `json:"direction"`
    GpsSpeed float32 `json:"gpsSpeed"`
    PosType int `json:"posType"`
}
