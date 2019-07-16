package borg

type BorgInfoRun struct {
	BorgRun
}

func NewBorgInfoRun() (*BorgInfoRun, error) {
	r := BorgRun{
		SubCommand: "info",
	}
	err := r.Prepare()
	if err != nil {
		return nil, err
	}


}

func (r *BorgInfoRun) ProcessResult() {
	
}
