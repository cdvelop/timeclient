// out: 2006-01-02 15:04:05
function currentDate() {

    const now_date = new Date();

    // fecha formato YYYY-MM-DD
    const year = now_date.getFullYear();
    const month = String(now_date.getMonth() + 1).padStart(2, '0');
    const day = String(now_date.getDate()).padStart(2, '0');

    const date = year + "-" + month + "-" + day;

    // Obtén la hora local en formato de 24 horas
    const hour = now_date.toLocaleTimeString('es-ES', { hour12: false });

    // console.log("NEW DATE:", date, " hora:", hour);

    return date+" "+hour
}