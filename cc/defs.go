package cc

import (
	"unsafe"
)

const (
	BUFSIZ     = 8192
	NSYMB      = 500
	NHASH      = 1024
	STRINGSZ   = 200
	HISTSZ     = 20
	YYMAXDEPTH = 500
	NTERM      = 10
	MAXALIGN   = 7
	BITS       = 5
	NVAR       = BITS * 32

	SIGNONE   = 0
	SIGDONE   = 1
	SIGINTERN = 2

	SIGNINTERN = 1729 * 325 * 1729

	Plan9 = 1 << iota
	Unix
	Windows
)

func SIGN(n uint64) uint64 {
	return 1 << (n - 1)
}
func MASK(n uint64) uint64 {
	return SIGN(n) | (SIGN(n) - 1)
}

type Bits struct {
	b [BITS]uint32
}
type Node struct {
	left    *Node
	right   *Node
	label   unsafe.Pointer
	pc      int32
	reg     int
	xoffset int32
	fconst  float64 /* fp constant */
	vconst  int64   /* non fp const */
	cstring []byte  /* character string */
	rstring []rune  /* rune string */

	Sym     Sym
	Type    *Type
	lineno  int32
	op      byte
	oldop   byte
	xcast   byte
	class   byte
	etype   byte
	complex byte
	addable byte
	scale   byte
	garb    byte
}

type Sym struct {
	link      *Sym
	Type      *Type
	suetag    *Type
	tenum     *Type
	macro     []byte
	varlineno int32
	offset    int32
	vconst    int64
	fconst    float64
	label     *Node
	lexical   uint16
	name      []byte
	block     uint16
	sueblock  uint16
	class     byte
	sym       byte
	aused     byte
	sig       byte
	dataflag  byte
}

type Decl struct {
	link      *Decl
	sym       *Sym
	Type      *Type
	varlineno int32
	offset    int32
	val       int16
	block     uint16
	class     byte
	aused     byte
}

type Type struct {
	sym    *Sym
	tag    *Sym
	funct  *Funct
	link   *Type
	down   *Type
	width  int32
	offset int32
	lineno int32
	shift  byte
	nbits  byte
	etype  byte
	garb   byte
	align  byte
}

//#define	NODECL	((void(*)(int, Type*, Sym*))0)

/* general purpose initialization */
type Init struct {
	code  int
	value uint32
	s     []byte
}

/*
EXTERN struct
{
	char*	p;
	int	c;
} fi;
*/

type Io struct {
	link *Io
	p    []byte
	b    [BUFSIZ]byte
	c    int16
	f    int16
}

type Hist struct {
	link   *Hist
	name   string
	line   int32
	offset int32
}

type Term struct {
	mult int64
	Node *Node
}

const (
	Axxx = iota
	Ael1
	Ael2
	Asu2
	Aarg0
	Aarg1
	Aarg2
	Aaut3
	NALIGN
)

