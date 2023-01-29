import React from "react";
import { Col, Container, Row, Card, Button } from "react-bootstrap";
import HomeMenu from "../components/HomeMenu";

function LoginHome() {
    return (
        <Container fluid>
            <Row>
                <HomeMenu />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "30rem" }}>
                        <Card.Title className="text-center mt-3 mb-5" as="h3">Welcome back!</Card.Title>
                        <div className="d-grid gap-2">
                            <Button variant="primary" size="lg" href="/login/employee">
                                Login as employee
                            </Button>
                            <Button variant="primary" size="lg" href="/login/employer">
                                Login as employer
                            </Button>
                            <Button variant="primary" size="lg" className="mb-5" href="/login/admin">
                                Login as admin
                            </Button>
                        </div>
                    </Card>
                </Col>
            </Row>
        </Container>
    )
}

export default LoginHome