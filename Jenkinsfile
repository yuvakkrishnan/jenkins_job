pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIALS = credentials('dockerhub-credentials')
        DOCKER_IMAGE_NAME = "yuvakkrishnans/golang-simple-service"
    }

    stages {
        stage('Build Docker Image') {
            steps {
                script {
                    sh "docker build -t ${DOCKER_IMAGE_NAME}:${BUILD_NUMBER} ."
                }
            }
        }

        stage('Login to Docker Hub') {
            steps {
                script {
                    withCredentials([usernamePassword(
                        credentialsId: 'dockerhub-credentials',
                        usernameVariable: 'DOCKER_USERNAME',
                        passwordVariable: 'DOCKER_PASSWORD'
                    )]) {
                        sh "echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin"
                    }
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    sh "docker push ${DOCKER_IMAGE_NAME}:${BUILD_NUMBER}"
                    sh "docker tag ${DOCKER_IMAGE_NAME}:${BUILD_NUMBER} ${DOCKER_IMAGE_NAME}:latest"
                    sh "docker push ${DOCKER_IMAGE_NAME}:latest"
                }
            }
        }

        stage('Run Docker Container') {
            steps {
                script {
                    sh "docker stop golang-service || true"
                    sh "docker rm golang-service || true"
                    // host port 8081 maps to container port 8080
                    sh "docker run -d -p 8081:8080 --name golang-service ${DOCKER_IMAGE_NAME}:${BUILD_NUMBER}"
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
        success {
            echo 'Pipeline finished successfully!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
