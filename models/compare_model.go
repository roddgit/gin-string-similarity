package models

type Eform_matching_treshold struct {
	// Id          primitive.ObjectID `json:"id,omitempty"`
	Logs_Id       string  `json:"logs_id,omitempty"`
	Name_PMO_Raw  string  `json:"name_pmo_raw,omitempty"`
	Name_Core_Raw string  `json:"name_core_raw,omitempty"`
	Name_PMO      string  `json:"name_pmo,omitempty"`
	Name_Core     string  `json:"name_core,omitempty"`
	Treshold      float64 `json:"treshold,omitempty"`
	Create_Date   string  `json:"create_date,omitempty"`
}
