package TransportData

func ParseScaleData(data *ScaleResponse) {
	/*

		   команда 0x95, ответ 0x7F - успешное подключение к устройству

		   команда 0x99 - запрос высоты
		   команда 0x88 - запрос ширины
		   команда 0x77 - запрос длины

		   TODO протокол измерения "0x2D 0x7F/0x7E 0x0B 0x64 0x7B"
		   TODO команда 99, ответ "0x2D - начало строки, 0x7F/0x7E - флаг готовности,  0xB - датчик, 0x64 - растояние, 0x7B - конец строки" все в 16ричной системе счисления


		   0x0B - ширина
		   0x16 - высота
		   0x21 - длина

		   0x7F - готов
		   0x7E - неготов
	   */
}

func ParseRulerData(data *RulerResponse) {

}