[![lang en](https://img.shields.io/badge/lang-en-red)](https://github.com/EmilioAugust/localFiles)
[![lang ru](https://img.shields.io/badge/lang-ru-blue)](github.com/EmilioAugust/localFiles/blob/main/README.ru.md)

# LAN File Sharing (Golang)
Простое веб-приложение для передачи файлов между устройствами в одной локальной сети (Wi-Fi).

---

## Функции
- Работает в LAN
- Передача файлов через HTTP
- Доступ через браузер (ПК, телефон и т.д.)

---

## Запуск через Docker

**1. Установить Docker и Docker Compose**
Убедитесь, что у вас установлены:
- Docker
- Docker Compose

**2. Склонировать проект**
```bash
git https://github.com/EmilioAugust/localFiles.git
cd localFiles
```

**3. Запустить контейнер**
```bash
docker-compose up --build
```

**4. Открыть веб-приложение**
```bash
http://localhost:8080
```

---

## Запуск без Docker

**1. Собрать проект**
```bash
go build -o server ./app/cmd/server
```

**2. Запустить**
```bash
./server
```
**3. Открыть**
```bash
http://localhost:8080
```

---

## Как зайти с другого устройства (телефон, ноутбук)

1. Узнайте IP адрес компьютера, где запущен сервер:
**Windows:**
```bash
ipconfig
```
Ищите:
```bash
IPv4 Address: 192.168.x.x
```
**macOS / Linux:**
```bash
ipconfig
```
или:
```bash
ip a
```
2. Откройте на другом устройстве:
```bash
http://<YOUR_IP>:8080
```
Пример::
```bash
http://192.168.1.104:8080
```
**⚠️ Важно:**
- Все устройства должны быть в одной Wi-Fi сети
- Firewall не должен блокировать порт 8080

---

## Хранение файлов
Все загруженные файлы сохраняются в папке:
```bash
./files
```
(через Docker volume → /app/uploads)

---

## Вклад

PR приветствуются!
Вы можете добавить новые источники, улучшить структуру или предложить новые функции.

---

## Лицензия

MIT — можно использовать в любом проекте.
