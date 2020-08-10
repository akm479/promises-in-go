This is the implementation of promises in golang. 
First a promise constructor is called  promise := PromiseConstructor() which creates a promise object. Then  wait until the Goroutines finishes execution. Finally the promise consumers(Then ,Catch and Finally) are called.

Below are the explaination of some functions

RequestMaker(): Represents a task, in this case makes a get request is made which returns a response body and error if any

SetValues(p *Promise): It sets the values  of the promise object  returned by the RequestMaker


Output Format: 

Inside Then/Catch Method

Resolved Data / Error Occured

"Response from the promise" (In my example a webpage)

Promise is made / Rejected
