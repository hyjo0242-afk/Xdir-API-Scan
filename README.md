```markdown

  ██╗  ██╗██████╗ ██╗██████╗ 
  ╚██╗██╔╝██╔══██╗██║██╔══██╗
   ╚███╔╝ ██║  ██║██║██████╔╝
   ██╔██╗ ██║  ██║██║██╔══██╗
  ██╔╝ ██╗██████╔╝██║██║  ██║
  ╚═╝  ╚═╝╚═════╝ ╚═╝╚═╝  ╚═╝
```

Xdir — Fast API Endpoint Extractor

---

Скачивает JavaScript-файлы с целевого сайта и автоматически извлекает из них скрытые эндпоинты API (/api/..., /v1/..., /graphql).

Установка

Убедитесь, что у вас установлен Go (версия 1.21 или выше).

```bash
git clone https://github.com/hyjo0242-afk/xdir.git
cd xdir
go mod tidy
go build -o xdir .
```

Использование

Базовый запуск:

```bash
./xdir -u https://example.com
```

Флаги:

· -u — URL цели (обязательно)
· -json — сохранить результат в JSON
· -o — имя файла для сохранения (по умолчанию xdir-output.json)
· -n — отключить анимацию баннера

Пример с сохранением в файл:

```bash
./xdir -u https://example.com -json -o result.json
```
