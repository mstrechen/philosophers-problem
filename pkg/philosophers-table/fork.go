package philosophers_table


type Fork struct {
	name string
	isFree chan bool
}