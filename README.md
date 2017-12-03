# testfake-template
Go templates for creating fakes for testing. This repo mostly exists for my reference. 

# Why Fake?
Go doesn't force you to put all your code in a class. That's good, depending on who you ask. Whether for or against procedural code, it's a fact that mocking OOP objects can forces abstracted implementations where is isn't necessariy a need for abstraction (i.e. no dynamic behavior). In my experience, that it also prevents you from doing good integration tests, as you have to swap out the entire object instead of just the piece that can't be tested easily. 

The sticking point that most devs have about fakes in go is that you have to define your function as a `var`:
`var MyFunc = func(){...}`

And someone is going to break that by inadvertently reassigning the function. Yeah, that would be bad, but that's all the more reason to support a working agreement to maintain good integration tests. I'm sure there's lots more things to discuss on this subject, but I'd just end up trying to make a case for C-style header files and I'm sure no one wants to hear that. 

# faketemplate.go 
This is just an example of my "most robust" setup for faking (meaning it's a lot of stuff you might not need. Use only the parts you need). Replace all the `Func1` and `Func2` stuff before using. 

`SetFake()` is used to reassign a real function to its fake analog. 
`SetReturns()` is used to set the return values that the fake function returns. 

Also in this repo, package `mypkg` shows how to set up your functions to fake. 

# fakeservertemplate.go
This is a setup for a fake HTTP server. Replace all the `ROUTE1, ROUTE2` stuff. Call `InitServer()` to start. Call `AddApiTransaction()` to add another anticipated call to a route prior to running the test. Check the data in the `Api` RouteCall array to verify the results.  

Also note, I used the enum with an array setup here, you could probably switch to a map which would make it easier to add routes. Your preference. 
