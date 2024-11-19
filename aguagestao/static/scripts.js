document.addEventListener('DOMContentLoaded', function() {
    const technicians = ['Mauro', 'Leonardo', 'Alan'];
    const supervisors = ['Ronaldo', 'Pasquale', 'Igor'];
    const clients = [
        { name: 'Abrava (Cortesia)', technician: 'Mauro', supervisor: 'Ronaldo' },
        { name: 'Alamo - Claro Corporate', technician: 'Leonardo', supervisor: 'Pasquale' },
    ];

    const technicianSelect = document.getElementById('technician');
    const supervisorSelect = document.getElementById('supervisor');
    const clientTable = document.getElementById('client-table');
    const searchInput = document.getElementById('search');

    function populateSelectOptions(selectElement, options) {
        options.forEach(option => {
            const opt = document.createElement('option');
            opt.value = option;
            opt.textContent = option;
            selectElement.appendChild(opt);
        });
    }

    function populateClientTable(clients) {
        clientTable.innerHTML = ''; // Clear existing rows
        clients.forEach(client => {
            const row = document.createElement('div');
            row.classList.add('table-row');
            row.innerHTML = `
                <div>${client.name}</div>
                <div>${client.technician}</div>
                <div>${client.supervisor}</div>
            `;
            clientTable.appendChild(row);
        });
    }

    function filterClients() {
        const selectedTechnician = technicianSelect.value;
        const selectedSupervisor = supervisorSelect.value;
        const searchTerm = searchInput.value.toLowerCase();

        const filteredClients = clients.filter(client => {
            const matchesTechnician = !selectedTechnician || client.technician === selectedTechnician;
            const matchesSupervisor = !selectedSupervisor || client.supervisor === selectedSupervisor;
            const matchesSearchTerm = client.name.toLowerCase().includes(searchTerm);
            return matchesTechnician && matchesSupervisor && matchesSearchTerm;
        });

        populateClientTable(filteredClients);
    }

    // Populate initial data
    populateSelectOptions(technicianSelect, technicians);
    populateSelectOptions(supervisorSelect, supervisors);
    populateClientTable(clients);

    // Add event listeners for filtering
    technicianSelect.addEventListener('change', filterClients);
    supervisorSelect.addEventListener('change', filterClients);
    searchInput.addEventListener('input', filterClients);

    // Redirect to /novoCliente on "Novo Cliente" button click
    document.getElementById('new-client').addEventListener('click', function() {
        window.location.href = '/novoCliente';
    });

    document.getElementById('cancel-button').addEventListener('click', function() {
        window.location.replace('/')
    });
});