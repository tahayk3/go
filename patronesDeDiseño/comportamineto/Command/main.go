package main

/*
Command es un patrón de diseño de comportamiento que convierte solicitudes u operaciones simples en objetos.
*/

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}
