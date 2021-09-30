package calculation
  
var Calculs int

func Do_add(num1 int, num2 int)(int) {
    sum := num1 + num2
    Calculs = sum
    return sum
}
