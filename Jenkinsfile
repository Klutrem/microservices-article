APPNAME = 'appname'

podTemplate(containers: [
    containerTemplate(name: 'docker', image: 'docker:dind', ttyEnabled: true, command: 'cat', privileged: true),
    containerTemplate(name: "golang", image: "klutrem/golang-java:1.1", command: "sleep", args: "99d", NodeSelector: "kubernetes.io/os=linux")
  ],
  volumes: [
    hostPathVolume(hostPath: '/var/run/docker.sock', mountPath: '/var/run/docker.sock')
  ], hostNetwork: true)  {
    if (BRANCH_NAME.contains('PR')) {
        BRANCH = CHANGE_BRANCH
    } else {
        BRANCH = BRANCH_NAME
    }
    node(POD_LABEL) {
      if (env.BRANCH_NAME == "main") {
        ENV = "prod"
      } else {
        ENV = env.BRANCH_NAME
      }
      stage('pull') {
        withCredentials([
          usernamePassword(
            credentialsId: 'GitBackendCreds',
            usernameVariable: 'GIT_LOGIN',
            passwordVariable: 'GIT_PASSWORD',
          )
        ]) 
        {
          sh "printenv"
          REPO_URL = "${env.GIT_SCHEME}$GIT_LOGIN:$GIT_PASSWORD@${env.GIT_HOST}/${env.GITEA_KAISER_ORG}/${APPNAME}.git"
          echo "Pulling $REPO_URL[$BRANCH_NAME]..."
          git url: "$REPO_URL", branch: BRANCH
        }
      }
      container("golang") {
      stage("build") {
        sh "go get ./..."
      }
      stage("test") {
        sh "go test ./... -coverprofile=coverage.out"
        }
      }
      if(BRANCH == 'develop' || BRANCH == 'staging' || BRANCH == 'main') {
      container('golang') {
       stage('SonarQube Analysis') {
        def scannerHome = tool 'main';
          withSonarQubeEnv() {
            sh "${scannerHome}/bin/sonar-scanner"
            sh "rm -rf coverage"
        }
       }
      }
      stage('build docker image') {
        container('docker') {
          IMG = "${env.DOCKER_REGISTRY_HOST}/${env.DOCKER_REGISTRY_ORGANISATION_NAME}/${APPNAME}_${ENV}:${BUILD_NUMBER}"
      		echo "Pushing $IMG to docker registry ${env.DOCKER_REGISTRY_HOST}"
          sh "docker build -t ${IMG} ."
        }
      }
      stage("push image") {
        withCredentials([
          usernamePassword(
            credentialsId: 'DockerRegistryCreds',
            usernameVariable: 'DOCKER_LOGIN',
            passwordVariable: 'DOCKER_PASSWORD',
          )
        ]) 
        {
          container('docker') {
            sh "docker login ${env.DOCKER_REGISTRY_HOST} -u $DOCKER_LOGIN -p $DOCKER_PASSWORD"
            sh "docker push $IMG"
          }
        }
      }
      stage ('cleanup'){
            cleanWs()
      }
    }
  }
}
  
  