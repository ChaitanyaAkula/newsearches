package newsearches

import (
	"log"

	"github.com/ChaitanyaAkula/gittyjobsdb"
)


func NewSearchKeyword(keyword string) {
	var Searchvalue string
	db:=dbconnection.Connection()
	defer db.Close()

	result, err := db.Query("select searchvalue from searchresults where searchvalue=?", keyword)
	if err != nil {

		log.Fatal(err)
	}
	for result.Next() {

		err1 := result.Scan(&Searchvalue)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	if keyword == Searchvalue {

	}
	if keyword != Searchvalue {
		_, err2 := db.Query("insert into searchresults(searchvalue,approval) values(?,?)", keyword, "progress")
		if err2 != nil {

			log.Fatal(err2)
		}

	}
}
func NewSearchCompany(keyword string) {
	var Searchvalue string
	db:=dbconnection.Connection()
	defer db.Close()

	result, err := db.Query("select companyname from searchcompany where companyname=?", keyword)
	if err != nil {

		log.Fatal(err)
	}
	for result.Next() {

		err1 := result.Scan(&Searchvalue)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	if keyword == Searchvalue {

	}
	if keyword != Searchvalue {
		_, err2 := db.Query("insert into searchcompany(companyname,approval) values(?,?)", keyword, "progress")
		if err2 != nil {

			log.Fatal(err2)
		}

	}
}
func NewSearchSkills(keyword string) {
	var Searchvalue string
	db:=dbconnection.Connection()
	defer db.Close()

	result, err := db.Query("select skills from searchskills where skills=?", keyword)
	if err != nil {

		log.Fatal(err)
	}
	for result.Next() {

		err1 := result.Scan(&Searchvalue)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	if keyword == Searchvalue {

	}
	if keyword != Searchvalue {
		_, err2 := db.Query("insert into searchskills(skills,approval) values(?,?)", keyword, "progress")
		if err2 != nil {

			log.Fatal(err2)
		}

	}
}
func NewSearchJobtitle(keyword string) {
	var Searchvalue string
	db:=dbconnection.Connection()
	defer db.Close()

	result, err := db.Query("select jobtitle from searchjobtitle where jobtitle=?", keyword)
	if err != nil {

		log.Fatal(err)
	}
	for result.Next() {

		err1 := result.Scan(&Searchvalue)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	if keyword == Searchvalue {

	}
	if keyword != Searchvalue {
		_, err2 := db.Query("insert into searchjobtitle(jobtitle,approval) values(?,?)", keyword, "progress")
		if err2 != nil {

			log.Fatal(err2)
		}

	}
}
