[![lang ru](https://img.shields.io/badge/lang-ru-blue)](https://github.com/EmilioAugust/localFiles/blob/main/README.ru.md)
[![lang en](https://img.shields.io/badge/lang-en-red)](https://github.com/EmilioAugust/localFiles)

# LAN File Sharing (Golang)

A simple web application for transferring files between devices within the same local network (Wi-Fi).

---

## Features
- Works in LAN
- File transfer via HTTP
- Accessible via browser (PC, phone, etc.)

---

## Run with Docker

**1. Install Docker & Docker Compose**
Make sure you have:
- Docker
- Docker Compose

**2. Clone the repository**
```bash
git clone https://github.com/EmilioAugust/localFiles.git
cd localFiles
```

**3. Start the container**
```bash
docker-compose up --build
```

**4. Open the app**
```bash
http://localhost:8080
```

---

## Run without Docker

**1. Build the project**
```bash
go build -o server ./app/cmd/server
```

**2. Run**
```bash
./server
```
**3. Open the app**
```bash
http://localhost:8080
```

---

## Access from other devices

1. Find your local IP address:
**Windows:**
```bash
ipconfig
```
Look for:
```bash
IPv4 Address: 192.168.x.x
```
**macOS / Linux:**
```bash
ipconfig
```
or:
```bash
ip a
```
2. Open in browser on another device:
```bash
http://<YOUR_IP>:8080
```
Example:
```bash
http://192.168.1.104:8080
```
**⚠️ Important:**
- All devices must be on the same Wi-Fi network
- Make sure port 8080 is not blocked by firewall

---

## File storage
Uploaded files are stored in:
```bash
./files
```
(Docker volume → /app/uploads)

---

## Contributing

PRs are welcome!
Feel free to add new sources or improve code structure.

---

## License

MIT — free for all use.
