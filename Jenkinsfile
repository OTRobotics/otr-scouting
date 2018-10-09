pipeline {
    agent any

    stages {
     
        stage('Compile Match Uploader') {
            steps {
                sh 'export GOROOT=/usr/loca/go'
                sh 'export GOHOME=/var/jenkins_home/go'
                sh 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH'
                echo 'MatchUploader Build for All OSs'
                sh 'cd matchUploader && export GOOS=windows && go build main.go'
                sh 'cd matchUploader && export GOOS=darwin && go build main.go'
                sh 'cd matchUploader && export GOOS=linux && go build main.go'
            }
        }
        stage('Deploy') {
            environment {
                GCP_APPENGINE = credentials('otr-scouting-appengine')
            }
            steps {
                echo 'Deploying to GCP'
                sh "/var/jenkins_home/google-cloud-sdk/bin/gcloud auth activate-service-account --key-file ${env.GCP_APPENGINE}"
                sh 'cd appengine && gcloud app deploy'
            }
        }
        stage('Upload MatchUploader to GCS') {
            environment {
                GCS_CREDS = credentials('otr-scouting-web-gcloud')
            }
            steps {
                sh "/var/jenkins_home/google-cloud-sdk/bin/gcloud auth activate-service-account --key-file ${env.GCS_CREDS}"
                sh '/var/jenkins_home/google-cloud-sdk/bin/gsutil cp *.pdf gs://staging.otr-scouting.appspot.com/matchUploader/'
            }
        }
    }
}