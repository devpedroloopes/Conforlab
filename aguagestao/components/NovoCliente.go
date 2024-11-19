package components

const NovoCliente = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Novo Cliente</title>
    <link rel="stylesheet" href="/static/styles.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
</head>
<body>
    <div class="form-container">
        <form class="form">
            <h2>Novo Cliente</h2>
            
            <div class="floating-label">
                <input type="text" id="contratante" class="input-field" placeholder=" ">
                <label for="contratante">Contratante</label>
            </div>

            <div class="floating-label">
                <input type="text" id="local" class="input-field" placeholder=" ">
                <label for="local">Local</label>
            </div>

            <div class="floating-label">
                <select id="supervisor" class="input-field">
                    <option value=""></option>
                    <option value="Geraldo">Geraldo</option>
                    <option value="Igor">Igor</option>
                    <option value="Pasquale">Pasquale</option>
                    <option value="Ronaldo">Ronaldo</option>
                </select>
                <label for="supervisor">Supervisor</label>
            </div>

            <div class="floating-label">
                <select id="tecnico" class="input-field">
                    <option value=""></option>
                    <option value="Alan">Alan</option>
                    <option value="Alcides">Alcides</option>
                    <option value="Leonardo">Leonardo</option>
                    <option value="Mauro">Mauro</option>
                    <option value="Severino">Severino</option>
                    <option value="Técnico BA">Técnico BA</option>
                    <option value="Técnico DF">Técnico DF</option>
                    <option value="Técnico MG">Técnico MG</option>
                    <option value="Técnico RS">Técnico RS</option>
                </select>
                <label for="tecnico">Técnico</label>
            </div>

            <div class="floating-label">
                <input type="text" id="logradouro" class="input-field" placeholder=" ">
                <label for="logradouro">Logradouro</label>
            </div>

            <div class="floating-label">
                <input type="text" id="bairro" class="input-field" placeholder=" ">
                <label for="bairro">Bairro</label>
            </div>

            <div class="floating-label">
                <input type="text" id="cidade" class="input-field" placeholder=" ">
                <label for="cidade">Cidade</label>
            </div>

            <div class="floating-label">
                <input type="text" id="uf" class="input-field" placeholder=" ">
                <label for="uf">UF</label>
            </div>

            <div class="floating-label">
                <input type="text" id="cep" class="input-field" placeholder=" ">
                <label for="cep">CEP</label>
            </div>

            <h2>Dados de Contato</h2>

            <div class="floating-label">
                <input type="text" id="nome" class="input-field" placeholder=" ">
                <label for="nome">Nome</label>
            </div>

            <div class="floating-label">
                <input type="text" id="telefone" class="input-field" placeholder=" ">
                <label for="telefone">Telefone</label>
            </div>

            <div class="floating-label">
                <input type="email" id="email" class="input-field" placeholder=" ">
                <label for="email">Email</label>
            </div>

            <div class="form-buttons">
                <button type="submit" class="confirm-button">
                    <i class="fas fa-check"></i>
                    <span>Confirmar</span>
                </button>
                <button type="button" class="cancel-button" id="cancel-button">
    					<i class="fas fa-times"></i>
    				<span>Cancelar</span>
				</button>
            </div>
        </form>
    </div>
    <script src="/static/script.js"></script>
</body>
</html>
`