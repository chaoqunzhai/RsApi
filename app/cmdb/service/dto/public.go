package dto

type CityTreeRow struct {
	Code     int           `json:"code"`
	Name     string        `json:"name"`
	Children []interface{} `json:"children"`
}
