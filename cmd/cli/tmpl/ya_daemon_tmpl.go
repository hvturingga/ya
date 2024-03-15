package tmpl

type YaDaemonPs1 struct {
	TaskName string
	Command  string
}

const YaDaemonTpl = `
$Action = New-ScheduledTaskAction -Execute 'powershell.exe' -Argument '-WindowStyle Hidden -Command "Start-Process -Verb runAs powershell -WindowStyle Hidden -ArgumentList ''{{.Command}}''"'
$Trigger = New-ScheduledTaskTrigger -AtLogOn
$Principal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -LogonType ServiceAccount -RunLevel Highest
Register-ScheduledTask -Action $Action -Trigger $Trigger -Principal $Principal -TaskName "{{.TaskName}}" -Description "Run {{.TaskName}} at startup"  -Force

Remove-Item $MyInvocation.MyCommand.Path -Force
`

type YaDaemonSystemd struct {
	User      string
	ExecStart string
}

const YaDaemonSystemdTpl = `[Unit]
Description=YA Daemon Service
After=network.target

[Service]
Type=simple
User={{.User}}
ExecStart={{.ExecStart}}
Restart=on-failure

[Install]
WantedBy=multi-user.target
`
