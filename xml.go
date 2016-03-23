type _ struct {
	Song struct {
		Version      string `xml:"version,attr"`
		CreatedIn    string `xml:"createdIn,attr"`
		ModifiedIn   string `xml:"modifiedIn,attr"`
		ModifiedDate string `xml:"modifiedDate,attr"`
		Properties   struct {
			Copyright     string   `xml:"copyright,omitempty"`
			CcliNo        string   `xml:"ccliNo,omitempty"`
			Released      string   `xml:"released,omitempty"`
			Transposition string   `xml:"transposition,omitempty"`
			Key           string   `xml:"key,omitempty"`
			Variant       string   `xml:"variant,omitempty"`
			Publisher     string   `xml:"publisher,omitempty"`
			Version       string   `xml:"version,omitempty"`
			Keywords      string   `xml:"keywords,omitempty"`
			VerseOrder    string   `xml:"verseOrder"`
			Comments      []string `xml:"comments,omitempty"`

			Title []struct {
				Original string `xml:"original,attr,omitempty"`
				Lang     string `xml:"lang,attr,omitempty"`
			} `xml:"titles>title"`

			Author []struct {
				Type string `xml:"type,attr,omitempty"`
				Lang string `xml:"lang,attr,omitempty"`
			} `xml:"authors>author,omitempty"`

			Tempo struct {
				Type  string `xml:"type,attr,omitempty"`
				Value string `xml:",innerxml"`
			} `xml:"tempo,omitempty"`

			Songbook []struct {
				Name  string `xml:"name,attr"`
				Entry string `xml:"entry,attr,omitempty"`
			} `xml:"songbooks>songbook,omitempty"`

			Theme []struct {
				Lang  string `xml:"lang,attr,omitempty"`
				Value string `xml:",innerxml"`
			} `xml:"themes>theme"`
		} `xml:"properties"`

		Verse []struct {
			Lang            string `xml:"lang,attr,omitempty"`
			transliteration string `xml:"translit,attr,omitempty"`
			Name            string `xml:"name,attr"`
			Lines           []struct {
				Part  string `xml"part,attr"`
				Value string `xml:",innerxml"`
			} `xml:"lines"`
			Comments []string `xml:"comments,omitempty"`
		} `xml:"lyrics>verse"`
	} `xml:"song"`
}

