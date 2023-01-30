import React from "react";
import { Container, Nav, Navbar } from "react-bootstrap";

function AdminMenu() {
    const logOut = () => {
        localStorage.setItem("jwt", "");
        localStorage.setItem("userId", "");
        localStorage.setItem("role", "");
    }

    return (
        <Navbar bg="dark" variant="dark">
            <Container>
                <Navbar.Brand href="/admin/employees">Jobity</Navbar.Brand>
                <Nav className="text-uppercase">
                    <Nav.Link href="/admin/employees">Employees</Nav.Link>
                    <Nav.Link href="/admin/employers">Employers</Nav.Link>
                    <Nav.Link href="/admin/ads">Ads</Nav.Link>
                    <Nav.Link href="/admin/reviews">Reviews</Nav.Link>
                    <Nav.Link href="/" onClick={logOut}>Log out</Nav.Link>
                </Nav>
            </Container>
        </Navbar>
    )
}

export default AdminMenu