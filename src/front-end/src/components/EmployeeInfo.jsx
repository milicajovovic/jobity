import React from "react";
import Modal from 'react-bootstrap/Modal';

function EmployeeInfo(props) {
    const convertDate = (posted) => {
        let date = new Date(posted);
        let year = date.getFullYear();
        let month = date.getMonth() + 1;
        let day = date.getDate();

        if (day < 10) {
            day = "0" + day;
        }
        if (month < 10) {
            month = "0" + month;
        }

        return day + '-' + month + '-' + year;
    };

    const convertArray = (array) => {
        if (array !== undefined) {
            return array.join(", ");
        }
        return "";
    }

    return (
        <Modal
            {...props}
            aria-labelledby="contained-modal-title-vcenter"
            centered>
            <Modal.Header closeButton>
                <Modal.Title>
                    {props.employee.FirstName} {props.employee.LastName}
                </Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <p><b>Email:</b> {props.employee.Email}</p>
                <p><b>Birthday:</b> {convertDate(props.employee.Birthday)}</p>
                <p><b>Education:</b> {props.employee.Education}</p>
                <p><b>Skills:</b> {convertArray(props.employee.Skills)}</p>
            </Modal.Body>
        </Modal>
    );
}

export default EmployeeInfo