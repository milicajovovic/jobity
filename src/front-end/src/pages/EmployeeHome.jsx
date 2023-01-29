import axios from "axios";
import React, { useEffect, useState } from "react";
import { Col, Container, Row, Card, Button, Navbar } from "react-bootstrap";
import AdCard from "../components/AdCard";
import SearchForm from "../components/SearchForm";
import EmployeeMenu from "../components/EmployeeMenu";

function EmployeeHome() {
    const [ads, setAds] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:3007/ads").then(res => {
            setAds(res.data);
        });
    }, []);

    const sortAds = () => {
        setAds([].concat(ads).sort((a, b) => new Date(b.Posted) - new Date(a.Posted)))
    }

    return (
        <Container fluid>
            <Row>
                <EmployeeMenu />
            </Row>
            <Row>
                <Navbar bg="light" className="searchBar">
                    <Container>
                        <SearchForm changeAds={setAds} />
                        <Button onClick={sortAds}>Sort by date</Button>
                    </Container>
                </Navbar>
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    {ads.length > 0 ?
                        ads.map(a => (
                            <AdCard ad={a} key={a.ID} />
                        ))
                        : <Card body style={{ width: "40rem" }} className="text-center">There is no ads for chosen criteria</Card>
                    }
                </Col>
            </Row>
        </Container>
    )
}

export default EmployeeHome