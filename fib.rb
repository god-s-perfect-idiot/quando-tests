def fibonacci(n)
  return [] if n <= 0
  return [0] if n == 1

  fib_sequence = [0, 1]
  
  (2...n).each do |i|
    fib_sequence << fib_sequence[i - 1] + fib_sequence[i - 2]
  end
  
  fib_sequence
end

n = 1000000
fib_sequence = fibonacci(n)