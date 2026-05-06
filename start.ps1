# Kill any process using port 1111
$port = 1111
$processInfo = netstat -ano | Select-String ":$port\s.*LISTENING"
if ($processInfo) {
    $procId = ($processInfo -split '\s+')[-1].Trim()
    Write-Host "Killing process on port $port (PID: $procId)..."
    taskkill /PID $procId /F 2>$null
    Start-Sleep -Milliseconds 700
}

# Start the server
go run src/main.go
