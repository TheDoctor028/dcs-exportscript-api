function exportDataPanelScreen()
	local raw = ExportScript.Tools.split(list_indication(2), "%c")
	local dataTxt = raw[9] .. raw[12] .. raw[15] .. raw[18] .. raw[21] .. raw[24]
	ExportScript.Tools.SendData(9999, string.format("%6d", dataTxt))
end