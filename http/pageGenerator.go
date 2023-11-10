package http

type PageData struct {
	*Server
	Hostname string
}

func Generate(s *Server) PageData {
	var pageData = PageData{
		Server:   s,
		Hostname: "Some hostname",
	}

	return pageData
}
