import axios from "axios";
import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Col, Container, Row, Card, Button, Table } from "react-bootstrap";
import AdminMenu from "../components/AdminMenu";
import Confirmation from "../components/Confirmation";

function AdminEmployers() {
    const [employers, setEmployers] = useState([]);
    const [question, setQuestion] = useState("");
    const [id, setId] = useState();
    const [action, setAction] = useState("");
    const [show, setShow] = useState(false);

    useEffect(() => {
        axios.get("http://localhost:3007/employers", { headers: { Authorization: localStorage.getItem("jwt") } }).then(res => {
            setEmployers(res.data);
        });
    }, [employers]);

    const deleteEmployer = (employer) => {
        setQuestion("Are you sure you want to delete " + employer.Name + "?");
        setId(employer.ID);
        setAction("employers/delete");
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
                        <Card.Title className="text-center mt-3" as="h3">Employers</Card.Title>
                        <Card.Body className="text-center">
                            <Table striped bordered>
                                <thead>
                                    <tr>
                                        <th>Email</th>
                                        <th>Name</th>
                                        <th>Address</th>
                                        <th>Delete</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {employers.map(employer => (
                                        <tr key={employer.ID}>
                                            <td>{employer.Email}</td>
                                            <td>{employer.Name}</td>
                                            <td>{employer.Address}</td>
                                            <td>
                                                <Button onClick={() => deleteEmployer(employer)} disabled={employer.Deleted}>
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

export default AdminEmployers