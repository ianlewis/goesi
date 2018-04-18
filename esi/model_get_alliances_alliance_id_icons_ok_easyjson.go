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

func easyjson17d0f88bDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetAlliancesAllianceIdIconsOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetAlliancesAllianceIdIconsOkList, 0, 2)
			} else {
				*out = GetAlliancesAllianceIdIconsOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetAlliancesAllianceIdIconsOk
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
func easyjson17d0f88bEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetAlliancesAllianceIdIconsOkList) {
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
func (v GetAlliancesAllianceIdIconsOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson17d0f88bEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdIconsOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson17d0f88bEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdIconsOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson17d0f88bDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdIconsOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson17d0f88bDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson17d0f88bDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetAlliancesAllianceIdIconsOk) {
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
		case "px64x64":
			out.Px64x64 = string(in.String())
		case "px128x128":
			out.Px128x128 = string(in.String())
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
func easyjson17d0f88bEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetAlliancesAllianceIdIconsOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Px64x64 != "" {
		const prefix string = ",\"px64x64\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Px64x64))
	}
	if in.Px128x128 != "" {
		const prefix string = ",\"px128x128\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Px128x128))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetAlliancesAllianceIdIconsOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson17d0f88bEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdIconsOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson17d0f88bEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdIconsOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson17d0f88bDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdIconsOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson17d0f88bDecodeGithubComAntihaxGoesiEsi1(l, v)
}