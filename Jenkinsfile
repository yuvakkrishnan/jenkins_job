pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIALS = credentials('dockerhub-credentials') // Jenkins credential ID
        DOCKER_IMAGE_NAME = "yuvakkrishnans/golang-simple-service" // Your Docker Hub username/repo
    }

    stages {
        // REMOVED: stage('Checkout') block is no longer needed
        // Jenkins automatically checks out the repository when using "Pipeline script from SCM"

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
                    withCredentials([usernamePassword(credentialsId: 'dockerhub-credentials', passwordVariable: 'DOCKER_PASSWORD', usernameVariable: 'DOCKER_USERNAME')]) {
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
                    // Stop and remove existing container if it exists
                    sh "docker stop golang-service || true"
                    sh "docker rm golang-service || true"
                    // Run the new container
                    sh "docker run -d -p 8080:8080 --name golang-service ${DOCKER_IMAGE_NAME}:${BUILD_NUMBER}"
                }
            }
        }
    }
    post {
        always {
            cleanWs() // Clean up workspace
        }
        success {
            echo 'Pipeline finished successfully!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
