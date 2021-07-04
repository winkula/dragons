
<#
This is just a helper script to easily generate new puzzles of different sizes and difficulties.
#>

$scriptPath = Join-Path $PSScriptRoot Generate-Puzzles.ps1

& $scriptPath -Difficulties easy   -Size  8 -Count 6

& $scriptPath -Difficulties easy   -Size 10 -Count 2

& $scriptPath -Difficulties easy   -Size 12 -Count 2
& $scriptPath -Difficulties medium -Size 12 -Count 2
& $scriptPath -Difficulties hard   -Size 12 -Count 2
