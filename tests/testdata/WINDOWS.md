# Windows Services Tests
## RUN SING-BOX IN THE BACKGROUND
```shell

Start-Process -Verb runAs powershell -WindowStyle Hidden -ArgumentList 'sing-box.exe run -c conf.json"'
```
## GET SING-BOX BACKGROUND PROCESS
```shell
Get-Process | Where-Object {$_.ProcessName -like "*sing-box*"}
```
## KILL SING-BOX BACKGROUND PROCESS
```shell
Start-Process powershell -Verb runAs -ArgumentList "Stop-Process -Name sing-box -Force"
```

## TASK

## SING BOX AUTO START
```shell
$taskExists = Get-ScheduledTask | Where-Object { $_.TaskName -eq "SingBoxAutoStart" }
if ($taskExists) {
    Unregister-ScheduledTask -TaskName "SingBoxAutoStart" -Confirm:$false
}

$Action = New-ScheduledTaskAction -Execute 'powershell.exe' -Argument '-WindowStyle Hidden -Command "Start-Process -Verb runAs powershell -WindowStyle Hidden -ArgumentList ''sing-box.exe run -c conf.json''"'
$Trigger = New-ScheduledTaskTrigger -AtLogOn
$Principal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -LogonType ServiceAccount -RunLevel Highest
Register-ScheduledTask -Action $Action -Trigger $Trigger -Principal $Principal -TaskName "SingBoxAutoStart" -Description "Run SingBoxAutoStart at startup"

Read-Host "Press Enter to exit"
```

## SING BOX AUTO START RUN
```shell
Start-Process  -Verb RunAs powershell -ArgumentList "-ExecutionPolicy Bypass", "-File `"`ps.ps1`""
```

## SING BOX AUTO START INFO
```shell
Get-ScheduledTask -TaskName "YaDaemon" | Get-ScheduledTaskInfo
```

## SING BOX AUTO START UNREG
```shell
Unregister-ScheduledTask -TaskName "YaDaemon"
```

## SING-BOX RELOAD
```shell
Start-Process -Verb runAs powershell -ArgumentList 'Stop-Process -Name sing-box -Force'

Start-Process -Verb runAs powershell -WindowStyle Hidden -ArgumentList 'sing-box.exe run -c conf.json'

Start-Process -Verb RunAs powershell -ArgumentList '-ExecutionPolicy Bypass', '-File ".\autostart.ps1"'
```

RUN DAEMON PROCESS
```shell
Start-Process powershell -Verb runAs -ArgumentList "-NoExit D:\dev\ya\.bin\ya-daemon.exe"
```