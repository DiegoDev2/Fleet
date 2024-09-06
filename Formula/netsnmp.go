package main

// NetSnmp - Implements SNMP v1, v2c, and v3, using IPv4 and IPv6
// Homepage: http://www.net-snmp.org/

import (
	"fmt"
	
	"os/exec"
)

func installNetSnmp() {
	// Método 1: Descargar y extraer .tar.gz
	netsnmp_tar_url := "https://downloads.sourceforge.net/project/net-snmp/net-snmp/5.9.4/net-snmp-5.9.4.tar.gz"
	netsnmp_cmd_tar := exec.Command("curl", "-L", netsnmp_tar_url, "-o", "package.tar.gz")
	err := netsnmp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netsnmp_zip_url := "https://downloads.sourceforge.net/project/net-snmp/net-snmp/5.9.4/net-snmp-5.9.4.zip"
	netsnmp_cmd_zip := exec.Command("curl", "-L", netsnmp_zip_url, "-o", "package.zip")
	err = netsnmp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netsnmp_bin_url := "https://downloads.sourceforge.net/project/net-snmp/net-snmp/5.9.4/net-snmp-5.9.4.bin"
	netsnmp_cmd_bin := exec.Command("curl", "-L", netsnmp_bin_url, "-o", "binary.bin")
	err = netsnmp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netsnmp_src_url := "https://downloads.sourceforge.net/project/net-snmp/net-snmp/5.9.4/net-snmp-5.9.4.src.tar.gz"
	netsnmp_cmd_src := exec.Command("curl", "-L", netsnmp_src_url, "-o", "source.tar.gz")
	err = netsnmp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netsnmp_cmd_direct := exec.Command("./binary")
	err = netsnmp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
