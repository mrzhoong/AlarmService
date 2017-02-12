{{define "navbar"}}
<a class="navbar-brand" href="/">告警告知</a>
<div>
    <ul class="nav navbar-nav">
        <li {{if .IsStrategy}}class="active"{{end}}><a href="/strategy">策略配置</a></li>
        <li {{if .IsEventLog}}class="active"{{end}}><a href="/eventlog">发送记录</a></li>
        <li {{if .IsOperationLog}}class="active"{{end}}><a href="/operationlog">操作日志</a></li>
        <li {{if .IsConfig}}class="active"{{end}}><a href="/config">系统配置</a></li>
    </ul>
</div>

<div class="pull-right">
    <ul class="nav navbar-nav">
        {{if .IsLogin}}
        <li><a href="/login?exit=true">退出登录</a></li>
        {{else}}
        <li><a href="/login">管理员登录</a></li>
        {{end}}
    </ul>
</div>
{{end}}