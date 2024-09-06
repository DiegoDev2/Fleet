package main

// Verapdf - Open-source industry-supported PDF/A validation
// Homepage: https://verapdf.org/home/

import (
	"fmt"
	
	"os/exec"
)

func installVerapdf() {
	// Método 1: Descargar y extraer .tar.gz
	verapdf_tar_url := "https://github.com/veraPDF/veraPDF-apps/archive/refs/tags/v1.26.2.tar.gz"
	verapdf_cmd_tar := exec.Command("curl", "-L", verapdf_tar_url, "-o", "package.tar.gz")
	err := verapdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	verapdf_zip_url := "https://github.com/veraPDF/veraPDF-apps/archive/refs/tags/v1.26.2.zip"
	verapdf_cmd_zip := exec.Command("curl", "-L", verapdf_zip_url, "-o", "package.zip")
	err = verapdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	verapdf_bin_url := "https://github.com/veraPDF/veraPDF-apps/archive/refs/tags/v1.26.2.bin"
	verapdf_cmd_bin := exec.Command("curl", "-L", verapdf_bin_url, "-o", "binary.bin")
	err = verapdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	verapdf_src_url := "https://github.com/veraPDF/veraPDF-apps/archive/refs/tags/v1.26.2.src.tar.gz"
	verapdf_cmd_src := exec.Command("curl", "-L", verapdf_src_url, "-o", "source.tar.gz")
	err = verapdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	verapdf_cmd_direct := exec.Command("./binary")
	err = verapdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
	fmt.Println("Instalando dependencia: openjdk@21")
	exec.Command("latte", "install", "openjdk@21").Run()
}
