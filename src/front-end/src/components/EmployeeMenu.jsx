import React from "react";
import { Container, Nav, Navbar } from "react-bootstrap";

function EmployeeMenu() {
    const logOut = () => {
        localStorage.setItem("jwt", "");
        localStorage.setItem("userId", "");
        localStorage.setItem("role", "");
    }

    return (
        <Navbar bg="dark" variant="dark">
            <Container>
                <Navbar.Brand href="/employee/home">Jobity</Navbar.Brand>
                <Nav className="text-uppercase">
                    <Nav.Link href="/employee/profile">Profile</Nav.Link>
                    <Nav.Link href="/employee/employers">Employers</Nav.Link>
                    <Nav.Link href="/" onClick={logOut}>Log out</Nav.Link>
                </Nav>
            </Container>
        </Navbar>
    )
}

export default EmployeeMenu