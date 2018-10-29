package main
import (
    "fmt"
)

var matrix1[100][100] int
var matrix2[100][100] int
var mat[100][100] int
var n int
 
func main(){
    fmt.Print("Enter number of ns: ")
    fmt.Scanln(&n)
    fmt.Print("Enter number of ns: ")
    fmt.Scanln(&n)
 
    fmt.Println()
    fmt.Println("========== Matrix1 =============")
    fmt.Println()
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf("Enter the element for Matrix1 %d %d :",i+1,j+1)
            fmt.Scanln(&matrix1[i][j])
        }
    }
 
    fmt.Println()
    fmt.Println("========== Matrix2 =============")
    fmt.Println()
 
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf("Enter the element for Matrix2 %d %d :",i+1,j+1)
            fmt.Scanln(&matrix2[i][j])
        }
    }

    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                mat[i][j]+=matrix1[i][k]*matrix2[k][j]
            }
        }
    }


    fmt.Println()
    fmt.Println("========== Matrix1 X Matrix2 =============")

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf(" %d ",mat[i][j])
        }
        fmt.Println("")
    }

    fmt.Println()

}