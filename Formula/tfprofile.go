package main

// TfProfile - CLI tool to profile Terraform runs
// Homepage: https://github.com/datarootsio/tf-profile

import (
	"fmt"
	
	"os/exec"
)

func installTfProfile() {
	// Método 1: Descargar y extraer .tar.gz
	tfprofile_tar_url := "https://github.com/datarootsio/tf-profile/archive/refs/tags/v0.4.0.tar.gz"
	tfprofile_cmd_tar := exec.Command("curl", "-L", tfprofile_tar_url, "-o", "package.tar.gz")
	err := tfprofile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tfprofile_zip_url := "https://github.com/datarootsio/tf-profile/archive/refs/tags/v0.4.0.zip"
	tfprofile_cmd_zip := exec.Command("curl", "-L", tfprofile_zip_url, "-o", "package.zip")
	err = tfprofile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tfprofile_bin_url := "https://github.com/datarootsio/tf-profile/archive/refs/tags/v0.4.0.bin"
	tfprofile_cmd_bin := exec.Command("curl", "-L", tfprofile_bin_url, "-o", "binary.bin")
	err = tfprofile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tfprofile_src_url := "https://github.com/datarootsio/tf-profile/archive/refs/tags/v0.4.0.src.tar.gz"
	tfprofile_cmd_src := exec.Command("curl", "-L", tfprofile_src_url, "-o", "source.tar.gz")
	err = tfprofile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tfprofile_cmd_direct := exec.Command("./binary")
	err = tfprofile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
