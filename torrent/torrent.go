package torrent

const port = 7777

type Info struct {
	Pieces      string
	PieceLength int
	Length      int
	Name        string
}

type Torrent struct {
	Announce string
	Info     Info
}

// func (to *Torrent) Download
