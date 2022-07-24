select
    e.nombre as equipo_local,
    e2.nombre as equipo_visitante,
    fecha_partido
from
    partidos p
    join equipos e on p.fk_equipo_local = e.id_equipos
    join equipos e2 on p.fk_equipo_visitante = e2.id_equipos
where
    p.fecha_partido in (?)