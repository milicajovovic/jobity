import axios from "axios";
import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Col, Container, Row, Card, Button, Table } from "react-bootstrap";
import EmployerMenu from "../components/EmployerMenu";
import Confirmation from "../components/Confirmation";

function EmployerReviews() {
    const [reviews, setReviews] = useState([]);
    const [question, setQuestion] = useState("");
    const [id, setId] = useState();
    const [action, setAction] = useState("");
    const [show, setShow] = useState(false);

    useEffect(() => {
        let jwt = localStorage.getItem("jwt");
        let userId = localStorage.getItem("userId");
        axios.get("http://localhost:3007/reviews/employer/" + userId, { headers: { Authorization: jwt } }).then(res => {
            setReviews(res.data);
        });
    }, [reviews]);

    const inappropriate = (review) => {
        setQuestion("Are you sure you want to mark this review as inappropriate?");
        setId(review.ID);
        setAction("reviews/inappropriate");
        setShow(true);
    };

    return (
        <Container fluid>
            <Row>
                <EmployerMenu />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "50rem" }}>
                        <Card.Title className="text-center mt-3" as="h3">Reviews</Card.Title>
                        <Card.Body className="text-center">
                            <Table striped bordered>
                                <thead>
                                    <tr>
                                        <th>Grade</th>
                                        <th>Comment</th>
                                        <th>Inapropriate</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {reviews.map(review => (
                                        <tr key={review.ID}>
                                            <td>{review.Grade}</td>
                                            <td>{review.Comment}</td>
                                            <td>
                                                <Button onClick={() => inappropriate(review)} disabled={review.Inappropriate}>
                                                    Inappropriate
                                                </Button>
                                            </td>
                                        </tr>
                                    ))}
                                </tbody>
                            </Table>
                        </Card.Body>
                    </Card>
                </Col>
            </Row>
            <Confirmation show={show} onHide={() => setShow(false)} question={question} id={id} action={action} close={setShow}/>
        </Container>
    )
}

export default EmployerReviews