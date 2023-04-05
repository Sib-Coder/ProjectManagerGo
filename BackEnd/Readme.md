 Создали докер для быстрого запуска.<br/>

Создаём образ.
 ```bash
 docker build -t my_websui .
 ```
 Запускаем контейнер.
 ```bash
  sudo docker run -d   -p 3001:3001 --name webproj  my_websui
 ```

 Запускаем на порту 3001 для корректной работы.
