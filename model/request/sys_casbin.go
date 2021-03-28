package request

type CasbinInfo struct {
	Path  string `json:"path"`
	Method string `json:"method"`
}

type CasbinInReceive struct {
	AuthorityId string `json:"authority_id"`
	CasbinInfos []CasbinInfo `json:"casbin_infos"`
}