param (
    [string[]]$Difficulties = @('easy', 'medium', 'hard'),
    [int]$Size = 8,
    [int]$Count = 3,
    [int]$DurationSolution = 3,
    [int]$DurationPuzzle = ($Size*$Size)/2,
    [string]$DatabaseFile = "database.txt",
    [string]$DragonsExe = ".\dragons.exe"
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
    New-Item -ItemType File -Name $DatabaseFile | Out-Null
}

# Compile binary
Invoke-Expression "go build ./cmd/dragons"

# Generate puzzles
(1..$count) | % {
    $i = $_ - 1
    $difficulty = $Difficulties[$i % $Difficulties.Count]
    Write-Host "Find puzzle for difficulty '$difficulty'..." 

    while ($true) {
        $solution = Dragons-GenerateSolution -Difficulty $difficulty -Size $Size
        $codeSolution = Extract-Code $solution
        Write-Host $codeSolution -ForegroundColor White
        if (Test-SolutionEntry $codeSolution) {
            Write-Host "Solution already exists." -ForegroundColor Yellow
            continue
        }

        
        $puzzle = Dragons-GeneratePuzzle -Difficulty $difficulty -Solution $codeSolution
        $codePuzzle = Extract-Code $puzzle
        $foundDifficulty = Extract-Difficulty $puzzle
        Write-Host $codePuzzle -ForegroundColor White

        if ($foundDifficulty -ne $difficulty) {
            Write-Host "Difficulty of puzzle does not match the requirement..." -ForegroundColor Yellow
            continue
        }

        if (Test-PuzzleEntry $codePuzzle) {
            Write-Host "Puzzle already exists." -ForegroundColor Yellow
            continue
        }
        
        Write-Host "New puzzle found (difficulty $foundDifficulty)..." -ForegroundColor Green

        $n = Get-LineNumbers + 1
        Write-DbEntry -Solution $codeSolution -Puzzle $codePuzzle -N $n -Difficulty $difficulty -Size $Size
            
        break
    }
}

# Render puzzles
Get-Content $DatabaseFile | % {
    $entry = Parse-DbEntry $_
    Dragons-Render -Filename "print/puzzles/$($entry.Size)-$($entry.Difficulty)-$($entry.Number)-solution" -Code $entry.Solution | Out-Null
    Dragons-Render -Filename "print/puzzles/$($entry.Size)-$($entry.Difficulty)-$($entry.Number)-puzzle" -Code $entry.Puzzle | Out-Null
}
