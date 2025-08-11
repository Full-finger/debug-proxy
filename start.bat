@echo off
title Debug-Proxy 一键启动
chcp 65001 > nul

:: 1. 按需编译
if not exist debug-proxy.exe (
    echo [INFO] 未找到 debug-proxy.exe，开始编译...
    where go >nul 2>nul
    if errorlevel 1 (
        echo [ERROR] 未检测到 go 环境，请先安装 Go!
        pause
        exit /b 1
    )
    go build -ldflags="-s -w" -o debug-proxy.exe ./cmd/debug-proxy
    if errorlevel 1 (
        echo [ERROR] 编译失败，请检查源代码或依赖!
        pause
        exit /b 1
    )
    echo [INFO] 编译完成
)

:: 2. 启动
echo [INFO] 启动 debug-proxy，默认端口 8080...
start debug-proxy.exe -addr :8080
timeout /t 2 >nul       :: 等 2 秒让服务起来
start http://127.0.0.1:8080

:: 3. 异常退出提示
echo.
echo [INFO] 程序已退出，按任意键关闭窗口...
pause > nul