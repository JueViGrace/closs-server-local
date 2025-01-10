-- +goose Up
CREATE TABLE IF NOT EXISTS closs_config(
    cnfg_idconfig TEXT NOT NULL,
    cnfg_clase TEXT NOT NULL DEFAULT '',
    cnfg_tipo TEXT NOT NULL,
    cnfg_valnum REAL NOT NULL DEFAULT 0.00,
    cnfg_valsino INTEGER NOT NULL DEFAULT 0,
    cnfg_valtxt TEXT NOT NULL,
    cnfg_lentxt INTEGER NOT NULL DEFAULT 0,
    cnfg_valfch TEXt NOT NULL DEFAULT '1000-01-01',
    cnfg_activa INTEGER NOT NULL DEFAULT 0,
    cnfg_etiq TEXT NOT NULL,
    cnfg_ttip TEXT NOT NULL,
    username TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL,
    PRIMARY KEY (cnfg_idconfig, username)
);

CREATE TABLE IF NOT EXISTS closs_customer(
    codigo TEXT NOT NULL PRIMARY KEY,
    nombre TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL UNIQUE DEFAULT '',
    direccion TEXT NOT NULL DEFAULT '',
    telefonos TEXT NOT NULL DEFAULT '',
    perscont TEXT NOT NULL DEFAULT '',
    vendedor TEXT NOT NULL DEFAULT '',
    contribespecial INTEGER NOT NULL DEFAULT 0,
    status INTEGER NOT NULL DEFAULT 0,
    sector TEXT NOT NULL DEFAULT '',
    subsector TEXT NOT NULL DEFAULT '',
    precio INTEGER NOT NULL DEFAULT 1,
    kne_activa INTEGER NOT NULL DEFAULT 1,
    kne_mtomin REAL NOT NULL DEFAULT 0.0,
    noemifac INTEGER NOT NULL DEFAULT 0,
    noeminota INTEGER NOT NULL DEFAULT 0,
    fchultvta TEXT NOT NULL DEFAULT '',
    mtoultvta REAL NOT NULL DEFAULT 0.0,
    prcdpagdia REAL NOT NULL DEFAULT 0.0,
    promdiasp REAL NOT NULL DEFAULT 0.0,
    riesgocrd REAL NOT NULL DEFAULT 0.0,
    cantdocs INTEGER NOT NULL DEFAULT 0,
    totmtodocs REAL NOT NULL DEFAULT 0.0,
    prommtodoc REAL NOT NULL DEFAULT 0.0,
    diasultvta REAL NOT NULL DEFAULT 0.0,
    promdiasvta REAL NOT NULL DEFAULT 0.0,
    limcred REAL NOT NULL DEFAULT 0.0,
    fchcrea TEXT NOT NULL DEFAULT '1000-01-01',
    dolarflete INTEGER NOT NULL DEFAULT 0,
    nodolarflete INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS closs_document(
    agencia TEXT NOT NULL DEFAULT '',
    tipodoc TEXT NOT NULL DEFAULT '',
    documento TEXT NOT NULL PRIMARY KEY,
    tipodocv TEXT NOT NULL DEFAULT '',
    codcliente TEXT NOT NULL DEFAULT '',
    nombrecli TEXT NOT NULL DEFAULT '',
    contribesp INTEGER NOT NULL DEFAULT 0,
    ruta_parme INTEGER NOT NULL DEFAULT 0,
    tipoprecio INTEGER NOT NULL DEFAULT 1,
    emision TEXT NOT NULL DEFAULT '1000-01-01',
    recepcion TEXT NOT NULL DEFAULT '1000-01-01',
    vence TEXT NOT NULL DEFAULT '1000-01-01',
    diascred INTEGER NOT NULL DEFAULT 0,
    estatusdoc INTEGER NOT NULL DEFAULT 0,
    dtotneto REAL NOT NULL DEFAULT 0.00,
    dtotimpuest REAL NOT NULL DEFAULT 0.00,
    dtotalfinal REAL NOT NULL DEFAULT 0.00,
    dtotpagos REAL NOT NULL DEFAULT 0.00,
    dtotdescuen REAL NOT NULL DEFAULT 0.00,
    dFlete REAL NOT NULL DEFAULT 0.00,
    dtotdev REAL NOT NULL DEFAULT 0.00,
    dvndmtototal REAL NOT NULL DEFAULT 0.00,
    dretencion REAL NOT NULL DEFAULT 0.00,
    dretencioniva REAL NOT NULL DEFAULT 0.00,
    vendedor TEXT NOT NULL DEFAULT '',
    codcoord TEXT NOT NULL DEFAULT '',
    aceptadev INTEGER NOT NULL DEFAULT 1,
    kti_negesp INTEGER NOT NULL DEFAULT 0,
    bsiva REAL NOT NULL DEFAULT 0.00,
    bsflete REAL NOT NULL DEFAULT 0.00,
    bsretencion REAL NOT NULL DEFAULT 0.00,
    bsretencioniva REAL NOT NULL DEFAULT 0.00,
    tasadoc REAL NOT NULL DEFAULT 0.00,
    mtodcto REAL NOT NULL DEFAULT 0.00,
    fchvencedcto TEXT NOT NULL DEFAULT '1000-01-01',
    tienedcto INTEGER NOT NULL DEFAULT 0,
    cbsret REAL NOT NULL DEFAULT 0.00,
    cdret REAL NOT NULL DEFAULT 0.00,
    cbsretiva REAL NOT NULL DEFAULT 0.00,
    cdretiva REAL NOT NULL DEFAULT 0.00,
    cbsrparme REAL NOT NULL DEFAULT 0.00,
    cdrparme REAL NOT NULL DEFAULT 0.00,
    cbsretflete REAL NOT NULL DEFAULT 0.00,
    cdretflete REAL NOT NULL DEFAULT 0.00,
    bsmtoiva REAL NOT NULL DEFAULT 0.00,
    bsmtofte REAL NOT NULL DEFAULT 0.00,
    retmun_mto REAL NOT NULL DEFAULT 0.00,
    dolarflete INTEGER NOT NULL DEFAULT 0,
    bsretflete REAL NOT NULL DEFAULT 0.00,
    dretflete REAL NOT NULL DEFAULT 0.00,
    dretmun_mto REAL NOT NULL DEFAULT 0.00,
    retivaoblig INTEGER NOT NULL DEFAULT 0,
    edoentrega INTEGER NOT NULL DEFAULT 0,
    dmtoiva REAL NOT NULL DEFAULT 0.00,
    prcdctoaplic REAL NOT NULL DEFAULT 0.00,
    montodctodol REAL NOT NULL DEFAULT 0.00,
    montodctobs REAL NOT NULL DEFAULT 0.00,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS closs_document_lines(
    agencia TEXT NOT NULL DEFAULT '',
    tipodoc TEXT NOT NULL DEFAULT '',
    documento TEXT NOT NULL,
    tipodocv TEXT NOT NULL DEFAULT '',
    grupo TEXT NOT NULL DEFAULT '',
    subgrupo TEXT NOT NULL DEFAULT '',
    origen INTEGER NOT NULL DEFAULT 0.00,
    codigo TEXT NOT NULL,
    codhijo TEXT NOT NULL DEFAULT '',
    pid TEXT NOT NULL DEFAULT '',
    nombre TEXT NOT NULL DEFAULT '',
    cantidad INTEGER NOT NULL DEFAULT 0,
    cntdevuelt INTEGER NOT NULL DEFAULT 0,
    vndcntdevuelt INTEGER NOT NULL DEFAULT 0.00,
    dvndmtototal REAL NOT NULL DEFAULT 0.00,
    dpreciofin REAL NOT NULL DEFAULT 0.00,
    dpreciounit REAL NOT NULL DEFAULT 0.00,
    dmontoneto REAL NOT NULL DEFAULT 0.00,
    dmontototal REAL NOT NULL DEFAULT 0.00,
    timpueprc REAL NOT NULL DEFAULT 0.00,
    unidevuelt INTEGER NOT NULL DEFAULT 0,
    fechadoc TEXT NOT NULL DEFAULT '1000-01-01',
    vendedor TEXT NOT NULL DEFAULT '',
    codcoord TEXT NOT NULL DEFAULT '',
    PRIMARY KEY (documento, codigo)
);

CREATE TABLE IF NOT EXISTS closs_group(
    codigo TEXT NOT NULL PRIMARY KEY,
    nombre TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS closs_managers(
    kng_codgcia TEXT NOT NULL DEFAULT '',
    kng_codcoord TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL,
    PRIMARY KEY (kng_codgcia, kng_codcoord)
);

CREATE TABLE IF NOT EXISTS closs_order(
    kti_ndoc TEXT NOT NULL PRIMARY KEY,
    kti_tdoc TEXT NOT NULL DEFAULT '',
    kti_codcli TEXT NOT NULL DEFAULT '',
    kti_nombrecli TEXT NOT NULL DEFAULT '',
    kti_codven TEXT NOT NULL DEFAULT '',
    kti_docsol TEXT NOT NULL DEFAULT '',
    kti_condicion TEXT NOT NULL DEFAULT '',
    kti_tipprec INTEGER NOT NULL DEFAULT 0,
    kti_totneto REAL NOT NULL DEFAULT 0.00,
    kti_status INTEGER NOT NULL DEFAULT 0,
    kti_nroped TEXT NOT NULL DEFAULT '',
    kti_fchdoc TEXT NOT NULL DEFAULT '1000-01-01 00:00:00',
    kti_negesp INTEGER NOT NULL DEFAULT 0,
    ke_pedstatus INTEGER NOT NULL DEFAULT 0,
    dolarflete INTEGER NOT NULL DEFAULT 0,
    complemento INTEGER NOT NULL DEFAULT 0,
    nro_complemento TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS closs_order_lines(
    kti_tdoc TEXT NOT NULL DEFAULT '',
    kti_ndoc TEXT NOT NULL DEFAULT '',
    kti_tipprec INTEGER NOT NULL DEFAULT 1,
    kmv_codart TEXT NOT NULL DEFAULT '',
    kmv_nombre TEXT NOT NULL DEFAULT '',
    kmv_cant INTEGER NOT NULL DEFAULT 0.00,
    kmv_artprec REAL NOT NULL DEFAULT 0.00,
    kmv_stot REAL NOT NULL DEFAULT 0.00,
    kmv_dctolin REAL NOT NULL DEFAULT 0.00,
    PRIMARY KEY (kti_ndoc, kmv_codart)
);

CREATE TABLE IF NOT EXISTS closs_product(
    codigo TEXT NOT NULL PRIMARY KEY,
    grupo TEXT NOT NULL DEFAULT '',
    subgrupo TEXT NOT NULL DEFAULT '',
    nombre TEXT NOT NULL DEFAULT '',
    referencia TEXT NOT NULL UNIQUE DEFAULT '',
    marca TEXT NOT NULL DEFAULT '',
    unidad TEXT NOT NULL DEFAULT '',
    discont INTEGER NOT NULL DEFAULT 0,
    existencia INTEGER NOT NULL DEFAULT 0,
    vta_max INTEGER NOT NULL DEFAULT 0,
    vta_min INTEGER NOT NULL DEFAULT 0,
    vta_minenx INTEGER NOT NULL DEFAULT 0,
    comprometido INTEGER NOT NULL DEFAULT 0,
    precio1 REAL NOT NULL DEFAULT 0.0,
    precio2 REAL NOT NULL DEFAULT 0.0,
    precio3 REAL NOT NULL DEFAULT 0.0,
    precio4 REAL NOT NULL DEFAULT 0.0,
    precio5 REAL NOT NULL DEFAULT 0.0,
    precio6 REAL NOT NULL DEFAULT 0.0,
    precio7 REAL NOT NULL DEFAULT 0.0,
    preventa INTEGER NOT NULL DEFAULT 0,
    dctotope REAL NOT NULL DEFAULT 0.0,
    vta_solofac INTEGER NOT NULL DEFAULT 0,
    vta_solone INTEGER NOT NULL DEFAULT 0,
    codbarras INTEGER NOT NULL DEFAULT 0,
    detalles TEXT NOT NULL DEFAULT '',
    cantbulto INTEGER NOT NULL DEFAULT 0,
    costo_prom REAL NOT NULL DEFAULT 0.0,
    util1 REAL NOT NULL DEFAULT 0.0,
    util2 REAL NOT NULL DEFAULT 0.0,
    util3 REAL NOT NULL DEFAULT 0.0,
    fchultcomp TEXT NOT NULL DEFAULT '',
    qtyultcomp INTEGER NOT NULL DEFAULT 0,
    images TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS closs_salesman(
    codigo TEXT NOT NULL PRIMARY KEY,
    nombre TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL UNIQUE DEFAULT '',
    telefono TEXT NOT NULL DEFAULT '',
    telefonos TEXT NOT NULL DEFAULT '',
    status INTEGER NOT NULL DEFAULT 0,
    supervpor TEXT NOT NULL DEFAULT '',
    sector TEXT NOT NULL DEFAULT '',
    subcodigo TEXT NOT NULL DEFAULT '',
    nivgcial INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS closs_salesman_statistic(
    codcoord TEXT NOT NULL DEFAULT '',
    nomcoord TEXT NOT NULL DEFAULT '',
    vendedor TEXT NOT NULL PRIMARY KEY,
    nombrevend TEXT NOT NULL DEFAULT '',
    cntpedidos INTEGER NOT NULL DEFAULT 0,
    mtopedidos REAL NOT NULL DEFAULT 0.00,
    cntfacturas INTEGER NOT NULL DEFAULT 0,
    mtofacturas REAL NOT NULL DEFAULT 0.00,
    metavend REAL NOT NULL DEFAULT 0.00,
    prcmeta REAL NOT NULL DEFAULT 0.00,
    cntclientes INTEGER NOT NULL DEFAULT 0,
    clivisit INTEGER NOT NULL DEFAULT 0,
    prcvisitas REAL NOT NULL DEFAULT 0.00,
    lom_montovtas REAL NOT NULL DEFAULT 0.00,
    lom_prcvtas REAL NOT NULL DEFAULT 0.00,
    lom_prcvisit REAL NOT NULL DEFAULT 0.00,
    rlom_montovtas REAL NOT NULL DEFAULT 0.00,
    rlom_prcvtas REAL NOT NULL DEFAULT 0.00,
    rlom_prcvisit REAL NOT NULL DEFAULT 0.00,
    fecha_estad TEXT NOT NULL DEFAULT '1000-01-01',
    ppgdol_totneto REAL NOT NULL DEFAULT 0.00,
    devdol_totneto REAL NOT NULL DEFAULT 0.00,
    defdol_totneto REAL NOT NULL DEFAULT 0.00,
    totdolcob REAL NOT NULL DEFAULT 0.00,
    cntrecl INTEGER NOT NULL DEFAULT 0.00,
    mtorecl REAL NOT NULL DEFAULT 0.00,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS closs_sector(
    codigo TEXT NOT NULL PRIMARY KEY,
    zona TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS closs_session(
    token TEXT NOT NULL,
    user_id TEXT NOT NULL PRIMARY KEY,
    FOREIGN KEY (user_id) REFERENCES closs_user(id)
);

CREATE TABLE IF NOT EXISTS closs_subgroup(
    codigo TEXT NOT NULL DEFAULT '',
    subcodigo TEXT NOT NULL UNIQUE,
    nombre TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL,
    PRIMARY KEY (codigo, subcodigo),
    FOREIGN KEY(codigo) REFERENCES closs_group(codigo)
);

CREATE TABLE IF NOT EXISTS closs_subsector(
    codigo TEXT NOT NULL DEFAULT '',
    subcodigo TEXT NOT NULL UNIQUE,
    subsector TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL,
    PRIMARY KEY (codigo, subcodigo),
    FOREIGN KEY (codigo) REFERENCES closs_sector(codigo)
);

CREATE TABLE IF NOT EXISTS closs_user(
    id TEXT NOT NULL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL DEFAULT '',
    codigo TEXT NOT NULL DEFAULT '',
    role TEXT NOT NULL DEFAULT 'SALESMAN',
    ult_sinc TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version TEXT NOT NULL DEFAULT '0.0.0',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL
);

-- +goose Down
DROP TABLE closs_company;
DROP TABLE closs_config;
DROP TABLE closs_customer;
DROP TABLE closs_document;
DROP TABLE closs_document_lines;
DROP TABLE closs_managers;
DROP TABLE closs_order;
DROP TABLE closs_order_lines;
DROP TABLE closs_product;
DROP TABLE closs_salesman;
DROP TABLE closs_salesman_statistic;
DROP TABLE closs_session;
DROP TABLE closs_user;

