package main

// Player ...
type Player struct {
	Name string `json:"name,omitempty"`
	Size int    `json:"size,string,omitempty"`
	// Color        rl.Color  `json:"Color,string,omitempty"`
	// ColorDefault rl.Color  `json:"ColorDefault,string,omitempty"`
	X            float32 `json:"x,string,omitempty"`
	Y            float32 `json:"y,string,omitempty"`
	Speed        float32 `json:"speed,string,omitempty"`
	SpeedDefault float32 `json:"speedDefault,string,omitempty"`
}

/*
// PlayerOne ...
var PlayerOne = Player{
    Name: "Player 1",
    Size: 50.0,
    // Color:        rl.SkyBlue,
    // ColorDefault: rl.SkyBlue,
    X:            float32(rl.GetScreenWidth() / 3),
    Y:            float32(rl.GetScreenHeight() / 2),
    Speed:        10.0,
    SpeedDefault: 10.0,
}

// PlayerTwo ...
var PlayerTwo = Player{
    Name: "Player 2",
    Size: 50.0,
    // Color:        rl.SkyBlue,
    // ColorDefault: rl.SkyBlue,
    X:            float32(rl.GetScreenWidth() / 2),
    Y:            float32(rl.GetScreenHeight() / 3),
    Speed:        10.0,
    SpeedDefault: 10.0,
}
*/
