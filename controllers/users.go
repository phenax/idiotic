
package controllers;

import (

);


func ProfilePage(ctx *Context) {

	var title string;

	if(ctx.params["name"] != "") {
		title = "<h1>Hey, " + ctx.params["name"] + "</h1>";
	} else {
		title = "<h1>This is cool</h1>";
	}

	ctx.Send(title);
}
