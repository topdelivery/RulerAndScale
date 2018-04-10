# RulerAndScale
Модуль для весов на базе терминала TBS-A для измерения габаритов коробки, состоит из 4х датчиков HC-SR04 подключеных в 1 пин
по этому способу [мануал](http://www.instructables.com/id/Hack-an-HC-SR04-to-a-3-pin-sensor/)

констранты которые необходимо указать:
Arduino:
1) WIDTH_MAX - растояние между боковыми датчиками для измерения ширины коробки
2) LENGTH_MAX - растояние от заднего датчика до стенки куда прикладывать коробку
3) TOP_MAX - растояние между верхним датчиком и весами
4) LEFT_PING_PIN, TOP_PING_PIN, BACK_PING_PIN, RIGHT_PING_PIN - пины в которые подключать дальномеры
Go: 
1) в файле SelectPort переменная portClass указывает какие типы портов сканировать а цикл for задает колво 
по умолчанию "/dev/ttyS", "/dev/ttyACM", "/dev/ttyUSB", от 0 до 9го порта
![Image alt](https://github.com/TrashPony/RulerAndScale/raw/master/image.png)
