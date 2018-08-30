## LispGo - A lisp implementation in golang
#### A semi-generic lisp implementation in golang.

### Goal 

The purpose of this project is to better understand how languages work. I have been playing around with programming language
implementation and wanted a somewhat easy language to implement.  

### Design choices

Some operation and function names are inspired from clojure, elixir and other languages. 
For example I don't like the lisp naming conventions of car (Contents of the Address part of Register number)
, cdr (Contents of the Decrement part of Register number). So I call them "head" and "tail", known from Elixir and Haskell, 
which describes the operations more intuitively. 

I'm still not convinced that the choice of making the token structure was worth it. To use an input - split - replace approach to tokenize might acutally be more readabe in the case of lisp. 

The test.clj is a real working program written in the language. It has the clojure extension for syntax highlighting in 
my text editor.


### Todo:
- [x] implement comments.
- [ ] implement strings.  
- [ ] better error handling and reporting for easier debugging.
- [ ] cleanup repl print. 
- [ ] macros.
- [ ] standard library. 
- [ ] testing for stability. 
- [ ] do blocks.
