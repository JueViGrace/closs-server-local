-- name: GetDocuments :many
select *
from closs_document
;

-- name: GetDocumentsWithLines :many
select *
from closs_document
left join
    closs_document_lines on closs_document.documento = closs_document_lines.documento
;

-- name: GetDocumentByCode :one
select *
from closs_document
where documento = ?
;

-- name: GetDocumentWithLinesByCode :many
select *
from closs_document
left join
    closs_document_lines on closs_document.documento = closs_document_lines.documento
where closs_document.documento = ?
;

-- name: GetDocumentsByManager :many
select closs_document.*
from closs_document
left join closs_managers on closs_document.codcoord = closs_managers.kng_codcoord
where closs_managers.kng_codgcia = ?
order by
    closs_document.vendedor asc,
    closs_document.emision desc,
    closs_document.codcliente asc
;

-- name: GetDocumentsBySalesman :many
select closs_document.*
from closs_document
where closs_document.vendedor = ?
order by closs_document.codcliente asc, closs_document.emision desc
;

-- name: GetDocumentsByCustomer :many
select closs_document.*
from closs_document
where closs_document.codcliente = ?
order by closs_document.emision desc
;

-- name: CreateDocument :one
INSERT INTO closs_document (
    agencia,
    tipodoc,
    documento,
    tipodocv,
    codcliente,
    nombrecli,
    contribesp,
    ruta_parme,
    tipoprecio,
    emision,
    recepcion,
    vence,
    diascred,
    estatusdoc,
    dtotneto,
    dtotimpuest,
    dtotalfinal,
    dtotpagos,
    dtotdescuen,
    dFlete,
    dtotdev,
    dvndmtototal,
    dretencion,
    dretencioniva,
    vendedor,
    codcoord,
    aceptadev,
    kti_negesp,
    bsiva,
    bsflete,
    bsretencion,
    bsretencioniva,
    tasadoc,
    mtodcto,
    fchvencedcto,
    tienedcto,
    cbsret,
    cdret,
    cbsretiva,
    cdretiva,
    cbsrparme,
    cdrparme,
    cbsretflete,
    cdretflete,
    bsmtoiva,
    bsmtofte,
    retmun_mto,
    dolarflete,
    bsretflete,
    dretflete,
    dretmun_mto,
    retivaoblig,
    edoentrega,
    dmtoiva,
    prcdctoaplic,
    montodctodol,
    montodctobs,
    created_at,
    updated_at
)
values (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: UpdateDocument :one
update closs_document set 
    agencia = ?,
    tipodoc = ?,
    tipodocv = ?,
    codcliente = ?,
    nombrecli = ?,
    contribesp = ?,
    ruta_parme = ?,
    tipoprecio = ?,
    emision = ?,
    recepcion = ?,
    vence = ?,
    diascred = ?,
    estatusdoc = ?,
    dtotneto = ?,
    dtotimpuest = ?,
    dtotalfinal = ?,
    dtotpagos = ?,
    dtotdescuen = ?,
    dFlete = ?,
    dtotdev = ?,
    dvndmtototal = ?,
    dretencion = ?,
    dretencioniva = ?,
    vendedor = ?,
    codcoord = ?,
    aceptadev = ?,
    kti_negesp = ?,
    bsiva = ?,
    bsflete = ?,
    bsretencion = ?,
    bsretencioniva = ?,
    tasadoc = ?,
    mtodcto = ?,
    fchvencedcto = ?,
    tienedcto = ?,
    cbsret = ?,
    cdret = ?,
    cbsretiva = ?,
    cdretiva = ?,
    cbsrparme = ?,
    cdrparme = ?,
    cbsretflete = ?,
    cdretflete = ?,
    bsmtoiva = ?,
    bsmtofte = ?,
    retmun_mto = ?,
    dolarflete = ?,
    bsretflete = ?,
    dretflete = ?,
    dretmun_mto = ?,
    retivaoblig = ?,
    edoentrega = ?,
    dmtoiva = ?,
    prcdctoaplic = ?,
    montodctodol = ?,
    montodctobs = ?,
    updated_at = ?
WHERE documento = ?
RETURNING *;

