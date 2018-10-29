package main
import (
    "fmt"
    "sync"
)

var matrix1[100][100] int
var matrix2[100][100] int
var mat[305][305] int
var n int
var z,k int
var wg sync.WaitGroup

func f2(i int, j int) {

	if i >= z {
		if j >= z {
			mat[i][j] = mat[i][j] + mat[i][j-z]*mat[i-z][j]
		}
	}

	defer wg.Done()
}

func f1(i int) {
    for j := 2*n - 1; j < 3*n - 1; j++ {
        go f2(i,j)
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

    // Fill the matrix
    for i := 0; i < 3*n; i++ {
    	for j := 0; j < 3*n; j++ {
    		mat[i][j] = 0
    	}
    	k++
    }

    k = 0
    for i := 3*n-2; i >= 2*n-1; i-- {
    	for j := k; j < k+n; j++ {
    		mat[i][j] = matrix1[i-2*n+1][j-k]
    	}
    	k++
    }

    k = 0
    for i := 3*n-2; i >= 2*n-1; i-- {
    	for j := k; j < k+n; j++ {
    		mat[j][i] = matrix2[j-k][i-2*n+1]
    	}
    	k++
    }
 	

    for k :=1; k <= 3*n; k++ {
    	z = k
    	wg.Add(n*n)
    	for i := 2*n - 1; i <= 3*n - 2; i++ {
    		go f1(i)
    	}
    	wg.Wait()
    }

    fmt.Println()

 	for i := 0; i < 3*n-1; i++ {
        for j := 0; j < 3*n-1; j++ {
            fmt.Printf(" %d ",mat[i][j])
        }
        fmt.Println("")
    }

    fmt.Println()
    fmt.Println("========== Matrix1 X Matrix2 =============")

    for i := 2*n-1; i < 3*n-1; i++ {
        for j := 2*n-1; j < 3*n-1; j++ {
            fmt.Printf(" %d ",mat[i][j])
        }
        fmt.Println("")
    }

    fmt.Println()


}