package domain

type celda struct {
	tablero          *Tablero
	estado           bool
	fila             int
	columna          int
	filaAnterior     int
	filaPosterior    int
	columnaAnterior  int
	columnaPosterior int
}

func (c *celda) vecinosVivos() int {
	vecinosVivos := 0

	vecinosVivos = vecinosVivos + c.contarVecino(c.filaAnterior, c.columnaAnterior)
	vecinosVivos = vecinosVivos + c.contarVecino(c.fila, c.columnaAnterior)
	vecinosVivos = vecinosVivos + c.contarVecino(c.filaPosterior, c.columnaAnterior)
	vecinosVivos = vecinosVivos + c.contarVecino(c.filaAnterior, c.columna)
	vecinosVivos = vecinosVivos + c.contarVecino(c.filaPosterior, c.columna)
	vecinosVivos = vecinosVivos + c.contarVecino(c.filaAnterior, c.columnaPosterior)
	vecinosVivos = vecinosVivos + c.contarVecino(c.fila, c.columnaPosterior)
	vecinosVivos = vecinosVivos + c.contarVecino(c.filaPosterior, c.columnaPosterior)

	return vecinosVivos
}

func (c *celda) contarVecino(fila int, col int) int {
	if c.tablero.Status(fila, col) {
		return 1
	}
	return 0
}

func (c *celda) tick() *celda {
	vecinosVivos := c.vecinosVivos()
	if c.estado {
		c.estado = vecinosVivos == 2 || vecinosVivos == 3
	} else {
		c.estado = vecinosVivos == 3
	}
	return c
}
