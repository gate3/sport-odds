package bookmaker

type SportApiModel struct {
	Key 		string		`json:"key"`
	Active		bool		`json:"active"`
	Group		string		`json:"group"`
	Details		string		`json:"details"`
	Title		string		`json:"title"`
}
