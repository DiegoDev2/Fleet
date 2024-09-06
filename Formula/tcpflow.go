package main

// Tcpflow - TCP/IP packet demultiplexer
// Homepage: https://github.com/simsong/tcpflow

import (
	"fmt"
	
	"os/exec"
)

func installTcpflow() {
	// Método 1: Descargar y extraer .tar.gz
	tcpflow_tar_url := "https://corp.digitalcorpora.org/downloads/tcpflow/tcpflow-1.6.1.tar.gz"
	tcpflow_cmd_tar := exec.Command("curl", "-L", tcpflow_tar_url, "-o", "package.tar.gz")
	err := tcpflow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcpflow_zip_url := "https://corp.digitalcorpora.org/downloads/tcpflow/tcpflow-1.6.1.zip"
	tcpflow_cmd_zip := exec.Command("curl", "-L", tcpflow_zip_url, "-o", "package.zip")
	err = tcpflow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcpflow_bin_url := "https://corp.digitalcorpora.org/downloads/tcpflow/tcpflow-1.6.1.bin"
	tcpflow_cmd_bin := exec.Command("curl", "-L", tcpflow_bin_url, "-o", "binary.bin")
	err = tcpflow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcpflow_src_url := "https://corp.digitalcorpora.org/downloads/tcpflow/tcpflow-1.6.1.src.tar.gz"
	tcpflow_cmd_src := exec.Command("curl", "-L", tcpflow_src_url, "-o", "source.tar.gz")
	err = tcpflow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcpflow_cmd_direct := exec.Command("./binary")
	err = tcpflow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
