// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func main() {
  nums := []int{30,20,50,40,10}
  quickSort(nums,0,len(nums)-1)
  fmt.Println(nums)
}

func partition(arr []int, low , high int)int{
    i, j := low+1, high
    pivot := arr[low]
    
    for i <= j{
		fmt.Println("i is ==> ", i, " j is ==>> ", j)
		fmt.Println("arr of i ==> ", arr[i], "arr of j ==>> ", arr[j])
        for i <= high && arr[i] <= pivot{
            i++
        }
        
        for j >= low && arr[j] > pivot{
            j--
        }
        
        if arr[i] < arr[j]{
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    
	arr[low], arr[j] = arr[j], arr[low]

    // arr[i], arr[j] = arr[low], pivot
    
    return j
}

func quickSort(nums []int, low, high int){
    if low < high{
        pivot := partition(nums, low, high)
        quickSort(nums,low, pivot-1)
        quickSort(nums, pivot+1, high)
    }
}