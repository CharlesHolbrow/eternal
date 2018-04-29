package eternal

import "fmt"

const chunkSize = 8

// Split decomposes an int into a chunk, tile coords
func split(i int) (chunk int, tile int) {
	tile = i % chunkSize
	if tile < 0 {
		tile = -tile
	}
	if i < 0 {
		i++
	}
	chunk = i / chunkSize
	return chunk, tile
}

func makeSubKey(mapName string, cx, cy int) string {
	return fmt.Sprintf("%s|%d|%d", mapName, cx, cy)
}
