package frame
import (
	"log"
	"os"
	tl "github.com/JoelOtter/termloop"
)

// Frame to struktura reprezentująca graficzną ramkę.
// Zagnieżdżamy *tl.Entity, co sprawia, że Frame automatycznie implementuje
// interfejsy tl.Drawable i tl.Tickable, dziedzicząc zachowanie Entity.
type Frame struct {
	*tl.Entity
}

// NewFrameFromFile tworzy nową ramkę, wczytując jej zawartość ASCII art z podanego pliku.
// Jeśli plik nie może zostać wczytany, program zakończy działanie z błędem.
func NewFrameFromFile(filePath string) *Frame {
	data, err := os.ReadFile("../utils/frame.txt")
	if err != nil {
		// Używamy log.Fatalf, ponieważ brak pliku ramki jest błędem krytycznym
		// i chcemy zatrzymać aplikację.
		log.Fatalf("Błąd wczytywania pliku ramki '%s': %v", filePath, err)
	}

	asciiArtString := string(data)
	// tl.CanvasFromString tworzy Canvas z podanego stringa.
	canvas := tl.CanvasFromString(asciiArtString)

	// Tworzymy nową encję (Entity) Termloop z utworzonego Canvasu.
	// Pozycja (0,0) oznacza lewy górny róg ekranu.
	entity := tl.NewEntityFromCanvas(0, 0, canvas)

	// Zwracamy wskaźnik do nowej instancji Frame, która zawiera naszą encję.
	return &Frame{Entity: entity}
}

// Możesz dodać więcej metod do typu Frame, jeśli chcesz, aby ramka miała
// specjalne zachowanie, np. animacje, przesuwanie itp.
// Na przykład, możesz dodać metodę do zmiany pozycji ramki:

func (f *Frame) SetPosition(x, y int) {
	f.Entity.SetPosition(x, y)
}
