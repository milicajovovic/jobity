import axios from "axios";
import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Col, Container, Row, Card, Button, Table } from "react-bootstrap";
import AdminMenu from "../components/AdminMenu";
import Confirmation from "../components/Confirmation";

function AdminAds() {
    const [ads, setAds] = useState([]);
    const [question, setQuestion] = useState("");
    const [id, setId] = useState();
    const [action, setAction] = useState("");
    const [show, setShow] = useState(false);

    useEffect(() => {
        axios.get("http://localhost:3007/ads", { headers: { Authorization: localStorage.getItem("jwt") } }).then(res => {
            setAds(res.data);
        });
    }, [ads]);

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

    const deleteAd = (ad) => {
        setQuestion("Are you sure you want to delete " + ad.Name + "?");
        setId(ad.ID);
        setAction("ads/delete");
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
                        <Card.Title className="text-center mt-3" as="h3">Ads</Card.Title>
                        <Card.Body className="text-center">
                            <Table striped bordered>
                                <thead>
                                    <tr>
                                        <th>Name</th>
                                        <th>Description</th>
                                        <th>Date posted</th>
                                        <th>Delete</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {ads.map(ad => (
                                        <tr key={ad.ID}>
                                            <td>{ad.Name}</td>
                                            <td>{ad.Description}</td>
                                            <td>{convertDate(ad.Posted)}</td>
                                            <td>
                                                <Button onClick={() => deleteAd(ad)} disabled={ad.Deleted}>
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

export default AdminAds