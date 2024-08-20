function isprime(number)
    for i = 2, number//2 + 1 do
      if number % i == 0 then return false end
    end
    return true
end
number = arg[1] or 17
print(number .. ' isprime ' .. tostring(isprime(number)))