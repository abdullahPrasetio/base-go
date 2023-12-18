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
	if start == 0 && end == 1 {
		return text
	}
	for i, r := range text {
		if i >= start && i <= end {
			newText += "*"
		} else {
			newText += string(r)
		}
	}
	return newText
}
