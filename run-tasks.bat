@echo off
setlocal

if "%1"=="" goto help

:run
if "%1"=="help" goto help
if "%1"=="test" goto test
if "%1"=="copy-example" goto copy-example
if "%1"=="paste-example" goto paste-example
echo Invalid target: %1
goto end

:test
go test -v ./... -count=1
goto end

:copy-example
go run examples/copy/copy.go
goto end

:paste-example
go run examples/paste/paste.go
goto end

:help
echo Usage: %0 [target]
echo.
echo    help: shows this help message
echo    test: runs unit tests
echo    copy-example: runs copy example
echo    paste-example: runs paste example
echo.

:end
endlocal
