How to install

`go get github.com/fahmiaz411/go-zipper`

This is an example program

```go
    import (
        zipper "github.com/fahmiaz411/go-zipper"
        "fmt"
    )

   func main() {

	arr1 := []interface{}{
		"a",2,"c",
	}

	arr2 := []interface{}{
		"d",5,"f",
	}

	for _, data := range zipper.Zip(arr1,arr2){

		a1 := data[0]
		a2 := data[1]

		fmt.Println(a1,a2)
	}
}
```

result:

```
a d
2 5
c f
```
