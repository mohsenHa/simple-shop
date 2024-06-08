package simpletransaction

type Transaction struct {
	rollbackChannel chan bool
	commitChannel   chan bool
}

func New() *Transaction {
	return &Transaction{
		rollbackChannel: make(chan bool),
		commitChannel:   make(chan bool),
	}
}

func (t *Transaction) GetRollbackChannel() <-chan bool {
	return t.rollbackChannel
}
func (t *Transaction) GetCommitChannel() <-chan bool {
	return t.commitChannel
}
func (t *Transaction) Rollback() {
	close(t.rollbackChannel)
}
func (t *Transaction) Commit() {
	close(t.commitChannel)
}
