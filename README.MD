# Приложение разработано в качестве вступительного задания в компанию QSOFT

## Задание:

### Часть 1 - приложение на фреймворке gin
Разработайте веб-приложение на фреймворке gin, которое содержит один маршрут /when/:year. В обработчике этого маршрута вам нужно вывести строку, которая покажет, сколько дней осталось или прошло до 1-го января года, указанного в параметре маршрута,

Например: /when/2000 Должно вывести сколько дней прошло с 01.01.2000 до сегодня

Например: /when/2030 Должно вывести сколько дней осталось до 01.01.2030

Статус HTTP-ответа должен быть 200.

Обработка ошибок
Ошибки работы приложения должны попадать в консольный лог.

Запуск приложения
Для запуска приложения создайте Makefile с двумя командами:

run - запускает приложение
build - компилирует приложение
### Часть 2 - middleware на фреймворке gin
Создайте свой собственный middleware в этом приложении, который проверяет наличие HTTP-заголовка. Если он содержит заголовок X-PING со значением ping, то в ответе вашего сервиса добавьте заголовок ответа X-PONG со значением pong.

### Необязательная Часть 3
Постарайтесь организовать структуру и код проекта, следуя лучшим практикам организации кода приложения и структуры приложения на языке Go.

## Реализовный функционал:

- HTTP сервер с реализацией Gracefully Shutdown (https://gin-gonic.com/docs/examples/graceful-restart-or-stop/)
- Использование Gin для end-point & middleware
- End-point по пути /when/:year, так же добавлены info и при ошибочном url открытие страницы 404.
- Сборка и запуск приложения через Makefile
- Структурирование приложения
- Простое логгирование и ошибки работы с помощью Zap.
- Переменные окружения .env