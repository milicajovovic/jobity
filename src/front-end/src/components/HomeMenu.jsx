import React from "react";
import { Container, Nav, Navbar } from "react-bootstrap";

function HomeMenu() {
    return (
        <Navbar bg="dark" variant="dark">
            <Container>
                <Navbar.Brand href="/">Jobity</Navbar.Brand>
                <Nav className="text-uppercase">
                    <Nav.Link href="/register">Register</Nav.Link>
                    <Nav.Link href="/login">Log in</Nav.Link>
                </Nav>
            </Container>
        </Navbar>
    )
}

export default HomeMenu