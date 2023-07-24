# A Comprehensive Guide to Data Structures in Go
You've probably heard of data structures and have used them in other programming languages, but do you know how to use them in Go?
As one of the fastest-growing programming languages in the industry, it is important for devs to understand how to utilize this vital feature in order to create scalable, reliable applications.

1. Arrays
An array is a collection of data of a specific type. It stores multiple values in a single variable where each element has an index to reference itself. Arrays come in handy when you need to keep more than one thing in a single location, like a list of people who attended an event or the age of students in a class.
  1.1 Creating and Array
  To create an array, we need to define its name, length, and type of values we will be storing:
  ```go
  var studentsAge [10] int
  ```
  In this code blog, we created an array named studentsAge, which can store a maximum of ten int values.
  1.2 Creating an Array from Literals
  You can create an array from literals, meaning you're assigning values to them at the point of creating.
  Let’s see how it can be used:
  ```go
  // Creating an array and assigning values later
  var studentsAge [10]int
  studentsAge = [10]int{1,2,3,4,5,6,7,8,9,10}

  // Creating an assigning values to an array
  var studentsAge = [10]int{1,2,3,4,5,6,7,8,9,10}

  // Creating and assigning values to an array without var keyword
  studentsAge := [10]int{1,2,3,4,5,6,7,8,9,10}
  ```
  1.3 Creating an Array of an Array
  You can create an array where every element is an individual array (nested array), like so:
  ```go
  // Creating a nested array
  nestedArray := \[3\] [5]int{
    {1,2,3,4,5},
    {6,7,8,9,10},
    {11,12,13,14,15},
  }
  ```
  1.4 Accessing the Values in an array
  Each element in array has an index that you can use to access and modify its value. The index of an array is always an integer and starts counting from zero:
  ```go
  // Creating an array of integers
  studentsAge := [10]int{1,2,3,4,5,6,7,8,9,10}

  // Accessing array values with their indexes
  fmt.printLn(studentsAge[0]) //1
  fmt.printLn(studentsAge[7]) //8
  fmt.printLn(studentsAge[4]) //5

  // Using a for loop to access an array
  for i := 0; i < 10; i++{
    fmt.Println(studentsAge[i])
  }

  // Using range to access an array
  for index, value := range studentsAge {
    fmt.Prinln(index, value) //it will show index and its own value
  }
  ```
  1.5 Modifying the Values in an Array
  Arrays are a mutable data structures, so it is possible to modify their values after creation.
  ```go
  // Creating an Array of Integeres
  studentsAge := [10]int{1,2,3,4,5,6,7,8,9,10}

  // Modifying Array values with their indexes
  studentsAge[7] = 22
  studentsAge[4] = 11

  fmt.Prinln(studentsAge)
  ```

  1.6 Getting the Length of an Array
  Go provides `len` function that you can use to get the Length of an Array.
  ```go
  //Creating an Getting the Length of an Array with a length of 10
  var arrayOfIntegers [10]int
  fmt.Println(len(arrayOfIntegers)) //10
  ```
  Note that it is impossible to change the length of an array because it becomes part of the type during creation.

2. Slices
Like arrays, slices allow you to store multiple values of the same type in a sinble variable and access them with indexes. The main difference between slices and arrays is that slices have dynamic lengths, while arrays are fixed.
  2.1 Creating a Slice
  To create a slice, we need to define its name and the type of values we will be storing:
  ```go
  var sliceOfIntegers []int
  ```
  We created a slice named `sliceOfIntegers``, which stores `int`` values.
  2.2 Creating a Slice from an Array
  In its original form, a slice is an extracted portion of an array. To create a slice from an array, we need to provide Go with the part to extract.
  ```go
  // Creating an Array of integers
  studentsAge := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  // Creating Slices from Arrays
  fiveStudents := studentsAge[0:5]
  fmt.Println(fiveStudents) // [1,2,3,4,5]
  fmt.Println(studentsAge[:2])
  fmt.Println(studentsAge[5:])
  fmt.Println(studentsAge[:])
  ```
  It is also possible to create slices form other slices withn the same format as arrays:
  ```go

  ```
  2.3 Creating a Slice with `make`

  LAST LINK: https://blog.logrocket.com/comprehensive-guide-data-structures-go/

## GO Samples [Link](https://go.dev/tour/welcome/1)
  ```go
  package main

  import "fmt"

  func main() {
    var studentsAge [10]int
  	// var nestedArray [4][9]int
    // nestedArray = [4][9]int{
    //	{1,2,3,4,5},
    //	{6,7,8,9,10},
    //	{11,12,13,14,15},
    //}
    
    studentsAge = [10]int{5,6,14,12,12}
    //var i int
    //fmt.Println(studentsAge[3])
    for i := 0; i < 10; i++ {
      fmt.Println(studentsAge[i] * 7)
    }
    
    studentsAge[7] = 99
    studentsAge[2] = 32
    
    for index, value := range studentsAge{
      fmt.Println(index, value)
    }
    
    threeStudents := studentsAge[:] //get all students
    fmt.Println(threeStudents[3:6]) //di slice
  }

  ```