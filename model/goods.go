package model


type Good struct {
    ID       int    `db:"id" json:"id"`
    Name     string `db:"name" json:"name"`
    Price    int    `db:"price" json:"price"`
    Number   int    `db:"number" json:"number"`
    ImgUrl   string `db:"img_url" json:"img_url"`

}