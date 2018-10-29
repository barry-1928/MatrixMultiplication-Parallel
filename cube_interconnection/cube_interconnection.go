package main
import (
    "fmt"
    "sync"
)

var matrix1[100][100] int
var matrix2[100][100] int
var matA[100][100][100] int
var matB[100][100][100] int
var matC[100][100] int
var n int
var z int
var wg sync.WaitGroup
 
func f3(i int, j int, k int) {
    matA[k][i][j] = matrix1[i][k]
    matB[k][j][i] = matrix2[k][i]
    defer wg.Done()
}

func f2(i int, k int) {
    for j:=0; j < n; j++ {
        go f3(i,j,k)
    }
}

func f1(k int) {
    for i := 0; i < n; i++ {
            go f2(i,k)
        }
}

func g3(i int, j int, k int) {
    matC[i][j] += matA[k][i][j] * matB[k][i][j]
    defer wg.Done()
}

func g2(i int, k int) {
    for j:=0; j < n; j++ {
        go g3(i,j,k)
    }
}

func g1(k int) {
    for i := 0; i < n; i++ {
            go g2(i,k)
        }
}

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

    wg.Add(n*n*n)
    
    for k := 0; k < n; k++ {
        go f1(k)
    }

    wg.Wait()
    wg.Add(n*n*n)

    for k := 0; k < n; k++ {
        go g1(k)
    }

    wg.Wait()

    
    

    fmt.Println()
    fmt.Println("========== Matrix1 X Matrix2 =============")
    fmt.Println()
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf("%d ",matC[i][j])
        }
        fmt.Printf("\n")
    }

}

/* export GOPATH=$HOME/work */