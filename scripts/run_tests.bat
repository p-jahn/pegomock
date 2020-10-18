@echo off

set PACKAGES_TO_SKIP="generate_test_mocks/xtools_go_loader,generate_test_mocks/gomock_reflect,generate_test_mocks/gomock_source"

cd $0\..

goto main

:cleanup
if exist mock_display_test.go del /q mock_display_test.go
if exist matchers\ rd /Q /S matchers
exit /b

:cleanexit
set lasterr=%errorlevel%
call:cleanup
exit /b %lasterr%

:main

call:cleanup
%GOPATH%\bin\ginkgo -succinct generate_test_mocks/xtools_go_loader
if %errorlevel% neq 0 goto cleanexit
%GOPATH%\bin\ginkgo -r -skipPackage=%PACKAGES_TO_SKIP% --randomizeAllSpecs --randomizeSuites --race --trace -cover
if %errorlevel% neq 0 goto cleanexit

call:cleanup
%GOPATH%\bin\ginkgo -succinct generate_test_mocks/gomock_reflect
if %errorlevel% neq 0 goto cleanexit
%GOPATH%\bin\ginkgo --randomizeAllSpecs --randomizeSuites --race --trace -cover
if %errorlevel% neq 0 goto cleanexit

call:cleanup
%GOPATH%\bin\ginkgo -succinct generate_test_mocks/gomock_source
if %errorlevel% neq 0 goto cleanexit
%GOPATH%\bin\ginkgo --randomizeAllSpecs --randomizeSuites --race --trace -cover
if %errorlevel% neq 0 goto cleanexit

call:cleanup