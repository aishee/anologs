package anologs

type options struct {
	InputModel  string `short:"i" long:"input"`
	OutputModel string `short:"o" long:"output"`
}

func readFile(fpath string, engine *anologs.Engine) error {
	fp, err := os.Open(fpath)
	if err != nil {
		log.Fatal("Fail to open file: ", fpath, " ", err)
		return err
	}
	defer fp.Close()
	s := bufio.NewScanner(fp)
  for s.Scan() {
    text := s.Text()
    if len(text) >0 {
      engine.Read(text)
    }
  }

}
