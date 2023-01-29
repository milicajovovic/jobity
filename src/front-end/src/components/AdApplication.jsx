import axios from "axios";
import React from "react";
import { useState } from "react";
import { Button, Container, Toast, ToastContainer } from "react-bootstrap";
import Modal from 'react-bootstrap/Modal';

function AdApplication(props) {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    
    const apply = () => {
        const jwt = localStorage.getItem("jwt");
        const employeeId = localStorage.getItem("userId");
        axios.post("http://localhost:3007/applications/apply/" + props.ad.ID + "/" + employeeId, null, { headers: { Authorization: jwt } }).then(res => {
            setMessage("successfully applied");
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
                        Application - {props.ad.Name}
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    Are you sure you want to apply for this job?
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={apply}>Apply</Button>
                </Modal.Footer>
            </Modal>
        </Container>
    );
}

export default AdApplication