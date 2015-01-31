Опис:

Програма призначена для створення заданої кількості файлів, заданого об'єму.

Можливості:

- Створення до 100 000 фалів
- Створення файлів об'ємом від 1 байту до 10 гігабайт
- Створення імені файлу за шаблоном
- Створення кількості символів в імені фалу згідно вказаного діапазоні
- Створення файлу з вказаним розширенням
- Файл налаштувань для збереження параметрів

Використання:

1) Запустити filegen.exe - створиться config.json
2) Відкрити будь-яким текстовим редактором config.json та ввести необхідні параметри, зберегти
3) Запустити filegen.exe
4) Переглянути результат роботи :)

Опис конфігураційного фалу:

"Path": "./test" - шлях до теки де будуть створені файли
"CountFiles": 10 - кількість файлів, що буде створена
"Separator": "_" - символ, що буде розділяти ім'я фалу за шаблоном та випадково згенероване
"Prefix": ""     - вказується шаблон імені файлу
"Size": 	 - кількість символів в імені файлу							
"Min": 8	 - мінімальне значення кількість символів в імені файлу	
			
"Max": 15  	 - максимальне значення кількість символів в імені файлу	
"Ext": ""	 - розширення файлу
"Size": 	 - об'єм створюваного файлу	
				
"Min": 10
	 - мінімальний об'єм створюваного файлу (в байтах)  			
"Max": 100 	 - максимальний об'єм створюваного файлу (в байтах). 


Description:

The program is designed to create the specified number of files with a specified size.

Features:

- Creation of up to 100 000 files
- Create a file size from 1 byte to 10 gigabyte
- Create a file name for the template
- Create a number of characters in the file name relative to the specified range
- Create a file with the specified extension
- Configuration file to save the settings

Use:

1) Run filegen.exe - will create config.json
2) Run any text editor and open config.json , enter the required parameters then save.
3) Run filegen.exe
4) View the result of the work :)

Description of configuration file:

"Path": "./test" - the path to the folder where the files will be created
"CountFiles": 10 - number of files that will be created
"Separator": "_" - a symbol that will divide file name on the template and randomly generated
"Prefix": ""     - specifies the filename template
"Size":          - the number of characters in the file name
"Min": 8         - the minimum number of characters in the file name
"Max": 15        - the maximum number of characters in the file name
"Ext": ""        - file extension
"Size":          - size of created file
"Min": 10        - the minimum size of created file (in bytes)
"Max": 100       - the maximum size of created file (in bytes).