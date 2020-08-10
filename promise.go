package main

import ("fmt"
		"io/ioutil"
		"net/http"
		"sync")


// Go does not have prototypes but interfaces. This custom Interface has three methods
// Catch, Then and Finally 

type PromiseInterface interface{
	Catch(func(error))
	Then(func(string),func(error))
	Finally(func(int))
}

// A Promise can throw an error or will be having some data and will be in one of the three states


type Promise struct{
	errors error
	data string
	state int
}

// Three states possible for a promise 
const (
	rejected = -1
	pending = 0
	fulfilled = 1
)

// Waitgroup which ensures that goroutines is completed before the completetion of the program
var wg  sync.WaitGroup




func SetValues (p *Promise){
	p.data,p.errors =   RequestMaker()
	p.state = pending
}

// Returns a  promise object  
func PromiseConstructor() *Promise{
	p:=&Promise{}
	wg.Add(1)
	go SetValues(p)

	return p
}

// Example task for which promise is made.  Can be replaced with any task

func RequestMaker() (string, error){

	resp, err := http.Get("http://example.com/")
	body, err := ioutil.ReadAll(resp.Body)
	wg.Done()

	return string(body), err

}


// Promise object calls this method. It takes two functions as arguments (onFullfiled and onRejected)

func (p *Promise) Then (onFullfiled func(string),onRejected func(error)){
	fmt.Println("Inside then method")
	if p.errors==nil {
		onFullfiled(p.data)
		p.state=fulfilled
		return
	}
	fmt.Println("Error occured")
	onRejected(p.errors)
	p.state=rejected
}

// Promise object calls Catch method. It  takes   one function as argument (onRejected)
func (p *Promise) Catch (onRejected func(error)){
	
	if p.errors!=nil {
		fmt.Println("Inside catch method")
		fmt.Println("Error occured")
		onRejected(p.errors)
		p.state=rejected
	}
	
}

// Promise object calls Finally method . It takes one function as argument (onFinally)
func (p *Promise) Finally (onFinally func(int)){
	
	onFinally(p.state)
	
}


func main(){

	promise := PromiseConstructor()
	wg.Wait()
	var promiseConsumer PromiseInterface = promise
	
	

	promiseConsumer.Catch(
		func(reject error) {	 
			fmt.Println(reject)
		})

	promiseConsumer.Then(
		func(resolve string) {
			fmt.Println("Resolved data:")
			fmt.Println(resolve)
		},
		func(reject error) {
			fmt.Println(reject)
		})

	promiseConsumer.Finally(
		func(state int) {
			if state == -1 {
				fmt.Println("Promise is Rejected")
			} 
			if state == 0 {
				fmt.Println("Promise is pending")
			} 
			if state == 1 {
				fmt.Println("Promise is made ")
			}
		})

 }