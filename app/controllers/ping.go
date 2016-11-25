package controllers

import "github.com/revel/revel"

type AppPing struct {
	App
}

type Response struct {
	Message string `json:"Message"`
}

func (c AppPing) Ping() revel.Result {
	data := make(map[string]interface{})
	data["error"] = nil
	response := Response{Message: "pong"}
	data["response"] = response
	return c.RenderJson(data)
}
