-- name: GetOrdersByUser :many
select *
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
where operti.tipodoc = 'PED' and operti.upickup = ? and operti.idcarrito = ''
;

