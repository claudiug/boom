a, b = 0, 1
10.times do |n|
  puts b unless a == 0
  a, b = b, a + b
end