const (
	DMARK = iota
	DAUTO
	DSUE
	DLABEL
)
const (
	OXXX = iota
	OADD
	OADDR
	OAND
	OANDAND
	OARRAY
	OAS
	OASI
	OASADD
	OASAND
	OASASHL
	OASASHR
	OASDIV
	OASHL
	OASHR
	OASLDIV
	OASLMOD
	OASLMUL
	OASLSHR
	OASMOD
	OASMUL
	OASOR
	OASSUB
	OASXOR
	OBIT
	OBREAK
	OCASE
	OCAST
	OCOMMA
	OCOND
	OCONST
	OCONTINUE
	ODIV
	ODOT
	ODOTDOT
	ODWHILE
	OENUM
	OEQ
	OEXREG
	OFOR
	OFUNC
	OGE
	OGOTO
	OGT
	OHI
	OHS
	OIF
	OIND
	OINDREG
	OINIT
	OLABEL
	OLDIV
	OLE
	OLIST
	OLMOD
	OLMUL
	OLO
	OLS
	OLSHR
	OLT
	OMOD
	OMUL
	ONAME
	ONE
	ONOT
	OOR
	OOROR
	OPOSTDEC
	OPOSTINC
	OPREDEC
	OPREINC
	OPREFETCH
	OPROTO
	OREGISTER
	ORETURN
	OSET
	OSIGN
	OSIZE
	OSTRING
	OLSTRING
	OSTRUCT
	OSUB
	OSWITCH
	OUNION
	OUSED
	OWHILE
	OXOR
	ONEG
	OCOM
	OPOS
	OELEM

	OTST /* used in some compilers */
	OINDEX
	OFAS
	OREGPAIR
	OROTL

	OEND
)
const (
	TXXX = iota
	TCHAR
	TUCHAR
	TSHORT
	TUSHORT
	TINT
	TUINT
	TLONG
	TULONG
	TVLONG
	TUVLONG
	TFLOAT
	TDOUBLE
	TIND
	TFUNC
	TARRAY
	TVOID
	TSTRUCT
	TUNION
	TENUM
	NTYPE

	TAUTO = NTYPE
	TEXTERN
	TSTATIC
	TTYPEDEF
	TTYPESTR
	TREGISTER
	TCONSTNT
	TVOLATILE
	TUNSIGNED
	TSIGNED
	TDOT
	TFILE
	TOLD
	NALLTYPES
)
const (
	CXXX = iota
	CAUTO
	CEXTERN
	CGLOBL
	CSTATIC
	CLOCAL
	CTYPEDEF
	CTYPESTR
	CPARAM
	CSELEM
	CLABEL
	CEXREG
	NCTYPES
)
const (
	GXXX     = 0
	GCONSTNT = 1 << iota
	GVOLATILE
	NGTYPES

	GINCOMPLETE = 1 << 2
)
const (
	BCHAR     int64 = 1 << TCHAR
	BUCHAR    int64 = 1 << TUCHAR
	BSHORT    int64 = 1 << TSHORT
	BUSHORT   int64 = 1 << TUSHORT
	BINT      int64 = 1 << TINT
	BUINT     int64 = 1 << TUINT
	BLONG     int64 = 1 << TLONG
	BULONG    int64 = 1 << TULONG
	BVLONG    int64 = 1 << TVLONG
	BUVLONG   int64 = 1 << TUVLONG
	BFLOAT    int64 = 1 << TFLOAT
	BDOUBLE   int64 = 1 << TDOUBLE
	BIND      int64 = 1 << TIND
	BFUNC     int64 = 1 << TFUNC
	BARRAY    int64 = 1 << TARRAY
	BVOID     int64 = 1 << TVOID
	BSTRUCT   int64 = 1 << TSTRUCT
	BUNION    int64 = 1 << TUNION
	BENUM     int64 = 1 << TENUM
	BFILE     int64 = 1 << TFILE
	BDOT      int64 = 1 << TDOT
	BCONSTNT  int64 = 1 << TCONSTNT
	BVOLATILE int64 = 1 << TVOLATILE
	BUNSIGNED int64 = 1 << TUNSIGNED
	BSIGNED   int64 = 1 << TSIGNED
	BAUTO     int64 = 1 << TAUTO
	BEXTERN   int64 = 1 << TEXTERN
	BSTATIC   int64 = 1 << TSTATIC
	BTYPEDEF  int64 = 1 << TTYPEDEF
	BTYPESTR  int64 = 1 << TTYPESTR
	BREGISTER int64 = 1 << TREGISTER

	BINTEGER = BCHAR | BUCHAR | BSHORT | BUSHORT | BINT | BUINT | BLONG | BULONG | BVLONG | BUVLONG
	BNUMBER  = BINTEGER | BFLOAT | BDOUBLE

	/* these can be overloaded with complex types */

	BCLASS = BAUTO | BEXTERN | BSTATIC | BTYPEDEF | BTYPESTR | BREGISTER
	BGARB  = BCONSTNT | BVOLATILE
)

type Funct struct {
	sym    [OEND]*Sym
	castto [NTYPE]*Sym
	castfr [NTYPE]*Sym
}

type Dynimp struct {
	local  []byte
	remote []byte
	path   []byte
}

type Dynexp struct {
	local  []byte
	remote []byte
}

/*
EXTERN	Dynimp	*dynimp;
EXTERN	int	ndynimp;
*/

