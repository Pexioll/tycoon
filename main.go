package main
import (
	"os"
	"log"
	tl "github.com/JoelOtter/termloop"
	"github.com/Pexioll/tycoon/scene/frame"
)

func main(){
	game := tl.NewGame()
	game.Screen().SetFps(30)

	data, err := os.ReadFile("duza_ramka_1_4.txt")
	if err != nil {
		log.Fatal(err)
	}
	backgroundEntity := tl.NewEntityFromCanvas(0,0,tl.CanvasFromString(string(data)))
	game.Screen().AddEntity(backgroundEntity)
	game.Start()

}
