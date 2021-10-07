package genshin

type ActionTicket struct {
	Code    int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Ticket string `json:"ticket"`
	} `json:"data"`
}

type UserGameRoles struct {
	Code    int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []UserGameRole `json:"list"`
	} `json:"data"`
}

type UserGameRole struct {
	GameBiz    string `json:"game_biz"`
	Region     string `json:"region"`
	RegionName string `json:"region_name"`
	IsOfficial bool   `json:"is_official"`
	GameUid    string `json:"game_uid"`
	Nickname   string `json:"nickname"`
	Level      int    `json:"level"`
}

type CookieAccount struct {
	Code    int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Uid         string `json:"uid"`
		CookieToken string `json:"cookie_token"`
	} `json:"data"`
}

type RewardInfo struct {
	Code    int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		TotalSignDay  int    `json:"total_sign_day"`
		Today         string `json:"today"`
		IsSign        bool   `json:"is_sign"`
		FirstBind     bool   `json:"first_bind"`
		IsSub         bool   `json:"is_sub"`
		MonthFirst    bool   `json:"month_first"`
		SignCntMissed int    `json:"sign_cnt_missed"`
	} `json:"data"`
}

type Sign struct {
	Code    int    `json:"retcode"`
	Message string `json:"message"`
}
