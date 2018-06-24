digits = {"-**----*--***--***---*---****--**--****--**---**--",
          "*--*--**-----*----*-*--*-*----*-------*-*--*-*--*-",
          "*--*---*---**---**--****-***--***----*---**---***-",
          "*--*---*--*-------*----*----*-*--*--*---*--*----*-",
          "-**---***-****-***-----*-***---**---*----**---**--",
          "--------------------------------------------------"}
w, h = 5, 6
for line in io.lines(arg[1]) do
  local r = {}
  for i = 1, h do r[i] = "" end
  for i in line:gmatch("%d") do
    for j = 1, h do
      local pos = (string.byte(i)-string.byte('0'))*w + 1
      r[j] = r[j] .. digits[j]:sub(pos, pos+w-1)
    end
  end
  print(table.concat(r, "\n"))
end
