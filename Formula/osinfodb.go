package main

// OsinfoDb - Osinfo database of operating systems for virtualization provisioning tools
// Homepage: https://libosinfo.org/

import (
	"fmt"
	
	"os/exec"
)

func installOsinfoDb() {
	// Método 1: Descargar y extraer .tar.gz
	osinfodb_tar_url := "https://releases.pagure.org/libosinfo/osinfo-db-20240701.tar.xz"
	osinfodb_cmd_tar := exec.Command("curl", "-L", osinfodb_tar_url, "-o", "package.tar.gz")
	err := osinfodb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osinfodb_zip_url := "https://releases.pagure.org/libosinfo/osinfo-db-20240701.tar.xz"
	osinfodb_cmd_zip := exec.Command("curl", "-L", osinfodb_zip_url, "-o", "package.zip")
	err = osinfodb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osinfodb_bin_url := "https://releases.pagure.org/libosinfo/osinfo-db-20240701.tar.xz"
	osinfodb_cmd_bin := exec.Command("curl", "-L", osinfodb_bin_url, "-o", "binary.bin")
	err = osinfodb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osinfodb_src_url := "https://releases.pagure.org/libosinfo/osinfo-db-20240701.tar.xz"
	osinfodb_cmd_src := exec.Command("curl", "-L", osinfodb_src_url, "-o", "source.tar.gz")
	err = osinfodb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osinfodb_cmd_direct := exec.Command("./binary")
	err = osinfodb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: osinfo-db-tools")
	exec.Command("latte", "install", "osinfo-db-tools").Run()
}
