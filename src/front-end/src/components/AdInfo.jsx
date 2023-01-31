import React from "react";
import Modal from 'react-bootstrap/Modal';

function AdInfo(props) {
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
                    {props.ad.Name}
                </Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <p><b>Description:</b> {props.ad.Description}</p>
                <p><b>Posted:</b> {convertDate(props.ad.Posted)}</p>
                <p><b>Job type:</b> {convertArray(props.ad.JobType)}</p>
                <p><b>Required skills:</b> {convertArray(props.ad.RequiredSkills)}</p>
            </Modal.Body>
        </Modal>
    );
}

export default AdInfo