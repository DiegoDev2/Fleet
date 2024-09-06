package main

// Ipmitool - Utility for IPMI control with kernel driver or LAN interface
// Homepage: https://codeberg.org/IPMITool/ipmitool

import (
	"fmt"
	
	"os/exec"
)

func installIpmitool() {
	// Método 1: Descargar y extraer .tar.gz
	ipmitool_tar_url := "https://codeberg.org/IPMITool/ipmitool/archive/IPMITOOL_1_8_19.tar.gz"
	ipmitool_cmd_tar := exec.Command("curl", "-L", ipmitool_tar_url, "-o", "package.tar.gz")
	err := ipmitool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipmitool_zip_url := "https://codeberg.org/IPMITool/ipmitool/archive/IPMITOOL_1_8_19.zip"
	ipmitool_cmd_zip := exec.Command("curl", "-L", ipmitool_zip_url, "-o", "package.zip")
	err = ipmitool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipmitool_bin_url := "https://codeberg.org/IPMITool/ipmitool/archive/IPMITOOL_1_8_19.bin"
	ipmitool_cmd_bin := exec.Command("curl", "-L", ipmitool_bin_url, "-o", "binary.bin")
	err = ipmitool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipmitool_src_url := "https://codeberg.org/IPMITool/ipmitool/archive/IPMITOOL_1_8_19.src.tar.gz"
	ipmitool_cmd_src := exec.Command("curl", "-L", ipmitool_src_url, "-o", "source.tar.gz")
	err = ipmitool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipmitool_cmd_direct := exec.Command("./binary")
	err = ipmitool_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
