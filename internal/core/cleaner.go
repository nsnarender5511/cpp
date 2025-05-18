package core

import (
	"bytes"
	"go/scanner"
	"go/token"
)

func RemoveComments(src []byte) ([]byte, error) {
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	var out bytes.Buffer
	var lastPos token.Pos

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}

		
		if file.Offset(pos) > file.Offset(lastPos) {
			out.Write(src[file.Offset(lastPos):file.Offset(pos)])
		}

		if tok != token.COMMENT {
			if lit != "" {
				out.WriteString(lit)
			} else {
				out.WriteString(tok.String())
			}
		}
		lastPos = pos + token.Pos(len(lit))
		if tok == token.ILLEGAL && lit == "" { 
			lastPos++
		} else if lit == "" && tok.String() == "" && tok != token.SEMICOLON { 
			
			
			
			lastPos = pos + 1
		} else if lit == "" && tok != token.SEMICOLON {
			lastPos = pos + token.Pos(len(tok.String()))
		}

	}
	
	if file.Offset(lastPos) < len(src) {
		out.Write(src[file.Offset(lastPos):])
	}

	return out.Bytes(), nil
}
