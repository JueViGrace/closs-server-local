-- +goose Up
CREATE TABLE IF NOT EXISTS articulo (
  id_empresa varchar(6) NOT NULL DEFAULT ' ',
  agencia varchar(3) NOT NULL DEFAULT ' ',
  codigo varchar(25) NOT NULL DEFAULT ' ',
  grupo varchar(6) NOT NULL DEFAULT ' ',
  subgrupo varchar(6) NOT NULL DEFAULT ' ',
  nombre char(150) NOT NULL DEFAULT '',
  nombrecorto varchar(20) NOT NULL DEFAULT ' ',
  referencia varchar(20) NOT NULL DEFAULT ' ',
  marca varchar(20) NOT NULL DEFAULT ' ',
  unidad varchar(15) NOT NULL DEFAULT ' ',
  para_web tinyint(1) NOT NULL DEFAULT '0',
  costo decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio7 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio8 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciofin1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciofin2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciofin3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciofin4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciofin5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciofin6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciofin7 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciofin8 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  existencia decimal(20,7) NOT NULL DEFAULT '0.0000000',
  rotacionvta decimal(20,7) NOT NULL DEFAULT '0.0000000',
  inactiva decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comprometido decimal(20,7) NOT NULL DEFAULT '0.0000000',
  minimo decimal(20,7) NOT NULL DEFAULT '0.0000000',
  maximo decimal(20,7) NOT NULL DEFAULT '0.0000000',
  origen tinyint(1) NOT NULL DEFAULT '0',
  impuesto tinyint(1) NOT NULL DEFAULT '0',
  comision1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comision2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comision3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comision4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comision5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comision6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comision7 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comision8 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  util1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  util2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  util3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  util4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  util5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  util6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  util7 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  util8 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  formato varchar(10) NOT NULL DEFAULT ' ',
  cfob decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cnacionali decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcnac decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cflete decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcflete decimal(20,7) NOT NULL DEFAULT '0.0000000',
  carancel decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcarancel decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cgastoper decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prccosoper decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cotroscos decimal(20,7) NOT NULL DEFAULT '0.0000000',
  ccostocyf decimal(20,7) NOT NULL DEFAULT '0.0000000',
  costoex decimal(20,7) NOT NULL DEFAULT '0.0000000',
  costoref decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio1ex decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio2ex decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio3ex decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio4ex decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio5ex decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio6ex decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio7ex decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio8ex decimal(20,7) NOT NULL DEFAULT '0.0000000',
  factor decimal(20,7) NOT NULL DEFAULT '0.0000000',
  garantia int(5) NOT NULL DEFAULT '0',
  unidadgrp varchar(15) NOT NULL DEFAULT ' ',
  cntgrp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  costgrp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio1grp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio2grp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio3grp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio4grp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio5grp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio6grp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio7grp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  precio8grp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcgrp1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcgrp2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcgrp3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcgrp4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcgrp5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcgrp6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcgrp7 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcgrp8 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  creadopor varchar(30) NOT NULL DEFAULT '',
  modifpor varchar(30) NOT NULL DEFAULT '',
  margenpor varchar(1) NOT NULL DEFAULT ' ',
  costo_ant decimal(20,7) NOT NULL DEFAULT '0.0000000',
  costo_prom decimal(20,2) NOT NULL DEFAULT '0.00',
  impuesto1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  rutafoto varchar(200) NOT NULL DEFAULT ' ',
  detalles text,
  contraindi text,
  usointerno tinyint(1) NOT NULL DEFAULT '0',
  aux1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  aux2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  aux3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  simbolo varchar(20) NOT NULL DEFAULT ' ',
  modelo varchar(20) NOT NULL DEFAULT ' ',
  discont tinyint(1) NOT NULL DEFAULT '0',
  flotante tinyint(1) NOT NULL DEFAULT '0',
  cimpuesto1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cimpuesto2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cimpuesto3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cimpuesto4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cimpuesto5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cimpuesto6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  fechacrea date NOT NULL DEFAULT '1901-01-01',
  cuentacont varchar(40) NOT NULL DEFAULT ' ',
  metodo tinyint(1) NOT NULL DEFAULT '0',
  sevence tinyint(1) NOT NULL DEFAULT '0',
  usalotes tinyint(1) NOT NULL DEFAULT '0',
  controlado tinyint(1) NOT NULL DEFAULT '0',
  volumen decimal(20,7) NOT NULL DEFAULT '0.0000000',
  peso decimal(20,7) NOT NULL DEFAULT '0.0000000',
  diametro varchar(15) NOT NULL DEFAULT ' ',
  usabalanza tinyint(1) NOT NULL DEFAULT '0',
  topeminvent decimal(20,7) NOT NULL DEFAULT '0.0000000',
  topemaxvent decimal(20,7) NOT NULL DEFAULT '0.0000000',
  forma varchar(15) NOT NULL DEFAULT ' ',
  esvehiculo tinyint(1) NOT NULL DEFAULT '0',
  aceptadscto tinyint(1) NOT NULL DEFAULT '0',
  aceptacred tinyint(1) NOT NULL DEFAULT '0',
  metodoround tinyint(1) NOT NULL DEFAULT '0',
  unidinamica tinyint(1) NOT NULL DEFAULT '0',
  campo1 varchar(100) NOT NULL DEFAULT ' ',
  campo2 varchar(100) NOT NULL DEFAULT ' ',
  campo3 varchar(100) NOT NULL DEFAULT ' ',
  campo4 varchar(100) NOT NULL DEFAULT ' ',
  campo5 varchar(100) NOT NULL DEFAULT ' ',
  campo6 varchar(100) NOT NULL DEFAULT ' ',
  campo7 varchar(100) NOT NULL DEFAULT ' ',
  campo8 varchar(100) NOT NULL DEFAULT ' ',
  campo9 varchar(100) NOT NULL DEFAULT ' ',
  campo10 varchar(100) NOT NULL DEFAULT ' ',
  imp_nacional decimal(20,7) NOT NULL DEFAULT '0.0000000',
  imp_producc decimal(20,7) NOT NULL DEFAULT '0.0000000',
  gimp_nacional decimal(20,7) NOT NULL DEFAULT '0.0000000',
  gimp_producc decimal(20,7) NOT NULL DEFAULT '0.0000000',
  ctacontinv varchar(40) NOT NULL DEFAULT ' ',
  ctacontcostos varchar(40) NOT NULL DEFAULT ' ',
  ctacontingresos varchar(40) NOT NULL DEFAULT ' ',
  ctacontdevolucion varchar(40) NOT NULL DEFAULT ' ',
  blocked decimal(4,0) NOT NULL DEFAULT '0',
  usatallacolor decimal(4,0) NOT NULL DEFAULT '0',
  deposito_in varchar(2) NOT NULL DEFAULT ' ',
  deposito_out varchar(2) NOT NULL DEFAULT ' ',
  und_simpleagrupado tinyint(1) NOT NULL DEFAULT '1',
  usa_servidor tinyint(1) NOT NULL DEFAULT '0',
  dadicart01 varchar(100) NOT NULL DEFAULT ' ',
  dadicart02 varchar(100) NOT NULL DEFAULT ' ',
  dadicart03 varchar(100) NOT NULL DEFAULT ' ',
  discontc tinyint(1) NOT NULL DEFAULT '0',
  artcolor decimal(10,0) NOT NULL DEFAULT '0',
  codarancel varchar(16) NOT NULL DEFAULT ' ',
  recargast tinyint(1) NOT NULL DEFAULT '0',
  exislr tinyint(1) NOT NULL DEFAULT '0',
  fechamodifi datetime NOT NULL DEFAULT '1901-01-01 00:00:01',
  esdinamico tinyint(1) NOT NULL DEFAULT '0',
  usa_escala tinyint(1) NOT NULL DEFAULT '0',
  usacanjepuntos tinyint(1) NOT NULL DEFAULT '0',
  usabulto tinyint(1) NOT NULL DEFAULT '0',
  cantbulto decimal(20,7) NOT NULL DEFAULT '0.0000000',
  nousaweb tinyint(1) NOT NULL DEFAULT '0',
  codcont_cmxp varchar(40) NOT NULL DEFAULT ' ',
  ctacontdevcomp varchar(40) NOT NULL DEFAULT ' ',
  garantiac decimal(5,0) NOT NULL DEFAULT '0',
  hashtag text,
  ctacontdescuen varchar(40) NOT NULL DEFAULT ' ',
  codigotipomed varchar(20) NOT NULL DEFAULT '',
  codigotipoprod varchar(20) NOT NULL DEFAULT '',
  usatasapercibido tinyint(1) NOT NULL DEFAULT '0',
  tipomedicamento tinyint(1) NOT NULL DEFAULT '0',
  generapuntos tinyint(1) NOT NULL DEFAULT '0',
  totcantmincomp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  totcantmaxcomp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  totcostmincomp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  totcostmaxcomp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  filtrarcomphashtag varchar(50) NOT NULL DEFAULT ' ',
  dscto_medicamentos tinyint(1) NOT NULL DEFAULT '0',
  dcosto decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dcosto_ant decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dcosto_prom decimal(20,2) NOT NULL DEFAULT '0.00',
  dprecio1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dprecio2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dprecio3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dprecio4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dprecio5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dprecio6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dprecio7 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dprecio8 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dpreciofin1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dpreciofin2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dpreciofin3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dpreciofin4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dpreciofin5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dpreciofin6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dpreciofin7 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  liquida tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (codigo),
  KEY nombre (nombre)
);

