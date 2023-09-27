# GoIP

Obtener la direccion IP publica consultando un servicio HTTP REST

https://api.ipify.org/?format=json

[Se publico una entrada sobre este código aquí.](https://80bits.blog/index.php/2023/09/26/golang-obtener-ip-publica-y-enviarla-por-medio-de-un-bot-a-telegram/)

## Compilacion

Para Raspberry Pi Model 3

```bash
GOOS=linux GOARCH=arm64 go build -ldflags "-w -s" -trimpath -o goIp main.go
```

```bash
go build -o goIp main.go
```

Se necesita un archivo llamado settings.json para leer las variables de configuración:

```json
{
    "ip_url": "https://api.ipify.org/?format=json",
    "data_filename": "data.txt",
    "bot_url": "https://api.telegram.org/botAQUI_VA_EL_ID_DEL_BOT/sendMessage",
    "chat_id": "ID_USUARIO_TELEGRAM",
    "telegram_message": "La ip en home ha cambiado a: "
}
```

Tiene que estar en la misma ruta que el binario.