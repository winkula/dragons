param (
    [string[]]$Difficulties = @('easy'),#@('easy', 'medium', 'hard'),
    [int]$Count = 1,
    [int]$Duration = 1,
    [int]$Size = 8,
    [string]$Filename = "database"
)

$dragons = "go run ./cmd/dragons"
$file = "$($Filename).txt"

if (-not (Test-Path $file)) {
    New-Item -ItemType File -Name $file | Out-Null
}

Function Get-LineNumbers { (Get-Content $file | Measure-Object -Line).Lines }

Function Extract-Code($str) { Extract-Group -Str $str -Pattern "Code: (?<code>[dfx_,]*)" -Group "code" }
Function Extract-Difficulty($str) { Extract-Group -Str $str -Pattern "Difficulty:\s*(?<difficulty>[a-z]*)" -Group "difficulty" }

Function Extract-Group($str, $pattern, $group) {
    $result = [regex]::Matches($str, $pattern)
    return $result[$result.Count - 1].Groups[$group].Value
}

Function Test-SolutionEntry($code) { Test-DbEntry -Solution $code }
Function Test-PuzzleEntry($code) { Test-DbEntry -Puzzle $code }

Function Get-DbEntry($N, $Difficulty, $Solution, $Puzzle) {
    Function Fallbacked($Value) { if ($Value) { $Value } else { ".*" } }
    $p = "n<$(Fallbacked($N))> d<$(Fallbacked($Difficulty))> s<$(Fallbacked($Solution))> p<$(Fallbacked($Puzzle))>"
    return Get-Content -Path $file | Select-String -Pattern $p
}

Function Test-DbEntry($N, $Difficulty, $Solution, $Puzzle) {
    (Get-DbEntry -N $N -Difficulty $Difficulty -Solution $Solution -Puzzle $Puzzle).Length -ge 1
}

Function Write-Entry($N, $Difficulty, $Solution, $Puzzle) { Add-Content -Path $file -Value "n<$N> d<$Difficulty> s<$Solution> p<$Puzzle>" -Force | Out-Null }

Function Dragons-GenerateSolution { Invoke-Expression "$dragons generate -duration $($duration)s -difficulty $difficulty -size $size -solution" }
Function Dragons-GeneratePuzzle($solution) { Invoke-Expression "$dragons generate -duration $($duration)s $solution" }
Function Dragons-Render($filename, $code) { Invoke-Expression "$dragons render -filename $filename $code" }

$difficulties | % {
	$difficulty = $_
	(1..$count) | % {

        $solution = Dragons-GenerateSolution
        $codeSolution = Extract-Code($solution)
        Write-Host $codeSolution -ForegroundColor White
        if (Test-SolutionEntry($codeSolution)) {
            Write-Host "Solution already exists." -ForegroundColor Yellow
            return
        }
        
        $puzzle = Dragons-GeneratePuzzle($codeSolution)
        $codePuzzle = Extract-Code($puzzle)
        $difficulty = Extract-Difficulty($puzzle)
        Write-Host $codePuzzle -ForegroundColor White
        if (Test-PuzzleEntry($codePuzzle)) {
            Write-Host "Puzzle already exists." -ForegroundColor Yellow
            return
        }
        
        Write-Host "New puzzle found (difficulty $difficulty)..." -ForegroundColor Green

        $n = Get-LineNumbers + 1
        Write-Entry -Solution $codeSolution -Puzzle $codePuzzle -N $n -Difficulty $difficulty

        Dragons-Render -Filename "puzzles/$size-$difficulty-$n-puzzle" -Code $codeSolution | Out-Null
        Dragons-Render -Filename "puzzles/$size-$difficulty-$n-solution" -Code $codePuzzle | Out-Null
	}
}
