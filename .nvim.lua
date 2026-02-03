vim.keymap.set({ "n", "t", "i", "x" }, "<c-cr>", function()
	Snacks.terminal.toggle("aoc run", { auto_close = false })
end, { desc = "AOC run example" })

vim.keymap.set({ "n", "t", "i", "x" }, "<c-s-cr>", function()
	Snacks.terminal.toggle("aoc run input", { auto_close = false })
end, { desc = "AOC run example" })
