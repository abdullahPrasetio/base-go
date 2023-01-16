package masking

/********************************************************************************
* Temancode Masking Package                       		  	                    *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

func Masking(text string, start int, end int) string {
	newText := ""
	var value string
	ends := len(text) - (end + 1)
	for i, r := range text {
		value = string(r)
		if i >= start && i <= ends {

			value = "*"
		}
		newText += value
	}

	return newText
}
