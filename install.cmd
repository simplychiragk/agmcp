@echo off
setlocal enabledelayedexpansion

set "BIN_DIR=%USERPROFILE%\.agmcp\bin"
if not exist "%BIN_DIR%" (
    mkdir "%BIN_DIR%"
)

echo Downloading agmcp.exe...
curl -fsSL "https://github.com/simplychiragk/agmcp/releases/latest/download/agmcp.exe" -o "%BIN_DIR%\agmcp.exe"
if %ERRORLEVEL% neq 0 (
    echo Error: Failed to download agmcp.exe.
    exit /b %ERRORLEVEL%
)

for /f "tokens=2*" %%A in ('reg query "HKCU\Environment" /v Path 2^>nul') do set "USER_PATH=%%B"

echo !USER_PATH! | findstr /I /C:"%BIN_DIR%" >nul
if %ERRORLEVEL% neq 0 (
    echo Adding %BIN_DIR% to user PATH...
    setx PATH "!USER_PATH!;%BIN_DIR%"
    echo Please restart your terminal to apply the PATH changes.
) else (
    echo agmcp is already in your PATH.
)

echo Installation complete! Try running 'agmcp list' in a new terminal.
