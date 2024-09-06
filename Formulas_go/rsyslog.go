package main

// Rsyslog - Enhanced, multi-threaded syslogd
// Homepage: https://www.rsyslog.com/

import (
	"fmt"
	
	"os/exec"
)

func installRsyslog() {
	// Método 1: Descargar y extraer .tar.gz
	rsyslog_tar_url := "https://www.rsyslog.com/files/download/rsyslog/rsyslog-8.2408.0.tar.gz"
	rsyslog_cmd_tar := exec.Command("curl", "-L", rsyslog_tar_url, "-o", "package.tar.gz")
	err := rsyslog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rsyslog_zip_url := "https://www.rsyslog.com/files/download/rsyslog/rsyslog-8.2408.0.zip"
	rsyslog_cmd_zip := exec.Command("curl", "-L", rsyslog_zip_url, "-o", "package.zip")
	err = rsyslog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rsyslog_bin_url := "https://www.rsyslog.com/files/download/rsyslog/rsyslog-8.2408.0.bin"
	rsyslog_cmd_bin := exec.Command("curl", "-L", rsyslog_bin_url, "-o", "binary.bin")
	err = rsyslog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rsyslog_src_url := "https://www.rsyslog.com/files/download/rsyslog/rsyslog-8.2408.0.src.tar.gz"
	rsyslog_cmd_src := exec.Command("curl", "-L", rsyslog_src_url, "-o", "source.tar.gz")
	err = rsyslog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rsyslog_cmd_direct := exec.Command("./binary")
	err = rsyslog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libestr")
exec.Command("latte", "install", "libestr")
	fmt.Println("Instalando dependencia: libfastjson")
exec.Command("latte", "install", "libfastjson")
}
