import axios from "axios";
import React, { useEffect, useState } from "react";
import { Col, Container, Row, Card, ToastContainer, Toast, Form, Button } from "react-bootstrap";
import EmployerMenu from "../components/EmployerMenu";

function EmployerProfile() {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    const [employer, setEmployer] = useState({});

    useEffect(() => {
        const jwt = localStorage.getItem("jwt");
        const userId = localStorage.getItem("userId");
        axios.get("http://localhost:3007/employers/employer/" + userId, { headers: { Authorization: jwt } }).then(res => {
            setEmployer(res.data);
        });
    }, []);

    const update = (event) => {
        event.preventDefault();

        let password = event.target.password.value;
        if (password === "Password") {
            password = employer.Password;
        }

        const updatedEmployer = {
            "ID": employer.ID,
            "Email": event.target.email.value,
            "Password": password,
            "Name": event.target.name.value,
            "Address": event.target.address.value,
        };

        let jwt = localStorage.getItem("jwt");

        axios.post("http://localhost:3007/employers/update", updatedEmployer, { headers: { Authorization: jwt } }).then(res => {
            setMessage("successfully updated");
            setShow(true);
        }).catch((err) => {
            setMessage(err.response.data);
            setShow(true);
        });
    };

    return (
        <Container fluid>
            <Row>
                <EmployerMenu />
            </Row>
            <ToastContainer position="top-center" className="text-center p-3">
                <Toast onClose={() => setShow(false)} show={show} delay={3000} autohide>
                    <Toast.Body>{message}</Toast.Body>
                </Toast>
            </ToastContainer>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "40rem" }}>
                        <Card.Title className="text-center mt-3 mb-3" as="h3">My profile</Card.Title>
                        <Form onSubmit={update}>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Email</Form.Label>
                                    <Form.Control type="text" name="email" defaultValue={employer.Email} disabled />
                                </Col>
                                <Col>
                                    <Form.Label>Password</Form.Label>
                                    <Form.Control type="password" name="password" defaultValue="Password" required />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Name</Form.Label>
                                    <Form.Control type="text" name="name" defaultValue={employer.Name} required />
                                </Col>
                                <Col>
                                    <Form.Label>Address</Form.Label>
                                    <Form.Control type="text" name="address" defaultValue={employer.Address} required />
                                </Col>
                            </Row>
                            <Row>
                                <Col className="d-grid">
                                    <Button variant="primary" type="submit" className="mb-2">
                                        Save changes
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

export default EmployerProfile