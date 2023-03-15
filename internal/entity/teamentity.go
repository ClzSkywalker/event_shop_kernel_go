package entity

type TeamItem struct {
	TeamId      string `json:"team_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Sort        int    `json:"sort,omitempty"`
}

type TeamResp struct {
	TeamList []TeamItem `json:"team_list"`
}

type TeamCreateReq struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Sort        int    `json:"sort,omitempty"`
}

type TeamCreateResp struct {
	TeamId string `json:"team_id"`
}

type TeamUpdateReq struct {
	TeamItem
}

type TeamDelReq struct {
	TeamId string `json:"team_id" binding:"required"`
}
