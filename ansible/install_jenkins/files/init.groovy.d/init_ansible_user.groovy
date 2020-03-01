import jenkins.model.*
import com.cloudbees.hudson.plugins.folder.*
import com.cloudbees.hudson.plugins.folder.properties.*
import com.cloudbees.hudson.plugins.folder.properties.FolderCredentialsProvider.FolderCredentialsProperty
import com.cloudbees.plugins.credentials.impl.*
import com.cloudbees.plugins.credentials.*
import com.cloudbees.plugins.credentials.domains.*
import com.cloudbees.plugins.credentials.common.*
import java.util.logging.Logger

final logger = Logger.getLogger("")
final jenkins = Jenkins.instance
final pass_env = "ANSIBLE_PASSWORD"

// From: https://github.com/hayderimran7/useful-jenkins-groovy-init-scripts/blob/master/useful_groovy_snippets.groovy#L107
def create_or_update_credentials = { username, password, description ->
    def global_domain = Domain.global()
    def credentials_store = Jenkins.instance
            .getExtensionList( 'com.cloudbees.plugins.credentials.SystemCredentialsProvider' )[0]
            .getStore()
    def credentials = new UsernamePasswordCredentialsImpl(CredentialsScope.GLOBAL, null, description, username, password)
    // From: https://github.com/hayderimran7/useful-jenkins-groovy-init-scripts/blob/master/useful_groovy_snippets.groovy#L14
    def existing_credentials = { username_param ->
            def username_matcher = CredentialsMatchers.withUsername(username_param)
            def available_credentials = CredentialsProvider.lookupCredentials(
                StandardUsernameCredentials.class,
                Jenkins.getInstance(),
                hudson.security.ACL.SYSTEM,
                new SchemeRequirement("ssh"))

            return CredentialsMatchers.firstOrNull(available_credentials, username_matcher)
        }(username)

    if(existing_credentials) {
        credentials_store.updateCredentials(global_domain, existing_credentials, credentials)
    } else {
        credentials_store.addCredentials(global_domain, credentials)
    }
}

def env = System.getenv()
if (!env["${pass_env}"]?.trim()) {
    logger.warn("Ansible user cannot be created. Password variable not found on this system.")
} else {
    final username = "ansible"
    final password = env["${pass_env}"]
    final description = "Credentials for SHH connection between Jenkins and Ansible clients"

    create_or_update_credentials(username, password, description)
}