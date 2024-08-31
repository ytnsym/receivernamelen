package a

type Example1 struct{}

func (e *Example1) F() {} // OK

type Example2 struct{}

func (ex *Example2) F() {} // OK

type Example3 struct{}

func (ex3 *Example3) F() {} // want `receiver variable names must be one or two characters in length`

type Example4 struct{}

func (example *Example4) F() {} // want `receiver variable names must be one or two characters in length`

func F() {} // OK
