// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type Articulo struct {
	IDEmpresa          string
	Agencia            string
	Codigo             string
	Grupo              string
	Subgrupo           string
	Nombre             string
	Nombrecorto        string
	Referencia         string
	Marca              string
	Unidad             string
	ParaWeb            bool
	Costo              string
	Precio1            string
	Precio2            string
	Precio3            string
	Precio4            string
	Precio5            string
	Precio6            string
	Precio7            string
	Precio8            string
	Preciofin1         string
	Preciofin2         string
	Preciofin3         string
	Preciofin4         string
	Preciofin5         string
	Preciofin6         string
	Preciofin7         string
	Preciofin8         string
	Existencia         string
	Rotacionvta        string
	Inactiva           string
	Comprometido       string
	Minimo             string
	Maximo             string
	Origen             bool
	Impuesto           bool
	Comision1          string
	Comision2          string
	Comision3          string
	Comision4          string
	Comision5          string
	Comision6          string
	Comision7          string
	Comision8          string
	Util1              string
	Util2              string
	Util3              string
	Util4              string
	Util5              string
	Util6              string
	Util7              string
	Util8              string
	Formato            string
	Cfob               string
	Cnacionali         string
	Prcnac             string
	Cflete             string
	Prcflete           string
	Carancel           string
	Prcarancel         string
	Cgastoper          string
	Prccosoper         string
	Cotroscos          string
	Ccostocyf          string
	Costoex            string
	Costoref           string
	Precio1ex          string
	Precio2ex          string
	Precio3ex          string
	Precio4ex          string
	Precio5ex          string
	Precio6ex          string
	Precio7ex          string
	Precio8ex          string
	Factor             string
	Garantia           int32
	Unidadgrp          string
	Cntgrp             string
	Costgrp            string
	Precio1grp         string
	Precio2grp         string
	Precio3grp         string
	Precio4grp         string
	Precio5grp         string
	Precio6grp         string
	Precio7grp         string
	Precio8grp         string
	Prcgrp1            string
	Prcgrp2            string
	Prcgrp3            string
	Prcgrp4            string
	Prcgrp5            string
	Prcgrp6            string
	Prcgrp7            string
	Prcgrp8            string
	Creadopor          string
	Modifpor           string
	Margenpor          string
	CostoAnt           string
	CostoProm          string
	Impuesto1          string
	Impuesto2          string
	Impuesto3          string
	Impuesto4          string
	Impuesto5          string
	Impuesto6          string
	Rutafoto           string
	Detalles           sql.NullString
	Contraindi         sql.NullString
	Usointerno         bool
	Aux1               string
	Aux2               string
	Aux3               string
	Simbolo            string
	Modelo             string
	Discont            bool
	Flotante           bool
	Cimpuesto1         string
	Cimpuesto2         string
	Cimpuesto3         string
	Cimpuesto4         string
	Cimpuesto5         string
	Cimpuesto6         string
	Fechacrea          time.Time
	Cuentacont         string
	Metodo             bool
	Sevence            bool
	Usalotes           bool
	Controlado         bool
	Volumen            string
	Peso               string
	Diametro           string
	Usabalanza         bool
	Topeminvent        string
	Topemaxvent        string
	Forma              string
	Esvehiculo         bool
	Aceptadscto        bool
	Aceptacred         bool
	Metodoround        bool
	Unidinamica        bool
	Campo1             string
	Campo2             string
	Campo3             string
	Campo4             string
	Campo5             string
	Campo6             string
	Campo7             string
	Campo8             string
	Campo9             string
	Campo10            string
	ImpNacional        string
	ImpProducc         string
	GimpNacional       string
	GimpProducc        string
	Ctacontinv         string
	Ctacontcostos      string
	Ctacontingresos    string
	Ctacontdevolucion  string
	Blocked            string
	Usatallacolor      string
	DepositoIn         string
	DepositoOut        string
	UndSimpleagrupado  bool
	UsaServidor        bool
	Dadicart01         string
	Dadicart02         string
	Dadicart03         string
	Discontc           bool
	Artcolor           string
	Codarancel         string
	Recargast          bool
	Exislr             bool
	Fechamodifi        time.Time
	Esdinamico         bool
	UsaEscala          bool
	Usacanjepuntos     bool
	Usabulto           bool
	Cantbulto          string
	Nousaweb           bool
	CodcontCmxp        string
	Ctacontdevcomp     string
	Garantiac          string
	Hashtag            sql.NullString
	Ctacontdescuen     string
	Codigotipomed      string
	Codigotipoprod     string
	Usatasapercibido   bool
	Tipomedicamento    bool
	Generapuntos       bool
	Totcantmincomp     string
	Totcantmaxcomp     string
	Totcostmincomp     string
	Totcostmaxcomp     string
	Filtrarcomphashtag string
	DsctoMedicamentos  bool
	Dcosto             string
	DcostoAnt          string
	DcostoProm         string
	Dprecio1           string
	Dprecio2           string
	Dprecio3           string
	Dprecio4           string
	Dprecio5           string
	Dprecio6           string
	Dprecio7           string
	Dprecio8           string
	Dpreciofin1        string
	Dpreciofin2        string
	Dpreciofin3        string
	Dpreciofin4        string
	Dpreciofin5        string
	Dpreciofin6        string
	Dpreciofin7        string
	Liquida            bool
}

