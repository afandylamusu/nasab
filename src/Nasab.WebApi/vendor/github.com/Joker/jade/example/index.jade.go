// Code generated by "jade.go"; DO NOT EDIT.

package main

import (
	"bytes"

	"github.com/Joker/hpp"
	pool "github.com/valyala/bytebufferpool"
)

const (
	index__0 = `<!DOCTYPE html><html lang="en"><head><title>`
	index__1 = `</title><script type="text/javascript">			if(question){
				answer(40 + 2)
			}</script></head><body><h1>Jade - template engine`
	index__2 = `</h1><div id="container" class="col">`
	index__3 = `<p>				Jade/Pug is a terse and simple
				templating language with
				a <strong>focus</strong> on performance 
				and powerful features.</p></div><footer><div class="footer">2018</div></footer></body></html>`
	index__4 = `<div id="cmd">Precompile jade templates to `
	index__5 = ` code.</div>`
	index__6 = `<p>You are amazing</p>`
	index__7 = `<p>Get on it!</p>`
)

func Index(pageTitle string, youAreUsingJade bool, buffer *pool.ByteBuffer) {

	buffer.WriteString(index__0)
	WriteEscString(pageTitle, buffer)
	buffer.WriteString(index__1)

	{
		var (
			golang = "Go"
		)

		buffer.WriteString(index__4)
		WriteEscString(golang, buffer)
		buffer.WriteString(index__5)
	}

	buffer.WriteString(index__2)

	if youAreUsingJade {
		buffer.WriteString(index__6)

	} else {
		buffer.WriteString(index__7)

	}
	buffer.WriteString(index__3)

	index__buffer := hpp.Print(bytes.NewReader(buffer.Bytes()))
	buffer.Reset()
	buffer.Write(index__buffer)

}
