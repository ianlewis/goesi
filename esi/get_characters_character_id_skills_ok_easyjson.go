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

func easyjson2298aef5DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdSkillsOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdSkillsOkList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdSkillsOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdSkillsOk
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
func easyjson2298aef5EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdSkillsOkList) {
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
func (v GetCharactersCharacterIdSkillsOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2298aef5EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSkillsOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2298aef5EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillsOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2298aef5DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillsOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2298aef5DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson2298aef5DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdSkillsOk) {
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
		case "skills":
			if in.IsNull() {
				in.Skip()
				out.Skills = nil
			} else {
				in.Delim('[')
				if out.Skills == nil {
					if !in.IsDelim(']') {
						out.Skills = make([]GetCharactersCharacterIdSkillsSkill, 0, 4)
					} else {
						out.Skills = []GetCharactersCharacterIdSkillsSkill{}
					}
				} else {
					out.Skills = (out.Skills)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdSkillsSkill
					easyjson2298aef5DecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.Skills = append(out.Skills, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "total_sp":
			out.TotalSp = int64(in.Int64())
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
func easyjson2298aef5EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdSkillsOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Skills) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"skills\":")
		if in.Skills == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Skills {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson2298aef5EncodeGithubComAntihaxGoesiEsi2(out, v6)
			}
			out.RawByte(']')
		}
	}
	if in.TotalSp != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"total_sp\":")
		out.Int64(int64(in.TotalSp))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdSkillsOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2298aef5EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSkillsOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2298aef5EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillsOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2298aef5DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillsOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2298aef5DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson2298aef5DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdSkillsSkill) {
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
		case "current_skill_level":
			out.CurrentSkillLevel = int32(in.Int32())
		case "skill_id":
			out.SkillId = int32(in.Int32())
		case "skillpoints_in_skill":
			out.SkillpointsInSkill = int64(in.Int64())
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
func easyjson2298aef5EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdSkillsSkill) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CurrentSkillLevel != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"current_skill_level\":")
		out.Int32(int32(in.CurrentSkillLevel))
	}
	if in.SkillId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"skill_id\":")
		out.Int32(int32(in.SkillId))
	}
	if in.SkillpointsInSkill != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"skillpoints_in_skill\":")
		out.Int64(int64(in.SkillpointsInSkill))
	}
	out.RawByte('}')
}