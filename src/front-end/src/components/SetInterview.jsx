import axios from "axios";
import React from "react";
import { useState } from "react";
import { Button, Container, Form, Toast, ToastContainer } from "react-bootstrap";
import Modal from 'react-bootstrap/Modal';

function SetInterview(props) {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");

    const save = (event) => {
        event.preventDefault();

        let updatedApplication = props.application;
        updatedApplication.Interview = new Date(event.target.interview.value).toISOString();
        updatedApplication.Status = 2;

        const jwt = localStorage.getItem("jwt");
        axios.post("http://localhost:3007/applications/update", updatedApplication, { headers: { Authorization: jwt } }).then(res => {
            setMessage("successfully setted");
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
                    <Modal.Title>
                        Set interview
                    </Modal.Title>
                </Modal.Header>
                <Form onSubmit={save}>
                    <Modal.Body>
                        <Form.Group className="mb-3">
                            <Form.Label>Interview date</Form.Label>
                            <Form.Control type="date" name="interview" required />
                        </Form.Group>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button type="submit">Save</Button>
                    </Modal.Footer>
                </Form>
            </Modal>
        </Container>
    );
}

export default SetInterview