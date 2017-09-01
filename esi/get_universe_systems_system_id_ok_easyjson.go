// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB8c84fd0DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseSystemsSystemIdOkList, 0, 1)
			} else {
				*out = GetUniverseSystemsSystemIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseSystemsSystemIdOk
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB8c84fd0EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseSystemsSystemIdOkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseSystemsSystemIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB8c84fd0EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemsSystemIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB8c84fd0EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB8c84fd0DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB8c84fd0DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonB8c84fd0DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdOk) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "constellation_id":
			out.ConstellationId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "planets":
			if in.IsNull() {
				in.Skip()
				out.Planets = nil
			} else {
				in.Delim('[')
				if out.Planets == nil {
					if !in.IsDelim(']') {
						out.Planets = make([]GetUniverseSystemsSystemIdPlanet, 0, 2)
					} else {
						out.Planets = []GetUniverseSystemsSystemIdPlanet{}
					}
				} else {
					out.Planets = (out.Planets)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetUniverseSystemsSystemIdPlanet
					(v4).UnmarshalEasyJSON(in)
					out.Planets = append(out.Planets, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "position":
			easyjsonB8c84fd0DecodeGithubComAntihaxGoesiEsi2(in, &out.Position)
		case "security_class":
			out.SecurityClass = string(in.String())
		case "security_status":
			out.SecurityStatus = float32(in.Float32())
		case "star_id":
			out.StarId = int32(in.Int32())
		case "stargates":
			if in.IsNull() {
				in.Skip()
				out.Stargates = nil
			} else {
				in.Delim('[')
				if out.Stargates == nil {
					if !in.IsDelim(']') {
						out.Stargates = make([]int32, 0, 16)
					} else {
						out.Stargates = []int32{}
					}
				} else {
					out.Stargates = (out.Stargates)[:0]
				}
				for !in.IsDelim(']') {
					var v5 int32
					v5 = int32(in.Int32())
					out.Stargates = append(out.Stargates, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "stations":
			if in.IsNull() {
				in.Skip()
				out.Stations = nil
			} else {
				in.Delim('[')
				if out.Stations == nil {
					if !in.IsDelim(']') {
						out.Stations = make([]int32, 0, 16)
					} else {
						out.Stations = []int32{}
					}
				} else {
					out.Stations = (out.Stations)[:0]
				}
				for !in.IsDelim(']') {
					var v6 int32
					v6 = int32(in.Int32())
					out.Stations = append(out.Stations, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "system_id":
			out.SystemId = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB8c84fd0EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseSystemsSystemIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConstellationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"constellation_id\":")
		out.Int32(int32(in.ConstellationId))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if len(in.Planets) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"planets\":")
		if in.Planets == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Planets {
				if v7 > 0 {
					out.RawByte(',')
				}
				(v8).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"position\":")
		easyjsonB8c84fd0EncodeGithubComAntihaxGoesiEsi2(out, in.Position)
	}
	if in.SecurityClass != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"security_class\":")
		out.String(string(in.SecurityClass))
	}
	if in.SecurityStatus != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"security_status\":")
		out.Float32(float32(in.SecurityStatus))
	}
	if in.StarId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"star_id\":")
		out.Int32(int32(in.StarId))
	}
	if len(in.Stargates) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"stargates\":")
		if in.Stargates == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Stargates {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v10))
			}
			out.RawByte(']')
		}
	}
	if len(in.Stations) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"stations\":")
		if in.Stations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Stations {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v12))
			}
			out.RawByte(']')
		}
	}
	if in.SystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"system_id\":")
		out.Int32(int32(in.SystemId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseSystemsSystemIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB8c84fd0EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemsSystemIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB8c84fd0EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB8c84fd0DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB8c84fd0DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjsonB8c84fd0DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdPosition) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x":
			out.X = float32(in.Float32())
		case "y":
			out.Y = float32(in.Float32())
		case "z":
			out.Z = float32(in.Float32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB8c84fd0EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetUniverseSystemsSystemIdPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"x\":")
		out.Float32(float32(in.X))
	}
	if in.Y != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"y\":")
		out.Float32(float32(in.Y))
	}
	if in.Z != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"z\":")
		out.Float32(float32(in.Z))
	}
	out.RawByte('}')
}
