# thumbnail-proxy
## Описание
- Сервис thumbnail-proxy используется для скачивания превью видео с YouTube.
- У thumbnail-proxy есть клиентское приложение (CLI утилита): https://github.com/sletkov/thumbnail-proxy-client
- Посмотреть документацию по API thumbnail proxy можно в директории /pkg/sdk/doc

## Установка и запуск
1. Склонируйте репозиторий на локальную машину командой `git clone [github.com/repository/path]`
2. Перейдите в директорию сервиса командой `cd [path/to/service]`
3. Создайте файл .env по примеру файла .env.example (обязательно укажите параметр YOUTUBE_API_KEY).
Чтобы создать ключ Youtube API перейдите по ссылке: https://console.cloud.google.com/apis/library/browse?q=google и выберете "Youtube Data API v3"
4. Запустите сервис с помощью docker-compose командой `docker-compose up`