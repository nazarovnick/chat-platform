package chat

import (
	validateproto "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"errors"
	validateimpl "github.com/bufbuild/protovalidate-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func convertImplViolationsToProtoViolations(vs []*validateimpl.Violation) []*validateproto.Violation {
	result := make([]*validateproto.Violation, len(vs))
	for i, v := range vs {
		var msg string
		if v.Proto != nil && v.Proto.Message != nil {
			msg = *v.Proto.Message
		}
		result[i] = &validateproto.Violation{
			Field:   v.Proto.Field,
			Message: &msg,
		}
	}

	return result
}

func protovalidateVialationsToGoogleViolations(vs []*validateproto.Violation) []*errdetails.BadRequest_FieldViolation {
	res := make([]*errdetails.BadRequest_FieldViolation, len(vs))
	for i, v := range vs {
		desc := ""
		if v.Message != nil {
			desc = *v.Message
		}
		res[i] = &errdetails.BadRequest_FieldViolation{
			Field:       v.Field.String(),
			Description: desc,
		}
	}
	return res
}

func convertProtovalidateValidationErrorToErrdetailsBadRequest(valErr *validateimpl.ValidationError) *errdetails.BadRequest {
	return &errdetails.BadRequest{
		FieldViolations: protovalidateVialationsToGoogleViolations(
			convertImplViolationsToProtoViolations(valErr.Violations)),
	}
}

func rpcValidationError(err error) error {
	if err == nil {
		return nil
	}

	var valErr *validateimpl.ValidationError
	if ok := errors.As(err, &valErr); ok {
		st, err := status.New(codes.InvalidArgument, codes.InvalidArgument.String()).
			WithDetails(convertProtovalidateValidationErrorToErrdetailsBadRequest(valErr))
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		return st.Err()
	}

	return status.Error(codes.Internal, err.Error())
}
