# POC-STRATEGY-GO

### TO REMEMBER
> Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.

> The Strategy pattern suggests that you take a class that does something specific in a lot of different ways and extract all of these algorithms into separate classes called strategies.

> The original class, called context, must have a field for storing a reference to one of the strategies. The context delegates the work to a linked strategy object instead of executing it on its own.

> The context isn’t responsible for selecting an appropriate algorithm for the job. Instead, the client passes the desired strategy to the context. In fact, the context doesn’t know much about strategies. It works with all strategies through the same generic interface, which only exposes a single method for triggering the algorithm encapsulated within the selected strategy. 

> This way the context becomes independent of concrete strategies, so you can add new algorithms or modify existing ones without changing the code of the context or other strategies. 

### POC DESCRIPTION
> It is very simple, based on a code received the app must return one phrase

Example:

> Codes: [GM, GA, GE]

> If code is GM the return must be "Good Morning!"
> 
> If code is GA the return must be "Good Afternoon!"
> 
> If code is GE the return must be "Good Evening!"
> 
> If code is not found the return must be "greeting for code [code_here] not found"