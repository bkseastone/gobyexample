package rpcdemo

type CalcService struct{}

type Args struct {
	A, B int
}

func (CalcService) Add(args Args, result *int) error {
	sum := args.A + args.B
	*result = sum
	return nil
}
