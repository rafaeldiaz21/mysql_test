with equipo as (
    select
        id_equipos,
        nombre
    from
        equipos e
    WHERE
        nombre = ?
),
local as(
    select
        sum(goles_local) as goles_local,
        id_equipos
    FROM
        partidos p
        join equipos e on p.fk_equipo_local = e.id_equipos
        and e.nombre = ?
),
visitante as (
    select
        sum(goles_visitante) as goles_visitante,
        id_equipos
    FROM
        partidos p
        join equipos e on p.fk_equipo_visitante = e.id_equipos
        and e.nombre = ?
)
select
    equipo.id_equipos,
    equipo.nombre,
    goles_local,
    goles_visitante
from
    equipo,
    local,
    visitante