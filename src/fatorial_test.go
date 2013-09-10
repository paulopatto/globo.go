package main

import "testing"

func Test_Fatorial(t *testing.T){
  if (Factorial(4) != 24 && Factorial(3) != 6){
    t.Error("func Factorial did not work as expected.")
  } else {
    t.Log("func Factorial test passed.")
  }
}

func Benchmark_FactorialFunction(b *testing.B){
  for i := 0; i < b.N; i++ {
    Factorial(4)
  }
}

func Test_Fatorial2(t *testing.T){
  if (Factorial2(4) != 24 && Factorial2(3) != 6){
    t.Error("func Factorial2 did not work as expected.")
  } else {
    t.Log("func Factorial2 test passed.")
  }
}
