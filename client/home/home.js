const dados = {
  missões: [
    {
      missão: "Salvar civis de um prédio em chamas",
      localização: "Nova York",
    },
    {
      missão: "Impedir um assalto a um banco",
      localização: "Gotham City",
    },
    {
      missão: "Capturar um vilão perigoso",
      localização: "Metropolis",
    },
    {
      missão: "Desvendar um mistério envolvendo superpoderes",
      localização: "Central City",
    },
  ],
};

document.addEventListener("DOMContentLoaded", function () {
  const tableBody = document.getElementById("table-body");
  dados.missões.forEach((missão) => {
    const row = document.createElement("tr");
    const missãoCell = document.createElement("td");
    const localizaçãoCell = document.createElement("td");
    missãoCell.textContent = missão.missão;
    localizaçãoCell.textContent = missão.localização;
    row.appendChild(missãoCell);
    row.appendChild(localizaçãoCell);
    tableBody.appendChild(row);
  });
});
