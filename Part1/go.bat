REM modified version of brainman's batch file at https://groups.google.com/forum/#!topic/golang-nuts/QVPKm7pbhds
 
REM setting goroot to my go installation folder.
set GOROOT=C:\go
 
REM setting gopath to my go playground folder
set GOPATH=C:\Users\wkrause\Documents\gocode\
Set GOBIN=%GOPATH%\bin
set PATH=%PATH%;c:\go\bin;%GOBIN%
 
REM finally make my playground source folder as my PWD, write your hello.go here and execute "go install" and check GOBIN folder for exe.
cd %GOPATH%\src
CMD
