import axios from "axios";
import React, { useEffect, useState } from "react";
import { FaStar } from "react-icons/fa";

function ReviewsInfo({ employerID }) {
    const [reviews, setReviews] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:3004/employer/" + employerID).then(res => {
            setReviews(res.data);
        });
        // eslint-disable-next-line
    }, []);

    return (
        <div>
            {reviews.map(review => (
                <div key={review.ID}>
                    <hr />
                    {[...Array(review.Grade)].map((v, i) => <FaStar key={i} />)}
                    <p>{review.Comment}</p>
                </div>
            ))}
        </div>
    )
}

export default ReviewsInfo