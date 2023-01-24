import React from "react";
import { Col, Container, Row, Card, Button } from "react-bootstrap";
import NavigationBar from "../components/NavigationBar";

function RegisterHome() {
    return (
        <Container fluid>
            <Row>
                <NavigationBar />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "30rem" }}>
                        <Card.Title className="text-center mt-3 mb-5" as="h3">Welcome!</Card.Title>
                        <div className="d-grid gap-2">
                            <Button variant="primary" size="lg" href="/register/employee">
                                Register as employee
                            </Button>
                            <Button variant="primary" size="lg" className="mb-5" href="/register/employer">
                                Register as employer
                            </Button>
                        </div>
                    </Card>
                </Col>
            </Row>
        </Container>
    )
}

export default RegisterHome