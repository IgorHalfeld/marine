package model

// MarineTraffic represents the marine struct data
type MarineTraffic struct {
	Extent    []float64  `json:"extent"`
	Positions []Position `json:"positions"`
}

// Position of the ship
type Position struct {
	Tst      string  `json:"tst"`
	ShipName string  `json:"ship_name"`
	Mmsi     string  `json:"mmsi"`
	Imo      string  `json:"imo"`
	Lat      string  `json:"lat"`
	Lon      string  `json:"lon"`
	Cog      float32 `json:"cog"`
	Sog      float32 `json:"sog"`
	Heading  string  `json:"heading"`
	Type     string  `json:"type"`
	Class    string  `json:"class"`
	Eta      string  `json:"eta"`
	Sources  string  `json:"sources"`
	Icon     int     `json:"icon"`
}

// Alert represents the marine struct data
type Alert struct {
	ImageURL    string `jsong:"image_url"`
	Type        string `jsong:"type"`
	Description string `jsong:"description"`
	FingerPrint string `jsong:"finger_print"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

// ResponseMessage represents the response message with a just a string detailing the result
type ResponseMessage struct {
	Message string `json:"message"`
}

// GeneralData represents the data from the porto de santos
type GeneralData struct {
	PesquisarTotalizadoresResult PesquisarTotalizadoresResult `json:"PesquisarTotalizadoresResult"`
}

// AcumuladosMes represents the acumulados mes from porto de santos
type AcumuladosMes struct {
	AnoAnterior   string `json:"AnoAnterior"`
	AnoAtual      string `json:"AnoAtual"`
	MesAnterior   string `json:"MesAnterior"`
	MesAtual      string `json:"MesAtual"`
	TotalAnterior int    `json:"TotalAnterior"`
	TotalAtual    int    `json:"TotalAtual"`
}

// PesquisarTotalizadoresResult represents the total from porto de santos
type PesquisarTotalizadoresResult struct {
	AcumuladosMes              AcumuladosMes `json:"AcumuladosMes"`
	EmOperacao                 int           `json:"EmOperacao"`
	NaviosAtracadosCaisPublico int           `json:"NaviosAtracadosCaisPublico"`
	NaviosAtracadosTerminais   int           `json:"NaviosAtracadosTerminais"`
	NaviosEsperados            int           `json:"NaviosEsperados"`
}
