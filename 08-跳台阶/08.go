func jumpFloor2(N int) int {
   if N <= 0 {
      return 0
   }

   if N == 1 || N == 2 {
      return N
   }

   return 2 * jumpFloor2(N-1)
}
