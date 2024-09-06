package main

// Openldap - Open source suite of directory software
// Homepage: https://www.openldap.org/software/

import (
	"fmt"
	
	"os/exec"
)

func installOpenldap() {
	// Método 1: Descargar y extraer .tar.gz
	openldap_tar_url := "https://www.openldap.org/software/download/OpenLDAP/openldap-release/openldap-2.6.8.tgz"
	openldap_cmd_tar := exec.Command("curl", "-L", openldap_tar_url, "-o", "package.tar.gz")
	err := openldap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openldap_zip_url := "https://www.openldap.org/software/download/OpenLDAP/openldap-release/openldap-2.6.8.tgz"
	openldap_cmd_zip := exec.Command("curl", "-L", openldap_zip_url, "-o", "package.zip")
	err = openldap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openldap_bin_url := "https://www.openldap.org/software/download/OpenLDAP/openldap-release/openldap-2.6.8.tgz"
	openldap_cmd_bin := exec.Command("curl", "-L", openldap_bin_url, "-o", "binary.bin")
	err = openldap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openldap_src_url := "https://www.openldap.org/software/download/OpenLDAP/openldap-release/openldap-2.6.8.tgz"
	openldap_cmd_src := exec.Command("curl", "-L", openldap_src_url, "-o", "source.tar.gz")
	err = openldap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openldap_cmd_direct := exec.Command("./binary")
	err = openldap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
