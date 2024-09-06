package main

// XmlrpcC - Lightweight RPC library (based on XML and HTTP)
// Homepage: https://xmlrpc-c.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installXmlrpcC() {
	// Método 1: Descargar y extraer .tar.gz
	xmlrpcc_tar_url := "https://downloads.sourceforge.net/project/xmlrpc-c/Xmlrpc-c%20Super%20Stable/1.59.03/xmlrpc-c-1.59.03.tgz"
	xmlrpcc_cmd_tar := exec.Command("curl", "-L", xmlrpcc_tar_url, "-o", "package.tar.gz")
	err := xmlrpcc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmlrpcc_zip_url := "https://downloads.sourceforge.net/project/xmlrpc-c/Xmlrpc-c%20Super%20Stable/1.59.03/xmlrpc-c-1.59.03.tgz"
	xmlrpcc_cmd_zip := exec.Command("curl", "-L", xmlrpcc_zip_url, "-o", "package.zip")
	err = xmlrpcc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmlrpcc_bin_url := "https://downloads.sourceforge.net/project/xmlrpc-c/Xmlrpc-c%20Super%20Stable/1.59.03/xmlrpc-c-1.59.03.tgz"
	xmlrpcc_cmd_bin := exec.Command("curl", "-L", xmlrpcc_bin_url, "-o", "binary.bin")
	err = xmlrpcc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmlrpcc_src_url := "https://downloads.sourceforge.net/project/xmlrpc-c/Xmlrpc-c%20Super%20Stable/1.59.03/xmlrpc-c-1.59.03.tgz"
	xmlrpcc_cmd_src := exec.Command("curl", "-L", xmlrpcc_src_url, "-o", "source.tar.gz")
	err = xmlrpcc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmlrpcc_cmd_direct := exec.Command("./binary")
	err = xmlrpcc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
