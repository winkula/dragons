[CmdletBinding(SupportsPaging = $true)]
param(
	[string] $Func = 'BenchmarkValidate$'
)

clear

& go test `
-cpuprofile $env:TEMP\cpu.prof `
-bench "^$Func" `
'github.com/winkula/dragons/pkg/model'

#& go tool pprof -http :6060 $env:TEMP\cpu.prof
#& go tool pprof -top -hide runtime $env:TEMP\cpu.prof
#& go tool pprof -list Validate $env:TEMP\cpu.prof

& go tool pprof -cum $env:TEMP\cpu.prof