CREATE TABLE IF NOT EXISTS operti (
  id_empresa varchar(6) NOT NULL DEFAULT '',
  agencia varchar(3) NOT NULL DEFAULT ' ',
  tipodoc varchar(3) NOT NULL DEFAULT ' ',
  subtipodoc varchar(4) NOT NULL DEFAULT '',
  moneda varchar(3) NOT NULL DEFAULT '',
  documento varchar(8) NOT NULL DEFAULT ' ',
  codcliente varchar(20) NOT NULL DEFAULT ' ',
  nombrecli varchar(100) NOT NULL DEFAULT '',
  contacto varchar(50) NOT NULL DEFAULT '',
  comprador varchar(20) NOT NULL DEFAULT '',
  rif varchar(25) NOT NULL DEFAULT '',
  nit varchar(25) NOT NULL DEFAULT '',
  direccion varchar(200) NOT NULL DEFAULT '',
  telefonos varchar(50) NOT NULL DEFAULT '',
  tipoprecio decimal(2,0) NOT NULL DEFAULT '0',
  orden varchar(15) NOT NULL DEFAULT '',
  emision date NOT NULL DEFAULT '0000-00-00',
  recepcion date NOT NULL DEFAULT '0000-00-00',
  vence date NOT NULL DEFAULT '0000-00-00',
  ult_interes date NOT NULL DEFAULT '0000-00-00',
  fechacrea date NOT NULL DEFAULT '0000-00-00',
  totcosto decimal(24,7) NOT NULL DEFAULT '0.0000000',
  totcomi decimal(24,7) NOT NULL DEFAULT '0.0000000',
  totbruto decimal(24,7) NOT NULL DEFAULT '0.0000000',
  totneto decimal(24,7) NOT NULL DEFAULT '0.0000000',
  totpagos decimal(24,7) NOT NULL DEFAULT '0.0000000',
  totalfinal decimal(24,7) NOT NULL DEFAULT '0.0000000',
  totimpuest decimal(24,7) NOT NULL DEFAULT '0.0000000',
  totdescuen decimal(24,7) NOT NULL DEFAULT '0.0000000',
  impuesto1 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  impuesto2 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  impuesto3 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  impuesto4 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  impuesto5 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  impuesto6 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  recargos decimal(24,7) NOT NULL DEFAULT '0.0000000',
  dsctoend1 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  dsctoend2 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  dsctolinea decimal(24,7) NOT NULL DEFAULT '0.0000000',
  notas text,
  estatusdoc varchar(1) NOT NULL DEFAULT '0',
  ultimopag date NOT NULL DEFAULT '0000-00-00',
  diascred decimal(5,0) NOT NULL DEFAULT '0',
  vendedor varchar(8) NOT NULL DEFAULT '',
  factorcamb decimal(24,7) NOT NULL DEFAULT '0.0000000',
  multi_div int(1) NOT NULL DEFAULT '0',
  factorreferencial decimal(24,7) NOT NULL DEFAULT '0.0000000',
  fechanul date NOT NULL DEFAULT '0000-00-00',
  uanulador varchar(30) NOT NULL DEFAULT '',
  uemisor varchar(30) NOT NULL DEFAULT '',
  estacion varchar(3) NOT NULL DEFAULT '',
  sinimpuest decimal(24,7) NOT NULL DEFAULT '0.0000000',
  almacen varchar(2) NOT NULL DEFAULT '',
  sector varchar(6) NOT NULL DEFAULT '',
  formafis int(1) NOT NULL DEFAULT '0',
  al_libro int(1) NOT NULL DEFAULT '0',
  codigoret varchar(3) NOT NULL DEFAULT '',
  retencion decimal(24,7) NOT NULL DEFAULT '0.0000000',
  aux1 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  aux2 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  aplicadoa varchar(11) NOT NULL DEFAULT '',
  referencia varchar(20) NOT NULL DEFAULT '',
  refmanual varchar(20) NOT NULL DEFAULT '',
  doc_class_id varchar(30) NOT NULL DEFAULT '',
  operac varchar(8) NOT NULL DEFAULT '',
  motanul text,
  seimporto int(1) NOT NULL DEFAULT '0',
  dbcr int(1) NOT NULL DEFAULT '0',
  horadocum varchar(5) NOT NULL DEFAULT '',
  ampm int(1) NOT NULL DEFAULT '0',
  tranferido int(1) NOT NULL DEFAULT '0',
  procedecre int(1) NOT NULL DEFAULT '0',
  entregado int(1) NOT NULL DEFAULT '0',
  vuelto decimal(24,7) NOT NULL DEFAULT '0.0000000',
  integrado int(1) NOT NULL DEFAULT '0',
  escredito int(1) NOT NULL DEFAULT '0',
  seq_nodo varchar(10) NOT NULL DEFAULT '',
  tipo_nc varchar(3) NOT NULL DEFAULT '',
  porbackord decimal(4,0) NOT NULL DEFAULT '0',
  chequedev decimal(4,0) NOT NULL DEFAULT '0',
  ordentrab varchar(8) NOT NULL DEFAULT '',
  compcont varchar(20) NOT NULL DEFAULT '',
  planillacob varchar(8) NOT NULL DEFAULT '',
  nodoremoto varchar(3) NOT NULL DEFAULT '',
  turno varchar(2) NOT NULL DEFAULT '',
  codvend_a varchar(8) NOT NULL DEFAULT '',
  codvend_b varchar(8) NOT NULL DEFAULT '',
  codvend_c varchar(8) NOT NULL DEFAULT '',
  codvend_d varchar(8) NOT NULL DEFAULT '',
  baseimpo1 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  baseimpo2 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  baseimpo3 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  iddocumento varchar(11) NOT NULL DEFAULT '',
  imp_nacional decimal(24,7) NOT NULL DEFAULT '0.0000000',
  imp_producc decimal(24,7) NOT NULL DEFAULT '0.0000000',
  retencioniva decimal(24,7) NOT NULL DEFAULT '0.0000000',
  fechayhora datetime NOT NULL DEFAULT '1901-01-01 00:00:01',
  tipopersona int(1) NOT NULL DEFAULT '0',
  idvalidacion varchar(12) NOT NULL DEFAULT '',
  nosujeto int(1) NOT NULL DEFAULT '0',
  serialprintf varchar(25) NOT NULL DEFAULT '',
  documentofiscal varchar(200) NOT NULL DEFAULT '',
  numeroz varchar(8) NOT NULL DEFAULT ' ',
  ubica varchar(30) NOT NULL DEFAULT '',
  usa_despacho int(1) NOT NULL DEFAULT '0',
  despachador varchar(15) NOT NULL DEFAULT '',
  despacho_nro varchar(8) NOT NULL DEFAULT '',
  checkin varchar(12) NOT NULL DEFAULT '',
  nureserva varchar(12) NOT NULL DEFAULT '',
  grandocnum varchar(8) NOT NULL DEFAULT ' ',
  agenciant varchar(3) NOT NULL DEFAULT ' ',
  tipodocant varchar(3) NOT NULL DEFAULT ' ',
  documant varchar(8) NOT NULL DEFAULT ' ',
  uemisorant varchar(30) NOT NULL DEFAULT '',
  estacioant varchar(3) NOT NULL DEFAULT '',
  emisionant date NOT NULL DEFAULT '0000-00-00',
  fchyhrant datetime NOT NULL DEFAULT '1901-01-01 00:00:01',
  frog int(1) NOT NULL DEFAULT '0',
  apa_nc int(1) NOT NULL DEFAULT '0',
  documentolocal varchar(200) NOT NULL DEFAULT '',
  comanda_movil int(1) NOT NULL DEFAULT '1',
  comanda_kmonitor int(1) NOT NULL DEFAULT '1',
  para_llevar int(1) NOT NULL DEFAULT '0',
  notimbrar int(1) NOT NULL DEFAULT '0',
  antipo varchar(3) NOT NULL DEFAULT ' ',
  antdoc varchar(8) NOT NULL DEFAULT ' ',
  xrequest text,
  xresponse text,
  parcialidad decimal(4,0) NOT NULL DEFAULT '0',
  cedcompra varchar(20) NOT NULL DEFAULT '',
  subcodigo varchar(6) NOT NULL DEFAULT '',
  cprefijoserie varchar(4) NOT NULL DEFAULT '',
  contingencia tinyint(1) NOT NULL DEFAULT '0',
  precta_movil tinyint(1) NOT NULL DEFAULT '0',
  tipodocfiscal tinyint(1) NOT NULL DEFAULT '0',
  cprefijodeserie varchar(4) NOT NULL DEFAULT ' ',
  cserie varchar(2) NOT NULL DEFAULT ' ',
  serieincluyeimpuesto tinyint(3) NOT NULL DEFAULT '0',
  serieauto tinyint(1) NOT NULL DEFAULT '0',
  opemail varchar(1024) NOT NULL DEFAULT '',
  refmanual2 varchar(30) NOT NULL DEFAULT '',
  baseimpo4 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  baseimpo5 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  baseimpo6 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  ke_pedstatus char(2) NOT NULL DEFAULT '',
  fechamodifi datetime NOT NULL DEFAULT '1901-01-01 00:00:01',
  no_aplica_cxc_div int(1) NOT NULL DEFAULT '0',
  prc_igtf_6687 decimal(7,2) NOT NULL DEFAULT '0.00',
  base_igtf_6687 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  monto_igtf_6687 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  base2_igtf_6687 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  monto2_igtf_6687 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  todo_igtf_6687 tinyint(4) NOT NULL DEFAULT '0',
  todo2_igtf_6687 tinyint(4) NOT NULL DEFAULT '0',
  KEY operti_ppal (agencia,tipodoc,documento,fechayhora,notimbrar),
  KEY operti_tipodoc_idx (tipodoc,documento),
  KEY codcliente (codcliente),
  KEY tipodoc (tipodoc,documento),
  KEY calcomprom (agencia,tipodoc,documento,emision),
  KEY recalculo (tipodoc,estatusdoc)
);

