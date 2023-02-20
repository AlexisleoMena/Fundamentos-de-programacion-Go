package main

import "fmt"

type Tarea struct {
	descripcion    string
	estaCompletada bool //No se pueden establecer valores por defecto
}

func (t *Tarea) GetTarea() { //Recordar que está mal visto mezclar ingles y español
	fmt.Printf("Completada: %v, Descripcion: %s \n", t.estaCompletada, t.descripcion)
}
func (t *Tarea) SetEstaCompletada(valor bool) {
	t.estaCompletada = valor
}

type ListaDeTareas struct {
	tareas []*Tarea //Lista de referencias
}

func (t *ListaDeTareas) AddTarea(tarea *Tarea) {
	t.tareas = append(t.tareas, tarea)
}
func (t *ListaDeTareas) DeleteTarea(tarea *Tarea) {
	for i := 0; i < len(t.tareas); i++ {
		if t.tareas[i].descripcion == tarea.descripcion {
			t.tareas = append(t.tareas[:i], t.tareas[i+1:]...)
		}
	}
}
func (t *ListaDeTareas) GetListaDeTareas() {
	for _, tarea := range t.tareas {
		tarea.GetTarea()
	}
}

func main() {
	t1 := Tarea{"Comprar vino", false}
	t2 := Tarea{"Comprar sal", false}
	t3 := Tarea{"Comprar limon", false}
	t3.SetEstaCompletada(true)

	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	fmt.Println("t3:", t3)

	fmt.Println()

	l1 := ListaDeTareas{}
	l1.AddTarea(&t1)
	l1.AddTarea(&t2)
	l1.AddTarea(&t3)

	l1.GetListaDeTareas()
	t1.SetEstaCompletada(true)
	l1.DeleteTarea(&t2)

	fmt.Println()

	l1.GetListaDeTareas()

}
