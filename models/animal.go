package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	Subid         string         `gorm:"type:varchar(32);comment:動物的收容編號"`
	AreaPkid      uint           `gorm:"type:smallint;comment:動物所屬縣市代碼"`
	ShelterPkid   uint           `gorm:"type:smallint;comment:動物所屬收容所代碼"`
	Place         string         `gorm:"type:varchar(128);comment:動物的實際所在地"`
	Kind          string         `gorm:"type:varchar(16);comment:動物的類型"`
	Variety       string         `gorm:"type:varchar(16);comment:動物的品種"`
	Sex           sql.NullString `gorm:"type:ENUM('M', 'F', 'N');default:'N';comment:動物性別"`
	Bodytype      sql.NullString `gorm:"type:ENUM('SMALL', 'MEDIUM', 'BIG')"`
	Colour        sql.NullString `gorm:"type:varchar(16);comment:動物毛色"`
	Age           sql.NullString `gorm:"ype:ENUM('CHILD', 'ADULT');comment:動物年紀"`
	Sterilization string         `gorm:"type:ENUM('T', 'F', 'N');default:'N';comment:是否絕育"`
	Bacterin      string         `gorm:"type:ENUM('T', 'F', 'N');default:'N';comment:是否施打狂犬病疫苗"`
	Foundplace    string         `gorm:"type:varchar(128);comment:動物尋獲地"`
	Title         string         `gorm:"type:varchar(128);comment:動物網頁標題"`
	Status        string         `gorm:"type:ENUM('NONE', 'OPEN', 'ADOPTED', 'OTHER','DEAD');default:'NONE';comment:動物狀態"`
	Remark        sql.NullString `gorm:"type:varchar(128);comment:資料備註"`
	Caption       string         `gorm:"type:varchar(128);comment:其他說明"`
	Opendate      string         `gorm:"type:date;column:date_column";comment:開放認養時間(起)`
	Closeddate    string         `gorm:"type:date;column:date_column";comment:開放認養時間(迄)`
	albumFile     string         `gorm:comment:圖片名稱`
}
