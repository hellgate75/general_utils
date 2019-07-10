package process // import "github.com/hellgate75/general_utils/process"


FUNCTIONS

func InitLogger()
func NewProcess(text string, pid PID, processState ProcessState) processError

TYPES

type PID int64

var REG_PID PID = 0
type ProcessChannel chan interface{}

type ProcessFunction func(ProcessManager, *ProcessChannel, *ProcessChannel) error

type ProcessManager interface {
	Start()
	Running() bool
	Status() ProcessState
	GetInChannel() *ProcessChannel
	GetOutChannel() *ProcessChannel
}

func NewProcessManager(fn ProcessFunction) ProcessManager
type ProcessState int

const (
	DEFAULT  ProcessState = 1
	STARTING ProcessState = 2
	RUNNING  ProcessState = 3
	PAUSED   ProcessState = 4
	RESUMING ProcessState = 5
	DONE     ProcessState = 6
	ERROR    ProcessState = 7
)
