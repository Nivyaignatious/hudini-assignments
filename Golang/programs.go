// count monkeys qn 1
package kata

func monkeyCount(n int) []int {
  var res []int
   for i:= 1; i<=n; i++{
     res = append(res,i)
   }
  return res
}

// print negative qn2
package kata

func MakeNegative(x int) int {
  if (x==0 || x<0){
    return x
  }else{
    return -x
  }
}

// multi qn3
func FindMultiples(integer, limit int) []int {
	var multi []int
	for i:=integer; i<=limit; i++{
	  if (i%integer==0){
		multi= append(multi,i)
	  }
	  }
	  return multi
  
  }

 // qn4
 func CountBy(x, n int) []int {
	var mult []int
	for i:=1; i<=n; i++{
	  a:= x*i
	  mult= append(mult,a)
	}
	return mult
  }
 
//qn5
import "math"
func PowersOfTwo(n int) []uint64 {
  var powers []uint64
  for i:=0 ; i<=n ;i++{
    power:= math.Pow(2,float64(i))
    powers= append(powers,uint64(power))
  }
  return powers
}

//qn6
func Points(games []string) int {
	score:=0
	for _, a:= range (games){
	  if a[0]> a[2]{
		score+=3
	  }
	  if a[0]== a[2]{
		score+=1
	  }
	}
	return score
  }

//qn7
func GetSum(a, b int) int {
	sum:=0
	if a==b {
	  return a
	}else if a<b {
	  for i:=a;i<=b;i++{
		sum+=i
	  }
	  return sum
	}else{
	  for i:=b;i<=a;i++{
		sum+=i
	}
	  return sum
  }
	}
	
//qn8
import "strings"
func FindShort(s string) int {
  str:= strings.Fields(s)
  length:=len(s)
  for i:=0;i<len(str);i++{
    if length>len(str[i]){
      length=len(str[i])
    }
  }
  return length
}

//qn9
func LeastLarger(a []int, i int) int {
	smallest:=1000
	index:= -1
	for j:=0; j<len(a); j++{
	  if a[j]>a[i]{
		if a[j]<smallest{
		  smallest = a[j]
		  index= j
		}
	  }
	}
   return index
   }

//qn 11
func SortMyString(s string) string {
	even:=""
	odd:=""
	for i:=0 ;i<len(s); i++{
	  if i==0{
		even+= string(s[i])
	  }else if i%2 ==0{
		even += string(s[i])
	  }else {
	  odd += string(s[i])
	}
	  }
	return even+" "+odd
  }
  
 //qn12
 func OddCount(n int) int{
	return n/2
  }

 //qn13
 func SquareSum(numbers []int) int {
    sums:=0
  for i:=0; i<len(numbers);i++{
    sum:=numbers[i]*numbers[i]
    sums+=sum 
  }
  return sums
}
 