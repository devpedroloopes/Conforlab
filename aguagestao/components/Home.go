package components

const Home = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Clientes</title>
    <link rel="stylesheet" type="text/css" href="/static/styles.css">
</head>
<body>
    <div class="container">
        <div class="header">
            <div style="flex: 1; display: flex;">
                <input type="text" id="search" placeholder="Pesquisar">
                <select id="technician">
                    <option value="">Técnico</option>
                </select>
                <select id="supervisor">
                    <option value="">Supervisor</option>
                </select>
            </div>
            <button id="new-client">Novo Cliente</button>
        </div>
        <div class="table">
            <div class="table-header">
                <div>Cliente</div>
                <div>Técnico</div>
                <div>Supervisor</div>
            </div>
            <div id="client-table" class="table-body"></div>
        </div>
    </div>
    <script src="/static/scripts.js"></script>
</body>
</html>
`