pipeline {
    agent any

    stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Go Version') {
            steps {
                sh '/usr/local/go/bin/go version'
            }
        }

        stage('Build') {
            steps {
                sh '/usr/local/go/bin/go build -o kubeops'
            }
        }

        stage('CLI Test') {
            steps {
                sh './kubeops version'
            }
        }

        stage('Docker Build & Push') {
    environment {
        IMAGE_NAME = "nikithatellatakula7685678587/kubeops-cli"
    }

    steps {
        withCredentials([usernamePassword(
            credentialsId: 'dockerhub',
            usernameVariable: 'DOCKER_USER',
            passwordVariable: 'DOCKER_PASS'
        )]) {

            sh '''
                echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin

                docker build -t $IMAGE_NAME:latest .

                docker push $IMAGE_NAME:latest

                docker logout
            '''
        }
    }
}

    post {
        success {
            echo 'Pipeline completed successfully!'
        }

        failure {
            echo 'Pipeline failed!'
        }
    }
}