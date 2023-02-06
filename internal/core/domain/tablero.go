package domain

type Tablero struct {
	celdas       [][]*celda
	cantFilas    int
	cantColumnas int
}

func (t *Tablero) Filas() int {
	return t.cantFilas
}

func (t *Tablero) Columnas() int {
	return t.cantColumnas
}

func (t *Tablero) Status(f int, c int) bool {
	if f < 0 || f >= t.cantFilas || c < 0 || c >= t.cantColumnas {
		return false
	}
	return t.celdas[f][c].estado
}

func (t *Tablero) SetStatus(f int, c int, status bool) {
	t.celdas[f][c].estado = status
}

func (t *Tablero) Vecinos(f int, c int) int {
	return t.celdas[f][c].vecinosVivos()
}

func (t *Tablero) newCelda(fila int, col int) *celda {
	return &celda{
		tablero:          t,
		estado:           false,
		fila:             fila,
		columna:          col,
		filaAnterior:     fila - 1,
		filaPosterior:    fila + 1,
		columnaAnterior:  col - 1,
		columnaPosterior: col + 1,
	}
}

func (t *Tablero) Tick() {
	tablero := NewTablero(t.cantFilas, t.cantColumnas)
	for i := 0; i < tablero.cantFilas; i++ {
		for j := 0; j < tablero.cantColumnas; j++ {
			tablero.celdas[i][j] = tablero.newCelda(i, j)
		}
	}
	t.celdas = tablero.celdas
}

func NewTablero(cantFilas int, cantColumnas int) *Tablero {
	tablero := &Tablero{cantFilas: cantFilas, cantColumnas: cantColumnas}
	celdas := make([][]*celda, cantFilas)
	for i := 0; i < cantFilas; i++ {
		columna := make([]*celda, cantColumnas)
		celdas[i] = columna
		for j := 0; j < cantColumnas; j++ {
			celdas[i][j] = tablero.newCelda(i, j)
		}
	}

	tablero.celdas = celdas
	return tablero
}
