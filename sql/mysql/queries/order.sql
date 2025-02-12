-- name: GetOrdersByUser :many
select
    operti.agencia,
    operti.tipodoc,
    operti.documento,
    operti.nombrecli,
    operti.codcliente,
    operti.emision,
    operti.upickup,
    operti.idcarrito,
    operti.almacen,
    operti.ke_pedstatus,
    opermv.agencia,
    opermv.tipodoc,
    opermv.documento,
    opermv.codigo,
    opermv.nombre,
    opermv.almacen,
    opermv.cantref,
    opermv.cantidad,
    kerutazonas.ruta_codigo,
    keruta.ruta_descrip,
    articulo.referencia,
    articulo.marca,
    articulo.unidad,
    articulo.rutafoto,
    articulo.fechacrea
from operti
left join
    opermv
    on operti.tipodoc = opermv.tipodoc
    and operti.agencia = opermv.agencia
    and operti.documento = opermv.documento
left join
    keol_opti2
    on operti.tipodoc = keol_opti2.opti2_tipodoc
    and operti.agencia = keol_opti2.opti2_agencia
    and operti.documento = keol_opti2.opti2_documento
left join articulo on opermv.codigo = articulo.codigo
left join
    kerutazonas
    on operti.sector = kerutazonas.codigo
    and operti.subcodigo = kerutazonas.subcodigo
left join keruta on kerutazonas.ruta_codigo = keruta.ruta_codigo
where operti.tipodoc = 'PED' and operti.upickup = ? and operti.ke_pedstatus = "14"
order by operti.emision asc, operti.documento asc, operti.almacen asc
;

-- name: GetOrderByCode :many
select
    operti.agencia,
    operti.tipodoc,
    operti.documento,
    operti.nombrecli,
    operti.codcliente,
    operti.emision,
    operti.upickup,
    operti.idcarrito,
    operti.almacen,
    operti.ke_pedstatus,
    opermv.agencia,
    opermv.tipodoc,
    opermv.documento,
    opermv.codigo,
    opermv.nombre,
    opermv.almacen,
    opermv.cantref,
    opermv.cantidad,
    kerutazonas.ruta_codigo,
    keruta.ruta_descrip,
    articulo.referencia,
    articulo.marca,
    articulo.unidad,
    articulo.rutafoto,
    articulo.fechacrea
from operti
left join
    opermv
    on operti.tipodoc = opermv.tipodoc
    and operti.agencia = opermv.agencia
    and operti.documento = opermv.documento
left join
    keol_opti2
    on operti.tipodoc = keol_opti2.opti2_tipodoc
    and operti.agencia = keol_opti2.opti2_agencia
    and operti.documento = keol_opti2.opti2_documento
left join articulo on opermv.codigo = articulo.codigo
left join
    kerutazonas
    on operti.sector = kerutazonas.codigo
    and operti.subcodigo = kerutazonas.subcodigo
left join keruta on kerutazonas.ruta_codigo = keruta.ruta_codigo
where operti.tipodoc = 'PED' and operti.upickup = ? and operti.documento = ?
;

-- name: GetOrderWithCart :one
select count(*)
from operti
where idcarrito = ?
;

-- name: UpdateOrderCart :exec
update operti set
    idcarrito = ?
where 
    tipodoc = 'PED'
    and documento = ? 
    and upickup = ?
;

