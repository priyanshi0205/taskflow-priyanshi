$ErrorActionPreference = "Stop"

$env:GOCACHE = Join-Path $PSScriptRoot ".gocache"
$env:GOMODCACHE = Join-Path $PSScriptRoot ".gomodcache"
$goExe = Join-Path (Split-Path $PSScriptRoot -Parent) "_local_go_sdk\bin\go.exe"

& $goExe mod tidy
& $goExe build .

Write-Host "Dependencies installed and build successful."
