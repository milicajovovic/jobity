import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { ToastContainer, Toast, Form, Button, Container, Row, Col, Card } from "react-bootstrap";
import NavigationBar from "../components/NavigationBar";

function RegisterEmployer() {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    const register = (event) => {
        event.preventDefault();

        const newEmployer = {
            "Email": event.target.email.value,
            "Password": event.target.password.value,
            "Name": event.target.name.value,
            "Address": event.target.address.value,
            "ProfilePicture": event.target.profilePicture.value,
        };

        if (validEmail(newEmployer.Email)) {
            axios.post("http://localhost:3003/register", newEmployer).then(res => {
                navigate("/home/employer");
            }).catch((err) => {
                setMessage(err.response.data);
                setShow(true);
            });
        } else {
            setMessage("email is not valid");
            setShow(true);
        }
    }

    const validEmail = (email) => {
        let emailRegex = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/;
        if (email.match(emailRegex)) {
            return true;
        }
        return false;
    }

    return (
        <Container fluid>
            <ToastContainer position="top-center" className="text-center p-3">
                <Toast onClose={() => setShow(false)} show={show} delay={3000} autohide>
                    <Toast.Body>{message}</Toast.Body>
                </Toast>
            </ToastContainer>
            <Row>
                <NavigationBar />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "40rem" }}>
                        <Card.Title className="text-center mt-3 mb-3" as="h3">Register as employer</Card.Title>
                        <Form onSubmit={register}>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Email</Form.Label>
                                    <Form.Control type="text" name="email" placeholder="Enter email" required />
                                </Col>
                                <Col>
                                    <Form.Label>Password</Form.Label>
                                    <Form.Control type="password" name="password" placeholder="Enter password" required />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Name</Form.Label>
                                    <Form.Control type="text" name="name" placeholder="Enter name" required />
                                </Col>
                                <Col>
                                    <Form.Label>Profile picture</Form.Label>
                                    <Form.Control type="file" name="profilePicture" />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Address</Form.Label>
                                    <Form.Control type="text" name="address" placeholder="Enter address" required />
                                </Col>
                            </Row>
                            <Row>
                                <Col className="d-grid">
                                    <Button variant="primary" type="submit" className="mb-2">
                                        Sign in
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

export default RegisterEmployer