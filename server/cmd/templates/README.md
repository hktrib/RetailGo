# Template Usage

This repository contains templates for RetailGo server. Follow the instructions below to call the templates in your code.

## Prerequisites

- Make sure you have RetailGo server installed on your machine.

## Usage

### Open templates folder
This will depend on where you start the server. If you run it from the `server` directory then it will look something like so 

    //Open Templates folder
    tmpl, err_file := template.ParseGlob("./cmd/templates/*")
    
    //Handle errors
    if err_file != nil {
        // Handle error (e.g., file not found)
        http.Error(w, "Failed to open email template folder: "+err_file.Error(), http.StatusInternalServerError)
        return
    }

### Use template
Each template has specific parameters that are required to be populated. 

For example the `email_invitation.html` template has the following parameters:
- Action_url: URL as string,
- Sender_name: string,

You must send these parameters to the template like so 

	templateData := HtmlTemplate{
		Sender_name: UserObj.FirstName + " " + UserObj.LastName,
		Action_url:  url.URL,
	}'
Then you can execute the template wherfe htmlBody is a buffer of bytes that is the email content
	htmlBody := new(bytes.Buffer)
    err_io := tmpl.ExecuteTemplate(htmlBody, "email_invitation.html", templateData)

# Current templates

## Email Invitation
Parameters:
```Email : string```
```Name : string```
```Store_name : string```
```Sender_name : string```
```Action_url : URL```

Template Name : ```email_invitation.html```

## Onboarding Template
Parameters:
```Sender_name : string```
```Action_url : URL```

Template Name : ```onboarding_template.html```

## Onboarding Success
Parameters:
```Sender_name : string```
```Action_url : URL```

Template Name: ```onboarding_success.html```