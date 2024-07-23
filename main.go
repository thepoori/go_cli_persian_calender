package main

import (
	"fmt"
	"time"
	"github.com/wsxiaoys/terminal/color"
	)	
	
var not_kabise = [12]int{ 0 , 31 , 59 , 90 , 120 , 151 , 181 , 212 , 243 , 273 , 304 , 334} 
	
var kabise = [12]int{0 , 31 , 60 , 91 , 121 , 152 , 182 , 213 , 244 , 274 , 305 , 335}
	
var persian_month = map[int]string{
		1 : "farvardin",2 : "ordibehesht",3 : "khordad",4 : "tir",5 : "mordad",6 : "shahrivar",7 : "mehr",8 : "aban",9 : "azar",10 : "dey" ,11 : "bahman",	12 : "esfand",
	}

/*var persian_weekday = map[string]string{
	"Saturday" : "Shanbe",
	"Sunday" : "Yekshanbe",
	"Monday" : "Doshanbe",
	"Tuesday" : "Seshanbe",
	"Wednesday" : "Chaharshanbe",
	"Thursday" : "Panjshanbe",
	"Friday" : "Jomee",

}*/

func main(){
	color.Println(print_str_time())

	
}	


func print_str_time()string{

	//wk := time.Time.Weekday(time.Now())
	var month_days int
	y , m , d := time.Now().Date()
	var year , month , day int
	if is_kabise(y){
		year , month , day = calc(y, int(m) , d , true)
	}else{
		year , month , day = calc(y, int(m) , d , false)
	}
	
	if month < 6{
		month_days = 31
	}else if month > 6{
		month_days = 30
	}else if month == 12 {
		if is_kabise(y) {
			month_days = 29
		}else{
			month_days = 30
		}}
		print_str := fmt.Sprintf(" %v %v %v \n" , day , persian_month[month] , year  )

		
		for i := 1 ; i<= month_days ; i++{
			if i == day{
				print_str += color.Sprintf("@b%2d ", i)

			}else{
				print_str += fmt.Sprintf("%2d ", i)

			}
			if i % 7 == 0{

				print_str += "\n"
			}

		}
		return print_str
}

func calc(m_year int, m_month int , m_day int ,is_kabise bool) (int,  int, int){
	var month int
	var year int
	var date int
	var diff_between_days int
	if !is_kabise{
		date = not_kabise[m_month - 1] + m_day
		if date > 79{
			date = date - 79
			if date <= 186{
				switch date % 31{
					case 0:
						month = date / 31
						date = 31
						
					default:
						month = (date /31) + 1
						date = (date % 31)
				}		
				year  = m_year - 621
			}else{
				date  = date - 186
				switch date % 30{
					case 0 :
						month = (date / 30) + 7
						date = date % 30
				}		
				year = m_year - 621

			}	
		}else{
			if (m_year > 1996) && (m_year % 4) == 1{
				diff_between_days = 11
			}else{
				diff_between_days = 10
			}	
			date = date + diff_between_days
			switch date % 30 {
				case 0:
					month = (date/30) + 9
					date = 30
				default:
					month = (date/30) + 10
					date = date % 30
			}		
			year = m_year - 622
		}	
	}else if is_kabise{
		date = kabise[m_month - 1] + m_day
		if(m_year >= 1996){
			diff_between_days = 79
		}else{
			diff_between_days = 80
		}	
		if date > diff_between_days{
			date = date - diff_between_days
			if date <= 186{
				switch date % 31{
				case 0:
					month = date / 31
					date = 31
				default:
					month = (date / 31) + 1
					date = date % 31
				}	
				year = m_year - 621
				
			}else{
				date = date - 186
				switch date % 30{
					case 0 :
						month = date / 30 + 7
						date = date % 30
					}	
					year = m_year - 621

			
			}		
		}else{
			date = date + 10
			switch date % 30{
			case 0 :
				month = date / 30 + 9
				date = 30
			default:
				month = (date / 30) +10
				date = date % 30
			}	
			year = m_year -622
		}	
		

	}	
	return year , month , date
}	




func is_kabise(year int) bool{
	if year %4 == 0{
		return true
	}else{
		return false}
}
