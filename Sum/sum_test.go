// sum_test.go
package sum

import "testing"

func TestSum1(t *testing.T) {
    if sum := Sum(1, 2); sum != 3 {
        t.Errorf("sum expected to be 3; got %d", sum)
    }
}
// for _, test := range tests {
//     if sum := Sum(test.values...); sum != test.want {
//         t.Errorf("%s: Sum() = %d, want %d", test.name, sum, test.want)
//     }
// }    
// for _, test := range tests {
//     t.Run(test.name, func(t *testing.T) {
//         assert.Equal(t, Sum(test.values...), test.want)
//     })
// }