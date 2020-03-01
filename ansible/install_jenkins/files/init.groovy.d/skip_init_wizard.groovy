// From: https://riptutorial.com/jenkins/example/24925/disable-setup-wizard
import jenkins.model.*
import hudson.util.*;
import jenkins.install.*;

def instance = Jenkins.getInstance()

instance.setInstallState(InstallState.INITIAL_SETUP_COMPLETED)