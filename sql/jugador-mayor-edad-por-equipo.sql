with CTE as (
    select
        (DATEDIFF(NOW(), fecha_nacimiento) / 365) edad,
        j.nombre,
        e.id_equipos,
        j.id_jugadores,
        e.nombre as equipo
    from
        jugadores j
        join equipos e on j.fk_equipos = e.id_equipos
)
select
    id_jugadores,
    nombre,
    MAX(edad) as edad,
    id_equipos,
    equipo
from
    CTE
group by
    id_equipos;