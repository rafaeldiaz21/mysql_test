select
    e.id_equipos,
    e.nombre as equipo,
    count(p.id_partidos) total_partidos_visitante
from
    equipos e
    left join partidos p on e.id_equipos = p.fk_equipo_visitante
group by
    e.id_equipos, e.nombre
order by
    nombre