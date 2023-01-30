import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { ToastContainer, Toast, Form, Button, Container, Row, Col, Card } from "react-bootstrap";
import HomeMenu from "../components/HomeMenu";

function LoginAdmin() {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    const login = (event) => {
        event.preventDefault();

        const loginDto = {
            "Email": event.target.email.value,
            "Password": event.target.password.value
        };

        axios.post("http://localhost:3007/admins/login", loginDto).then(res => {
            localStorage.setItem("jwt", res.data.Jwt)
            localStorage.setItem("userId", res.data.UserId);
            localStorage.setItem("role", "admin");
            navigate("/admin/employees");
        }).catch((err) => {
            setMessage(err.response.data);
            setShow(true);
        });
    };

    return (
        <Container fluid>
            <ToastContainer position="top-center" className="text-center p-3">
                <Toast onClose={() => setShow(false)} show={show} delay={3000} autohide>
                    <Toast.Body>{message}</Toast.Body>
                </Toast>
            </ToastContainer>
            <Row>
                <HomeMenu />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "30rem" }}>
                        <Card.Title className="text-center mt-3 mb-3" as="h3">Login as employee</Card.Title>
                        <Form onSubmit={login}>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Email</Form.Label>
                                    <Form.Control type="text" name="email" placeholder="Enter email" required />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Password</Form.Label>
                                    <Form.Control type="password" name="password" placeholder="Enter password" required />
                                </Col>
                            </Row>
                            <Row>
                                <Col className="d-grid">
                                    <Button variant="primary" type="submit" className="mb-2">
                                        Log in
                                    </Button>
                                </Col>
                            </Row>
                        </Form>
                    </Card>
                </Col>
            </Row>
        </Container>
    )
}

export default LoginAdmin