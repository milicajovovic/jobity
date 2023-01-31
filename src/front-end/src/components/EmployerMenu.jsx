import React from "react";
import { Container, Nav, Navbar } from "react-bootstrap";

function EmployerMenu() {
    const logOut = () => {
        localStorage.setItem("jwt", "");
        localStorage.setItem("userId", "");
        localStorage.setItem("role", "");
    }

    return (
        <Navbar bg="dark" variant="dark">
            <Container>
                <Navbar.Brand href="/employer/home">Jobity</Navbar.Brand>
                <Nav className="text-uppercase">
                    <Nav.Link href="/employer/profile">Profile</Nav.Link>
                    <Nav.Link href="/employer/ad">New add</Nav.Link>
                    <Nav.Link href="/employer/reviews">Reviews</Nav.Link>
                    <Nav.Link href="/employer/applications">Applications</Nav.Link>
                    <Nav.Link href="/employer/interviews">Interviews</Nav.Link>
                    <Nav.Link href="/" onClick={logOut}>Log out</Nav.Link>
                </Nav>
            </Container>
        </Navbar>
    )
}

export default EmployerMenu