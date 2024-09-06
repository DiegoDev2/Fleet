package main

// SyslogNg - Log daemon with advanced processing pipeline and a wide range of I/O methods
// Homepage: https://www.syslog-ng.com

import (
	"fmt"
	
	"os/exec"
)

func installSyslogNg() {
	// Método 1: Descargar y extraer .tar.gz
	syslogng_tar_url := "https://github.com/syslog-ng/syslog-ng/releases/download/syslog-ng-4.8.0/syslog-ng-4.8.0.tar.gz"
	syslogng_cmd_tar := exec.Command("curl", "-L", syslogng_tar_url, "-o", "package.tar.gz")
	err := syslogng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	syslogng_zip_url := "https://github.com/syslog-ng/syslog-ng/releases/download/syslog-ng-4.8.0/syslog-ng-4.8.0.zip"
	syslogng_cmd_zip := exec.Command("curl", "-L", syslogng_zip_url, "-o", "package.zip")
	err = syslogng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	syslogng_bin_url := "https://github.com/syslog-ng/syslog-ng/releases/download/syslog-ng-4.8.0/syslog-ng-4.8.0.bin"
	syslogng_cmd_bin := exec.Command("curl", "-L", syslogng_bin_url, "-o", "binary.bin")
	err = syslogng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	syslogng_src_url := "https://github.com/syslog-ng/syslog-ng/releases/download/syslog-ng-4.8.0/syslog-ng-4.8.0.src.tar.gz"
	syslogng_cmd_src := exec.Command("curl", "-L", syslogng_src_url, "-o", "source.tar.gz")
	err = syslogng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	syslogng_cmd_direct := exec.Command("./binary")
	err = syslogng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: grpc")
exec.Command("latte", "install", "grpc")
	fmt.Println("Instalando dependencia: hiredis")
exec.Command("latte", "install", "hiredis")
	fmt.Println("Instalando dependencia: ivykis")
exec.Command("latte", "install", "ivykis")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: libdbi")
exec.Command("latte", "install", "libdbi")
	fmt.Println("Instalando dependencia: libmaxminddb")
exec.Command("latte", "install", "libmaxminddb")
	fmt.Println("Instalando dependencia: libnet")
exec.Command("latte", "install", "libnet")
	fmt.Println("Instalando dependencia: libpaho-mqtt")
exec.Command("latte", "install", "libpaho-mqtt")
	fmt.Println("Instalando dependencia: librdkafka")
exec.Command("latte", "install", "librdkafka")
	fmt.Println("Instalando dependencia: mongo-c-driver")
exec.Command("latte", "install", "mongo-c-driver")
	fmt.Println("Instalando dependencia: net-snmp")
exec.Command("latte", "install", "net-snmp")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: rabbitmq-c")
exec.Command("latte", "install", "rabbitmq-c")
	fmt.Println("Instalando dependencia: riemann-client")
exec.Command("latte", "install", "riemann-client")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
