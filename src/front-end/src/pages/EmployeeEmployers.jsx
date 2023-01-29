import axios from "axios";
import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Col, Container, Row, Card, Button, Table } from "react-bootstrap";
import EmployeeMenu from "../components/EmployeeMenu";
import ReviewForm from "../components/ReviewForm";

function EmployeeEmployers() {
    const [employers, setEmployers] = useState([]);
    const [reviewed, setReviewed] = useState({});
    const [showReview, setShowReview] = useState(false);

    useEffect(() => {
        axios.get("http://localhost:3007/employers").then(res => {
            setEmployers(res.data);
        });
    }, []);

    const addReview = (employer) => {
        setReviewed(employer);
        setShowReview(true);
    }

    return (
        <Container fluid>
            <Row>
                <EmployeeMenu />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "30rem" }}>
                        <Card.Title className="text-center mt-3" as="h3">My employers</Card.Title>
                        <Card.Body className="text-center">
                            <Table striped bordered>
                                <thead>
                                    <tr>
                                        <th>Name</th>
                                        <th>Address</th>
                                        <th>Review</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {employers.map(employer => (
                                        <tr key={employer.ID}>
                                            <td>{employer.Name}</td>
                                            <td>{employer.Address}</td>
                                            <td><Button onClick={() => addReview(employer)}>Add</Button></td>
                                        </tr>
                                    ))}
                                </tbody>
                            </Table>
                        </Card.Body>
                    </Card>
                </Col>
            </Row>
            <ReviewForm show={showReview} onHide={() => setShowReview(false)} employer={reviewed} close={setShowReview} />
        </Container>
    )
}

export default EmployeeEmployers