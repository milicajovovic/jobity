use axum::{
    response::{Html, IntoResponse},
    routing::{get, post},
    Json, Router,
};
use dotenv::dotenv; // import the dotenv crate for parsing the `.env file`
use serde::{Deserialize, Serialize};
use std::net::SocketAddr; // import serde for parsing our struct to and from Json
                          //import the email library needed modules
use lettre::transport::smtp::authentication::Credentials;
use lettre::{Message, SmtpTransport, Transport};

/// define a structure that maps to the format of our HTTP request body
/// derive the Debug trait, this will allow, printing the struct in stdout
/// derive the Serializing trait, this will allow building up JSON
/// derive the Deserializing trait
#[derive(Debug, Serialize, Deserialize)]
struct EmailPayload {
    fullname: String,
    email: String,
    message: String,
}

//mount the tokio runtime to allow our main function to support asynchronous execution
#[tokio::main]
async fn main() {
    dotenv().ok();
    // build our application with a route
    let app = Router::new()
        .route("/", get(handler))
        //mount the handle to a path, using the HTTP POST verb
        .route("/send-email", post(dispatch_email));

    // run it
    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    println!("listening on http://{}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn handler() -> Html<&'static str> {
    Html("<h1>Hello, World!</h1>")
}

/// define an email handler, the handler will take the user's email address and the message body
/// the parsed payload will be fed into the `lettre` library and finally, a response will be returned
async fn dispatch_email(Json(payload): Json<EmailPayload>) -> impl IntoResponse {
    // println!("{:#?}", payload);
    //destructure the HTTP request body
    let EmailPayload {
        email,
        message,
        fullname,
    } = &payload;

    //contruct emil config
    let from_address = String::from("You <jobityntp@gmail.com>");
    let to_address = format!("{fullname} <{email}>");
    let reply_to = String::from("You <jobityntp@gmail.com>");
    let email_subject = "Axum Rust tutorial";

    let email = Message::builder()
        .from(from_address.parse().unwrap())
        .reply_to(reply_to.parse().unwrap())
        .to(to_address.parse().unwrap())
        .subject(email_subject)
        .body(String::from(message))
        .unwrap();

    let creds = Credentials::new("jobityntp@gmail.com".to_string(), "Jobity22!".to_string()); 

    // Open a remote connection to SMTP server
    let mailer = SmtpTransport::relay("smtp.gmail.com")
        .unwrap()
        .credentials(creds)
        .build();

    // Send the email
    match mailer.send(&email) {
        Ok(_) => println!("Email sent successfully!"),
        Err(e) => panic!("Could not send email: {:?}", e),
    }
}