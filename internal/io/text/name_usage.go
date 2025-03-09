package text

func (t *text) importNamesUsage() error {
	return nil
}

// 	var err error
// 	chIn := make(chan string)
// 	chOut := make(chan coldp.NameUsage)
// 	var wg1, wg2 sync.WaitGroup
// 	wg1.Add(1)
// 	wg2.Add(1)
//
// 	go t.reader(chIn, chOut, *wg1)
// 	go t.writer(chOut, &wg2)
//
// 	err = loadLines(chIn, t.textPath)
// 	if err != nil {
// 		return err
// 	}
// 	close(chIn)
// 	return nil
// }
//
// func (t *text) reader(
// 	chIn <-chan string,
// 	chOut chan<- coldp.NameUsage,
// 	wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	code := "any"
// 	switch t.cfg.NomCode {
// 	case coldp.Botanical, coldp.Cultivars:
// 		code = "botanical"
// 	}
// 	p := <-t.parserPool[code]
// 	defer func() {
// 		t.parserPool[code] <- p
// 	}()
//
// 	for name := range chIn {
// 		prsd := p.ParseName(name).Flatten()
// 		nu := coldp.NameUsage{
// 			ScientificNameString: name,
// 			ScientificName:       prsd.CanonicalFull,
// 			Authorship:           prsd.Authorship,
// 			Cardinality:          coldp.ToInt(prsd.Cardinality),
// 		}
// 		chOut <- nu
// 	}
//
// }
//
// func loadLines(chIn chan<- string, path string) error {
// 	f, err := os.Open(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
//
// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		chIn <- scanner.Text()
// 	}
//
// 	err = scanner.Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
