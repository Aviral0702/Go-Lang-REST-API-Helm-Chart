// Generating Random Numbers in Golang 
package main 

import ( 
	"fmt"
	"math/rand"
	"time"
) 

func main() { 

	// Intn returns, as an 
	// int, a non-negative 
	// pseudo-random number in 
	// [0,n) from the default Source. 
	// i.e. simply call Intn to 
	// get the next random integer. 
	fmt.Println(rand.Intn(200)) 
	fmt.Println(rand.Intn(200)) 
	fmt.Println(rand.Intn(200)) 
	fmt.Println() 

	// Float64 returns, as 
	// a float64, a pseudo-random 
	// number in [0.0,1.0) 
	// from the default Source. 
	fmt.Println(rand.Float64()) 

	// By default, it uses the value 1. 
	fmt.Println((rand.Float64() * 8) + 7) 
	fmt.Println((rand.Float64() * 8) + 7) 
	fmt.Println() 

	// Seeding - Go provides a method, 
	// Seed(see int64), that allows you 
	// to initialize this default sequence. 
	
	// Implementation is slow 
	// to make it faster 
	// rand.Seed(time.Now().UnixNano()) 
	// is added. Seed is the current time, 
	// converted to int64 by UnixNano. 
	// Gives constantly changing numbers 
	x1 := rand.NewSource(time.Now().UnixNano()) 
	y1 := rand.New(x1) 
	
	fmt.Println(y1.Intn(200)) 
	fmt.Println(y1.Intn(200)) 
	fmt.Println() 

	x2 := rand.NewSource(55) 
	y2 := rand.New(x2) 
	fmt.Println(y2.Intn(200)) 
	fmt.Println(y2.Intn(200)) 
	fmt.Println() 
	
	x3 := rand.NewSource(5) 
	y3 := rand.New(x3) 
	fmt.Println(y3.Intn(200)) 
	fmt.Println(y3.Intn(200)) 
} 
