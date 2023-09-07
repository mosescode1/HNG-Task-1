"use strict";

// Get references to HTML elements with IDs 'currentDay' and 'currentTime'
let dayElement = document.getElementById('currentDay');
let timeElement = document.getElementById('currentTime');

let currentDay;
let currentTime;

// Function to get and display the current day and time
function getCurrentDay() {
    let currentDate = Date.now();

    // Get the current day as a string (e.g., "Tuesday")
    currentDay = new Date(currentDate).toLocaleDateString(undefined, { weekday: 'long' });

    // Get the current time as a string (e.g., "14:30:00")
    currentTime = new Date(currentDate).toLocaleTimeString();

    // Display the current day and time in the respective HTML elements
    dayElement.textContent = currentDay;
    timeElement.textContent = currentTime;
}

// Attach the getCurrentDay function to the window.onload event
window.onload = getCurrentDay;
