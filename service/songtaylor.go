package service 

type SongTaylorService interface{
	AddSong(name string, LinkSong string) (bool, error)
}

func AddSong(name string, LinkSong string) (bool, error){
	
}