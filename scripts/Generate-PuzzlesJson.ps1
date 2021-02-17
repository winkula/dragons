param (
    [string]$DatabaseFile = "database.txt",
    [string]$TargetFile = ".\website\src\assets\data\grids.json"
)



# ========== Functions ==========

Function Get-LineNumbers { (Get-Content $DatabaseFile | Measure-Object -Line).Lines }

Function Extract-Code($str) { Extract-Group -Str $str -Pattern "Code: (?<code>[dfx_,]*)" -Group "code" }
Function Extract-Difficulty($str) { Extract-Group -Str $str -Pattern "Difficulty:\s*(?<difficulty>[a-z]*)" -Group "difficulty" }
Function Extract-Group($Str, $Pattern, $Group) {
    $result = [regex]::Matches($str, $pattern)
    return $result[$result.Count - 1].Groups[$group].Value
}

Function Test-SolutionEntry($code) { Test-DbEntry -Solution $code }
Function Test-PuzzleEntry($code) { Test-DbEntry -Puzzle $code }

Function Write-DbEntry($N, $Size, $Difficulty, $Solution, $Puzzle) { Add-Content -Path $DatabaseFile -Value "n<$N> x<$Size> d<$Difficulty> s<$Solution> p<$Puzzle>" -Force | Out-Null }

Function Get-DbEntry($N, $Size, $Difficulty, $Solution, $Puzzle) {
    Function Fallbacked($Value) { if ($Value) { $Value } else { ".*" } }
    $p = "n<$(Fallbacked($N))> x<$(Fallbacked($Size))> d<$(Fallbacked($Difficulty))> s<$(Fallbacked($Solution))> p<$(Fallbacked($Puzzle))>"
    return Get-Content -Path $DatabaseFile | Select-String -Pattern $p
}

Function Test-DbEntry($N, $Difficulty, $Solution, $Puzzle) {
    (Get-DbEntry -N $N -Difficulty $Difficulty -Solution $Solution -Puzzle $Puzzle).Length -ge 1
}

Function Parse-DbEntry($Line) {
    @{
        Number     = Parse-DbField -Line $Line -Name "n"
        Size       = Parse-DbField -Line $Line -Name "x"
        Difficulty = Parse-DbField -Line $Line -Name "d"
        Solution   = Parse-DbField -Line $Line -Name "s"
        Puzzle     = Parse-DbField -Line $Line -Name "p"
    }
}
Function Parse-DbField($Line, $Name) { Extract-Group -Str $Line -Pattern "$Name<(?<value>[^>]*)>" -Group "value" }

Function Dragons-GenerateSolution($Difficulty, $Size) { Dragons-Run "generate -duration $($DurationSolution)s -difficulty $Difficulty -size $Size -solution" }
Function Dragons-GeneratePuzzle($Difficulty, $Solution) { Dragons-Run "generate -duration $($DurationPuzzle)s -difficulty $Difficulty $Solution" }
Function Dragons-Render($Filename, $Code) { Dragons-Run "render -filename $Filename $Code" }
Function Dragons-Run($Command) {
    $Command = "$DragonsExe $Command"
    Write-Host $Command -ForegroundColor DarkGray
    return Invoke-Expression "$Command"
}



# ========== Main script ==========

if (-not (Test-Path $DatabaseFile)) {
    Write-Error "Database file not existing"
}

$all = @{
    easy   = @()
    medium = @()
    hard   = @()
}

Get-Content $DatabaseFile | % {
    $entry = Parse-DbEntry $_


    $all[$entry.Difficulty] += @{
        size     = [int] $entry.Size
        puzzle   = $entry.Puzzle
        solution = $entry.Solution
    }
}

ConvertTo-Json $all | Set-Content -Path $TargetFile
