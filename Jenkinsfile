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

        stage('Docker Build') {
            steps {
                sh 'docker build -t kubeops-cli:latest .'
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