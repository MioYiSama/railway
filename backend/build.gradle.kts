plugins {
    idea
    alias(libs.plugins.kotlin.jvm) apply false
}

idea {
    module {
        isDownloadSources = true
    }
}

allprojects {
    group = "net.mioyi.railway"
    version = "1.0.0"

    repositories {
        mavenCentral()
    }
}