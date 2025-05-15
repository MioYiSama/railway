package net.mioyi.railway.user

import jakarta.ws.rs.GET
import jakarta.ws.rs.Path

@Path("/")
class MainResource {
    @GET
    fun index() = "user"
}