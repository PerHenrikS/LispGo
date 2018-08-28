;Using clojure as a test file for syntax highlight in vscode
(defn somenumber 29)            ;define variable somenumber 

(defun incr (x) (+ x 1))        ;increments by one 

(print (incr somenumber))       ;prints 30

(print (tail (list 1 2 3 4)))   ;prints [2 3 4]

;Had some small problems with two comments after each other and at EOF 
;Not crashing !