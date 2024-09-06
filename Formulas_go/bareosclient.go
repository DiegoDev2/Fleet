package main

// BareosClient - Client for Bareos (Backup Archiving REcovery Open Sourced)
// Homepage: https://www.bareos.org/

import (
	"fmt"
	
	"os/exec"
)

func installBareosClient() {
	// Método 1: Descargar y extraer .tar.gz
	bareosclient_tar_url := "https://github.com/bareos/bareos/archive/refs/tags/Release/23.0.3.tar.gz"
	bareosclient_cmd_tar := exec.Command("curl", "-L", bareosclient_tar_url, "-o", "package.tar.gz")
	err := bareosclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bareosclient_zip_url := "https://github.com/bareos/bareos/archive/refs/tags/Release/23.0.3.zip"
	bareosclient_cmd_zip := exec.Command("curl", "-L", bareosclient_zip_url, "-o", "package.zip")
	err = bareosclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bareosclient_bin_url := "https://github.com/bareos/bareos/archive/refs/tags/Release/23.0.3.bin"
	bareosclient_cmd_bin := exec.Command("curl", "-L", bareosclient_bin_url, "-o", "binary.bin")
	err = bareosclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bareosclient_src_url := "https://github.com/bareos/bareos/archive/refs/tags/Release/23.0.3.src.tar.gz"
	bareosclient_cmd_src := exec.Command("curl", "-L", bareosclient_src_url, "-o", "source.tar.gz")
	err = bareosclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bareosclient_cmd_direct := exec.Command("./binary")
	err = bareosclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: jansson")
exec.Command("latte", "install", "jansson")
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: acl")
exec.Command("latte", "install", "acl")
}