CREATE TABLE IF NOT EXISTS opermv (
  id_empresa varchar(6) NOT NULL DEFAULT '',
  agencia varchar(3) NOT NULL DEFAULT '',
  tipodoc varchar(3) NOT NULL DEFAULT '',
  documento varchar(8) NOT NULL DEFAULT '',
  grupo varchar(6) NOT NULL DEFAULT '',
  subgrupo varchar(6) NOT NULL DEFAULT '',
  origen int(1) NOT NULL DEFAULT '0',
  codigo varchar(25) NOT NULL DEFAULT '',
  codhijo varchar(25) NOT NULL DEFAULT '',
  pid varchar(12) NOT NULL DEFAULT '',
  nombre varchar(150) NOT NULL DEFAULT '',
  costounit decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciounit decimal(20,7) NOT NULL DEFAULT '0.0000000',
  diferencia decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dsctounit decimal(20,7) NOT NULL DEFAULT '0.0000000',
  coddescuento varchar(6) NOT NULL DEFAULT '',
  dsctoprc decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dsctoendm decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dsctomtolinea decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dsctoendp decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciofin decimal(20,7) NOT NULL DEFAULT '0.0000000',
  prcrecargo decimal(20,7) NOT NULL DEFAULT '0.0000000',
  recargounit decimal(20,7) NOT NULL DEFAULT '0.0000000',
  preciooriginal decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cantidad decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cntdevuelt decimal(20,7) NOT NULL DEFAULT '0.0000000',
  unidevuelt decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cntentrega decimal(20,7) NOT NULL DEFAULT '0.0000000',
  tallas varchar(3) NOT NULL DEFAULT '',
  colores varchar(3) NOT NULL DEFAULT '',
  montoneto decimal(20,7) NOT NULL DEFAULT '0.0000000',
  montototal decimal(20,7) NOT NULL DEFAULT '0.0000000',
  almacen varchar(2) NOT NULL DEFAULT '01',
  proveedor varchar(20) NOT NULL DEFAULT '',
  fechadoc date NOT NULL DEFAULT '1901-01-01',
  impuesto1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto4 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto5 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impuesto6 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  timpueprc decimal(20,7) NOT NULL DEFAULT '0.0000000',
  impu_mto decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comision decimal(20,7) NOT NULL DEFAULT '0.0000000',
  comisprc decimal(20,7) NOT NULL DEFAULT '0.0000000',
  vendedor varchar(8) NOT NULL DEFAULT '',
  emisor varchar(30) NOT NULL DEFAULT '',
  usaserial decimal(2,0) NOT NULL DEFAULT '0',
  tipoprecio decimal(2,0) NOT NULL DEFAULT '0',
  unidad varchar(25) NOT NULL DEFAULT '',
  agrupado int(1) NOT NULL DEFAULT '0',
  cntagrupada decimal(20,7) NOT NULL DEFAULT '0.0000000',
  seimporto int(1) NOT NULL DEFAULT '0',
  desdeimpor varchar(11) NOT NULL DEFAULT '',
  notas text,
  oferta int(1) NOT NULL DEFAULT '0',
  compuesto int(1) NOT NULL DEFAULT '0',
  usaexist int(1) NOT NULL DEFAULT '0',
  marca decimal(4,0) NOT NULL DEFAULT '0',
  aux1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  estacion varchar(3) NOT NULL DEFAULT '',
  estacionfac varchar(3) NOT NULL DEFAULT '',
  clasifica decimal(2,0) NOT NULL DEFAULT '0',
  cuentacont varchar(40) NOT NULL DEFAULT '',
  placa varchar(20) NOT NULL DEFAULT '',
  placaod varchar(20) NOT NULL DEFAULT '',
  udinamica int(1) NOT NULL DEFAULT '0',
  cantref decimal(20,7) NOT NULL DEFAULT '0.0000000',
  unidadref varchar(15) NOT NULL DEFAULT '',
  baseimpo1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  baseimpo2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  baseimpo3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  lote varchar(30) NOT NULL DEFAULT '',
  imp_nacional decimal(20,7) NOT NULL DEFAULT '0.0000000',
  imp_producc decimal(20,7) NOT NULL DEFAULT '0.0000000',
  se_imprimio int(1) NOT NULL DEFAULT '0',
  ofertaconvenio decimal(20,7) NOT NULL DEFAULT '0.0000000',
  cod_servidor varchar(8) NOT NULL DEFAULT '',
  prc_comi_servidor decimal(20,7) NOT NULL DEFAULT '0.0000000',
  mto_comi_servidor decimal(20,7) NOT NULL DEFAULT '0.0000000',
  checkin varchar(12) NOT NULL DEFAULT '',
  habi_nro varchar(12) NOT NULL DEFAULT '',
  idvalidacion varchar(12) NOT NULL DEFAULT '',
  idoferta varchar(8) NOT NULL DEFAULT '',
  agenciant varchar(3) NOT NULL DEFAULT ' ',
  tipodocant varchar(3) NOT NULL DEFAULT ' ',
  documant varchar(8) NOT NULL DEFAULT ' ',
  uemisorant varchar(30) NOT NULL DEFAULT '',
  estacioant varchar(3) NOT NULL DEFAULT '',
  fechadocant date NOT NULL DEFAULT '1901-01-01',
  fechayhora datetime NOT NULL DEFAULT '1901-01-01 00:00:01',
  frog int(1) NOT NULL DEFAULT '0',
  documentolocal varchar(20) NOT NULL DEFAULT ' ',
  linbloq tinyint(3) NOT NULL DEFAULT '0',
  enviado_kmonitor tinyint(3) NOT NULL DEFAULT '1',
  se_guardo tinyint(3) NOT NULL DEFAULT '0',
  baseimpo4 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  baseimpo5 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  baseimpo6 decimal(24,7) NOT NULL DEFAULT '0.0000000',
  dpreciofin decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dpreciounit decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dprecioorig decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dmontoneto decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dmontototal decimal(20,7) NOT NULL DEFAULT '0.0000000',
  ddctounit decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dmtodctolin decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dcostounit decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dbaseimpo1 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dbaseimpo2 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  dbaseimpo3 decimal(20,7) NOT NULL DEFAULT '0.0000000',
  KEY opermv_ppal (agencia,tipodoc,documento,codigo,codhijo,pid),
  KEY tipodoc (tipodoc,documento,codigo,codhijo,pid),
  KEY opermv_documento_idx (documento,tipodoc),
  KEY codigoart (codigo),
  KEY cliente (proveedor),
  KEY id_empresa (agencia,tipodoc,documento,grupo,codigo,codhijo,pid),
  KEY key_hoteles (agencia,codigo,checkin,fechadoc),
  KEY opermvnewppal (agencia,tipodoc,documento),
  KEY opmvcodhijo (codigo,codhijo),
  KEY opmvcodhijo2 (codhijo,codigo)
);

