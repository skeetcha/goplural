package main

type Member struct {
	Avatar   string     `json:"avatar"`
	Name     string     `json:"name"`
	Pronouns string     `json:"pronouns"`
	Proxy    []ProxyTag `json:"proxy"`
}

type ProxyTag struct {
	Prefix *string `json:"prefix,omitempty"`
	Suffix *string `json:"suffix,omitempty"`
}
