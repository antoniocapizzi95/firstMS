package repository

var AddressBookDb AddressBookRepo

var AddressBookQueue AddressBookRepo

func InitModule(abrDb *AddressBookRepo, abdQueue *AddressBookRepo) {
	AddressBookDb = *abrDb
	AddressBookQueue = *abdQueue
}
