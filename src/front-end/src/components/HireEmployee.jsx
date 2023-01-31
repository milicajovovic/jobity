import axios from "axios";
import React from "react";
import { useState } from "react";
import { Button, Container, Toast, ToastContainer } from "react-bootstrap";
import Modal from 'react-bootstrap/Modal';

function HireEmployee(props) {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    
    const decide = (status) => {
        let updatedApplication = props.application;
        updatedApplication.Status = status;

        const jwt = localStorage.getItem("jwt");
        axios.post("http://localhost:3007/applications/update", updatedApplication, { headers: { Authorization: jwt } }).then(res => {
            setMessage("successfully saved decision");
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
                        Hire employee
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    Do you want to hire employee?
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={() => decide(3)}>Yes</Button>
                    <Button onClick={() => decide(1)}>No</Button>
                </Modal.Footer>
            </Modal>
        </Container>
    );
}

export default HireEmployee