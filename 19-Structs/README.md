# Structs
São coleções tipadas de campos, que são úteis para agrupar dados juntos para formar registros. Elas permitem que você organize dados relacionados de maneira estruturada.

## Estrutura de uma Struct

O tipo de struct chamado `person` tem dois campos: `name` (nome) e `age` (idade).

```go
type person struct {
    name string
    age  int
}
```

## Função Construtora
A função newPerson constrói uma nova struct do tipo person com o nome fornecido:

```go
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}
```

Go é uma linguagem com coleta de lixo (garbage collected), o que significa que você pode retornar um ponteiro para uma variável local de forma segura. Essa variável só será liberada pelo coletor de lixo quando não houver referências ativas a ela.

## Uso da Struct
No main, a criação de uma nova struct é feita com a seguinte sintaxe:

```go
fmt.Println(person{"Bob", 20})
Você pode nomear os campos ao inicializar uma struct:

```go
fmt.Println(person{name: "Alice", age: 30})
Se um campo for omitido, ele será inicializado com seu valor zero:

```go
fmt.Println(person{name: "Fred"})
```
O prefixo & retorna um ponteiro para a struct:

```go
fmt.Println(&person{name: "Ann", age: 40})
```
É idiomático encapsular a criação de novas structs em funções construtoras:

```go
fmt.Println(newPerson("Jon"))
```
## Acesso aos Campos da Struct
Os campos da struct são acessados usando o operador ponto (.):

```go
s := person{name: "Sean", age: 50}
fmt.Println(s.name)
```
Você também pode usar o operador ponto com ponteiros para structs — os ponteiros são automaticamente desreferenciados:

```go
sp := &s
fmt.Println(sp.age)
```
As structs são mutáveis, o que significa que você pode alterar seus campos:

```go
sp.age = 51
fmt.Println(sp.age)
```
## Structs Anônimas
Se um tipo de struct é usado apenas para um único valor, não precisamos dar um nome a ele. O valor pode ter um tipo de struct anônimo. Essa técnica é comumente usada em testes baseados em tabelas:

```go
dog := struct {
    name   string
    isGood bool
}{
    "Rex",
    true,
}
fmt.Println(dog)
```
### Saída do Código
Quando você executa o programa, a saída é a seguinte:

```bash
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
{Rex true}
```

## Executar o código
```bash
go run structs.go
```