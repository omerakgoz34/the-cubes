package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gorilla/websocket"
)

// Launch Flags
var flagName = flag.String("name", "", "Your player name. (Required)")
var flagType = flag.String("type", "", "Sets the multiplayer connection type. [host/peer] (Required) ")
var flagAddr = flag.String("addr", "localhost:3434", "The adress of the server.")

func main() {
	flag.Parse()
	log.SetFlags(0)

	// Flag Checks
	if *flagName == "" {
		fmt.Println("Settings:")
		flag.PrintDefaults()
		os.Exit(1)
	} else if *flagType == "" {
		fmt.Println("Settings:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Reassign Flags
	name := string(*flagName)
	mode := string(*flagType)
	addr := string(*flagAddr)
	url := url.URL{Scheme: "ws", Host: addr, Path: "/"}

	// Init Raylib
	fmt.Println("starting raylib...")

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(700, 500, "The Cubes")
	rl.SetTargetFPS(60)

	// Players
	playerOne := player{
		Name: name,
		Size: 50.0,
		// Color:        rl.SkyBlue,
		// ColorDefault: rl.SkyBlue,
		X:            float32(rl.GetScreenWidth() / 3),
		Y:            float32(rl.GetScreenHeight() / 2),
		Speed:        10.0,
		SpeedDefault: 10.0,
	}

	playerTwo := player{
		Name: "Player 2",
		Size: 50.0,
		// Color:        rl.SkyBlue,
		// ColorDefault: rl.SkyBlue,
		X:            float32(rl.GetScreenWidth() / 2),
		Y:            float32(rl.GetScreenHeight() / 3),
		Speed:        10.0,
		SpeedDefault: 10.0,
	}

	// Multiplayer Socket Connection
	connected := false
	if mode == "host" {
		upgrader := websocket.Upgrader{}
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			c, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Print("Server upgrade error: ", err)
				return
			}

			defer c.Close()

			go func() {
				for {
					// Send Data
					err = c.WriteJSON(playerOne)
					if err != nil {
						log.Println("Write: ", err)
						return
					}

					time.Sleep(time.Millisecond * 16)
				}
			}()

			// time.Sleep(time.Millisecond * 16)

			for {
				// Recv Data
				_, recvMsg, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					return
				}
				err = json.Unmarshal(recvMsg, &playerTwo)
				if err != nil {
					fmt.Println("JSON Decoding Error: ", err)
				}

				connected = true
				time.Sleep(time.Millisecond * 16)
			}
		})

		go http.ListenAndServe(addr, nil)

	} else if mode == "peer" {
		c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
		if err != nil {
			log.Fatal("dial error:", err)
		}

		go func() {
			defer c.Close()

			go func() {
				for {
					// Recv Data
					_, recvMsg, err := c.ReadMessage()
					if err != nil {
						log.Println("read:", err)
						return
					}
					err = json.Unmarshal(recvMsg, &playerTwo)
					if err != nil {
						fmt.Println("JSON Decoding Error: ", err)
					}

					connected = true
					time.Sleep(time.Millisecond * 15)
				}
			}()

			// time.Sleep(time.Millisecond * 16)

			for {
				// Send Data
				err = c.WriteJSON(playerOne)
				if err != nil {
					log.Println("Write: ", err)
					return
				}

				time.Sleep(time.Millisecond * 16)
			}
		}()
	}

	fmt.Println("starting game loop...")

	// Game Loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Background
		rl.ClearBackground(rl.Gray)

		// Mouse Clicks
		/*
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(playerOne.X), float32(playerOne.Y), float32(playerOne.Size), float32(playerOne.Size))) {
				// playerOne.Color = rl.Red
				playerOne.DamageTimer = time.Now()
				playerOne.Score--
			} else if rl.IsMouseButtonPressed(rl.MouseRightButton) && rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(playerOne.X), float32(playerOne.Y), float32(playerOne.Size), float32(playerOne.Size))) {
				// playerOne.Color = rl.Green
				playerOne.DamageTimer = time.Now()
				playerOne.Score++
			}
		*/

		// Controls
		if rl.IsKeyDown(rl.KeyRight) && playerOne.X < float32(rl.GetScreenWidth()-playerOne.Size) {
			playerOne.X = playerOne.X + playerOne.Speed
		}
		if rl.IsKeyDown(rl.KeyLeft) && playerOne.X > 0 {
			playerOne.X = playerOne.X - playerOne.Speed
		}
		if rl.IsKeyDown(rl.KeyUp) && playerOne.Y > 0 {
			playerOne.Y = playerOne.Y - playerOne.Speed
		}
		if rl.IsKeyDown(rl.KeyDown) && playerOne.Y < float32(rl.GetScreenHeight()-playerOne.Size) {
			playerOne.Y = playerOne.Y + playerOne.Speed
		}

		/*
			if time.Since(playerOne.DamageTimer).Milliseconds() >= 50 {
				playerOne.Color = rl.SkyBlue
			}
		*/

		// playerOne
		rl.DrawRectangle(int32(playerOne.X), int32(playerOne.Y), int32(playerOne.Size), int32(playerOne.Size), rl.SkyBlue)
		rl.DrawText(playerOne.Name, int32(playerOne.X)-5, int32(playerOne.Y-float32(playerOne.Size/2)), 15, rl.White)

		// playerTwo
		rl.DrawRectangle(int32(playerTwo.X), int32(playerTwo.Y), int32(playerTwo.Size), int32(playerTwo.Size), rl.SkyBlue)
		rl.DrawText(playerTwo.Name, int32(playerTwo.X)-5, int32(playerTwo.Y-float32(playerTwo.Size/2)), 15, rl.White)

		// Connection Error
		if connected == false {
			rl.DrawText("No Connection", 8, int32(rl.GetScreenHeight()-30), 30, rl.Red)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
