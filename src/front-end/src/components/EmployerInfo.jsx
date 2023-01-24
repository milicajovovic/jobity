import React from "react";
import Modal from 'react-bootstrap/Modal';
import GradesInfo from "./GradesInfo";

function EmployerInfo(props) {
    return (
        <Modal
            {...props}
            aria-labelledby="contained-modal-title-vcenter"
            centered>
            <Modal.Header closeButton>
                <Modal.Title>
                    {props.employer.Name}
                </Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <p>Address: {props.employer.Address}</p>
                <GradesInfo employerID={props.employer.ID}/>
            </Modal.Body>
        </Modal>
    );
}

export default EmployerInfo