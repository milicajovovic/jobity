import axios from "axios";
import React from "react";
import { useState } from "react";
import { Button, Container, Toast, ToastContainer } from "react-bootstrap";
import Modal from 'react-bootstrap/Modal';

function Confirmation(props) {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    
    const confirm = () => {
        const jwt = localStorage.getItem("jwt");
        axios.post("http://localhost:3007/" + props.action + "/" + props.id, null, { headers: { Authorization: jwt } }).then(res => {
            setMessage("successfully " + props.action.split('/')[1] + "ed");
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
                        Confirmation
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    {props.question}
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={confirm}>Confirm</Button>
                </Modal.Footer>
            </Modal>
        </Container>
    );
}

export default Confirmation