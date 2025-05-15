plugins {
    alias(libs.plugins.kotlin.jvm)
    alias(libs.plugins.kotlin.serialization)
    alias(libs.plugins.quarkus)
}

dependencies {
    implementation(enforcedPlatform(libs.quarkus.bom))
    implementation(libs.bundles.quarkus.kotlin.rest)
    implementation(libs.quarkus.smallrye.jwt)

    testImplementation(libs.bundles.quarkus.kotlin.rest.test)
}
