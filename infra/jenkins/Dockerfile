FROM jenkins/jenkins:alpine

ENV JENKINS_USER admin \
    JENKINS_PASS admin \
    JAVA_OPTS -Djenkins.install.runSetupWizard=false

ADD --chown=jenkins plugins.txt /usr/share/jenkins/plugins.txt

RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/plugins.txt
