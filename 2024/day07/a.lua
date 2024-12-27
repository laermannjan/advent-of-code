-- generator for the cartesian product of ops x ops, with len count
local function generate_combinations(ops, count)
	return coroutine.wrap(function()
		local combo = {}
		local function generate(depth)
			if depth > count then
				coroutine.yield({ table.unpack(combo) })
				return
			end
			for _, op in ipairs(ops) do
				combo[depth] = op
				generate(depth + 1)
			end
		end
		generate(1)
	end)
end

local operators = {
	function(a, b)
		return a + b
	end,
	function(a, b)
		return a * b
	end,
}

local total = 0

for line in io.lines() do
	local res, nums_str = line:match("([^:]+):%s*(.+)")
	res = tonumber(res)
	local nums = {}
	for num in nums_str:gmatch("%d+") do
		table.insert(nums, tonumber(num))
	end

	print(string.format("res=%d, nums=%s, len(nums)=%d", res, table.concat(nums, " "), #nums))

	for op_combo in generate_combinations(operators, #nums - 1) do
		local result = nums[1]
		for i, op in ipairs(op_combo) do
			result = op(result, nums[i + 1])
		end
		print(string.format("  result=%d", result))

		if result == res then
			total = total + res
			print("    matched")
			break
		end
	end
end

io.stderr:write(total .. "\n")