type KeSession struct {
	Token  string
	UserID string
}

type Opermv struct {
	IDEmpresa       string
	Agencia         string
	Tipodoc         string
	Documento       string
	Grupo           string
	Subgrupo        string
	Origen          int32
	Codigo          string
	Codhijo         string
	Pid             string
	Nombre          string
	Costounit       string
	Preciounit      string
	Diferencia      string
	Dsctounit       string
	Coddescuento    string
	Dsctoprc        string
	Dsctoendm       string
	Dsctomtolinea   string
	Dsctoendp       string
	Preciofin       string
	Prcrecargo      string
	Recargounit     string
	Preciooriginal  string
	Cantidad        string
	Cntdevuelt      string
	Unidevuelt      string
	Cntentrega      string
	Tallas          string
	Colores         string
	Montoneto       string
	Montototal      string
	Almacen         string
	Proveedor       string
	Fechadoc        time.Time
	Impuesto1       string
	Impuesto2       string
	Impuesto3       string
	Impuesto4       string
	Impuesto5       string
	Impuesto6       string
	Timpueprc       string
	ImpuMto         string
	Comision        string
	Comisprc        string
	Vendedor        string
	Emisor          string
	Usaserial       string
	Tipoprecio      string
	Unidad          string
	Agrupado        int32
	Cntagrupada     string
	Seimporto       int32
	Desdeimpor      string
	Notas           sql.NullString
	Oferta          int32
	Compuesto       int32
	Usaexist        int32
	Marca           string
	Aux1            string
	Estacion        string
	Estacionfac     string
	Clasifica       string
	Cuentacont      string
	Placa           string
	Placaod         string
	Udinamica       int32
	Cantref         string
	Unidadref       string
	Baseimpo1       string
	Baseimpo2       string
	Baseimpo3       string
	Lote            string
	ImpNacional     string
	ImpProducc      string
	SeImprimio      int32
	Ofertaconvenio  string
	CodServidor     string
	PrcComiServidor string
	MtoComiServidor string
	Checkin         string
	HabiNro         string
	Idvalidacion    string
	Idoferta        string
	Agenciant       string
	Tipodocant      string
	Documant        string
	Uemisorant      string
	Estacioant      string
	Fechadocant     time.Time
	Fechayhora      time.Time
	Frog            int32
	Documentolocal  string
	Linbloq         int8
	EnviadoKmonitor int8
	SeGuardo        int8
	Baseimpo4       string
	Baseimpo5       string
	Baseimpo6       string
	Dpreciofin      string
	Dpreciounit     string
	Dprecioorig     string
	Dmontoneto      string
	Dmontototal     string
	Ddctounit       string
	Dmtodctolin     string
	Dcostounit      string
	Dbaseimpo1      string
	Dbaseimpo2      string
	Dbaseimpo3      string
}

