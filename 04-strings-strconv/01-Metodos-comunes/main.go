package main

import "strings"

func main() {

	strings.ToUpper("Alexis")                         //ALEXIS
	strings.ToLower("Mena")                           //mena
	strings.Count("Manzana", "a")                     //3
	strings.Contains("El alfajor de chocolate", "de") //true
	strings.Split("hola a todos!", " ")               //["hola", "a", "todos!"]
	strings.Title("alexis leonardo mena")             //Alexis Leonardo Mena
	strings.Trim("¡Hola! y ¡Adios!", "!¡")            //"Hola! y ¡Adios"
	strings.TrimSpace("           Hola  ")            //"Hola"
	strings.EqualFold("aLEXIS", "AlexiS")             //true
	strings.Index("Hola a todos", "tod")              //7
	strings.IndexAny("Hola a todos", "d")             //9

	strings.Fields("   Pera Manzana     Naranja")             //["Pera", "Manzana", "Naranja"]
	strings.Join([]string{"Pera", "Manzana", "Banana"}, ", ") //"Pera, Manzana, Banana"

	strings.Replace("hola, hola, hola amigos!", "hola", "adios", 1)  //adios, hola, hola amigos!"
	strings.Replace("hola, hola, hola amigos!", "hola", "adios", 2)  //"adios, adios, hola amigos!"
	strings.Replace("hola, hola, hola amigos!", "hola", "adios", -1) //"adios, adios, adios, amigos!"

	strings.Compare("A", "B") //-1
	strings.Compare("A", "A") //0
	strings.Compare("B", "A") //1

	strings.ContainsAny("ABCDJKL", "E")  //false
	strings.ContainsAny("ABCDJKL", "EJ") //true

	strings.HasPrefix("Hola a todos", "Ho") //true
	strings.HasSuffix("Hola a todos", "os") //true

	strings.Repeat("na", 5) //nanananana

}
