package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"gol-tdd/internal/core/domain"
)

var _ = Describe("Tablero", func() {
	var tablero *domain.Tablero
	BeforeEach(func() {
		tablero = domain.NewTablero(3, 5)
	})
	Describe("Crear Tablero rectangular", func() {
		Context("Tamaño 3x5", func() {
			It("Deberia tener el ancho correcto", func() {
				gomega.Expect(tablero.Filas()).To(gomega.Equal(3))
			})
			It("Deberia tener el largo correcto", func() {
				gomega.Expect(tablero.Columnas()).To(gomega.Equal(5))
			})
			It("Deberia poder consultar el estado de una celda", func() {
				gomega.Expect(tablero.Status(1, 3)).To(gomega.BeFalse())
			})
			It("Deberia poder consultar el estado de una celda despues de cambiarlo", func() {
				tablero.SetStatus(1, 3, true)
				gomega.Expect(tablero.Status(1, 3)).To(gomega.BeTrue())
			})
		})
	})

})

var _ = Describe("Reglas", func() {
	var tablero *domain.Tablero
	BeforeEach(func() {
		tablero = domain.NewTablero(10, 10)
		tablero.SetStatus(5, 7, true)
		tablero.SetStatus(6, 6, true)
		tablero.SetStatus(7, 7, true)
	})
	Describe("Crear Tablero rectangular", func() {
		Context("Tamaño 10x10", func() {
			It("Deberia poder consultar cantidad de celdas vivas adyacentes a una detereminada", func() {
				gomega.Expect(tablero.Vecinos(1, 3)).To(gomega.Equal(0))
			})

			It("Deberia poder consultar cantidad de celdas vivas adyacentes a una detereminada", func() {
				gomega.Expect(tablero.Vecinos(6, 7)).To(gomega.Equal(3))
			})
			It("Celda inicial deberia estar muerta", func() {
				gomega.Expect(tablero.Status(6, 7)).To(gomega.BeFalse())
			})
			It("Despues de un tick, la celda inicial deberia estar viva", func() {
				tablero.Tick()
				gomega.Expect(tablero.Status(6, 7)).To(gomega.BeFalse())
			})
		})
	})

})
