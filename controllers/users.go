
package controllers;

import (

);


func ProfilePage(ctx *Context) {

	var title string;

	if(ctx.Params["name"] != "") {
		title = "<h1>Hey, " + ctx.Params["name"] + "</h1>";
	} else {
		title = "<h1>This is cool</h1>";
	}

	ctx.Send(title, &ResponseConfig{
		ContentType: "text/html",
	});
}
