// 1. Inject sliders into the form
const sliders = [
    { id: 'sleep_quality',   label: 'Sleep Quality (-3 to 3)', valId: 'val-sleep' },
    { id: 'physical_energy', label: 'Physical Energy',         valId: 'val-energy' },
    { id: 'focus',           label: 'Focus',                   valId: 'val-focus' },
    { id: 'motivation',      label: 'Motivation',              valId: 'val-motivation' },
    { id: 'past_view',       label: 'Past View',               valId: 'val-past' },
    { id: 'social_activity', label: 'Social Activity',         valId: 'val-social' },
];

const form = document.getElementById('metricsForm');
const submitBtn = document.getElementById('submitBtn');

const dateDiv = document.createElement('div');
dateDiv.className = 'input-group';
dateDiv.innerHTML = `<label for="entry_date">Date:</label>
    <input type="date" id="entry_date" required>`;
form.insertBefore(dateDiv, submitBtn);

sliders.forEach(({ id, label, valId }) => {
    const div = document.createElement('div');
    div.className = 'slider-container';
    div.innerHTML = `<label>${label}: <span id="${valId}">0</span></label>
        <input type="range" id="${id}" min="-3" max="3" value="0">`;
    form.insertBefore(div, submitBtn);
});

const valIdMap = Object.fromEntries(sliders.map(s => [s.id, s.valId]));

document.querySelectorAll('input[type="range"]').forEach(slider => {
    slider.addEventListener('input', (e) => {
        document.getElementById(valIdMap[e.target.id]).innerText = e.target.value;
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

// 3. The Chart Logic
let healthChart; // We store the chart here so we can refresh it

async function loadChartData() {
    const response = await fetch('/logs');
    const data = await response.json();

    // If there is no data, don't try to draw a chart
    if (!data || data.length === 0) return;

    // Pre-fill form with last entry values
    const last = data[data.length - 1];
    document.getElementById('entry_date').value = last.entry_date.split('T')[0];
    sliders.forEach(({ id, valId }) => {
        const val = last[id] ?? 0;
        document.getElementById(id).value = val;
        document.getElementById(valId).innerText = val;
    });

    // Prepare labels (Dates) and datasets (Scores)
    const labels = data.map(item => item.entry_date.split('T')[0]); // Clean up the date string
    const sleepData = data.map(item => item.sleep_quality);
    const focusData = data.map(item => item.focus);
    const energyData = data.map(item => item.physical_energy);
    const motivationData = data.map(item => item.motivation);
    const pastViewData = data.map(item => item.past_view);  
    const socialData = data.map(item => item.social_activity);


    const ctx = document.getElementById('healthChart').getContext('2d');

    // If a chart already exists, destroy it before making a new one (to avoid glitches)
    if (healthChart) healthChart.destroy();

    healthChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: labels,
            datasets: [
            {
                    label: 'Sleep Quality',
                    data: sleepData,
                    borderColor: '#2ecc71',
                    tension: 0.3,
                    fill: false
                },
                {
                    label: 'Focus',
                    data: focusData,
                    borderColor: '#3498db',
                    tension: 0.3,
                    fill: false
                },
                {
                    label: 'Physical Energy',
                    data: energyData,
                    borderColor: '#e74c3c',
                    tension: 0.3,
                    fill: false
                },
                {
                    label: 'Motivation',
                    data: motivationData,
                    borderColor: '#9b59b6',
                    tension: 0.3,
                    fill: false
                },
                {
                    label: 'Past View',
                    data: pastViewData,
                    borderColor: '#f39c12',
                    tension: 0.3,
                    fill: false
                },
                {
                    label: 'Social Activity',
                    data: socialData,
                    borderColor: '#16a085',
                    tension: 0.3,
                    fill: false
                }

                
            ]
        },
        options: {
            scales: {
                y: {
                    min: -3,
                    max: 3,
                    ticks: { stepSize: 1 }
                }
            }
        }
    });
}

// Run this when the page first loads
loadChartData();