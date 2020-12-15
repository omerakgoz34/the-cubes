package main

type player struct {
	Name string `json:"name,omitempty"`
	Size int    `json:"size,string,omitempty"`
	// Color        rl.Color  `json:"Color,string,omitempty"`
	// ColorDefault rl.Color  `json:"ColorDefault,string,omitempty"`
	X            float32 `json:"x,string,omitempty"`
	Y            float32 `json:"y,string,omitempty"`
	Speed        float32 `json:"speed,string,omitempty"`
	SpeedDefault float32 `json:"speedDefault,string,omitempty"`
}
