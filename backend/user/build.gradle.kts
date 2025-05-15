plugins {
    alias(libs.plugins.kotlin.jvm)
    alias(libs.plugins.quarkus)
}

dependencies {
    implementation(enforcedPlatform(libs.quarkus.bom))
    implementation(libs.quarkus.kotlin)
    implementation(libs.bundles.quarkus.jwt)
}
