// Code generated by "jade.go"; DO NOT EDIT.

package jade

import (
	"github.com/valyala/bytebufferpool"
)

func tpl_conditionals(buffer *bytebufferpool.ByteBuffer) {

	var user = struct {
		description, name string
		isAnonymous       bool
	}{description: "foo bar baz", name: "zxc"}
	var authorised = false
	buffer.WriteString(`<div id="user">`)
	if len(user.description) > 0 {
		buffer.WriteString(`<h2 class="green">Description</h2><p class="description">`)
		WriteEscString(user.description, buffer)
		buffer.WriteString(`</p>`)
	} else if authorised {
		buffer.WriteString(`<h2 class="blue">Description</h2><p class="description">      User has no description,
      why not add one...</p>`)

	} else {
		buffer.WriteString(`<h2 class="red">Description</h2><p class="description">User has no description</p>`)

	}
	buffer.WriteString(`</div>`)
	if !user.isAnonymous {
		buffer.WriteString(`<p>You're logged in as `)
		WriteEscString(user.name, buffer)
		buffer.WriteString(`</p>`)
	}
	if !user.isAnonymous {
		buffer.WriteString(`<p>You're logged in as `)
		WriteEscString(user.name, buffer)
		buffer.WriteString(`</p>`)
	}

}
