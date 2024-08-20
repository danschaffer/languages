# lua

## day 1
```
brew install lua
lua
5.4.4
> print 'hello'
> ='hello'
> = 'hello' .. ' ' .. 'world'
> return 'hello ' .. ' ' .. 'world'
> = true
> # 'hello'
5
> 2 ^ 3
8.0
> = 1899 % 100
99
> not (( true or false) and false)
true
> = true or nothing()
true
> = false or nothing()
stdin:1: attempt to call a nil value (global 'nothing')
```

```
function triple(num) return 3 * num end
```

```
=(function(num) return 3 * num end)(2)
6
```

```
function call_twice(f)
  ff = function(num) 
    return f(f(num))
  end
  return ff
end
```

```
function reverse(s, t)
  if #s < 1 then return t end
  first = string.sub(s, 1, 1)
  rest = string.sub(s, 2, -1)
  return reverse(rest, first .. t)
end

large = string.rep('hello ', 5000)
print(reverse(large, ''))
```

```
function weapons()
  return 'bullwhip', 'revolver'
end
w1 = weapons()
w1,w2 = weapons()
```

```
table = {small=5, medium=7, large=10}
= table.small
```

```
film = 'Raiders'
if film == 'Raiders' then print('Good')
elseif film == 'Temple' then print('Meh')
else print('Huh') end
```

```
for i = 1,5 do print (i) end
```

```
while math.random(100) < 50 do print('Tails; flipping again') end
```

```
function hyp(a,b)
  local a2 = a*a
  local b2 = b*b
  return math.sqrt(a2+b2)
end
```

os.exit(1)
dofile("lib1.lua")
lua -e "print(math.sin(12))"
lua -i -l a.lua -e "x=10"
arg[0] script.lua
arg[1] 
arg[2]

page = [[
<html>
<head><title>title</title></head>
<body>
page body
</body>
</html>
]]