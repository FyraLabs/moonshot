#!/usr/bin/env osascript -l JavaScript

// Standing on the sholders of giants :)
// This is all we borrow, huge thanks to the Etcher team... even though we're trying to replace it :P
// https://github.com/balena-io/etcher/blob/master/lib/shared/sudo/sudo-askpass.osascript-en.js

ObjC.import("stdlib");

const app = Application.currentApplication();
app.includeStandardAdditions = true;

const result = app.displayDialog(
  "Moonshot needs privileged access in order to flash your drive.\n\nType your password to allow this.",
  {
    defaultAnswer: "",
    withIcon: "caution",
    buttons: ["Cancel", "Ok"],
    defaultButton: "Ok",
    hiddenAnswer: true,
  },
);

if (result.buttonReturned === "Ok") {
  result.textReturned;
} else {
  $.exit(255);
}
