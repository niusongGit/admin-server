package schema

//request

type SportTypeIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type SportTypeListRequest struct {
	Name   string  `json:"name" validate:"string|min_len:1|max_len:40"`
	Status []int64 `json:"status"  validate:"ints"`
}

type SportTypeAddRequest struct {
	Name                                 string  `json:"name" validate:"string|min_len:1|max_len:40"`
	PostMaxPoints                        float64 `json:"post_max_points" validate:"float"`
	PostMinPoints                        float64 `json:"post_min_points" validate:"float"`
	TemplateCode                         string  `json:"template_code"  validate:"required|alpha_dash"`
	Template                             string  `json:"template"`
	CompetitionFinishDisableEditTemplate int64   `json:"competition_finish_disable_edit_template"  validate:"in:0,1"`
	CompetitionAddDisableEditTemplate    int64   `json:"competition_add_disable_edit_template"  validate:"in:0,1"`
	TeamDictionary                       string  `json:"team_dictionary"`
	PostContentTemplate                  string  `json:"post_content_template"`
	Sort                                 int64   `json:"sort" validate:"int"`
	Status                               int64   `json:"status"  validate:"in:0,1"`
}

type SportTypeUpdateRequest struct {
	Id                                   int64   `json:"id" validate:"required|gt:0"`
	Name                                 string  `json:"name" validate:"string|min_len:1|max_len:40"`
	Sort                                 int64   `json:"sort" validate:"int"`
	PostMaxPoints                        float64 `json:"post_max_points" validate:"float"`
	PostMinPoints                        float64 `json:"post_min_points" validate:"float"`
	TemplateCode                         string  `json:"template_code"  validate:"required|alpha_dash"`
	Template                             string  `json:"template"`
	CompetitionFinishDisableEditTemplate int64   `json:"competition_finish_disable_edit_template"  validate:"in:0,1"`
	CompetitionAddDisableEditTemplate    int64   `json:"competition_add_disable_edit_template"  validate:"in:0,1"`
	TeamDictionary                       string  `json:"team_dictionary"`
	PostContentTemplate                  string  `json:"post_content_template"`
	Status                               int64   `json:"status"  validate:"in:0,1"`
}
