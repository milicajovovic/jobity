import axios from "axios";
import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Col, Container, Row, Card, Button, Table } from "react-bootstrap";
import AdminMenu from "../components/AdminMenu";
import Confirmation from "../components/Confirmation";

function AdminReviews() {
    const [reviews, setReviews] = useState([]);
    const [question, setQuestion] = useState("");
    const [id, setId] = useState();
    const [action, setAction] = useState("");
    const [show, setShow] = useState(false);

    useEffect(() => {
        axios.get("http://localhost:3007/reviews", { headers: { Authorization: localStorage.getItem("jwt") } }).then(res => {
            setReviews(res.data);
        });
    }, [reviews]);

    const appropriateReview = (review) => {
        setQuestion("Are you sure you want to mark this review as appropriate?");
        setId(review.ID);
        setAction("reviews/appropriate");
        setShow(true);
    };

    const deleteReview = (review) => {
        setQuestion("Are you sure you want to delete this review?");
        setId(review.ID);
        setAction("reviews/delete");
        setShow(true);
    };

    return (
        <Container fluid>
            <Row>
                <AdminMenu />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "50rem" }}>
                        <Card.Title className="text-center mt-3" as="h3">Employees</Card.Title>
                        <Card.Body className="text-center">
                            <Table striped bordered>
                                <thead>
                                    <tr>
                                        <th>Grade</th>
                                        <th>Comment</th>
                                        <th>Apropriate</th>
                                        <th>Delete</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {reviews.map(review => (
                                        <tr key={review.ID}>
                                            <td>{review.Grade}</td>
                                            <td>{review.Comment}</td>
                                            <td>
                                                <Button onClick={() => appropriateReview(review)} disabled={!review.Inappropriate || review.Deleted}>
                                                    Appropriate
                                                </Button>
                                            </td>
                                            <td>
                                                <Button onClick={() => deleteReview(review)} disabled={review.Deleted}>
                                                    Delete
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

export default AdminReviews