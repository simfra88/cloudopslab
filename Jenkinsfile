pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git branch: 'produccion', url: 'https://github.com/simfra88/cloudopslab.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t cloudopslab .'
            }
        }

        stage('Deploy Container') {
            steps {
                sh '''
                docker stop cloudopslab || true
                docker rm cloudopslab || true
                docker run -d -p 8080:8080 --name cloudopslab cloudopslab
                '''
            }
        }
    }
}