type Operti struct {
	IDEmpresa            string
	Agencia              string
	Tipodoc              string
	Subtipodoc           string
	Moneda               string
	Documento            string
	Codcliente           string
	Nombrecli            string
	Contacto             string
	Comprador            string
	Rif                  string
	Nit                  string
	Direccion            string
	Telefonos            string
	Tipoprecio           string
	Orden                string
	Emision              time.Time
	Recepcion            time.Time
	Vence                time.Time
	UltInteres           time.Time
	Fechacrea            time.Time
	Totcosto             string
	Totcomi              string
	Totbruto             string
	Totneto              string
	Totpagos             string
	Totalfinal           string
	Totimpuest           string
	Totdescuen           string
	Impuesto1            string
	Impuesto2            string
	Impuesto3            string
	Impuesto4            string
	Impuesto5            string
	Impuesto6            string
	Recargos             string
	Dsctoend1            string
	Dsctoend2            string
	Dsctolinea           string
	Notas                sql.NullString
	Estatusdoc           string
	Ultimopag            time.Time
	Diascred             string
	Vendedor             string
	Factorcamb           string
	MultiDiv             int32
	Factorreferencial    string
	Fechanul             time.Time
	Uanulador            string
	Uemisor              string
	Estacion             string
	Sinimpuest           string
	Almacen              string
	Sector               string
	Formafis             int32
	AlLibro              int32
	Codigoret            string
	Retencion            string
	Aux1                 string
	Aux2                 string
	Aplicadoa            string
	Referencia           string
	Refmanual            string
	DocClassID           string
	Operac               string
	Motanul              sql.NullString
	Seimporto            int32
	Dbcr                 int32
	Horadocum            string
	Ampm                 int32
	Tranferido           int32
	Procedecre           int32
	Entregado            int32
	Vuelto               string
	Integrado            int32
	Escredito            int32
	SeqNodo              string
	TipoNc               string
	Porbackord           string
	Chequedev            string
	Ordentrab            string
	Compcont             string
	Planillacob          string
	Nodoremoto           string
	Turno                string
	CodvendA             string
	CodvendB             string
	CodvendC             string
	CodvendD             string
	Baseimpo1            string
	Baseimpo2            string
	Baseimpo3            string
	Iddocumento          string
	ImpNacional          string
	ImpProducc           string
	Retencioniva         string
	Fechayhora           time.Time
	Tipopersona          int32
	Idvalidacion         string
	Nosujeto             int32
	Serialprintf         string
	Documentofiscal      string
	Numeroz              string
	Ubica                string
	UsaDespacho          int32
	Despachador          string
	DespachoNro          string
	Checkin              string
	Nureserva            string
	Grandocnum           string
	Agenciant            string
	Tipodocant           string
	Documant             string
	Uemisorant           string
	Estacioant           string
	Emisionant           time.Time
	Fchyhrant            time.Time
	Frog                 int32
	ApaNc                int32
	Documentolocal       string
	ComandaMovil         int32
	ComandaKmonitor      int32
	ParaLlevar           int32
	Notimbrar            int32
	Antipo               string
	Antdoc               string
	Xrequest             sql.NullString
	Xresponse            sql.NullString
	Parcialidad          string
	Cedcompra            string
	Subcodigo            string
	Cprefijoserie        string
	Contingencia         bool
	PrectaMovil          bool
	Tipodocfiscal        bool
	Cprefijodeserie      string
	Cserie               string
	Serieincluyeimpuesto int8
	Serieauto            bool
	Opemail              string
	Refmanual2           string
	Baseimpo4            string
	Baseimpo5            string
	Baseimpo6            string
	KePedstatus          string
	Fechamodifi          time.Time
	NoAplicaCxcDiv       int32
	PrcIgtf6687          string
	BaseIgtf6687         string
	MontoIgtf6687        string
	Base2Igtf6687        string
	Monto2Igtf6687       string
	TodoIgtf6687         int8
	Todo2Igtf6687        int8
}
