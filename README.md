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

{
	"Path": "./test",				- the path to the folder where the files will be created
	"CountFiles": 10,				- number of files that will be created
	"File": {
		"Name": {
			"Separator": "_",		- a symbol that will divide file name on the template and randomly generated
			"Prefix": "",			- specifies the filename template
			"RandomPart": {
				"Size": {			- the number of characters in the file name
					"Min": 8,		- the minimum number of characters in the file name
					"Max": 15		- the maximum number of characters in the file name
				}
			},
			"Ext": ""				- file extension
		},
		"Data": {
			"Size": {				- size of created file
				"Min": 10,			- the minimum size of created file (in bytes)
				"Max": 100			- the maximum size of created file (in bytes)
			}
		}
	}
}
