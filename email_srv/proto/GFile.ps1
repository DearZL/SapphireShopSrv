$scriptPath = $MyInvocation.MyCommand.Path
$scriptDirectory = Split-Path -Path $scriptPath -Parent
Set-Location $scriptDirectory
Remove-Item srv/*.go
protoc  --go_out=./srv --go_opt=paths=source_relative --go-grpc_out=./srv --go-grpc_opt=paths=source_relative *.proto
Write-Output Finish!
exit