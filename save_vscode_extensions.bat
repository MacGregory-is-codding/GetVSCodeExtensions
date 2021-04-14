@echo off

rem TODO: add flags: -p %path%, -e

rem ####################################################################
rem #                                                                  # 
rem #      Saves VS Code extensions to %DEST_PATH% in %FILE_NAME%      #       
rem #  if %DEST_PATH% is not specified with -p flag - saves to Desktop #
rem #                                                                  #
rem ####################################################################

:main
    cls
    call :check_rights
    call :set_vars
    call :print_extensions_to_file
exit

:check_rights
    net session >nul 2>&1
    if not %errorlevel% == 0 call :print_fail_not_admin
goto :EOF

:set_vars
    rem TODO: if %DEST_PATH% == "" 
    set DEST_PATH=%userprofile%\desktop
    set FILE_NAME=%username%extensions%date%.txt
    set FULL_PATH=%DEST_PATH%\%FILE_NAME%
    set ERROR_COLOR = 4
    set SUC_COLOR = 2
goto :EOF

:print_result_to_file
    rem code --list-extensions
    rem echo. > "extensions%DATE%.txt"
    rem goto :EOF
    rem:success
    rem @echo on
    rem echo "File %FILE_NAME% in %DEST_PATH% created"

    call :print_suc
goto :EOF

:print_suc
    cls
    color %SUC_COLOR%
    echo ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    echo %USERNAME%'s VSCode extensions were saved
    echo in '%FULL_PATH%'
    echo at %date% %time%
    echo Good luck!
    echo ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
pause
exit

:print_fail_invalid_path
    cls
    color %ERROR_COLOR%
    echo ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    echo ERROR: File %DEST_PATH% is invalid
    echo ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
pause
exit

:print_fail_not_admin
    cls
    color %ERROR_COLOR%
    echo ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    echo ERROR: Please, run this script with admin rights
    echo ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
pause
exit