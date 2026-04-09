// 1. Make the slider numbers update live
document.querySelectorAll('input[type="range"]').forEach(slider => {
    slider.addEventListener('input', (e) => {
        // This finds the span next to the label and updates the text
        const spanId = "val-" + e.target.id.split('_')[0]; 
        // Note: I used a shortcut here, but for your specific IDs:
        if(e.target.id === "past_view") document.getElementById("val-past").innerText = e.target.value;
        else if(e.target.id === "social_activity") document.getElementById("val-social").innerText = e.target.value;
        else if(e.target.id === "physical_energy") document.getElementById("val-energy").innerText = e.target.value;
        else document.getElementById("val-" + e.target.id.split('_')[0]).innerText = e.target.value;
    });
});

// 2. Handle the Form Submission
const metricsForm = document.getElementById('metricsForm');

metricsForm.addEventListener('submit', async (e) => {
    e.preventDefault(); // Stop the page from refreshing!

    // Gather data from the form
    const formData = {
        entry_date: document.getElementById('entry_date').value,
        sleep_quality: parseInt(document.getElementById('sleep_quality').value),
        physical_energy: parseInt(document.getElementById('physical_energy').value),
        focus: parseInt(document.getElementById('focus_quality')?.value || document.getElementById('focus').value),
        motivation: parseInt(document.getElementById('motivation').value),
        past_view: parseInt(document.getElementById('past_view').value),
        social_activity: parseInt(document.getElementById('social_activity').value)
    };

    // Send it to our Go Backend
    const response = await fetch('/submit', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
    });

    if (response.ok) {
        alert("Saved successfully!");
        // We will call the chart update function here later!
    } else {
        alert("Error saving data. Check console.");
    }
});