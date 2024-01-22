// out: 2006-01-02 15:04:05
function currentDate() {

    const now_date = new Date();

    // fecha formato YYYY-MM-DD
    const year = now_date.getFullYear();
    const month = String(now_date.getMonth() + 1).padStart(2, '0');
    const day = String(now_date.getDate()).padStart(2, '0');

    const date = year + "-" + month + "-" + day;


    // Opciones formato de 24 horas HH:MM:SS
    const hour_format = { hour: '2-digit', minute: '2-digit', second: '2-digit', hour12: false };

    // Formatea la hora en el formato HH:MM:SS
    const hour = new Intl.DateTimeFormat('es-ES', hour_format).format(now_date);

    // console.log("NEW DATE:", date, " hora:", hour);

    return date + " " + hour
}