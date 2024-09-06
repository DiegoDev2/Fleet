package main

// Acl - Commands for manipulating POSIX access control lists
// Homepage: https://savannah.nongnu.org/projects/acl/

import (
	"fmt"
	
	"os/exec"
)

func installAcl() {
	// Método 1: Descargar y extraer .tar.gz
	acl_tar_url := "https://download.savannah.nongnu.org/releases/acl/acl-2.3.2.tar.gz"
	acl_cmd_tar := exec.Command("curl", "-L", acl_tar_url, "-o", "package.tar.gz")
	err := acl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	acl_zip_url := "https://download.savannah.nongnu.org/releases/acl/acl-2.3.2.zip"
	acl_cmd_zip := exec.Command("curl", "-L", acl_zip_url, "-o", "package.zip")
	err = acl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	acl_bin_url := "https://download.savannah.nongnu.org/releases/acl/acl-2.3.2.bin"
	acl_cmd_bin := exec.Command("curl", "-L", acl_bin_url, "-o", "binary.bin")
	err = acl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	acl_src_url := "https://download.savannah.nongnu.org/releases/acl/acl-2.3.2.src.tar.gz"
	acl_cmd_src := exec.Command("curl", "-L", acl_src_url, "-o", "source.tar.gz")
	err = acl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	acl_cmd_direct := exec.Command("./binary")
	err = acl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: attr")
exec.Command("latte", "install", "attr")
}
