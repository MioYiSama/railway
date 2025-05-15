plugins {
    alias(libs.plugins.kotlin.jvm)
    alias(libs.plugins.kotlin.serialization)
    alias(libs.plugins.quarkus)
}

dependencies {
    implementation(enforcedPlatform(libs.quarkus.bom))
    implementation(libs.quarkus.kotlin)
    implementation(libs.quarkus.rest.kotlin.serialization)
    implementation(libs.bundles.quarkus.jwt)

    testImplementation(libs.bundles.quarkus.test)
    testImplementation(libs.kotlin.test)
}
