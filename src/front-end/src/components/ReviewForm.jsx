import axios from "axios";
import React from "react";
import { useState } from "react";
import { Button, Container, Form, Toast, ToastContainer } from "react-bootstrap";
import Modal from 'react-bootstrap/Modal';

function ReviewForm(props) {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");

    const add = (event) => {
        event.preventDefault();

        const jwt = localStorage.getItem("jwt");
        const employeeId = parseInt(localStorage.getItem("userId"));

        const newReview = {
            "EmployerID": props.employer.ID,
            "EmployeeID": employeeId,
            "Grade": parseInt(event.target.grade.value),
            "Comment": event.target.comment.value,
        };

        axios.post("http://localhost:3007/reviews/create", newReview, { headers: { Authorization: jwt } }).then(res => {
            setMessage("successfully added");
            setShow(true);
        }).catch((err) => {
            setMessage(err.response.data);
            setShow(true);
        });

        props.close(false);
    }

    return (
        <Container>
            <ToastContainer position="top-center" className="text-center p-3">
                <Toast onClose={() => setShow(false)} show={show} delay={3000} autohide>
                    <Toast.Body>{message}</Toast.Body>
                </Toast>
            </ToastContainer>
            <Modal
                {...props}
                aria-labelledby="contained-modal-title-vcenter"
                centered
            >
                <Modal.Header closeButton>
                    <Modal.Title>Review - {props.employer.Name}</Modal.Title>
                </Modal.Header>

                <Form onSubmit={add}>
                    <Modal.Body>
                        <Form.Group className="mb-3">
                            <Form.Label>Grade</Form.Label>
                            <Form.Control type="number" name="grade" min="0" max="5" />
                        </Form.Group>
                        <Form.Group>
                            <Form.Label>Comment</Form.Label>
                            <Form.Control as="textarea" name="comment" rows={2} />
                        </Form.Group>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button type="submit">Add</Button>
                    </Modal.Footer>
                </Form>
            </Modal>
        </Container>
    );
}

export default ReviewForm