import React from "react";
import { Container, Nav, Navbar } from "react-bootstrap";

function NavigationBar() {
    return (
        <Navbar bg="dark" variant="dark">
            <Container>
                <Navbar.Brand href="/">Jobity</Navbar.Brand>
                <Nav className="text-uppercase">
                    <Nav.Link href="/register">Register</Nav.Link>
                    <Nav.Link href="/login">Login</Nav.Link>
                </Nav>
            </Container>
        </Navbar>
    )
}

export default NavigationBar