@echo off
title Sonya Gopher Server
set GOPATH=%~dp0
:A
C:\Go\bin\go.exe run main.go
goto A