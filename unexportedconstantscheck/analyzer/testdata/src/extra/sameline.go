package main

const fooSameLine, _nameSameLine = "bar", "John" // want `unexported constant "fooSameLine" should start with "_"`

const unexportedConstant, unexportedConstant2 = "one", "two" // want `unexported constant "unexportedConstant" should start with "_"`
