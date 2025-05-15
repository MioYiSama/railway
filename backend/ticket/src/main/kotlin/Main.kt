package net.mioyi.railway.ticket

import jakarta.ws.rs.GET
import jakarta.ws.rs.Path

@Path("/")
class MainResource {
    @GET
    fun index() = "ticket"
}