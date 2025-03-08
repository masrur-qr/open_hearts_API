package structs

import (
	"github.com/dgrijalva/jwt-go"
)

type LangForUser struct{
	Name       string

}


type UserStruct struct {

	Id         string `bson:"_id"`
	Photo string `json:"photo"`
	Ru LangForUser `json:"ru"`
	En LangForUser `json:"en"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	Permission string `json:"permission"`
	Code uint64 `json:"code"`
}

type Verify struct {
	Id   string `bson:"_id"`
	Code  uint64 `json:"code"`
	Email string `json:"email"`
	User_Id string `json:"user_id"`
}
type CheckPassword struct {
	Id   string `bson:"_id"`
	Email string `json:"email"`
	Password string `json:"password"`
}




type Tokenclaim struct{
	jwt.StandardClaims
	UserId     string `bson:"_id"`
	Phone      string `json:"phone"`
	Permission string `json:"permission"`
}









type LangForStatistic struct{
	Description	string `json:"description"`
}

type AddStatistic struct{
	 Id	string `bson:"_id"`
	 Quantity	int `json:"quantity"`
	 Ru LangForStatistic `json:"ru"`
	 En LangForStatistic `json:"en"`
}


type LangForProjectStatistic struct{
	Description	string `json:"description"`
	Name string `json:"name"`
}

type AddStatisticForCenter struct{
	 Id	string `bson:"_id"`
	 Quantity	int `json:"quantity"`
	 Ru LangForProjectStatistic `json:"ru"`
	 En LangForProjectStatistic `json:"en"`
}




type ChangNumber struct{
	 Id	string `bson:"_id"`
	 Quantity	int `json:"quantity"`
}







type LangForPatient struct{
	Full_Name string `json:"full_name"`
	Description []string `json:"description"`
	Quot string `json:"quot"`
}
type Patient_story struct{
	Id string `bson:"_id"`
	Photo string  `json:"photo"`
	Smallphoto string  `json:"smallphoto"`
	En LangForPatient `json:"en"`
	Ru LangForPatient`json:"ru"`
}
type DeletePatientStory struct{
	Id string `bson:"_id"`
	FolderName string `json:"folderName"`
	ImageId string 
}








type Partner struct{
	Id string `bson:"_id"`
	Logo string 
}
type DeletePartner struct{
	Id string `bson:"_id"`
	LogoId string `json:"logoId"`
	FolderName string `json:"folderName"`
}








type RuDiract struct{
	Title string `json:"title"`
	Short_info string `json:"short_info"`
}
type EnDiract struct{
	Title string `json:"title"`
	Short_info string `json:"short_info"`
} 
type Diraction struct{
	Id string `bson:"_id"`
	Ru RuDiract `json:"ru"`
	En EnDiract `json:"en"`
}







type LangForDiraction struct{
	Title string `json:"title"`
	Description string `json:"description"`
	MainServices [] string  `json:"mainServices"`
}

type MainDiraction struct{
	Id string `bson:"_id"`
	Photo string `json:"photo"`
	Phone string `json:"phone"`
	Ru LangForDiraction `json:"ru"`
	En LangForDiraction `json:"en"`
}
type DeleteMainDiraction struct{
	Id string `bson:"_id"`
	PhotoId string
	FolderName string `json:"foldername"`
}










type LangForProject struct{
	Title string `json:"name"`
	Description string `json:"description"`
	Address string `json:"adress"`
}

type Project struct{
	Id string `bson:"_id"`
	Photo string
	Phone string `json:"phone"`
	Email string `json:"email"`
	En LangForProject `json:"en"`
	Ru LangForProject `json:"ru"`
}
type DeleteProject struct{
	Id string `bson:"_id"`
	FolderName string `json:"folderName"`
	PhotoId string `json:"photoid"`
}




type Time_For_Team struct{
		Days_Of_Week [] string `json:"days_of_week"`
		Start_Time string `json:"start_time"`
		End_Time string `json:"end_time"`
}


type LangForTeam struct{
	Full_Name string `json:"full_name"`
	Job_Title string `json:"job_title"`
	Education string `json:"education"`
	Expirence string `json:"expirence"`
	Place string `json:"place"`
}



type Team struct{
	Id string `bson:"_id"`
	Ru LangForTeam `json:"ru"`
	En LangForTeam `json:"en"`
	Photo string  `json:"photo"`

}
type DeleteTeam struct{
	Id string `bson:"_id"`
	FolderName string `json:"folderName"`
	PhotoId string `json:"photoid"`
}


type Program struct{
	Id string `bson:"_id"`
	Photo string `json:"photo"`
	Ru ProgramLang  `json:"ru"`
	En ProgramLang  `json:"en"`
}
type ProgramLang struct{
	Title string `json:"title"`
	Description string `json:"description"`
}
type Services struct{
	Id string `bson:"_id"`
	Photo string `json:"photo"`
	Ru ServiceLang  `json:"ru"`
	En ServiceLang  `json:"en"`
}
type ServiceLang struct{
	Title string `json:"name"`
	Maindescription string `json:"maindescription"`
	Servisec []string `json:"servisec"`
	LastDescription string `json:"lastdescription"`
}



type UpdatePassword struct{
	Email string `json:"email"`
	Code int `json:"code"`
}


type RuReport struct {
	Title string `json:"title"`
}
type EnReport struct {
	Title string `json:"title"`
}



type AddReport struct{
	Id string `bson:"_id"`
	Ru RuReport `json:"ru"`
	En EnReport `json:"en"`
	File string `json:"file"`
	Date string `json:"date"`
}

type DeleteReport struct{
	Id string `bson:"_id"`
	FolderName string `json:"folderName"`
	FileId string `json:"fileId"`
}
