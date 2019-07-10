package utils // import "github.com/hellgate75/general_utils/utils"


FUNCTIONS

func CorrectInput(input string) string
func CreateFile(file string) (*os.File, error)
func CreateFileAndUse(file string, consumer func(*os.File) (interface{}, error)) (interface{}, error)
func CreateFileIfNotExists(file string) error
func DecodeBytes(encodedByteArray []byte) []byte
func DeleteIfExists(file string) error
func EncodeBytes(decodedByteArray []byte) []byte
func FileExists(file string) bool
func InitLogger()
func IntToString(n int) string
func MakeFolderIfNotExists(folder string) error
func StringToInt(s string) (int, error)

TYPES

type ArrayNav interface {
	Get() common.Type
	// Has unexported methods.
}

func NewArrayNav(arr []common.Type) ArrayNav
type BoolArrayNav interface {
	Get() bool
	// Has unexported methods.
}

func NewBoolArrayNav(arr []bool) BoolArrayNav
type BoolNavAttr struct {
	// Has unexported fields.
}

func (nav *BoolNavAttr) Get() bool
func (nav *BoolNavAttr) Len() int
func (nav *BoolNavAttr) Next() bool
func (nav *BoolNavAttr) Position() int
func (nav *BoolNavAttr) Prev() bool
type FloatArrayNav interface {
	Get() float64
	// Has unexported methods.
}

func NewFloatArrayNav(arr []float64) FloatArrayNav
type FloatNavAttr struct {
	// Has unexported fields.
}

func (nav *FloatNavAttr) Get() float64
func (nav *FloatNavAttr) Len() int
func (nav *FloatNavAttr) Next() bool
func (nav *FloatNavAttr) Position() int
func (nav *FloatNavAttr) Prev() bool
type IntArrayNav interface {
	Get() int
	// Has unexported methods.
}

func NewIntArrayNav(arr []int) IntArrayNav
type IntNavAttr struct {
	// Has unexported fields.
}

func (nav *IntNavAttr) Get() int
func (nav *IntNavAttr) Len() int
func (nav *IntNavAttr) Next() bool
func (nav *IntNavAttr) Position() int
func (nav *IntNavAttr) Prev() bool
type NavAttr struct {
	// Has unexported fields.
}

func (nav *NavAttr) Get() common.Type
func (nav *NavAttr) Len() int
func (nav *NavAttr) Next() bool
func (nav *NavAttr) Position() int
func (nav *NavAttr) Prev() bool
func (nav *NavAttr) Print()
