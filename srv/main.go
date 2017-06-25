package srv

import "context"

type Say struct {}

func (s *Say) Hello(ctx context.Context)