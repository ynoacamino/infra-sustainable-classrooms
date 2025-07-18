// Code generated by goa v3.21.1, DO NOT EDIT.
//
// profiles HTTP server encoders and decoders
//
// Command:
// $ goa gen
// github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/design/api
// -o ./services/profiles/

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"

	profiles "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateStudentProfileResponse returns an encoder for responses returned
// by the profiles CreateStudentProfile endpoint.
func EncodeCreateStudentProfileResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*profiles.StudentProfileResponse)
		enc := encoder(ctx, w)
		body := NewCreateStudentProfileResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateStudentProfileRequest returns a decoder for requests sent to the
// profiles CreateStudentProfile endpoint.
func DecodeCreateStudentProfileRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateStudentProfileRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateStudentProfileRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			sessionToken string
			c            *http.Cookie
		)
		c, err = r.Cookie("session")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("session_token", "cookie"))
		} else {
			sessionToken = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateStudentProfilePayload(&body, sessionToken)

		return payload, nil
	}
}

// EncodeCreateStudentProfileError returns an encoder for errors returned by
// the CreateStudentProfile profiles endpoint.
func EncodeCreateStudentProfileError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid_input":
			var res profiles.InvalidInput
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "profile_already_exists":
			var res profiles.ProfileAlreadyExists
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "unauthorized":
			var res profiles.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateTeacherProfileResponse returns an encoder for responses returned
// by the profiles CreateTeacherProfile endpoint.
func EncodeCreateTeacherProfileResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*profiles.TeacherProfileResponse)
		enc := encoder(ctx, w)
		body := NewCreateTeacherProfileResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateTeacherProfileRequest returns a decoder for requests sent to the
// profiles CreateTeacherProfile endpoint.
func DecodeCreateTeacherProfileRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateTeacherProfileRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateTeacherProfileRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			sessionToken string
			c            *http.Cookie
		)
		c, err = r.Cookie("session")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("session_token", "cookie"))
		} else {
			sessionToken = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateTeacherProfilePayload(&body, sessionToken)

		return payload, nil
	}
}

// EncodeCreateTeacherProfileError returns an encoder for errors returned by
// the CreateTeacherProfile profiles endpoint.
func EncodeCreateTeacherProfileError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid_input":
			var res profiles.InvalidInput
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "profile_already_exists":
			var res profiles.ProfileAlreadyExists
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "unauthorized":
			var res profiles.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetCompleteProfileResponse returns an encoder for responses returned
// by the profiles GetCompleteProfile endpoint.
func EncodeGetCompleteProfileResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*profiles.CompleteProfileResponse)
		enc := encoder(ctx, w)
		body := NewGetCompleteProfileResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetCompleteProfileRequest returns a decoder for requests sent to the
// profiles GetCompleteProfile endpoint.
func DecodeGetCompleteProfileRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			sessionToken string
			err          error
			c            *http.Cookie
		)
		c, err = r.Cookie("session")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("session_token", "cookie"))
		} else {
			sessionToken = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetCompleteProfilePayload(sessionToken)

		return payload, nil
	}
}

// EncodeGetCompleteProfileError returns an encoder for errors returned by the
// GetCompleteProfile profiles endpoint.
func EncodeGetCompleteProfileError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "profile_not_found":
			var res profiles.ProfileNotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res profiles.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetPublicProfileByIDResponse returns an encoder for responses returned
// by the profiles GetPublicProfileById endpoint.
func EncodeGetPublicProfileByIDResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*profiles.PublicProfileResponse)
		enc := encoder(ctx, w)
		body := NewGetPublicProfileByIDResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetPublicProfileByIDRequest returns a decoder for requests sent to the
// profiles GetPublicProfileById endpoint.
func DecodeGetPublicProfileByIDRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			userID int64
			err    error

			params = mux.Vars(r)
		)
		{
			userIDRaw := params["user_id"]
			v, err2 := strconv.ParseInt(userIDRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("user_id", userIDRaw, "integer"))
			}
			userID = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetPublicProfileByIDPayload(userID)

		return payload, nil
	}
}

// EncodeGetPublicProfileByIDError returns an encoder for errors returned by
// the GetPublicProfileById profiles endpoint.
func EncodeGetPublicProfileByIDError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid_input":
			var res profiles.InvalidInput
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "profile_not_found":
			var res profiles.ProfileNotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateProfileResponse returns an encoder for responses returned by the
// profiles UpdateProfile endpoint.
func EncodeUpdateProfileResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*profiles.ProfileResponse)
		enc := encoder(ctx, w)
		body := NewUpdateProfileResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateProfileRequest returns a decoder for requests sent to the
// profiles UpdateProfile endpoint.
func DecodeUpdateProfileRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateProfileRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateProfileRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			sessionToken string
			c            *http.Cookie
		)
		c, err = r.Cookie("session")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("session_token", "cookie"))
		} else {
			sessionToken = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateProfilePayload(&body, sessionToken)

		return payload, nil
	}
}

// EncodeUpdateProfileError returns an encoder for errors returned by the
// UpdateProfile profiles endpoint.
func EncodeUpdateProfileError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid_input":
			var res profiles.InvalidInput
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "profile_not_found":
			var res profiles.ProfileNotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res profiles.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
