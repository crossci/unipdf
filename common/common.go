//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package common contains common properties used by the subpackages.
package common ;import (_e "fmt";_dd "io";_c "os";_d "path/filepath";_b "runtime";_ab "time";);

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };func (_gdb WriterLogger )logToWriter (_da _dd .Writer ,_fg string ,_gf string ,_dfd ...interface{}){_ee (_da ,_fg ,_gf ,_dfd );};

// Error logs error message.
func (_bcb WriterLogger )Error (format string ,args ...interface{}){if _bcb .LogLevel >=LogLevelError {_gc :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_bcb .logToWriter (_bcb .Output ,_gc ,format ,args ...);};};const _egd ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";


// DummyLogger does nothing.
type DummyLogger struct{};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// LogLevel is the verbosity level for logging.
type LogLevel int ;const _gbe =2025;

// Debug logs debug message.
func (_dce ConsoleLogger )Debug (format string ,args ...interface{}){if _dce .LogLevel >=LogLevelDebug {_f :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_dce .output (_c .Stdout ,_f ,format ,args ...);};};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// Notice logs notice message.
func (_df WriterLogger )Notice (format string ,args ...interface{}){if _df .LogLevel >=LogLevelNotice {_de :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_df .logToWriter (_df .Output ,_de ,format ,args ...);};};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// Info logs info message.
func (_def WriterLogger )Info (format string ,args ...interface{}){if _def .LogLevel >=LogLevelInfo {_fc :="\u005bI\u004e\u0046\u004f\u005d\u0020";_def .logToWriter (_def .Output ,_fc ,format ,args ...);};};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_g string ,_ge ...interface{});Warning (_bd string ,_ae ...interface{});Notice (_gg string ,_bc ...interface{});Info (_ea string ,_ac ...interface{});Debug (_cf string ,_ce ...interface{});Trace (_acf string ,_dc ...interface{});
IsLogLevel (_eg LogLevel )bool ;};

// Notice logs notice message.
func (_bdd ConsoleLogger )Notice (format string ,args ...interface{}){if _bdd .LogLevel >=LogLevelNotice {_cc :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_bdd .output (_c .Stdout ,_cc ,format ,args ...);};};

// Trace logs trace message.
func (_dge WriterLogger )Trace (format string ,args ...interface{}){if _dge .LogLevel >=LogLevelTrace {_fe :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_dge .logToWriter (_dge .Output ,_fe ,format ,args ...);};};func (_ceb ConsoleLogger )output (_af _dd .Writer ,_eae string ,_ccd string ,_fb ...interface{}){_ee (_af ,_eae ,_ccd ,_fb ...);
};var Log Logger =DummyLogger {};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _dd .Writer ;};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_ag ConsoleLogger )IsLogLevel (level LogLevel )bool {return _ag .LogLevel >=level };

// Info logs info message.
func (_cd ConsoleLogger )Info (format string ,args ...interface{}){if _cd .LogLevel >=LogLevelInfo {_gd :="\u005bI\u004e\u0046\u004f\u005d\u0020";_cd .output (_c .Stdout ,_gd ,format ,args ...);};};

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat (t _ab .Time )string {return t .Format (_egd )+"\u0020\u0055\u0054\u0043"};const _aae =6;

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// Warning logs warning message.
func (_ba ConsoleLogger )Warning (format string ,args ...interface{}){if _ba .LogLevel >=LogLevelWarning {_aa :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ba .output (_c .Stdout ,_aa ,format ,args ...);};};

// Debug logs debug message.
func (_fbe WriterLogger )Debug (format string ,args ...interface{}){if _fbe .LogLevel >=LogLevelDebug {_abg :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_fbe .logToWriter (_fbe .Output ,_abg ,format ,args ...);};};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};const _fdg =23;const _ad =30;var ReleasedAt =_ab .Date (_gbe ,_aae ,_fdg ,_feb ,_ad ,0,0,_ab .UTC );

// Trace logs trace message.
func (_dg ConsoleLogger )Trace (format string ,args ...interface{}){if _dg .LogLevel >=LogLevelTrace {_fd :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_dg .output (_c .Stdout ,_fd ,format ,args ...);};};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};const Version ="\u0034\u002e\u0031.\u0030";const _feb =15;

// Warning logs warning message.
func (_fda WriterLogger )Warning (format string ,args ...interface{}){if _fda .LogLevel >=LogLevelWarning {_ged :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_fda .logToWriter (_fda .Output ,_ged ,format ,args ...);};};

// Error logs error message.
func (_agc ConsoleLogger )Error (format string ,args ...interface{}){if _agc .LogLevel >=LogLevelError {_cfb :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_agc .output (_c .Stdout ,_cfb ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_gb WriterLogger )IsLogLevel (level LogLevel )bool {return _gb .LogLevel >=level };

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _dd .Writer )*WriterLogger {_eb :=WriterLogger {Output :writer ,LogLevel :logLevel };return &_eb ;};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;
LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);func _ee (_cce _dd .Writer ,_dgg string ,_gfc string ,_ef ...interface{}){_ ,_aga ,_be ,_ec :=_b .Caller (3);if !_ec {_aga ="\u003f\u003f\u003f";_be =0;}else {_aga =_d .Base (_aga );};_aag :=_e .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_dgg ,_aga ,_be )+_gfc +"\u000a";
_e .Fprintf (_cce ,_aag ,_ef ...);};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };