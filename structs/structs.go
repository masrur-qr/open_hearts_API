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
	Description string `json:"description"`
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




type LangForTeam struct{
	Full_Name string `json:"full_name"`
	Job_Title string `json:"job_title"`
	Education string `json:"education"`
	Expirence string `json:"expirence"`
	
}



type Team struct{
	Id string `bson:"_id"`
	Place string `json:"place"`
	Ru LangForTeam `json:"ru"`
	En LangForTeam `json:"en"`
	Photo string  `json:"photo"`

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
	Title string `json:"title"`
	Description string `json:"description"`
}



type UpdatePassword struct{
	Email string `json:"email"`
	Code int `json:"code"`
}