/*
EXTERN	Dynexp	*dynexp;
EXTERN	int	ndynexp;
*/

// EXTERN struct
// {
// 	Type*	tenum;		/* type of entire enum 
// 	Type*	cenum;		/* type of current enum run */
// 	vlong	lastenum;	/* value of current enum */
// 	double	floatenum;	/* value of current enum */
// } en;

// EXTERN	int	autobn;
// EXTERN	int32	autoffset;
// EXTERN	int	blockno;
// EXTERN	Decl*	dclstack;
// EXTERN	int	debug[256];
// EXTERN	Hist*	ehist;
// EXTERN	int32	firstbit;
// EXTERN	Sym*	firstarg;
// EXTERN	Type*	firstargtype;
// EXTERN	Decl*	firstdcl;
// EXTERN	int	fperror;
// EXTERN	Sym*	hash[NHASH];
// EXTERN	char*	hunk;
// EXTERN	char**	include;
// EXTERN	Io*	iofree;
// EXTERN	Io*	ionext;
// EXTERN	Io*	iostack;
// EXTERN	int32	lastbit;
// EXTERN	char	lastclass;
// EXTERN	Type*	lastdcl;
// EXTERN	int32	lastfield;
// EXTERN	Type*	lasttype;
// EXTERN	int32	lineno;
// EXTERN	int32	nearln;
// EXTERN	int	nerrors;
// EXTERN	int	newflag;
// EXTERN	int32	nhunk;
// EXTERN	int	ninclude;
// EXTERN	Node*	nodproto;
// EXTERN	Node*	nodcast;
// EXTERN	int32	nsymb;
// EXTERN	Biobuf	outbuf;
// EXTERN	Biobuf	diagbuf;
// EXTERN	char*	outfile;
// EXTERN	char*	pathname;
// EXTERN	int	peekc;
// EXTERN	int32	stkoff;
// EXTERN	Type*	strf;
// EXTERN	Type*	strl;
// EXTERN	char*	symb;
// EXTERN	Sym*	symstring;
// EXTERN	int	taggen;
// EXTERN	Type*	tfield;
// EXTERN	Type*	tufield;
// EXTERN	int	thechar;
// EXTERN	char*	thestring;
// EXTERN	Type*	thisfn;
// EXTERN	int32	thunk;
// EXTERN	Type*	types[NALLTYPES];
// EXTERN	Type*	fntypes[NALLTYPES];
// EXTERN	Node*	initlist;
// EXTERN	Term	term[NTERM];
// EXTERN	int	nterm;
// EXTERN	int	packflg;
// EXTERN	int	fproundflg;
// EXTERN	int	textflag;
// EXTERN	int	dataflag;
// EXTERN	int	flag_largemodel;
// EXTERN	int	ncontin;
// EXTERN	int	canreach;
// EXTERN	int	warnreach;
// EXTERN	Bits	zbits;

// extern	char	*onames[], *tnames[], *gnames[];
// extern	char	*cnames[], *qnames[], *bnames[];
// extern	uchar	tab[NTYPE][NTYPE];
// extern	uchar	comrel[], invrel[], logrel[];
// extern	int32	ncast[], tadd[], tand[];
// extern	int32	targ[], tasadd[], tasign[], tcast[];
// extern	int32	tdot[], tfunct[], tindir[], tmul[];
// extern	int32	tnot[], trel[], tsub[];

// extern	uchar	typeaf[];
// extern	uchar	typefd[];
// extern	uchar	typei[];
// extern	uchar	typesu[];
// extern	uchar	typesuv[];
// extern	uchar	typeu[];
// extern	uchar	typev[];
// extern	uchar	typec[];
// extern	uchar	typeh[];
// extern	uchar	typeil[];
// extern	uchar	typeilp[];
// extern	uchar	typechl[];
// extern	uchar	typechlv[];
// extern	uchar	typechlvp[];
// extern	uchar	typechlp[];
// extern	uchar	typechlpfd[];

// EXTERN	uchar*	typeword;
// EXTERN	uchar*	typecmplx;

// extern	uint32	thash1;
// extern	uint32	thash2;
// extern	uint32	thash3;
// extern	uint32	thash[];

// extern	schar	ewidth[];
