package main

// AmdatuBootstrap - Bootstrapping OSGi development
// Homepage: https://bitbucket.org/amdatuadm/amdatu-bootstrap/

import (
	"fmt"
	
	"os/exec"
)

func installAmdatuBootstrap() {
	// Método 1: Descargar y extraer .tar.gz
	amdatubootstrap_tar_url := "https://bitbucket.org/amdatuadm/amdatu-bootstrap/downloads/bootstrap-bin-r9.zip"
	amdatubootstrap_cmd_tar := exec.Command("curl", "-L", amdatubootstrap_tar_url, "-o", "package.tar.gz")
	err := amdatubootstrap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	amdatubootstrap_zip_url := "https://bitbucket.org/amdatuadm/amdatu-bootstrap/downloads/bootstrap-bin-r9.zip"
	amdatubootstrap_cmd_zip := exec.Command("curl", "-L", amdatubootstrap_zip_url, "-o", "package.zip")
	err = amdatubootstrap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	amdatubootstrap_bin_url := "https://bitbucket.org/amdatuadm/amdatu-bootstrap/downloads/bootstrap-bin-r9.zip"
	amdatubootstrap_cmd_bin := exec.Command("curl", "-L", amdatubootstrap_bin_url, "-o", "binary.bin")
	err = amdatubootstrap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	amdatubootstrap_src_url := "https://bitbucket.org/amdatuadm/amdatu-bootstrap/downloads/bootstrap-bin-r9.zip"
	amdatubootstrap_cmd_src := exec.Command("curl", "-L", amdatubootstrap_src_url, "-o", "source.tar.gz")
	err = amdatubootstrap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	amdatubootstrap_cmd_direct := exec.Command("./binary")
	err = amdatubootstrap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@8")
exec.Command("latte", "install", "openjdk@8")
}
