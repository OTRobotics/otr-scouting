node {
    withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}", "GOROOT=/usr/local/go", "CLOUDSDK_CORE_PROJECT=otr-scouting"]) {
        env.PATH="${GOPATH}/bin:${GOROOT}/bin:$PATH"
        stage('Pre Test'){
            echo 'Pulling Dependencies'
            sh 'go version'
            sh 'go get .'
            sh 'go get -u github.com/otrobotics/otr-scouting'
        }

        stage('Compile Match Uploader') {
            echo 'MatchUploader Build for All OSs'
            sh 'cd matchUploader && export GOOS=windows && go build main.go'
            sh 'cd matchUploader && export GOOS=darwin && go build main.go'
            sh 'cd matchUploader && export GOOS=linux && go build main.go'
        }
        stage('Deploy') {
            withCredentials([file(credentialsId: 'otr-scouting-appengine', variable: 'GCP_CREDS')]) {
                echo 'Deploying to GCP'
                sh "/var/jenkins_home/google-cloud-sdk/bin/gcloud auth activate-service-account --key-file $GCP_CREDS"
                sh 'cd appengine && /var/jenkins_home/google-cloud-sdk/bin/gcloud app deploy'
            }
        }
        stage('Archive Match Uploader') {
            archiveArtifacts allowEmptyArchive: true, artifacts: 'matchUploader/main*', excludes: 'matchUploader/main.go'
        }
    }
}