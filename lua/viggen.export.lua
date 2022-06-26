function exportDataPanelScreen()
	local raw = ExportScript.Tools.split(list_indication(2), "%c")
	local dataTxt = raw[9] .. raw[12] .. raw[15] .. raw[18] .. raw[21] .. raw[24]
	ExportScript.Tools.SendData(900, string.format("%6d", dataTxt))
end

--[[
	C14,3510,1 to the the data input wheel
]--