CREATE TABLE IF NOT EXISTS ke_wusuarios (
  username varchar(25) NOT NULL DEFAULT '',
  userid varchar(15) NOT NULL DEFAULT '',
  desactivo tinyint(1) NOT NULL DEFAULT 0,
  sevence tinyint(1) NOT NULL DEFAULT 0,
  vigenciadesde date NOT NULL DEFAULT '0000-00-00',
  vigenciahasta date NOT NULL DEFAULT '0000-00-00',
  cedula varchar(14) NOT NULL DEFAULT '',
  fechanac date NOT NULL DEFAULT '0000-00-00',
  nombre varchar(50) NOT NULL DEFAULT '',
  direccion varchar(170) NOT NULL DEFAULT '',
  telefono varchar(30) NOT NULL DEFAULT '',
  email varchar(30) NOT NULL DEFAULT '',
  mail_servidor varchar(30) NOT NULL DEFAULT '',
  mail_puerto int(4) NOT NULL DEFAULT 25,
  mail_login varchar(30) NOT NULL DEFAULT '',
  mail_password varchar(20) NOT NULL DEFAULT '',
  mail_autenticacion tinyint(1) NOT NULL DEFAULT 0,
  mail_cifrada tinyint(1) NOT NULL DEFAULT 0,
  cliente varchar(20) NOT NULL DEFAULT '',
  vendedor varchar(8) NOT NULL DEFAULT '',
  almacen varchar(2) NOT NULL DEFAULT '',
  agenciausu varchar(3) NOT NULL DEFAULT '',
  formatofac varchar(25) NOT NULL DEFAULT '',
  copiasfac tinyint(1) NOT NULL DEFAULT 1,
  formatofac2 varchar(25) NOT NULL DEFAULT '',
  copiasfac2 tinyint(1) NOT NULL DEFAULT 0,
  rutaquery varchar(254) NOT NULL DEFAULT '',
  rutareporte varchar(254) NOT NULL DEFAULT '',
  password_app varchar(32) NOT NULL DEFAULT '' COMMENT 'Contraseña de la APP Movil. Encriptada en MD5 (con Base 64).',
  tipo int(11) NOT NULL DEFAULT 0 COMMENT 'Tipo de Usuario, segun listado combinado (Por los momentos solo va a existir:1=Pickeador)',
  fechamodifi datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (username)
);

-- +goose Down


