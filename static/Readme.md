Overview:

The Groupie Tracker website provides a platform for users to explore information about various music artists and their relationships. It offers a user-friendly interface with features for searching artists, browsing through a paginated list, and accessing detailed information about individual artists.

 Design and Functionality:

Home Page:
   
Upon visiting the website, users are presented with a paginated list of artists.
Each artist is displayed as a card containing basic information such as their name, creation date, and first album.
Users can navigate between pages using pagination links located at the bottom of the page.
A search form at the top allows users to search for specific artists by name.

Artist Pages:
   
Clicking on an artist card navigates the user to a dedicated page displaying detailed information about that artist.
Information includes the artist's name, creation date, first album, members, and relationships (dates and locations).

Navigation
   
The website allows seamless navigation between the home page and individual artist pages.
Users can easily go back to the previous page using the "Go Back" button on artist detail pages.

API Integration
   
Artist information and relationships are fetched from a remote API (https://groupietrackers.herokuapp.com/api/artists and https://groupietrackers.herokuapp.com/api/relation/).
This ensures that the data displayed on the website is always up-to-date and accurate.

Styling:
   
The website utilizes CSS stylesheets for enhancing visual presentation, providing an aesthetically pleasing user experience.
Each page layout is designed to be responsive and accessible across different devices.

Footer
   
A footer is included on every page, containing information about the Groupie Tracker project and links to its social media profiles.
It also displays copyright information.

Technologies Used:

Backend: Go (Golang)
Frontend: HTML, CSS
API Integration: HTTP requests to remote API endpoints
Template Engine HTML templates with Go's html/template package

Sources:

 https://www.w3schools.com/css/
 https://codingnepalweb.com/demos/draggable-card-slider-html-css-javascript/
 https://rapidapi.com/guides/call-apis-go
 https://youtu.be/Su6zn1_blw0?si=KDX9m8PaVImRg3hY
 https://youtu.be/FjkSJ1iXKpg?si=MbiQUfPnE2Z6cYIr



