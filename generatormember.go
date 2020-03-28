package main

func fakeMemberGenerator(name string, shopid string, memberChan chan member) {
	_createupdatetime := randomDateTime(2000, 2020)
	fakeMember := member{
		Name:        name,
		Memo:        "",
		Gender:      randmoGener(),
		Birthday:    randomDate(1980, 2020),
		Phone:       randomPhone(),
		Cardid:      randomCardid(),
		Point:       0,
		CheckinShop: "",
		ShopsList:   []string{shopid},
		Created_at:  _createupdatetime,
		Updated_at:  _createupdatetime,
	}

	memberChan <- fakeMember
}
