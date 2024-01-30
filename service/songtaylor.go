package service 

type SongTaylorService interface{
	AddSong(nome string, LinkSong string) (bool, error)
}

func AddSong(nome string, LinkSong string) (bool, error){
	
}