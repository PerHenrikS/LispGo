## LispGo - A lisp implementation in golang
#### A semi-generic lisp implementation in golang.

### Goal 

The purpose of this project is to better understand how languages work. I have been playing around with programming language
implementation and wanted a somewhat easy language to implement. The ultimate goal for this project is to have a language 
complete enough that it is possible to implement a simple http server and host my own site on it.

### Design choices

Some operation and function names are inspired from clojure, elixir and other languages. 
For example I don't like the lisp naming conventions of car (Contents of the Address part of Register number)
, cdr (Contents of the Decrement part of Register number). So I call them "head" and "tail", known from Elixir and Haskell, 
which describes the operations more intuitively. 

### Todo:
- [ ] implement comments.
- [ ] implement strings. 
- [ ] a net/http library based on golangs net/http. 
- [ ] better error handling and reporting for easier debugging.
- [ ] standard library. 
