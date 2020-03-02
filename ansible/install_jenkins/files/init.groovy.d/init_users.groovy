import java.io.File

// From: https://github.com/hayderimran7/useful-jenkins-groovy-init-scripts/blob/master/useful_groovy_snippets.groovy#L33
users = [
    "admin": [email: "kergeodeta@gmail.com", full_name: "Halász Miklós"],
    "chemaxon": [email: "teszt.elek@teszt.com", full_name: "Teszt Elek"]
]

users.each { username, v ->
     email = v["email"]
     full_name = v["full_name"]

     def user = hudson.model.User.get(username)
     user.setFullName(full_name)

     def email_param = new hudson.tasks.Mailer.UserProperty(email)
     user.addProperty(email_param)

     def password = UUID.randomUUID().toString().replaceAll("-", "")
     def pw_param = hudson.security.HudsonPrivateSecurityRealm.Details.fromPlainPassword(password)
     user.addProperty(pw_param)

     println """

        ########################
        ${username}: ${password}

     """

     user.save()
}