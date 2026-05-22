$ErrorActionPreference = 'Stop'

$binDir = Join-Path $env:USERPROFILE ".agmcp\bin"
if (!(Test-Path $binDir)) {
    New-Item -ItemType Directory -Force -Path $binDir | Out-Null
}

$exePath = Join-Path $binDir "agmcp.exe"
$downloadUrl = "https://github.com/simplychiragk/agmcp/releases/latest/download/agmcp.exe"

Write-Host "Downloading agmcp.exe..." -ForegroundColor Cyan
Invoke-WebRequest -Uri $downloadUrl -OutFile $exePath -UseBasicParsing

$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($currentPath -split ';' -notcontains $binDir) {
    Write-Host "Adding $binDir to user PATH..." -ForegroundColor Cyan
    $newPath = $currentPath
    if (!$newPath.EndsWith(';')) {
        $newPath += ';'
    }
    $newPath += $binDir
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "Please restart your terminal to apply the PATH changes." -ForegroundColor Green
} else {
    Write-Host "agmcp is already in your PATH." -ForegroundColor Green
}

Write-Host "Installation complete! Try running 'agmcp list' in a new terminal." -ForegroundColor Green
