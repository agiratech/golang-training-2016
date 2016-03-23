package main
import "fmt"

var great int

func main(){
    //----------variable Declaration----------------
    var n int
    //var x [10]int
     
    fmt.Println("Enter the size:")
    fmt.Scanf("%d", &n)
    var x = make([]int,n)
    //---------input values
    fmt.Println("Enter the Set of", n, "values:")
    
	    for i :=0; i < n; i++ {
	    
	     fmt.Scanf("%d",&x[i])
	    } 
    //Check that values which one is greatest 
	        great = x[0]
		    for i :=0; i < len(x); i++ {
		
	           if great < x[i]{
	           	great = x[i]
	           }

		}
        
   
    fmt.Println("Greatest Number:",great)
}
	