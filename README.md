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

1) Run filegen - will create config.json
2) Run any text editor and open config.json , enter the required parameters then save.
3) Run filegen.exe
4) View the result of the work :)

Description of configuration file:

Path - the path to the folder where the files will be created
CountFiles - number of files that will be created
File.Name.Separator - a symbol that will divide file name on the template and randomly generated
File.Name.Prefix - specifies the filename template
File.Name.RandomPart.Size - the number of characters in the file name
File.Name.RandomPart.Size.Min - the minimum number of characters in the file name
File.Name.RandomPart.Size.Max	 - the maximum number of characters in the file name
File.Name.Ext - file extension
File.Data.Size - size of created file
File.Data.Size.Min - the minimum size of created file (in bytes)
File.Data.Size.Max - the maximum size of created file (in bytes)
