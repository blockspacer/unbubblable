package main

// MapTile определяет тайл карты
type MapTile struct {
	// Type определяет тип тайла
	Type int
	// Name определяет тип тайла
	Name string
	// Speed определяет коэффициент скорости движения по тайлу
	Speed float64
	// Texture определяет номер текстуры в атласе
	Texture int
}

var Tiles = []MapTile{
	{1, "Grass", 1.0, 1},         // Трава
	{2, "Stone", 1.0, 2},         // Камень
	{3, "Dirt", 1.0, 3},          // Земля
	{17, "Cobblestone", 1.0, 17}, // Булыжник
	{18, "Bedrock", 1.0, 18},     // Коренная порода
	{19, "Sand", 1.0, 19},        // Песок
	{20, "Gravel", 1.0, 20},      // Гравий
	{1, "", 1.0, 1},
	{1, "", 1.0, 1},
	{1, "", 1.0, 1},
	{1, "", 1.0, 1},
	{1, "", 1.0, 1},
	{1, "", 1.0, 1},
}
