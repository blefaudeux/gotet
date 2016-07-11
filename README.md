# gotet
[**Work in progress**] 

# What is it ?
A "Go SDK" to [The Eye Tribe](http://theeyetribe.com/) eye tracker, implementing most (if all) the API calls described [here](http://dev.theeyetribe.com/api/). 
No graphical client for now, yet to come. 

# How to use it ?
You'll need to install *dealer*, a super small socket client handling json requests.
> go get github.com/blefaudeux/dealer/src/dealer

Afterwards, the syntax should be pretty self-explanatory, but something like 
> import "gotet"
>
> client := gotet.Client{}
> client.Connect("localhost", "6555")

should be enough to get you running



