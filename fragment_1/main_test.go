package main

import (
    "testing"
)

func TestAbs(t *testing.T) {
    tests := []struct {
        name  string
        value float64
        want  float64
    }{
        {name: "simple negative value", value: -10, want: 10},
        {name: "simple positive value", value: 10, want: 10},
        {name: "zero", value: -0, want: 0},
        {name: "small value", value: -0.000000001, want: 0.000000001},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Abs(tt.value); got != tt.want {
                t.Errorf("Abs() = %v, want %v", got, tt.want)
            }
        })
    }
}