import axios from "axios";
import React from "react";
import { useState } from "react";
import { Button, Container, Toast, ToastContainer } from "react-bootstrap";
import Modal from 'react-bootstrap/Modal';

function DeclineApplication(props) {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    
    const decline = () => {
        let updatedApplication = props.application;
        updatedApplication.Status = 1;

        const jwt = localStorage.getItem("jwt");
        axios.post("http://localhost:3007/applications/update", updatedApplication, { headers: { Authorization: jwt } }).then(res => {
            setMessage("successfully declined");
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
                        Decline application
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    Are you sure you want to decline this application?
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={decline}>Decline</Button>
                </Modal.Footer>
            </Modal>
        </Container>
    );
}

export default DeclineApplication