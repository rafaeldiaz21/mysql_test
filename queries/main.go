
package queries
// Automatically generated queries 
// ALL EDITS WILL BE LOST
const (
	JuegosPorFechas = `select e.nombre as equipo_local, e2.nombre as equipo_visitante, fecha_partido from partidos p join equipos e on p.fk_equipo_local = e.id_equipos join equipos e2 on p.fk_equipo_visitante = e2.id_equipos where p.fecha_partido in (?)`
	JugadorMayorEdadPorEquipo = `with CTE as (select (DATEDIFF(NOW(), fecha_nacimiento)/ 365)edad, j.nombre, e.id_equipos, j.id_jugadores, e.nombre as equipo from jugadores j join equipos e on j.fk_equipos = e.id_equipos )select id_jugadores, nombre, MAX(edad)as edad, id_equipos, equipo from CTE group by id_equipos;`
	PartidosVisitantePorEquipo = `select e.id_equipos, e.nombre as equipo, count(p.id_partidos)total_partidos_visitante from equipos e left join partidos p on e.id_equipos = p.fk_equipo_visitante group by e.id_equipos, e.nombre order by nombre`
	TotalesGolesPorEquipo = `with equipo as (select id_equipos, nombre from equipos e WHERE nombre = ? ), local as(select sum(goles_local)as goles_local, id_equipos FROM partidos p join equipos e on p.fk_equipo_local = e.id_equipos and e.nombre = ? ), visitante as (select sum(goles_visitante)as goles_visitante, id_equipos FROM partidos p join equipos e on p.fk_equipo_visitante = e.id_equipos and e.nombre = ? )select equipo.id_equipos, equipo.nombre, goles_local, goles_visitante from equipo, local, visitante` 
)
