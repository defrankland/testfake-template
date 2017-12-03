# testfake-template
Go templates for creating fakes for testing. This repo mostly exists for my reference. 

# Why Fake?
Fakes can be used instead of mocks. The idea is to replace a function rather than a whole class. Sometimes you want don't want to remove all of your imported code for writing integration tests around a module, but just the parts that would be impossible to test. Sometimes you don't want to write maintain interfaces and a factory (or other creational pattern) for something that you know will never need that level of abstraction (i.e. no dynamic behavior). Sometimes a problem domain is better solved via procedural code than OOP.  

The sticking point that most devs have about fakes in go is that you have to define your function as a `var`:

`var MyFunc = func(){...}`

Someone might inadvertently reassigning the function -- all the more reason to have good integration tests where everything isn't mocked out one level down!! Personally I prefer having header files to implement fakes as in C, but hey, sometimes you have to get all javascript-y. 

# Files 
The sections below describe the two fake templates in this repo. They're not really templates, they're examples.

## faketemplate.go 
This is just an example of my "most robust" setup for faking (meaning it's a lot of stuff you might not need. Use only the parts you need). Replace all the `Func1` and `Func2` stuff before using. 

`SetFake()` is used to reassign a real function to its fake alternative. 
`SetReturns()` is used to set the return values that the fake function returns. 

Also in this repo, package `mypkg` shows how to set up your functions to fake. 

# fakeservertemplate.go
This is a setup for a fake HTTP server. Replace all the `ROUTE1, ROUTE2` stuff. Call `InitServer()` to start. Call `AddApiTransaction()` to add another anticipated call to a route prior to running the test. Check the data in the `Api` RouteCall array to verify the results.  

Also note, I used the enum with an array setup here, you could probably switch to a map which would make it easier to add routes. Your preference. 
